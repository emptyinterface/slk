package slk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"time"

	"golang.org/x/net/websocket"
)

type (
	RTMClient struct {
		c *Client

		options  url.Values
		Response *RTMStartResponse

		ws *websocket.Conn

		Events struct {
			AccountsChanged                chan *AccountsChangedEvent
			BotAdded                       chan *BotAddedEvent
			BotChanged                     chan *BotChangedEvent
			ChannelArchive                 chan *ChannelArchiveEvent
			ChannelCreated                 chan *ChannelCreatedEvent
			ChannelDeleted                 chan *ChannelDeletedEvent
			ChannelHistoryChanged          chan *ChannelHistoryChangedEvent
			ChannelJoined                  chan *ChannelJoinedEvent
			ChannelLeft                    chan *ChannelLeftEvent
			ChannelMarked                  chan *ChannelMarkedEvent
			ChannelRename                  chan *ChannelRenameEvent
			ChannelUnarchive               chan *ChannelUnarchiveEvent
			CommandsChanged                chan *CommandsChangedEvent
			DndUpdated                     chan *DndUpdatedEvent
			DndUpdatedUser                 chan *DndUpdatedUserEvent
			EmailDomainChanged             chan *EmailDomainChangedEvent
			EmojiChanged                   chan *EmojiChangedEvent
			FileChange                     chan *FileChangeEvent
			FileCommentAdded               chan *FileCommentAddedEvent
			FileCommentDeleted             chan *FileCommentDeletedEvent
			FileCommentEdited              chan *FileCommentEditedEvent
			FileCreated                    chan *FileCreatedEvent
			FileDeleted                    chan *FileDeletedEvent
			FilePrivate                    chan *FilePrivateEvent
			FilePublic                     chan *FilePublicEvent
			FileShared                     chan *FileSharedEvent
			FileUnshared                   chan *FileUnsharedEvent
			GroupArchive                   chan *GroupArchiveEvent
			GroupClose                     chan *GroupCloseEvent
			GroupHistoryChanged            chan *GroupHistoryChangedEvent
			GroupJoined                    chan *GroupJoinedEvent
			GroupLeft                      chan *GroupLeftEvent
			GroupMarked                    chan *GroupMarkedEvent
			GroupOpen                      chan *GroupOpenEvent
			GroupRename                    chan *GroupRenameEvent
			GroupUnarchive                 chan *GroupUnarchiveEvent
			Hello                          chan *HelloEvent
			ImClose                        chan *ImCloseEvent
			ImCreated                      chan *ImCreatedEvent
			ImHistoryChanged               chan *ImHistoryChangedEvent
			ImMarked                       chan *ImMarkedEvent
			ImOpen                         chan *ImOpenEvent
			ManualPresenceChange           chan *ManualPresenceChangeEvent
			Message                        chan *MessageEvent
			PinAdded                       chan *PinAddedEvent
			PinRemoved                     chan *PinRemovedEvent
			PrefChange                     chan *PrefChangeEvent
			PresenceChange                 chan *PresenceChangeEvent
			ReactionAdded                  chan *ReactionAddedEvent
			ReactionRemoved                chan *ReactionRemovedEvent
			ReconnectUrl                   chan *ReconnectUrlEvent
			StarAdded                      chan *StarAddedEvent
			StarRemoved                    chan *StarRemovedEvent
			SubteamCreated                 chan *SubteamCreatedEvent
			SubteamSelfAdded               chan *SubteamSelfAddedEvent
			SubteamSelfRemoved             chan *SubteamSelfRemovedEvent
			SubteamUpdated                 chan *SubteamUpdatedEvent
			TeamDomainChange               chan *TeamDomainChangeEvent
			TeamJoin                       chan *TeamJoinEvent
			TeamMigrationStarted           chan *TeamMigrationStartedEvent
			TeamPlanChange                 chan *TeamPlanChangeEvent
			TeamPrefChange                 chan *TeamPrefChangeEvent
			TeamProfileChange              chan *TeamProfileChangeEvent
			TeamProfileDelete              chan *TeamProfileDeleteEvent
			TeamProfileReorder             chan *TeamProfileReorderEvent
			TeamRename                     chan *TeamRenameEvent
			UserChange                     chan *UserChangeEvent
			UserTyping                     chan *UserTypingEvent
			MessageSubtypeBotMessage       chan *MessageSubtypeBotMessageEvent
			MessageSubtypeMeMessage        chan *MessageSubtypeMeMessageEvent
			MessageSubtypeMessageChanged   chan *MessageSubtypeMessageChangedEvent
			MessageSubtypeMessageDeleted   chan *MessageSubtypeMessageDeletedEvent
			MessageSubtypeChannelJoin      chan *MessageSubtypeChannelJoinEvent
			MessageSubtypeChannelLeave     chan *MessageSubtypeChannelLeaveEvent
			MessageSubtypeChannelTopic     chan *MessageSubtypeChannelTopicEvent
			MessageSubtypeChannelPurpose   chan *MessageSubtypeChannelPurposeEvent
			MessageSubtypeChannelName      chan *MessageSubtypeChannelNameEvent
			MessageSubtypeChannelArchive   chan *MessageSubtypeChannelArchiveEvent
			MessageSubtypeChannelUnarchive chan *MessageSubtypeChannelUnarchiveEvent
			MessageSubtypeGroupJoin        chan *MessageSubtypeGroupJoinEvent
			MessageSubtypeGroupLeave       chan *MessageSubtypeGroupLeaveEvent
			MessageSubtypeGroupTopic       chan *MessageSubtypeGroupTopicEvent
			MessageSubtypeGroupPurpose     chan *MessageSubtypeGroupPurposeEvent
			MessageSubtypeGroupName        chan *MessageSubtypeGroupNameEvent
			MessageSubtypeGroupArchive     chan *MessageSubtypeGroupArchiveEvent
			MessageSubtypeGroupUnarchive   chan *MessageSubtypeGroupUnarchiveEvent
			MessageSubtypeFileShare        chan *MessageSubtypeFileShareEvent
			MessageSubtypeFileComment      chan *MessageSubtypeFileCommentEvent
			MessageSubtypeFileMention      chan *MessageSubtypeFileMentionEvent
			MessageSubtypePinnedItem       chan *MessageSubtypePinnedItemEvent
			MessageSubtypeUnpinnedItem     chan *MessageSubtypeUnpinnedItemEvent
			Unrecognized                   chan UnrecognizedEvent
			Error                          chan error
		}
	}
)

func (c *Client) NewRTMClient(options url.Values) *RTMClient {
	rtm := &RTMClient{
		options: options,
		c:       c,
	}
	rtm.Events.AccountsChanged = make(chan *AccountsChangedEvent, 8)
	rtm.Events.BotAdded = make(chan *BotAddedEvent, 8)
	rtm.Events.BotChanged = make(chan *BotChangedEvent, 8)
	rtm.Events.ChannelArchive = make(chan *ChannelArchiveEvent, 8)
	rtm.Events.ChannelCreated = make(chan *ChannelCreatedEvent, 8)
	rtm.Events.ChannelDeleted = make(chan *ChannelDeletedEvent, 8)
	rtm.Events.ChannelHistoryChanged = make(chan *ChannelHistoryChangedEvent, 8)
	rtm.Events.ChannelJoined = make(chan *ChannelJoinedEvent, 8)
	rtm.Events.ChannelLeft = make(chan *ChannelLeftEvent, 8)
	rtm.Events.ChannelMarked = make(chan *ChannelMarkedEvent, 8)
	rtm.Events.ChannelRename = make(chan *ChannelRenameEvent, 8)
	rtm.Events.ChannelUnarchive = make(chan *ChannelUnarchiveEvent, 8)
	rtm.Events.CommandsChanged = make(chan *CommandsChangedEvent, 8)
	rtm.Events.DndUpdated = make(chan *DndUpdatedEvent, 8)
	rtm.Events.DndUpdatedUser = make(chan *DndUpdatedUserEvent, 8)
	rtm.Events.EmailDomainChanged = make(chan *EmailDomainChangedEvent, 8)
	rtm.Events.EmojiChanged = make(chan *EmojiChangedEvent, 8)
	rtm.Events.FileChange = make(chan *FileChangeEvent, 8)
	rtm.Events.FileCommentAdded = make(chan *FileCommentAddedEvent, 8)
	rtm.Events.FileCommentDeleted = make(chan *FileCommentDeletedEvent, 8)
	rtm.Events.FileCommentEdited = make(chan *FileCommentEditedEvent, 8)
	rtm.Events.FileCreated = make(chan *FileCreatedEvent, 8)
	rtm.Events.FileDeleted = make(chan *FileDeletedEvent, 8)
	rtm.Events.FilePrivate = make(chan *FilePrivateEvent, 8)
	rtm.Events.FilePublic = make(chan *FilePublicEvent, 8)
	rtm.Events.FileShared = make(chan *FileSharedEvent, 8)
	rtm.Events.FileUnshared = make(chan *FileUnsharedEvent, 8)
	rtm.Events.GroupArchive = make(chan *GroupArchiveEvent, 8)
	rtm.Events.GroupClose = make(chan *GroupCloseEvent, 8)
	rtm.Events.GroupHistoryChanged = make(chan *GroupHistoryChangedEvent, 8)
	rtm.Events.GroupJoined = make(chan *GroupJoinedEvent, 8)
	rtm.Events.GroupLeft = make(chan *GroupLeftEvent, 8)
	rtm.Events.GroupMarked = make(chan *GroupMarkedEvent, 8)
	rtm.Events.GroupOpen = make(chan *GroupOpenEvent, 8)
	rtm.Events.GroupRename = make(chan *GroupRenameEvent, 8)
	rtm.Events.GroupUnarchive = make(chan *GroupUnarchiveEvent, 8)
	rtm.Events.Hello = make(chan *HelloEvent, 8)
	rtm.Events.ImClose = make(chan *ImCloseEvent, 8)
	rtm.Events.ImCreated = make(chan *ImCreatedEvent, 8)
	rtm.Events.ImHistoryChanged = make(chan *ImHistoryChangedEvent, 8)
	rtm.Events.ImMarked = make(chan *ImMarkedEvent, 8)
	rtm.Events.ImOpen = make(chan *ImOpenEvent, 8)
	rtm.Events.ManualPresenceChange = make(chan *ManualPresenceChangeEvent, 8)
	rtm.Events.Message = make(chan *MessageEvent, 8)
	rtm.Events.PinAdded = make(chan *PinAddedEvent, 8)
	rtm.Events.PinRemoved = make(chan *PinRemovedEvent, 8)
	rtm.Events.PrefChange = make(chan *PrefChangeEvent, 8)
	rtm.Events.PresenceChange = make(chan *PresenceChangeEvent, 8)
	rtm.Events.ReactionAdded = make(chan *ReactionAddedEvent, 8)
	rtm.Events.ReactionRemoved = make(chan *ReactionRemovedEvent, 8)
	rtm.Events.ReconnectUrl = make(chan *ReconnectUrlEvent, 8)
	rtm.Events.StarAdded = make(chan *StarAddedEvent, 8)
	rtm.Events.StarRemoved = make(chan *StarRemovedEvent, 8)
	rtm.Events.SubteamCreated = make(chan *SubteamCreatedEvent, 8)
	rtm.Events.SubteamSelfAdded = make(chan *SubteamSelfAddedEvent, 8)
	rtm.Events.SubteamSelfRemoved = make(chan *SubteamSelfRemovedEvent, 8)
	rtm.Events.SubteamUpdated = make(chan *SubteamUpdatedEvent, 8)
	rtm.Events.TeamDomainChange = make(chan *TeamDomainChangeEvent, 8)
	rtm.Events.TeamJoin = make(chan *TeamJoinEvent, 8)
	rtm.Events.TeamMigrationStarted = make(chan *TeamMigrationStartedEvent, 8)
	rtm.Events.TeamPlanChange = make(chan *TeamPlanChangeEvent, 8)
	rtm.Events.TeamPrefChange = make(chan *TeamPrefChangeEvent, 8)
	rtm.Events.TeamProfileChange = make(chan *TeamProfileChangeEvent, 8)
	rtm.Events.TeamProfileDelete = make(chan *TeamProfileDeleteEvent, 8)
	rtm.Events.TeamProfileReorder = make(chan *TeamProfileReorderEvent, 8)
	rtm.Events.TeamRename = make(chan *TeamRenameEvent, 8)
	rtm.Events.UserChange = make(chan *UserChangeEvent, 8)
	rtm.Events.UserTyping = make(chan *UserTypingEvent, 8)
	rtm.Events.MessageSubtypeBotMessage = make(chan *MessageSubtypeBotMessageEvent, 8)
	rtm.Events.MessageSubtypeMeMessage = make(chan *MessageSubtypeMeMessageEvent, 8)
	rtm.Events.MessageSubtypeMessageChanged = make(chan *MessageSubtypeMessageChangedEvent, 8)
	rtm.Events.MessageSubtypeMessageDeleted = make(chan *MessageSubtypeMessageDeletedEvent, 8)
	rtm.Events.MessageSubtypeChannelJoin = make(chan *MessageSubtypeChannelJoinEvent, 8)
	rtm.Events.MessageSubtypeChannelLeave = make(chan *MessageSubtypeChannelLeaveEvent, 8)
	rtm.Events.MessageSubtypeChannelTopic = make(chan *MessageSubtypeChannelTopicEvent, 8)
	rtm.Events.MessageSubtypeChannelPurpose = make(chan *MessageSubtypeChannelPurposeEvent, 8)
	rtm.Events.MessageSubtypeChannelName = make(chan *MessageSubtypeChannelNameEvent, 8)
	rtm.Events.MessageSubtypeChannelArchive = make(chan *MessageSubtypeChannelArchiveEvent, 8)
	rtm.Events.MessageSubtypeChannelUnarchive = make(chan *MessageSubtypeChannelUnarchiveEvent, 8)
	rtm.Events.MessageSubtypeGroupJoin = make(chan *MessageSubtypeGroupJoinEvent, 8)
	rtm.Events.MessageSubtypeGroupLeave = make(chan *MessageSubtypeGroupLeaveEvent, 8)
	rtm.Events.MessageSubtypeGroupTopic = make(chan *MessageSubtypeGroupTopicEvent, 8)
	rtm.Events.MessageSubtypeGroupPurpose = make(chan *MessageSubtypeGroupPurposeEvent, 8)
	rtm.Events.MessageSubtypeGroupName = make(chan *MessageSubtypeGroupNameEvent, 8)
	rtm.Events.MessageSubtypeGroupArchive = make(chan *MessageSubtypeGroupArchiveEvent, 8)
	rtm.Events.MessageSubtypeGroupUnarchive = make(chan *MessageSubtypeGroupUnarchiveEvent, 8)
	rtm.Events.MessageSubtypeFileShare = make(chan *MessageSubtypeFileShareEvent, 8)
	rtm.Events.MessageSubtypeFileComment = make(chan *MessageSubtypeFileCommentEvent, 8)
	rtm.Events.MessageSubtypeFileMention = make(chan *MessageSubtypeFileMentionEvent, 8)
	rtm.Events.MessageSubtypePinnedItem = make(chan *MessageSubtypePinnedItemEvent, 8)
	rtm.Events.MessageSubtypeUnpinnedItem = make(chan *MessageSubtypeUnpinnedItemEvent, 8)
	rtm.Events.Unrecognized = make(chan UnrecognizedEvent, 8)
	rtm.Events.Error = make(chan error, 8)

	return rtm

}

func (rtm *RTMClient) Run() (err error) {

	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("panic! %v", e)
		}
	}()

	rtm.Response, err = rtm.c.RTMStart(rtm.options)
	if err != nil {
		return err
	}

	rtm.ws, err = websocket.Dial(rtm.Response.URL, "", "http://localhost")
	if err != nil {
		return err
	}
	defer rtm.ws.Close()

	errs := make(chan error)
	done := make(chan struct{})
	defer close(done)

	// ping loop
	const PingTick = 5 * time.Second
	var pingMessage = []byte("hello")
	go func() {
		tick := time.NewTicker(PingTick)
		defer tick.Stop()
		for {
			select {
			case <-tick.C:
				w, err := rtm.ws.NewFrameWriter(websocket.PingFrame)
				if err != nil {
					errs <- fmt.Errorf("ping: %s", err)
					return
				}
				if _, err := w.Write(pingMessage); err != nil {
					errs <- fmt.Errorf("ping: %s", err)
					return
				}
				if err = w.Close(); err != nil {
					errs <- fmt.Errorf("ping: %s", err)
					return
				}
			case <-done:
				return
			}
			if rtm.c.Debug {
				fmt.Println(">>> " + websocketFrameType(websocket.PingFrame))
			}
		}
	}()

	// frame reader loop
	datas := make(chan []byte)
	pongs := make(chan struct{})
	go func() {
		for {
			r, err := rtm.ws.NewFrameReader()
			if err != nil {
				errs <- fmt.Errorf("broken connection: %s", err)
				return
			}
			switch r.PayloadType() {
			case websocket.PingFrame:
				w, err := rtm.ws.NewFrameWriter(websocket.PongFrame)
				if err != nil {
					errs <- fmt.Errorf("pong back: %s", err)
					return
				}
				pingData, err := ioutil.ReadAll(r)
				if err != nil {
					errs <- fmt.Errorf("pong back: %s", err)
					return
				}
				if len(pingData) == 0 {
					pingData = pingMessage
				}
				if _, err := w.Write(pingMessage); err != nil {
					errs <- fmt.Errorf("pong back: %s", err)
					return
				}
				if err = w.Close(); err != nil {
					errs <- fmt.Errorf("pong back: %s", err)
					return
				}
				if rtm.c.Debug {
					fmt.Println(">>> " + websocketFrameType(websocket.PongFrame))
				}
			case websocket.PongFrame:
				io.Copy(ioutil.Discard, r)
				pongs <- struct{}{}
			case websocket.TextFrame:
				data, err := ioutil.ReadAll(r)
				if err != nil {
					errs <- fmt.Errorf("text frame: %s", err)
					return
				}
				datas <- data
			case websocket.CloseFrame:
				errs <- errors.New("websocket closed by server")
				return
			default:
				data, _ := ioutil.ReadAll(r)
				fmt.Printf("Unrecognized frame type: %d, %q", r.PayloadType(), data)
			}
			if rtm.c.Debug {
				fmt.Println("<<< "+websocketFrameType(r.PayloadType()), r.Len())
			}
		}
	}()

	// consumer/timeout loop
	const Timeout = 2 * PingTick
	for {
		select {
		case <-pongs:
		case <-time.After(Timeout):
			return fmt.Errorf("websocket timeout: %s", Timeout)
		case err := <-errs:
			return fmt.Errorf("websocket error: %s", err)
		case data := <-datas:
			if err = rtm.parseEvent(data); err != nil {
				select {
				case rtm.Events.Error <- fmt.Errorf("Error parsing event: %s\n%s", err, data):
				default:
				}
			}
		}
	}

}

func websocketFrameType(n byte) string {
	switch n {
	case websocket.ContinuationFrame:
		return "continuation frame"
	case websocket.TextFrame:
		return "text frame"
	case websocket.BinaryFrame:
		return "binary frame"
	case websocket.CloseFrame:
		return "close frame"
	case websocket.PingFrame:
		return "ping frame"
	case websocket.PongFrame:
		return "pong frame"
	case websocket.UnknownFrame:
		return "unknown frame"
	default:
		return "bad frame type"
	}
}

func (rtm *RTMClient) Stop() error {
	return rtm.ws.Close()
}

func (rtm *RTMClient) parseEvent(data []byte) error {

	var typesniff struct {
		Type    string `json:"type"`
		SubType string `json:"sub_type"`
	}

	if err := json.Unmarshal(data, &typesniff); err != nil {
		return err
	}

	// consider something saner
	switch typesniff.Type {
	case AccountsChangedEventCode:
		v := &AccountsChangedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.AccountsChanged <- v:
		default:
		}
	case BotAddedEventCode:
		v := &BotAddedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.BotAdded <- v:
		default:
		}
	case BotChangedEventCode:
		v := &BotChangedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.BotChanged <- v:
		default:
		}
	case ChannelArchiveEventCode:
		v := &ChannelArchiveEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.ChannelArchive <- v:
		default:
		}
	case ChannelCreatedEventCode:
		v := &ChannelCreatedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.ChannelCreated <- v:
		default:
		}
	case ChannelDeletedEventCode:
		v := &ChannelDeletedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.ChannelDeleted <- v:
		default:
		}
	case ChannelHistoryChangedEventCode:
		v := &ChannelHistoryChangedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.ChannelHistoryChanged <- v:
		default:
		}
	case ChannelJoinedEventCode:
		v := &ChannelJoinedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.ChannelJoined <- v:
		default:
		}
	case ChannelLeftEventCode:
		v := &ChannelLeftEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.ChannelLeft <- v:
		default:
		}
	case ChannelMarkedEventCode:
		v := &ChannelMarkedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.ChannelMarked <- v:
		default:
		}
	case ChannelRenameEventCode:
		v := &ChannelRenameEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.ChannelRename <- v:
		default:
		}
	case ChannelUnarchiveEventCode:
		v := &ChannelUnarchiveEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.ChannelUnarchive <- v:
		default:
		}
	case CommandsChangedEventCode:
		v := &CommandsChangedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.CommandsChanged <- v:
		default:
		}
	case DndUpdatedEventCode:
		v := &DndUpdatedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.DndUpdated <- v:
		default:
		}
	case DndUpdatedUserEventCode:
		v := &DndUpdatedUserEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.DndUpdatedUser <- v:
		default:
		}
	case EmailDomainChangedEventCode:
		v := &EmailDomainChangedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.EmailDomainChanged <- v:
		default:
		}
	case EmojiChangedEventCode:
		v := &EmojiChangedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.EmojiChanged <- v:
		default:
		}
	case FileChangeEventCode:
		v := &FileChangeEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.FileChange <- v:
		default:
		}
	case FileCommentAddedEventCode:
		v := &FileCommentAddedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.FileCommentAdded <- v:
		default:
		}
	case FileCommentDeletedEventCode:
		v := &FileCommentDeletedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.FileCommentDeleted <- v:
		default:
		}
	case FileCommentEditedEventCode:
		v := &FileCommentEditedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.FileCommentEdited <- v:
		default:
		}
	case FileCreatedEventCode:
		v := &FileCreatedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.FileCreated <- v:
		default:
		}
	case FileDeletedEventCode:
		v := &FileDeletedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.FileDeleted <- v:
		default:
		}
	case FilePrivateEventCode:
		v := &FilePrivateEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.FilePrivate <- v:
		default:
		}
	case FilePublicEventCode:
		v := &FilePublicEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.FilePublic <- v:
		default:
		}
	case FileSharedEventCode:
		v := &FileSharedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.FileShared <- v:
		default:
		}
	case FileUnsharedEventCode:
		v := &FileUnsharedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.FileUnshared <- v:
		default:
		}
	case GroupArchiveEventCode:
		v := &GroupArchiveEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.GroupArchive <- v:
		default:
		}
	case GroupCloseEventCode:
		v := &GroupCloseEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.GroupClose <- v:
		default:
		}
	case GroupHistoryChangedEventCode:
		v := &GroupHistoryChangedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.GroupHistoryChanged <- v:
		default:
		}
	case GroupJoinedEventCode:
		v := &GroupJoinedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.GroupJoined <- v:
		default:
		}
	case GroupLeftEventCode:
		v := &GroupLeftEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.GroupLeft <- v:
		default:
		}
	case GroupMarkedEventCode:
		v := &GroupMarkedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.GroupMarked <- v:
		default:
		}
	case GroupOpenEventCode:
		v := &GroupOpenEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.GroupOpen <- v:
		default:
		}
	case GroupRenameEventCode:
		v := &GroupRenameEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.GroupRename <- v:
		default:
		}
	case GroupUnarchiveEventCode:
		v := &GroupUnarchiveEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.GroupUnarchive <- v:
		default:
		}
	case HelloEventCode:
		v := &HelloEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.Hello <- v:
		default:
		}
	case ImCloseEventCode:
		v := &ImCloseEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.ImClose <- v:
		default:
		}
	case ImCreatedEventCode:
		v := &ImCreatedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.ImCreated <- v:
		default:
		}
	case ImHistoryChangedEventCode:
		v := &ImHistoryChangedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.ImHistoryChanged <- v:
		default:
		}
	case ImMarkedEventCode:
		v := &ImMarkedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.ImMarked <- v:
		default:
		}
	case ImOpenEventCode:
		v := &ImOpenEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.ImOpen <- v:
		default:
		}
	case ManualPresenceChangeEventCode:
		v := &ManualPresenceChangeEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.ManualPresenceChange <- v:
		default:
		}
	case MessageEventCode:
		switch typesniff.SubType {
		case MessageSubtypeBotMessageEventCode:
			v := &MessageSubtypeBotMessageEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypeBotMessage <- v:
			default:
			}
		case MessageSubtypeMeMessageEventCode:
			v := &MessageSubtypeMeMessageEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypeMeMessage <- v:
			default:
			}
		case MessageSubtypeMessageChangedEventCode:
			v := &MessageSubtypeMessageChangedEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypeMessageChanged <- v:
			default:
			}
		case MessageSubtypeMessageDeletedEventCode:
			v := &MessageSubtypeMessageDeletedEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypeMessageDeleted <- v:
			default:
			}
		case MessageSubtypeChannelJoinEventCode:
			v := &MessageSubtypeChannelJoinEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypeChannelJoin <- v:
			default:
			}
		case MessageSubtypeChannelLeaveEventCode:
			v := &MessageSubtypeChannelLeaveEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypeChannelLeave <- v:
			default:
			}
		case MessageSubtypeChannelTopicEventCode:
			v := &MessageSubtypeChannelTopicEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypeChannelTopic <- v:
			default:
			}
		case MessageSubtypeChannelPurposeEventCode:
			v := &MessageSubtypeChannelPurposeEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypeChannelPurpose <- v:
			default:
			}
		case MessageSubtypeChannelNameEventCode:
			v := &MessageSubtypeChannelNameEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypeChannelName <- v:
			default:
			}
		case MessageSubtypeChannelArchiveEventCode:
			v := &MessageSubtypeChannelArchiveEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypeChannelArchive <- v:
			default:
			}
		case MessageSubtypeChannelUnarchiveEventCode:
			v := &MessageSubtypeChannelUnarchiveEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypeChannelUnarchive <- v:
			default:
			}
		case MessageSubtypeGroupJoinEventCode:
			v := &MessageSubtypeGroupJoinEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypeGroupJoin <- v:
			default:
			}
		case MessageSubtypeGroupLeaveEventCode:
			v := &MessageSubtypeGroupLeaveEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypeGroupLeave <- v:
			default:
			}
		case MessageSubtypeGroupTopicEventCode:
			v := &MessageSubtypeGroupTopicEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypeGroupTopic <- v:
			default:
			}
		case MessageSubtypeGroupPurposeEventCode:
			v := &MessageSubtypeGroupPurposeEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypeGroupPurpose <- v:
			default:
			}
		case MessageSubtypeGroupNameEventCode:
			v := &MessageSubtypeGroupNameEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypeGroupName <- v:
			default:
			}
		case MessageSubtypeGroupArchiveEventCode:
			v := &MessageSubtypeGroupArchiveEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypeGroupArchive <- v:
			default:
			}
		case MessageSubtypeGroupUnarchiveEventCode:
			v := &MessageSubtypeGroupUnarchiveEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypeGroupUnarchive <- v:
			default:
			}
		case MessageSubtypeFileShareEventCode:
			v := &MessageSubtypeFileShareEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypeFileShare <- v:
			default:
			}
		case MessageSubtypeFileCommentEventCode:
			v := &MessageSubtypeFileCommentEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypeFileComment <- v:
			default:
			}
		case MessageSubtypeFileMentionEventCode:
			v := &MessageSubtypeFileMentionEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypeFileMention <- v:
			default:
			}
		case MessageSubtypePinnedItemEventCode:
			v := &MessageSubtypePinnedItemEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypePinnedItem <- v:
			default:
			}
		case MessageSubtypeUnpinnedItemEventCode:
			v := &MessageSubtypeUnpinnedItemEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.MessageSubtypeUnpinnedItem <- v:
			default:
			}
		default:
			v := &MessageEvent{}
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			select {
			case rtm.Events.Message <- v:
			default:
			}
		}
	case PinAddedEventCode:
		v := &PinAddedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.PinAdded <- v:
		default:
		}
	case PinRemovedEventCode:
		v := &PinRemovedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.PinRemoved <- v:
		default:
		}
	case PrefChangeEventCode:
		v := &PrefChangeEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.PrefChange <- v:
		default:
		}
	case PresenceChangeEventCode:
		v := &PresenceChangeEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.PresenceChange <- v:
		default:
		}
	case ReactionAddedEventCode:
		v := &ReactionAddedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.ReactionAdded <- v:
		default:
		}
	case ReactionRemovedEventCode:
		v := &ReactionRemovedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.ReactionRemoved <- v:
		default:
		}
	case ReconnectUrlEventCode:
		v := &ReconnectUrlEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.ReconnectUrl <- v:
		default:
		}
	case StarAddedEventCode:
		v := &StarAddedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.StarAdded <- v:
		default:
		}
	case StarRemovedEventCode:
		v := &StarRemovedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.StarRemoved <- v:
		default:
		}
	case SubteamCreatedEventCode:
		v := &SubteamCreatedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.SubteamCreated <- v:
		default:
		}
	case SubteamSelfAddedEventCode:
		v := &SubteamSelfAddedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.SubteamSelfAdded <- v:
		default:
		}
	case SubteamSelfRemovedEventCode:
		v := &SubteamSelfRemovedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.SubteamSelfRemoved <- v:
		default:
		}
	case SubteamUpdatedEventCode:
		v := &SubteamUpdatedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.SubteamUpdated <- v:
		default:
		}
	case TeamDomainChangeEventCode:
		v := &TeamDomainChangeEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.TeamDomainChange <- v:
		default:
		}
	case TeamJoinEventCode:
		v := &TeamJoinEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.TeamJoin <- v:
		default:
		}
	case TeamMigrationStartedEventCode:
		v := &TeamMigrationStartedEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.TeamMigrationStarted <- v:
		default:
		}
	case TeamPlanChangeEventCode:
		v := &TeamPlanChangeEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.TeamPlanChange <- v:
		default:
		}
	case TeamPrefChangeEventCode:
		v := &TeamPrefChangeEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.TeamPrefChange <- v:
		default:
		}
	case TeamProfileChangeEventCode:
		v := &TeamProfileChangeEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.TeamProfileChange <- v:
		default:
		}
	case TeamProfileDeleteEventCode:
		v := &TeamProfileDeleteEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.TeamProfileDelete <- v:
		default:
		}
	case TeamProfileReorderEventCode:
		v := &TeamProfileReorderEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.TeamProfileReorder <- v:
		default:
		}
	case TeamRenameEventCode:
		v := &TeamRenameEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.TeamRename <- v:
		default:
		}
	case UserChangeEventCode:
		v := &UserChangeEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.UserChange <- v:
		default:
		}
	case UserTypingEventCode:
		v := &UserTypingEvent{}
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		select {
		case rtm.Events.UserTyping <- v:
		default:
		}
	default:
		v := UnrecognizedEvent{}
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		select {
		case rtm.Events.Unrecognized <- v:
		default:
		}
	}

	return nil

}
