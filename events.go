package slk

type (

	// https://api.slack.com/events/hello
	// {
	//     "type": "hello"
	// }
	HelloEvent struct {
		Type string `json:"type"`
	}

	// https://api.slack.com/events/user_typing
	// {
	//     "type": "user_typing",
	//     "channel": "C02ELGNBH",
	//     "user": "U024BE7LH"
	// }
	UserTypingEvent struct {
		Type    string `json:"type"`
		Channel string `json:"channel"`
		User    string `json:"user"`
	}

	// https://api.slack.com/events/channel_marked
	// {
	//     "type": "channel_marked",
	//     "channel": "C024BE91L",
	//     "ts": "1401383885.000061"
	// }
	ChannelMarkedEvent struct {
		Type      string `json:"type"`
		Channel   string `json:"channel"`
		Timestamp Time   `json:"ts"`
	}

	// https://api.slack.com/events/channel_created
	// {
	//     "type": "channel_created",
	//     "channel": {
	//         "id": "C024BE91L",
	//         "name": "fun",
	//         "created": 1360782804,
	//         "creator": "U024BE7LH"
	//     }
	// }
	ChannelCreatedEvent struct {
		Type    string `json:"type"`
		Channel struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Created Time   `json:"created"`
			Creator string `json:"creator"`
		} `json:"channel"`
	}

	// https://api.slack.com/events/channel_joined
	// {
	//     "type": "channel_joined",
	//     "channel": {
	//         ...
	//     }
	// }
	ChannelJoinedEvent struct {
		Type    string   `json:"type"`
		Channel *Channel `json:"channel"`
	}

	// https://api.slack.com/events/channel_left
	// {
	//     "type": "channel_left",
	//     "channel": "C024BE91L"
	// }
	ChannelLeftEvent struct {
		Type    string `json:"type"`
		Channel string `json:"channel"`
	}

	// https://api.slack.com/events/channel_deleted
	// {
	//     "type": "channel_deleted",
	//     "channel": "C024BE91L"
	// }
	ChannelDeletedEvent struct {
		Type    string `json:"type"`
		Channel string `json:"channel"`
	}

	// https://api.slack.com/events/channel_rename
	// {
	//     "type": "channel_rename",
	//     "channel": {
	//         "id":"C02ELGNBH",
	//         "name":"new_name",
	//         "created":1360782804
	//     }
	// }
	ChannelRenameEvent struct {
		Type    string `json:"type"`
		Channel struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Created Time   `json:"created"`
		} `json:"channel"`
	}

	// https://api.slack.com/events/channel_archive
	// {
	//     "type": "channel_archive",
	//     "channel": "C024BE91L",
	//     "user": "U024BE7LH"
	// }
	ChannelArchiveEvent struct {
		Type    string `json:"type"`
		Channel string `json:"channel"`
		User    string `json:"user"`
	}

	// https://api.slack.com/events/channel_unarchive
	// {
	//     "type": "channel_unarchive",
	//     "channel": "C024BE91L",
	//     "user": "U024BE7LH"
	// }
	ChannelUnarchiveEvent struct {
		Type    string `json:"type"`
		Channel string `json:"channel"`
		User    string `json:"user"`
	}

	// https://api.slack.com/events/channel_history_changed
	// {
	//     "type": "channel_history_changed",
	//     "latest": "1358877455.000010",
	//     "ts": "1361482916.000003",
	//     "event_ts": "1361482916.000004"
	// }
	ChannelHistoryChangedEvent struct {
		Type           string `json:"type"`
		Latest         Time   `json:"latest"`
		Timestamp      Time   `json:"ts"`
		EventTimestamp Time   `json:"event_ts"`
	}

	// https://api.slack.com/events/dnd_updated
	// {
	//     "type": "dnd_updated",
	//     "user": "U1234",
	//     "dnd_status": {
	//         "dnd_enabled": true,
	//         "next_dnd_start_ts": 1450387800,
	//         "next_dnd_end_ts": 1450423800,
	//         "snooze_enabled": true,
	//         "snooze_endtime": 1450373897
	//     }
	// }
	DndUpdatedEvent struct {
		Type      string `json:"type"`
		User      string `json:"user"`
		DndStatus struct {
			DndEnabled            bool `json:"dnd_enabled"`
			NextDndStartTimestamp Time `json:"next_dnd_start_ts"`
			NextDndEndTimestamp   Time `json:"next_dnd_end_ts"`
			SnoozeEnabled         bool `json:"snooze_enabled"`
			SnoozeEndtime         Time `json:"snooze_endtime"`
		} `json:"dnd_status"`
	}

	// https://api.slack.com/events/dnd_updated_user
	// {
	//     "type": "dnd_updated_user",
	//     "user": "U1234",
	//     "dnd_status": {
	//         "dnd_enabled": true,
	//         "next_dnd_start_ts": 1450387800,
	//         "next_dnd_end_ts": 1450423800
	//     }
	// }
	DndUpdatedUserEvent struct {
		Type      string `json:"type"`
		User      string `json:"user"`
		DndStatus struct {
			DndEnabled            bool `json:"dnd_enabled"`
			NextDndStartTimestamp Time `json:"next_dnd_start_ts"`
			NextDndEndTimestamp   Time `json:"next_dnd_end_ts"`
		} `json:"dnd_status"`
	}

	// https://api.slack.com/events/im_open
	// {
	//     "type": "im_open",
	//     "user": "U024BE7LH",
	//     "channel": "D024BE91L"
	// }
	ImOpenEvent struct {
		Type    string `json:"type"`
		User    string `json:"user"`
		Channel string `json:"channel"`
	}

	// https://api.slack.com/events/im_close
	// {
	//     "type": "im_close",
	//     "user": "U024BE7LH",
	//     "channel": "D024BE91L"
	// }
	ImCloseEvent struct {
		Type    string `json:"type"`
		User    string `json:"user"`
		Channel string `json:"channel"`
	}

	// https://api.slack.com/events/im_marked
	// {
	//     "type": "im_marked",
	//     "channel": "D024BE91L",
	//     "ts": "1401383885.000061"
	// }
	ImMarkedEvent struct {
		Type      string `json:"type"`
		Channel   string `json:"channel"`
		Timestamp Time   `json:"ts"`
	}

	// https://api.slack.com/events/im_history_changed
	// {
	//     "type": "im_history_changed",
	//     "latest": "1358877455.000010",
	//     "ts": "1361482916.000003",
	//     "event_ts": "1361482916.000004"
	// }
	ImHistoryChangedEvent struct {
		Type           string `json:"type"`
		Latest         Time   `json:"latest"`
		Timestamp      Time   `json:"ts"`
		EventTimestamp Time   `json:"event_ts"`
	}

	// https://api.slack.com/events/group_open
	// {
	//     "type": "group_open",
	//     "user": "U024BE7LH",
	//     "channel": "G024BE91L"
	// }
	GroupOpenEvent struct {
		Type    string `json:"type"`
		User    string `json:"user"`
		Channel string `json:"channel"`
	}

	// https://api.slack.com/events/group_close
	// {
	//     "type": "group_close",
	//     "user": "U024BE7LH",
	//     "channel": "G024BE91L"
	// }
	GroupCloseEvent struct {
		Type    string `json:"type"`
		User    string `json:"user"`
		Channel string `json:"channel"`
	}

	// https://api.slack.com/events/group_archive
	// {
	//     "type": "group_archive",
	//     "channel": "G024BE91L"
	// }
	GroupArchiveEvent struct {
		Type    string `json:"type"`
		Channel string `json:"channel"`
	}

	// https://api.slack.com/events/group_unarchive
	// {
	//     "type": "group_unarchive",
	//     "channel": "G024BE91L"
	// }
	GroupUnarchiveEvent struct {
		Type    string `json:"type"`
		Channel string `json:"channel"`
	}

	// https://api.slack.com/events/group_rename
	// {
	//     "type": "group_rename",
	//     "channel": {
	//         "id":"G02ELGNBH",
	//         "name":"new_name",
	//         "created":1360782804
	//     }
	// }
	GroupRenameEvent struct {
		Type    string `json:"type"`
		Channel struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Created Time   `json:"created"`
		} `json:"channel"`
	}

	// https://api.slack.com/events/group_marked
	// {
	//     "type": "group_marked",
	//     "channel": "G024BE91L",
	//     "ts": "1401383885.000061"
	// }
	GroupMarkedEvent struct {
		Type      string `json:"type"`
		Channel   string `json:"channel"`
		Timestamp Time   `json:"ts"`
	}

	// https://api.slack.com/events/group_history_changed
	// {
	//     "type": "group_history_changed",
	//     "latest": "1358877455.000010",
	//     "ts": "1361482916.000003",
	//     "event_ts": "1361482916.000004"
	// }
	GroupHistoryChangedEvent struct {
		Type           string `json:"type"`
		Latest         Time   `json:"latest"`
		Timestamp      Time   `json:"ts"`
		EventTimestamp Time   `json:"event_ts"`
	}

	// https://api.slack.com/events/file_private
	// {
	//     "type": "file_private",
	//     "file": "F2147483862"
	// }
	FilePrivateEvent struct {
		Type string `json:"type"`
		File string `json:"file"`
	}

	// https://api.slack.com/events/file_deleted
	// {
	//     "type": "file_deleted",
	//     "file_id": "F2147483862",
	//     "event_ts": "1361482916.000004"
	// }
	FileDeletedEvent struct {
		Type           string `json:"type"`
		FileID         string `json:"file_id"`
		EventTimestamp Time   `json:"event_ts"`
	}

	// https://api.slack.com/events/presence_change
	// {
	//     "type": "presence_change",
	//     "user": "U024BE7LH",
	//     "presence": "away"
	// }
	PresenceChangeEvent struct {
		Type     string   `json:"type"`
		User     string   `json:"user"`
		Presence Presence `json:"presence"`
	}

	// https://api.slack.com/events/manual_presence_change
	// {
	//     "type": "manual_presence_change",
	//     "presence": "away"
	// }
	ManualPresenceChangeEvent struct {
		Type     string   `json:"type"`
		Presence Presence `json:"presence"`
	}

	// https://api.slack.com/events/pref_change
	// {
	//     "type": "pref_change",
	//     "name": "messages_theme",
	//     "value": "dense"
	// }
	PrefChangeEvent struct {
		Type  string      `json:"type"`
		Name  string      `json:"name"`
		Value interface{} `json:"value"`
	}

	// https://api.slack.com/events/emoji_changed
	// {
	//     "type": "emoji_changed",
	//     "event_ts" : "1361482916.000004"
	// }
	EmojiChangedEvent struct {
		Type           string `json:"type"`
		EventTimestamp Time   `json:"event_ts"`
	}

	// https://api.slack.com/events/commands_changed
	// {
	//     "type": "commands_changed",
	//     "event_ts" : "1361482916.000004"
	// }
	CommandsChangedEvent struct {
		Type           string `json:"type"`
		EventTimestamp Time   `json:"event_ts"`
	}

	// https://api.slack.com/events/team_plan_change
	// {
	//     "type": "team_plan_change",
	//     "plan": "std"
	// }
	TeamPlanChangeEvent struct {
		Type string `json:"type"`
		Plan string `json:"plan"`
	}

	// https://api.slack.com/events/team_pref_change
	// {
	//     "type": "team_pref_change",
	//     "name": "slackbot_responses_only_admins",
	//     "value": true
	// }
	TeamPrefChangeEvent struct {
		Type  string `json:"type"`
		Name  string `json:"name"`
		Value bool   `json:"value"`
	}

	// https://api.slack.com/events/team_rename
	// {
	//     "type": "team_rename",
	//     "name": "New Team Name Inc."
	// }
	TeamRenameEvent struct {
		Type string `json:"type"`
		Name string `json:"name"`
	}

	// https://api.slack.com/events/team_domain_change
	// {
	//     "type": "team_domain_change",
	//     "url": "https://my.slack.com",
	//     "domain": "my"
	// }
	TeamDomainChangeEvent struct {
		Type   string `json:"type"`
		URL    string `json:"url"`
		Domain string `json:"domain"`
	}

	// https://api.slack.com/events/email_domain_changed
	// {
	//     "type": "email_domain_changed",
	//     "email_domain":"example.com",
	//     "event_ts": "1360782804.083113"
	// }
	EmailDomainChangedEvent struct {
		Type           string `json:"type"`
		EmailDomain    string `json:"email_domain"`
		EventTimestamp Time   `json:"event_ts"`
	}

	// https://api.slack.com/events/bot_added
	// {
	//     "type": "bot_added",
	//     "bot": {
	//         "id": "B024BE7LH",
	//         "name": "hugbot",
	//         "icons": {
	//             "image_48": "https:\/\/slack.com\/path\/to\/hugbot_48.png"
	//         }
	//     }
	// }
	BotAddedEvent struct {
		Type string `json:"type"`
		Bot  struct {
			ID    string            `json:"id"`
			Name  string            `json:"name"`
			Icons map[string]string `json:"icons"`
		} `json:"bot"`
	}

	// https://api.slack.com/events/bot_changed
	// {
	//     "type": "bot_changed",
	//     "bot": {
	//         "id": "B024BE7LH",
	//         "name": "hugbot",
	//         "icons": {
	//             "image_48": "https:\/\/slack.com\/path\/to\/hugbot_48.png"
	//         }
	//     }
	// }
	BotChangedEvent struct {
		Type string `json:"type"`
		Bot  struct {
			ID    string            `json:"id"`
			Name  string            `json:"name"`
			Icons map[string]string `json:"icons"`
		} `json:"bot"`
	}

	// https://api.slack.com/events/accounts_changed
	// {
	//     "type": "accounts_changed"
	// }
	AccountsChangedEvent struct {
		Type string `json:"type"`
	}

	// https://api.slack.com/events/team_migration_started
	// {
	//     "type": "team_migration_started"
	// }
	TeamMigrationStartedEvent struct {
		Type string `json:"type"`
	}

	// https://api.slack.com/events/reconnect_url
	// {
	//     "type": "reconnect_url"
	// }
	ReconnectUrlEvent struct {
		Type string `json:"type"`
	}

	// https://api.slack.com/events/subteam_created
	// {
	//     "type": "subteam_created",
	//     "subteam": {
	//         "id": "S0615G0KT",
	//         "team_id": "T060RNRCH",
	//         "is_usergroup": true,
	//         "name": "Marketing Team",
	//         "description": "Marketing gurus, PR experts and product advocates.",
	//         "handle": "marketing-team",
	//         "is_external": false,
	//         "date_create": 1446746793,
	//         "date_update": 1446746793,
	//         "date_delete": 0,
	//         "auto_type": null,
	//         "created_by": "U060RNRCZ",
	//         "updated_by": "U060RNRCZ",
	//         "deleted_by": null,
	//         "prefs": {
	//             "channels": [

	//             ],
	//             "groups": [

	//             ]
	//         },
	//         "user_count": "0"
	//     }
	// }
	SubteamCreatedEvent struct {
		Type    string `json:"type"`
		Subteam struct {
			ID          string                 `json:"id"`
			TeamID      string                 `json:"team_id"`
			IsUsergroup bool                   `json:"is_usergroup"`
			Name        string                 `json:"name"`
			Description string                 `json:"description"`
			Handle      string                 `json:"handle"`
			IsExternal  bool                   `json:"is_external"`
			DateCreated Time                   `json:"date_create"`
			DateUpdated Time                   `json:"date_update"`
			DateDeleted Time                   `json:"date_delete"`
			AutoType    string                 `json:"auto_type"`
			CreatedBy   string                 `json:"created_by"`
			UpdatedBy   string                 `json:"updated_by"`
			DeletedBy   string                 `json:"deleted_by"`
			Prefs       map[string]interface{} `json:"prefs"`
			UserCount   int                    `json:"user_count"`
		} `json:"subteam"`
	}

	// https://api.slack.com/events/subteam_updated
	// {
	//     "type": "subteam_updated",
	//     "subteam": {
	//        "id": "S0614TZR7",
	//        "team_id": "T060RNRCH",
	//        "is_usergroup": true,
	//        "name": "Team Admins",
	//        "description": "A group of all Administrators on your team.",
	//        "handle": "admins",
	//        "is_external": false,
	//        "date_create": 1446598059,
	//        "date_update": 1446670362,
	//        "date_delete": 0,
	//        "auto_type": "admin",
	//        "created_by": "USLACKBOT",
	//        "updated_by": "U060RNRCZ",
	//        "deleted_by": null,
	//        "prefs": {
	//            "channels": [

	//            ],
	//            "groups": [

	//            ]
	//        },
	//        "users": [
	//            "U060RNRCZ",
	//            "U060ULRC0",
	//            "U06129G2V",
	//            "U061309JM"
	//        ],
	//        "user_count": "4"
	//     }
	// }
	SubteamUpdatedEvent struct {
		Type    string `json:"type"`
		Subteam struct {
			ID          string                 `json:"id"`
			TeamID      string                 `json:"team_id"`
			IsUsergroup bool                   `json:"is_usergroup"`
			Name        string                 `json:"name"`
			Description string                 `json:"description"`
			Handle      string                 `json:"handle"`
			IsExternal  bool                   `json:"is_external"`
			DateCreated Time                   `json:"date_create"`
			DateUpdated Time                   `json:"date_update"`
			DateDeleted Time                   `json:"date_delete"`
			AutoType    string                 `json:"auto_type"`
			CreatedBy   string                 `json:"created_by"`
			UpdatedBy   string                 `json:"updated_by"`
			DeletedBy   string                 `json:"deleted_by"`
			Prefs       map[string]interface{} `json:"prefs"`
			UserCount   int                    `json:"user_count"`
			Users       []string               `json:"users"`
		} `json:"subteam"`
	}

	// https://api.slack.com/events/subteam_self_added
	// {
	//     "type": "subteam_self_added",
	//     "subteam_id": "S0615G0KT"
	// }
	SubteamSelfAddedEvent struct {
		Type      string `json:"type"`
		SubteamID string `json:"subteam_id"`
	}

	// https://api.slack.com/events/subteam_self_removed
	// {
	//     "type": "subteam_self_removed",
	//     "subteam_id": "S0615G0KT"
	// }
	SubteamSelfRemovedEvent struct {
		Type      string `json:"type"`
		SubteamID string `json:"subteam_id"`
	}

	// https://api.slack.com/events/im_created
	// {
	//     "type": "im_created",
	//     "user": "U024BE7LH",
	//     "channel": {}
	// }
	ImCreatedEvent struct {
		Type    string   `json:"type"`
		Channel *Channel `json:"channel"`
		User    string   `json:"user"`
	}

	// https://api.slack.com/events/group_joined
	// {
	//     "type": "group_joined",
	//     "channel": {}
	// }
	GroupJoinedEvent struct {
		Type    string   `json:"type"`
		Channel *Channel `json:"channel"`
	}

	// https://api.slack.com/events/group_left
	// {
	//     "type": "group_left",
	//     "channel": {}
	// }
	GroupLeftEvent struct {
		Type    string   `json:"type"`
		Channel *Channel `json:"channel"`
	}

	// https://api.slack.com/events/file_created
	// {
	//     "type": "file_created",
	//     "file": {}
	// }
	FileCreatedEvent struct {
		Type string `json:"type"`
		File *File  `json:"file"`
	}

	// https://api.slack.com/events/file_shared
	// {
	//     "type": "file_shared",
	//     "file": {}
	// }
	FileSharedEvent struct {
		Type string `json:"type"`
		File *File  `json:"file"`
	}

	// https://api.slack.com/events/file_unshared
	// {
	//     "type": "file_unshared",
	//     "file": {}
	// }
	FileUnsharedEvent struct {
		Type string `json:"type"`
		File *File  `json:"file"`
	}

	// https://api.slack.com/events/file_public
	// {
	//     "type": "file_public",
	//     "file": {}
	// }
	FilePublicEvent struct {
		Type string `json:"type"`
		File *File  `json:"file"`
	}

	// https://api.slack.com/events/file_change
	// {
	//     "type": "file_change",
	//     "file": {}
	// }
	FileChangeEvent struct {
		Type string `json:"type"`
		File *File  `json:"file"`
	}

	// https://api.slack.com/events/file_comment_added
	// {
	//     "type": "file_comment_added",
	//     "file": {},
	//     "comment": {}
	// }
	FileCommentAddedEvent struct {
		Type    string   `json:"type"`
		File    *File    `json:"file"`
		Comment *Comment `json:"comment"`
	}

	// https://api.slack.com/events/file_comment_edited
	// {
	//     "type": "file_comment_edited",
	//     "file": {},
	//     "comment": {}
	// }
	FileCommentEditedEvent struct {
		Type    string   `json:"type"`
		File    *File    `json:"file"`
		Comment *Comment `json:"comment"`
	}

	// https://api.slack.com/events/file_comment_deleted
	// {
	//     "type": "file_comment_deleted",
	//     "file": {},
	//     "comment": "Fc67890"
	// }
	FileCommentDeletedEvent struct {
		Type    string `json:"type"`
		File    *File  `json:"file"`
		Comment string `json:"comment"`
	}

	// https://api.slack.com/events/pin_added
	// {
	//     "type": "pin_added",
	//     "user": "U024BE7LH",
	//     "channel_id": "C02ELGNBH",
	//     "item": {},
	//     "event_ts": "1360782804.083113"
	// }
	PinAddedEvent struct {
		Type           string `json:"type"`
		User           string `json:"user"`
		ChannelID      string `json:"channel_id"`
		EventTimestamp Time   `json:"event_ts"`
		Item           *Item  `json:"item"`
	}

	// https://api.slack.com/events/pin_removed
	// {
	//     "type": "pin_removed",
	//     "user": "U024BE7LH",
	//     "channel_id": "C02ELGNBH",
	//     "item": {},
	//     "has_pins": false,
	//     "event_ts": "1360782804.083113"
	// }
	PinRemovedEvent struct {
		Type           string `json:"type"`
		User           string `json:"user"`
		ChannelID      string `json:"channel_id"`
		Item           *Item  `json:"item"`
		HasPins        bool   `json:"has_pins"`
		EventTimestamp Time   `json:"event_ts"`
	}

	// https://api.slack.com/events/user_change
	// {
	//     "type": "user_change",
	//     "user": {}
	// }
	UserChangeEvent struct {
		Type string `json:"type"`
		User *User  `json:"user"`
	}

	// https://api.slack.com/events/team_join
	// {
	//     "type": "team_join",
	//     "user": {}
	// }
	TeamJoinEvent struct {
		Type string `json:"type"`
		User *User  `json:"user"`
	}

	// https://api.slack.com/events/reaction_added
	// {
	//     "type": "reaction_added",
	//     "user": "U024BE7LH",
	//     "name": "thumbsup",
	//     "item": {"a":"b"},
	//     "event_ts": "1360782804.083113"
	// }
	ReactionAddedEvent struct {
		Type           string `json:"type"`
		User           string `json:"user"`
		Name           string `json:"name"`
		Item           *Item  `json:"item"`
		EventTimestamp Time   `json:"event_ts"`
	}

	// https://api.slack.com/events/reaction_removed
	// {
	//     "type": "reaction_removed",
	//     "user": "U024BE7LH",
	//     "name": "thumbsup",
	//     "item": {"a":"b"},
	//     "event_ts": "1360782804.083113"
	// }
	ReactionRemovedEvent struct {
		Type           string            `json:"type"`
		User           string            `json:"user"`
		Name           string            `json:"name"`
		Item           map[string]string `json:"item"`
		EventTimestamp Time              `json:"event_ts"`
	}

	// https://api.slack.com/events/team_profile_change
	// {
	//     "type": "team_profile_change",
	//     "profile": {
	//         "fields": [{"a":"b"}]
	//     }
	// }
	TeamProfileChangeEvent struct {
		Type    string `json:"type"`
		Profile struct {
			Fields []map[string]interface{} `json:"profile"`
		} `json:"profile"`
	}

	// https://api.slack.com/events/team_profile_delete
	// {
	//     "type": "team_profile_delete",
	//     "profile": {
	//         "fields": [ "Xf06054AAA"]
	//     }
	// }
	TeamProfileDeleteEvent struct {
		Type    string `json:"type"`
		Profile struct {
			Fields []string `json:"fields"`
		} `json:"profile"`
	}

	// https://api.slack.com/events/team_profile_reorder
	// {
	//     "type": "team_profile_reorder",
	//     "profile": {
	//         "fields": [{"a":"b"}]
	//     }
	// }
	TeamProfileReorderEvent struct {
		Type    string `json:"type"`
		Profile struct {
			Fields []map[string]string `json:"profile"`
		} `json:"profile"`
	}

	// https://api.slack.com/events/star_added
	// {
	//     "type": "star_added",
	//     "user": "U024BE7LH",
	//     "item": {},
	//     "event_ts": "1360782804.083113"
	// }
	StarAddedEvent struct {
		Type           string `json:"type"`
		User           string `json:"user"`
		Item           *Item  `json:"item"`
		EventTimestamp Time   `json:"event_ts"`
	}

	// https://api.slack.com/events/star_removed
	// {
	//     "type": "star_removed",
	//     "user": "U024BE7LH",
	//     "item": {},
	//     "event_ts": "1360782804.083113"
	// }
	StarRemovedEvent struct {
		Type           string `json:"type"`
		User           string `json:"user"`
		Item           *Item  `json:"item"`
		EventTimestamp Time   `json:"event_ts"`
	}

	// https://api.slack.com/events/message
	// {
	//     "type": "message",
	//     "channel": "C2147483705",
	//     "user": "U2147483697",
	//     "text": "Hello world",
	//     "ts": "1355517523.000005"
	//     "is_starred": true,
	//     "pinned_to": ["C024BE7LT"],
	//     "reactions": [
	//         {
	//             "name": "astonished",
	//             "count": 3,
	//             "users": [ "U1", "U2", "U3" ]
	//         },
	//         {
	//             "name": "facepalm",
	//             "count": 1034,
	//             "users": [ "U1", "U2", "U3", "U4", "U5" ]
	//         }
	//     ],
	//     "edited": {
	//         "user": "U2147483697",
	//         "ts": "1355517536.000001"
	//     }
	// }
	MessageEvent struct {
		Type      string   `json:"type"`
		Channel   string   `json:"channel"`
		User      string   `json:"user"`
		Text      string   `json:"text"`
		Timestamp Time     `json:"ts"`
		IsStarred bool     `json:"is_starred"`
		PinnedTo  []string `json:"pinned_to"`
		Reactions []struct {
			Count int      `json:"count"`
			Name  string   `json:"name"`
			Users []string `json:"users"`
		} `json:"reactions"`
		Edited struct {
			User      string `json:"user"`
			Timestamp Time   `json:"ts"`
		}
	}

	// https://api.slack.com/events/message/bot_message
	// {
	//     "type": "message",
	//     "subtype": "bot_message",
	//     "ts": "1358877455.000010",
	//     "text": "Pushing is the answer",
	//     "bot_id": "BB12033",
	//     "username": "github",
	//     "icons": {}
	// }
	MessageSubtypeBotMessageEvent struct {
		Type      string            `json:"type"`
		Subtype   string            `json:"subtype"`
		Timestamp Time              `json:"ts"`
		Text      string            `json:"text"`
		BotID     string            `json:"bot_id"`
		Username  string            `json:"username"`
		Icons     map[string]string `json:"icons"`
	}

	// https://api.slack.com/events/message/me_message
	// {
	//     "type": "message",
	//     "subtype": "me_message",
	//     "channel": "C2147483705",
	//     "user": "U2147483697",
	//     "text": "is doing that thing",
	//     "ts": "1355517523.000005"
	// }
	MessageSubtypeMeMessageEvent struct {
		Type      string `json:"type"`
		Subtype   string `json:"subtype"`
		Channel   string `json:"channel"`
		User      string `json:"user"`
		Text      string `json:"text"`
		Timestamp Time   `json:"ts"`
	}

	// https://api.slack.com/events/message/message_changed
	// {
	//     "type": "message",
	//     "subtype": "message_changed",
	//     "hidden": true,
	//     "channel": "C2147483705",
	//     "ts": "1358878755.000001",
	//     "message": {
	//         "type": "message",
	//         "user": "U2147483697",
	//         "text": "Hello, world!",
	//         "ts": "1355517523.000005",
	//         "edited": {
	//             "user": "U2147483697",
	//             "ts": "1358878755.000001"
	//         }
	//     }
	// }
	MessageSubtypeMessageChangedEvent struct {
		Type      string `json:"type"`
		Subtype   string `json:"subtype"`
		Hidden    bool   `json:"hidden"`
		Channel   string `json:"channel"`
		Timestamp Time   `json:"ts"`
		Message   struct {
			Type      string `json:"type"`
			User      string `json:"user"`
			Text      string `json:"text"`
			Timestamp Time   `json:"ts"`
			Edited    struct {
				Timestamp Time   `json:"ts"`
				User      string `json:"user"`
			} `json:"edited"`
		} `json:"message"`
	}

	// https://api.slack.com/events/message/message_deleted
	// {
	//     "type": "message",
	//     "subtype": "message_deleted",
	//     "hidden": true,
	//     "channel": "C2147483705",
	//     "ts": "1358878755.000001",
	//     "deleted_ts": "1358878749.000002"
	// }
	MessageSubtypeMessageDeletedEvent struct {
		Type             string `json:"type"`
		Subtype          string `json:"subtype"`
		Hidden           bool   `json:"hidden"`
		Channel          string `json:"channel"`
		Timestamp        Time   `json:"ts"`
		DeletedTimestamp Time   `json:"deleted_ts"`
	}

	// https://api.slack.com/events/message/channel_join
	// {
	//     "type": "message",
	//     "subtype": "channel_join",
	//     "ts": "1358877458.000011",
	//     "user": "U2147483828",
	//     "text": "<@U2147483828|cal> has joined the channel"
	// }
	MessageSubtypeChannelJoinEvent struct {
		Type      string `json:"type"`
		Subtype   string `json:"subtype"`
		Timestamp Time   `json:"ts"`
		User      string `json:"user"`
		Text      string `json:"text"`
	}

	// https://api.slack.com/events/message/channel_leave
	// {
	//     "type": "message",
	//     "subtype": "channel_leave",
	//     "ts": "1358877455.000010",
	//     "user": "U2147483828",
	//     "text": "<@U2147483828|cal> has left the channel"
	// }
	MessageSubtypeChannelLeaveEvent struct {
		Type      string `json:"type"`
		Subtype   string `json:"subtype"`
		Timestamp Time   `json:"ts"`
		User      string `json:"user"`
		Text      string `json:"text"`
	}

	// https://api.slack.com/events/message/channel_topic
	// {
	//     "type": "message",
	//     "subtype": "channel_topic",
	//     "ts": "1358877455.000010",
	//     "user": "U2147483828",
	//     "topic": "hello world",
	//     "text": "<@U2147483828|cal> set the channel topic: hello world"
	// }
	MessageSubtypeChannelTopicEvent struct {
		Type      string `json:"type"`
		Subtype   string `json:"subtype"`
		Timestamp Time   `json:"ts"`
		User      string `json:"user"`
		Topic     string `json:"topic"`
		Text      string `json:"text"`
	}

	// https://api.slack.com/events/message/channel_purpose
	// {
	//     "type": "message",
	//     "subtype": "channel_purpose",
	//     "ts": "1358877455.000010",
	//     "user": "U2147483828",
	//     "purpose": "whatever",
	//     "text": "<@U2147483828|cal> set the channel purpose: whatever"
	// }
	MessageSubtypeChannelPurposeEvent struct {
		Type      string `json:"type"`
		Subtype   string `json:"subtype"`
		Timestamp Time   `json:"ts"`
		User      string `json:"user"`
		Purpose   string `json:"purpose"`
		Text      string `json:"text"`
	}

	// https://api.slack.com/events/message/channel_name
	// {
	//     "type": "message",
	//     "subtype": "channel_name",
	//     "ts": "1358877455.000010",
	//     "user": "U2147483828",
	//     "old_name": "random",
	//     "name": "watercooler",
	//     "text": "<@U2147483828|cal> has renamed the channek from \"random\" to \"watercooler\""
	// }
	MessageSubtypeChannelNameEvent struct {
		Type      string `json:"type"`
		Subtype   string `json:"subtype"`
		Timestamp Time   `json:"ts"`
		User      string `json:"user"`
		OldName   string `json:"old_name"`
		NewName   string `json:"name"`
		Text      string `json:"text"`
	}

	// https://api.slack.com/events/message/channel_unarchive
	// {
	//     "type": "message",
	//     "subtype": "channel_unarchive",
	//     "ts": "1361482916.000003",
	//     "text": "<U1234|@cal> un-archived the channel",
	//     "user": "U1234"
	// }
	MessageSubtypeChannelUnarchiveEvent struct {
		Type      string `json:"type"`
		Subtype   string `json:"subtype"`
		Timestamp Time   `json:"ts"`
		Text      string `json:"text"`
		User      string `json:"user"`
	}

	// https://api.slack.com/events/message/group_join
	// {
	//     "type": "message",
	//     "subtype": "group_join",
	//     "ts": "1358877458.000011",
	//     "user": "U2147483828",
	//     "text": "<@U2147483828|cal> has joined the group"
	// }
	MessageSubtypeGroupJoinEvent struct {
		Type      string `json:"type"`
		Subtype   string `json:"subtype"`
		Timestamp Time   `json:"ts"`
		User      string `json:"user"`
		Text      string `json:"text"`
	}

	// https://api.slack.com/events/message/group_leave
	// {
	//     "type": "message",
	//     "subtype": "group_leave",
	//     "ts": "1358877455.000010",
	//     "user": "U2147483828",
	//     "text": "<@U2147483828|cal> has left the group"
	// }
	MessageSubtypeGroupLeaveEvent struct {
		Type      string `json:"type"`
		Subtype   string `json:"subtype"`
		Timestamp Time   `json:"ts"`
		User      string `json:"user"`
		Text      string `json:"text"`
	}

	// https://api.slack.com/events/message/group_topic
	// {
	//     "type": "message",
	//     "subtype": "group_topic",
	//     "ts": "1358877455.000010",
	//     "user": "U2147483828",
	//     "topic": "hello world",
	//     "text": "<@U2147483828|cal> set the group topic: hello world"
	// }
	MessageSubtypeGroupTopicEvent struct {
		Type      string `json:"type"`
		Subtype   string `json:"subtype"`
		Timestamp Time   `json:"ts"`
		User      string `json:"user"`
		Topic     string `json:"topic"`
		Text      string `json:"text"`
	}

	// https://api.slack.com/events/message/group_purpose
	// {
	//     "type": "message",
	//     "subtype": "group_purpose",
	//     "ts": "1358877455.000010",
	//     "user": "U2147483828",
	//     "purpose": "whatever",
	//     "text": "<@U2147483828|cal> set the group purpose: whatever"
	// }
	MessageSubtypeGroupPurposeEvent struct {
		Type      string `json:"type"`
		Subtype   string `json:"subtype"`
		Timestamp Time   `json:"ts"`
		User      string `json:"user"`
		Purpose   string `json:"purpose"`
		Text      string `json:"text"`
	}

	// https://api.slack.com/events/message/group_name
	// {
	//     "type": "message",
	//     "subtype": "group_name",
	//     "ts": "1358877455.000010",
	//     "user": "U2147483828",
	//     "old_name": "random",
	//     "name": "watercooler",
	//     "text": "<@U2147483828|cal> has renamed the group from \"random\" to \"watercooler\""
	// }
	MessageSubtypeGroupNameEvent struct {
		Type      string `json:"type"`
		Subtype   string `json:"subtype"`
		Timestamp Time   `json:"ts"`
		User      string `json:"user"`
		OldName   string `json:"old_name"`
		NewName   string `json:"name"`
		Text      string `json:"text"`
	}

	// https://api.slack.com/events/message/group_unarchive
	// {
	//     "type": "message",
	//     "subtype": "group_unarchive",
	//     "ts": "1361482916.000003",
	//     "text": "<U1234|@cal> un-archived the group",
	//     "user": "U1234"
	// }
	MessageSubtypeGroupUnarchiveEvent struct {
		Type      string `json:"type"`
		Subtype   string `json:"subtype"`
		Timestamp Time   `json:"ts"`
		Text      string `json:"text"`
		User      string `json:"user"`
	}

	// https://api.slack.com/events/message/file_comment
	// {
	//     "type": "message",
	//     "subtype": "file_comment",
	//     "ts": "1361482916.000003",
	//     "text": "<@cal> commented on a file: ...",
	//     "file": {},
	//     "comment": {}
	// }
	MessageSubtypeFileCommentEvent struct {
		Type      string   `json:"type"`
		Subtype   string   `json:"subtype"`
		Timestamp Time     `json:"ts"`
		Text      string   `json:"text"`
		File      *File    `json:"file"`
		Comment   *Comment `json:"comment"`
	}

	// https://api.slack.com/events/message/channel_archive
	//  {
	//      "type": "message",
	//      "subtype": "channel_archive",
	//      "ts": "1361482916.000003",
	//      "text": "<U1234|@cal> archived the channel",
	//      "user": "U1234",
	//      "members": ["U1234", "U5678"]
	//  }
	MessageSubtypeChannelArchiveEvent struct {
		Type      string   `json:"type"`
		Subtype   string   `json:"subtype"`
		Timestamp Time     `json:"ts"`
		Text      string   `json:"text"`
		User      string   `json:"user"`
		Members   []string `json:"members"`
	}

	// https://api.slack.com/events/message/group_archive
	//   {
	//      "type": "message",
	//      "subtype": "group_archive",
	//      "ts": "1361482916.000003",
	//      "text": "<U1234|@cal> archived the group",
	//      "user": "U1234",
	//      "members": ["U1234", "U5678"]
	//  }
	MessageSubtypeGroupArchiveEvent struct {
		Type      string   `json:"type"`
		Subtype   string   `json:"subtype"`
		Timestamp Time     `json:"ts"`
		Text      string   `json:"text"`
		User      string   `json:"user"`
		Members   []string `json:"members"`
	}

	// https://api.slack.com/events/message/file_share
	//   {
	//      "type": "message",
	//      "subtype": "file_share",
	//      "ts": "1358877455.000010",
	//      "text": "<@cal> uploaded a file: <https:...7.png|7.png>",
	//      "file": {},
	//      "user": "U2147483697",
	//      "upload": true
	//  }
	MessageSubtypeFileShareEvent struct {
		Type      string `json:"type"`
		Subtype   string `json:"subtype"`
		Timestamp Time   `json:"ts"`
		Text      string `json:"text"`
		File      *File  `json:"file"`
		User      string `json:"user"`
		Upload    bool   `json:"upload"`
	}

	// https://api.slack.com/events/message/file_mention
	//   {
	//      "type": "message",
	//      "subtype": "file_mention",
	//      "ts": "1358877455.000010",
	//      "text": "<@cal> mentioned a file: <https:...7.png|7.png>",
	//      "file": {},
	//      "user": "U2147483697"
	//  }
	MessageSubtypeFileMentionEvent struct {
		Type      string `json:"type"`
		Subtype   string `json:"subtype"`
		Timestamp Time   `json:"ts"`
		Text      string `json:"text"`
		File      *File  `json:"file"`
		User      string `json:"user"`
	}

	// https://api.slack.com/events/message/pinned_item
	//   {
	//      "type": "message",
	//      "subtype": "pinned_item",
	//      "user": "U024BE7LH",
	//      "item_type": "F",
	//      "text": "<@U024BE7LH|cal> pinned their Image <https:...7.png|7.png> to this channel.",
	//      "item": {},
	//      "channel": "C02ELGNBH",
	//      "ts": "1360782804.083113"
	//  }
	MessageSubtypePinnedItemEvent struct {
		Type      string `json:"type"`
		Subtype   string `json:"subtype"`
		User      string `json:"user"`
		ItemType  string `json:"item_type"`
		Text      string `json:"text"`
		Item      *Item  `json:"item"`
		Channel   string `json:"channel"`
		Timestamp Time   `json:"ts"`
	}

	// https://api.slack.com/events/message/unpinned_item
	//   {
	//      "type": "message",
	//      "subtype": "unpinned_item",
	//      "user": "USLACKBOT",
	//      "item_type": "G",
	//      "text": "<@U024BE7LH|cal> unpinned the message you pinned to the secretplans group.",
	//      "item": {},
	//      "channel": "G024BE91L",
	//      "ts": "1360782804.083113"
	//  }
	MessageSubtypeUnpinnedItemEvent struct {
		Type      string `json:"type"`
		Subtype   string `json:"subtype"`
		User      string `json:"user"`
		ItemType  string `json:"item_type"`
		Text      string `json:"text"`
		Item      *Item  `json:"item"`
		Channel   string `json:"channel"`
		Timestamp Time   `json:"ts"`
	}

	UnrecognizedEvent map[string]interface{}
)

const (
	AccountsChangedEventCode       = `accounts_changed`        // The list of accounts a user is signed into has changed
	BotAddedEventCode              = `bot_added`               // An integration bot was added
	BotChangedEventCode            = `bot_changed`             // An integration bot was changed
	ChannelArchiveEventCode        = `channel_archive`         // A team channel was archived
	ChannelCreatedEventCode        = `channel_created`         // A team channel was created
	ChannelDeletedEventCode        = `channel_deleted`         // A team channel was deleted
	ChannelHistoryChangedEventCode = `channel_history_changed` // Bulk updates were made to a channel's history
	ChannelJoinedEventCode         = `channel_joined`          // You joined a channel
	ChannelLeftEventCode           = `channel_left`            // You left a channel
	ChannelMarkedEventCode         = `channel_marked`          // Your channel read marker was updated
	ChannelRenameEventCode         = `channel_rename`          // A team channel was renamed
	ChannelUnarchiveEventCode      = `channel_unarchive`       // A team channel was unarchived
	CommandsChangedEventCode       = `commands_changed`        // A team slash command has been added or changed
	DndUpdatedEventCode            = `dnd_updated`             // Do not Disturb settings changed for the current user
	DndUpdatedUserEventCode        = `dnd_updated_user`        // Do not Disturb settings changed for a team member
	EmailDomainChangedEventCode    = `email_domain_changed`    // The team email domain has changed
	EmojiChangedEventCode          = `emoji_changed`           // A team custom emoji has been added or changed
	FileChangeEventCode            = `file_change`             // A file was changed
	FileCommentAddedEventCode      = `file_comment_added`      // A file comment was added
	FileCommentDeletedEventCode    = `file_comment_deleted`    // A file comment was deleted
	FileCommentEditedEventCode     = `file_comment_edited`     // A file comment was edited
	FileCreatedEventCode           = `file_created`            // A file was created
	FileDeletedEventCode           = `file_deleted`            // A file was deleted
	FilePrivateEventCode           = `file_private`            // A file was made private
	FilePublicEventCode            = `file_public`             // A file was made public
	FileSharedEventCode            = `file_shared`             // A file was shared
	FileUnsharedEventCode          = `file_unshared`           // A file was unshared
	GroupArchiveEventCode          = `group_archive`           // A private group was archived
	GroupCloseEventCode            = `group_close`             // You closed a group channel
	GroupHistoryChangedEventCode   = `group_history_changed`   // Bulk updates were made to a group's history
	GroupJoinedEventCode           = `group_joined`            // You joined a private group
	GroupLeftEventCode             = `group_left`              // You left a private group
	GroupMarkedEventCode           = `group_marked`            // A private group read marker was updated
	GroupOpenEventCode             = `group_open`              // You opened a group channel
	GroupRenameEventCode           = `group_rename`            // A private group was renamed
	GroupUnarchiveEventCode        = `group_unarchive`         // A private group was unarchived
	HelloEventCode                 = `hello`                   // The client has successfully connected to the server
	ImCloseEventCode               = `im_close`                // You closed a direct message channel
	ImCreatedEventCode             = `im_created`              // A direct message channel was created
	ImHistoryChangedEventCode      = `im_history_changed`      // Bulk updates were made to a DM channel's history
	ImMarkedEventCode              = `im_marked`               // A direct message read marker was updated
	ImOpenEventCode                = `im_open`                 // You opened a direct message channel
	ManualPresenceChangeEventCode  = `manual_presence_change`  // You manually updated your presence
	MessageEventCode               = `message`                 // A message was sent to a channel
	PinAddedEventCode              = `pin_added`               // A pin was added to a channel
	PinRemovedEventCode            = `pin_removed`             // A pin was removed from a channel
	PrefChangeEventCode            = `pref_change`             // You have updated your preferences
	PresenceChangeEventCode        = `presence_change`         // A team member's presence changed
	ReactionAddedEventCode         = `reaction_added`          // A team member has added an emoji reaction to an item
	ReactionRemovedEventCode       = `reaction_removed`        // A team member removed an emoji reaction
	ReconnectUrlEventCode          = `reconnect_url`           // Experimental
	StarAddedEventCode             = `star_added`              // A team member has starred an item
	StarRemovedEventCode           = `star_removed`            // A team member removed a star
	SubteamCreatedEventCode        = `subteam_created`         // A user group has been added to the team
	SubteamSelfAddedEventCode      = `subteam_self_added`      // You have been added to a user group
	SubteamSelfRemovedEventCode    = `subteam_self_removed`    // You have been removed from a user group
	SubteamUpdatedEventCode        = `subteam_updated`         // An existing user group has been updated or its members changed
	TeamDomainChangeEventCode      = `team_domain_change`      // The team domain has changed
	TeamJoinEventCode              = `team_join`               // A new team member has joined
	TeamMigrationStartedEventCode  = `team_migration_started`  // The team is being migrated between servers
	TeamPlanChangeEventCode        = `team_plan_change`        // The team billing plan has changed
	TeamPrefChangeEventCode        = `team_pref_change`        // A team preference has been updated
	TeamProfileChangeEventCode     = `team_profile_change`     // Team profile fields have been updated
	TeamProfileDeleteEventCode     = `team_profile_delete`     // Team profile fields have been deleted
	TeamProfileReorderEventCode    = `team_profile_reorder`    // Team profile fields have been reordered
	TeamRenameEventCode            = `team_rename`             // The team name has changed
	UserChangeEventCode            = `user_change`             // A team member's data has changed
	UserTypingEventCode            = `user_typing`             // A channel member is typing a message

	// message subtype
	MessageSubtypeBotMessageEventCode       = `bot_message`       // A message was posted by an integration
	MessageSubtypeMeMessageEventCode        = `me_message`        // A /me message was sent
	MessageSubtypeMessageChangedEventCode   = `message_changed`   // A message was changed
	MessageSubtypeMessageDeletedEventCode   = `message_deleted`   // A message was deleted
	MessageSubtypeChannelJoinEventCode      = `channel_join`      // A team member joined a channel
	MessageSubtypeChannelLeaveEventCode     = `channel_leave`     // A team member left a channel
	MessageSubtypeChannelTopicEventCode     = `channel_topic`     // A channel topic was updated
	MessageSubtypeChannelPurposeEventCode   = `channel_purpose`   // A channel purpose was updated
	MessageSubtypeChannelNameEventCode      = `channel_name`      // A channel was renamed
	MessageSubtypeChannelArchiveEventCode   = `channel_archive`   // A channel was archived
	MessageSubtypeChannelUnarchiveEventCode = `channel_unarchive` // A channel was unarchived
	MessageSubtypeGroupJoinEventCode        = `group_join`        // A team member joined a group
	MessageSubtypeGroupLeaveEventCode       = `group_leave`       // A team member left a group
	MessageSubtypeGroupTopicEventCode       = `group_topic`       // A group topic was updated
	MessageSubtypeGroupPurposeEventCode     = `group_purpose`     // A group purpose was updated
	MessageSubtypeGroupNameEventCode        = `group_name`        // A group was renamed
	MessageSubtypeGroupArchiveEventCode     = `group_archive`     // A group was archived
	MessageSubtypeGroupUnarchiveEventCode   = `group_unarchive`   // A group was unarchived
	MessageSubtypeFileShareEventCode        = `file_share`        // A file was shared into a channel
	MessageSubtypeFileCommentEventCode      = `file_comment`      // A comment was added to a file
	MessageSubtypeFileMentionEventCode      = `file_mention`      // A file was mentioned in a channel
	MessageSubtypePinnedItemEventCode       = `pinned_item`       // An item was pinned in a channel
	MessageSubtypeUnpinnedItemEventCode     = `unpinned_item`     // An item was unpinned from a channel
)
