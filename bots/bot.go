package bots

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/emptyinterface/slk"
)

type (
	Bot struct {
		Client *slk.Client
		User   *slk.User

		// either Channel or Group will be non-nil
		Channel *slk.Channel
		Group   *slk.Group

		authedUsers      []*slk.User
		commands         map[string]*Command
		intervalCommands []*IntervalCommand
		stopIntervals    chan struct{}
		listeners        map[*regexp.Regexp]*Listener

		rtm *slk.RTMClient
	}

	IntervalCommand struct {
		command     string
		description string
		dur         time.Duration
		interval    *time.Ticker
		handler     func(w io.Writer)
		lastRun     time.Time
	}

	Command struct {
		command     string
		description string
		handler     func(w io.Writer, msg *slk.MessageEvent, from *slk.User, args ...string)
		lastRun     string
	}

	Listener struct {
		pattern     *regexp.Regexp
		description string
		handler     func(w io.Writer, msg *slk.MessageEvent, from *slk.User)
	}
)

const (
	ColorSuccess = "#5cb85c"
	ColorInfo    = "#5bc0de"
	ColorWarning = "#f0ad4e"
	ColorError   = "#d9534f"

	helpDateTimeFormat = `Mon, Jan 02, 3:04:05 PM MST 2006`
)

func NewBot(token, channel string) (*Bot, error) {

	bot := &Bot{
		Client:    slk.NewClient(token),
		commands:  make(map[string]*Command),
		listeners: make(map[*regexp.Regexp]*Listener),
	}

	authResp, err := bot.Client.AuthTest()
	if err != nil {
		return nil, fmt.Errorf("auth test failed: %s", err)
	}
	userResp, err := bot.Client.UsersInfo(authResp.UserID)
	if err != nil {
		return nil, err
	}
	bot.User = userResp.User

	// try private channels first
	groupListResp, err := bot.Client.GroupsList(url.Values{"exclude_archived": {"1"}})
	if err != nil {
		return nil, err
	}
	for _, g := range groupListResp.Groups {
		if g.Name == channel {
			bot.Group = g
			break
		}
	}
	if bot.Group == nil {
		chanListResp, err := bot.Client.ChannelsList(url.Values{"exclude_archived": {"1"}})
		if err != nil {
			return nil, err
		}
		for _, c := range chanListResp.Channels {
			if c.Name == channel {
				bot.Channel = c
				break
			}
		}
		if bot.Channel == nil {
			return nil, fmt.Errorf("Unable to locate %s channel/group", channel)
		}
	}

	bot.AddCommand("help", "shows all commands",
		func(_ io.Writer, _ *slk.MessageEvent, _ *slk.User, _ ...string) {
			var names []string
			for name, _ := range bot.commands {
				names = append(names, name)
			}
			sort.Strings(names)
			var buf bytes.Buffer
			fmt.Fprintln(&buf, "*Help*")
			fmt.Fprintln(&buf, ">*Commands:*")
			for _, name := range names {
				fmt.Fprintf(&buf, ">\t`%s`: %s\n", name, bot.commands[name].description)
				if lastrun := bot.commands[name].lastRun; len(lastrun) > 0 {
					fmt.Fprintf(&buf, ">\t\t\t\t_Last run: %s_\n", bot.commands[name].lastRun)
				}
			}
			fmt.Fprintln(&buf, ">*Interval Commands:*")
			for _, ic := range bot.intervalCommands {
				fmt.Fprintf(&buf, ">\t`Every %s`: %s\n", ic.dur, ic.description)
				if !ic.lastRun.IsZero() {
					fmt.Fprintf(&buf, ">\t\t\t\t_Last run: %s_\n", ic.lastRun.Format(helpDateTimeFormat))
				}
			}
			fmt.Fprintln(&buf, ">*Bot Listeners:*")
			for r, l := range bot.listeners {
				fmt.Fprintf(&buf, ">\t`%s`: %s\n", r.String(), l.description)
			}
			bot.SendMessage(buf.String())
		},
	)

	return bot, nil

}

func (bot *Bot) AddAuthorizedUser(userids ...string) error {
	for _, id := range userids {
		resp, err := bot.Client.UsersInfo(id)
		if err != nil {
			return err
		}
		bot.authedUsers = append(bot.authedUsers, resp.User)
	}
	return nil
}

func (bot *Bot) IsAuthorizedUserID(id string) bool {
	for _, user := range bot.authedUsers {
		if user.ID == id {
			return true
		}
	}
	return false
}

func (bot *Bot) AddCommand(command, description string, handler func(w io.Writer, msg *slk.MessageEvent, from *slk.User, args ...string)) {
	bot.commands[command] = &Command{
		command:     command,
		description: description,
		handler:     handler,
	}
}

func (bot *Bot) AddIntervalCommand(interval time.Duration, description string, handler func(w io.Writer)) {
	bot.intervalCommands = append(bot.intervalCommands, &IntervalCommand{
		description: description,
		dur:         interval,
		interval:    time.NewTicker(interval),
		handler:     handler,
	})
}

func (bot *Bot) AddListener(pattern *regexp.Regexp, description string, handler func(w io.Writer, msg *slk.MessageEvent, from *slk.User)) {
	bot.listeners[pattern] = &Listener{
		pattern:     pattern,
		description: description,
		handler:     handler,
	}
}

func (bot *Bot) ChannelID() string {
	if bot.Channel != nil {
		return bot.Channel.ID
	}
	return bot.Group.ID
}

func (bot *Bot) SendMessage(text string) error {
	args := url.Values{}
	args.Set("as_user", "true")
	args.Set("link_names", "1")
	_, err := bot.Client.ChatPostMessage(bot.ChannelID(), text, args)
	if err != nil {
		log.Println("post message err:", err)
	}
	return err
}

func (bot *Bot) SendColorMessage(text, color string) error {
	args := url.Values{}
	args.Set("as_user", "true")
	args.Set("link_names", "1")
	data, err := json.Marshal([]*slk.MessageAttachment{{
		Color:      color,
		Text:       text,
		MarkdownIn: []string{"text"},
	}})
	if err != nil {
		return err
	}
	args.Set("attachments", string(data))
	_, err = bot.Client.ChatPostMessage(bot.ChannelID(), "", args)
	if err != nil {
		log.Println("post message err:", err, "message:", text)
	}
	return err
}

// opt:username // My Bot // Optional //  Name of bot.
// opt:as_user // true // Optional //  Pass true to post the message as the authed user, instead of as a bot
// opt:parse // full // Optional //  Change how messages are treated. See below.
// opt:link_names // 1 // Optional //  Find and link channel names and usernames.
// opt:attachments // [{"pretext": "pre-hello", "text": "text-world"}] // Optional //  Structured message attachments.
// opt:unfurl_links // true // Optional //  Pass true to enable unfurling of primarily text-based content.
// opt:unfurl_media // false // Optional //  Pass false to disable unfurling of media content.
// opt:icon_url // http://lorempixel.com/48/48 // Optional //  URL to an image to use as the icon for this message
// opt:icon_emoji // :chart_with_upwards_trend: // Optional //  emoji to use as the icon for this message. Overridesicon_url
func (bot *Bot) SendMessageOptions(text string, options url.Values) error {
	_, err := bot.Client.ChatPostMessage(bot.ChannelID(), text, options)
	if err != nil {
		log.Println("post message err:", err)
	}
	return err
}

func (bot *Bot) AddReaction(msg *slk.MessageEvent, reaction string) error {
	_, err := bot.Client.ReactionsAdd(reaction, url.Values{
		"channel":   {bot.ChannelID()},
		"timestamp": {msg.Timestamp.SlackString()},
	})
	return err
}

func (bot *Bot) startIntervalCommands() {

	var cases []reflect.SelectCase

	for _, ic := range bot.intervalCommands {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ic.interval.C),
		})
	}

	// add signal for exiting
	bot.stopIntervals = make(chan struct{})
	cases = append(cases, reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(bot.stopIntervals),
	})
	stopI := len(cases) - 1

	for {
		i, _, _ := reflect.Select(cases)
		if i == stopI {
			return
		}
		go func(ic *IntervalCommand) {
			r, w := io.Pipe()
			lines := bufio.NewScanner(r)

			go func() {
				defer func() {
					if e := recover(); e != nil {
						w.CloseWithError(fmt.Errorf("panic: %v", e))
					} else {
						w.Close()
					}
				}()
				ic.handler(w)
			}()

			for lines.Scan() {
				bot.SendMessage(fmt.Sprintf(">%s", lines.Text()))
			}
			if err := lines.Err(); err != nil {
				bot.SendColorMessage(err.Error(), ColorError)
			}
			ic.lastRun = time.Now()
		}(bot.intervalCommands[i])
	}

}

func (bot *Bot) stopIntervalCommands() {
	close(bot.stopIntervals)
}

func (bot *Bot) Run() error {

	go bot.startIntervalCommands()
	defer bot.stopIntervalCommands()

	bot.rtm = bot.Client.NewRTMClient(nil)

	rtmerr := make(chan error)
	go func() { rtmerr <- bot.rtm.Run() }()

	for {
		select {
		case err := <-rtmerr:
			return err
		case msg := <-bot.rtm.Events.Message:

			// ignore messages from other sources
			if msg.Channel != bot.ChannelID() {
				continue
			}

			// skip messages sent by bot
			if msg.User == bot.User.ID {
				continue
			}

			botprefix := fmt.Sprintf("<@%s>:", bot.User.ID)

			// if has prefix, run command after
			if strings.HasPrefix(msg.Text, botprefix) {

				if strings.Contains(msg.Text, "\n") {
					bot.SendMessage("multi-line commands not supported")
					continue
				}

				args := parseDoubleQuotedFields(strings.TrimPrefix(msg.Text, botprefix))
				if len(args) == 0 {
					bot.SendMessage("no command")
					continue
				}

				c, exists := bot.commands[args[0]]
				if !exists {
					bot.SendMessage("unrecognized command")
					continue
				}

				go func() {

					cleanCommand := slk.EscapeMessage(strings.Join(args, " "))

					bot.SendMessage(fmt.Sprintf("_running command_ `%s`", cleanCommand))

					r, w := io.Pipe()
					lines := bufio.NewScanner(r)
					start := time.Now()

					go func() {
						defer func() {
							if err := recover(); err != nil {
								w.CloseWithError(fmt.Errorf("panic: %v", err))
							} else {
								w.Close()
							}
						}()
						if len(bot.authedUsers) > 0 && !bot.IsAuthorizedUserID(msg.User) {
							bot.SendColorMessage("Unauthorized user", ColorWarning)
							return
						}
						resp, err := bot.Client.UsersInfo(msg.User)
						if err != nil {
							w.CloseWithError(fmt.Errorf("Unable to retrieve user: %s", err))
							return
						}
						c.handler(w, msg, resp.User, args[1:]...)
						c.lastRun = fmt.Sprintf("<@%s|%s> ran `%s` on %s",
							resp.User.ID,
							resp.User.Name,
							cleanCommand,
							time.Now().Local().Format(helpDateTimeFormat),
						)
					}()

					for lines.Scan() {
						bot.SendMessage(fmt.Sprintf(">%s", lines.Text()))
					}
					if err := lines.Err(); err != nil {
						bot.SendColorMessage(err.Error(), ColorError)
					}
					bot.SendMessage(fmt.Sprintf("_completed in %s_", time.Since(start)))

				}()

			} else {
				for pattern, l := range bot.listeners {
					if pattern.MatchString(msg.Text) {
						go func(l *Listener) {
							r, w := io.Pipe()
							lines := bufio.NewScanner(r)
							go func() {
								defer func() {
									if err := recover(); err != nil {
										w.CloseWithError(fmt.Errorf("panic: %v", err))
									} else {
										w.Close()
									}
								}()
								resp, err := bot.Client.UsersInfo(msg.User)
								if err != nil {
									fmt.Println(err)
									return
								}
								l.handler(w, msg, resp.User)
							}()
							for lines.Scan() {
								bot.SendMessage(lines.Text())
							}
							if err := lines.Err(); err != nil {
								bot.SendColorMessage(err.Error(), ColorError)
							}
						}(l)
					}
				}
			}

		case err := <-bot.rtm.Events.Error:
			fmt.Println("ERROR", err)
		}
	}

}

func (bot *Bot) Stop() error {
	return bot.rtm.Stop()
}

func parseDoubleQuotedFields(input string) []string {
	var in, escape bool
	output := strings.FieldsFunc(input, func(c rune) bool {
		if escape {
			escape = false
		} else if c == '\\' {
			escape = true
		} else if c == '"' {
			in = !in
			return true
		} else if !in && c == ' ' {
			return true
		}
		return false
	})
	for i, s := range output {
		output[i] = strings.Replace(s, `\"`, `"`, -1)
	}
	return output
}
