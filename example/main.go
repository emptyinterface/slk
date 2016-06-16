package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/emptyinterface/slk"
	"github.com/emptyinterface/slk/bots"
	"github.com/emptyinterface/sq"
)

func main() {

	token, exists := os.LookupEnv("SLACK_TOKEN")
	if !exists {
		log.Fatal("missing SLACK_TOKEN environment variable")
	}

	channel, exists := os.LookupEnv("SLACK_CHANNEL")
	if !exists {
		log.Fatal("missing SLACK_CHANNEL environment variable")
	}

	uids, exists := os.LookupEnv("SLACK_AUTHORIZED_USER_IDS")
	if !exists {
		log.Fatal("missing SLACK_AUTHORIZED_USER_IDS environment variable")
	}

	bot, err := bots.NewBot(token, channel)
	if err != nil {
		log.Fatal(err)
	}

	if err := bot.AddAuthorizedUser(strings.Split(uids, ",")...); err != nil {
		log.Fatal(err)
	}

	bot.AddListener(regexp.MustCompile(`\b`+bot.User.Name+`\b`), "listens for bot's name",
		func(w io.Writer, msg *slk.MessageEvent, from *slk.User) {
			reactions := [...]string{
				"stuck_out_tongue_winking_eye",
				"heart",
			}
			if err := bot.AddReaction(msg, reactions[rand.Intn(len(reactions))]); err != nil {
				fmt.Println(err)
			}
		},
	)

	bot.AddCommand("date", "returns the output of `date` on the local machine",
		func(w io.Writer, _ *slk.MessageEvent, _ *slk.User, _ ...string) {
			cmd := exec.Command("date")
			cmd.Stdout = w
			cmd.Stderr = w
			cmd.Env = []string{}
			if err := cmd.Run(); err != nil {
				fmt.Fprint(w, err)
			}
		},
	)

	bot.AddIntervalCommand(30*time.Second, "check aws status", func() func(io.Writer) {
		problemServices := map[string]string{}
		return func(_ io.Writer) {
			resp, err := http.Get("http://status.aws.amazon.com/")
			if err != nil {
				bot.SendColorMessage(fmt.Sprintf("aws status GET failed: %s", err), "warning")
				return
			}
			if resp.StatusCode != http.StatusOK {
				bot.SendColorMessage(fmt.Sprintf("aws status GET failed: %s", resp.Status), "warning")
				return
			}
			defer resp.Body.Close()
			var page struct {
				Services []struct {
					Name   string `sq:"td:nth-child(2) | text"`
					Status string `sq:"td:nth-child(3) | text"`
				} `sq:"#NA_block tbody tr"`
			}
			for _, err := range sq.Scrape(&page, resp.Body) {
				fmt.Println("Scrape error:", err)
			}
			for _, svc := range page.Services {
				if strings.Contains(svc.Status, "operating normally") {
					if _, exists := problemServices[svc.Name]; exists {
						bot.SendColorMessage(fmt.Sprintf("%s: %s", svc.Name, svc.Status), "good")
						delete(problemServices, svc.Name)
					}
				} else if status, exists := problemServices[svc.Name]; !exists || svc.Status != status {
					bot.SendColorMessage(fmt.Sprintf("%s: %s", svc.Name, svc.Status), "#FF0000")
					problemServices[svc.Name] = svc.Status
				}
			}
		}
	}())

	fmt.Println("starting up")

	for {
		log.Printf("Bot error: %s\n", bot.Run())
		time.Sleep(2 * time.Second)
	}

}
