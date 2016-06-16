package slk

type (

	// https://api.slack.com/methods/api.test
	// {
	//     "ok": true,
	//     "args": {
	//         "foo": "bar"
	//     }
	// }
	// {
	//     "ok": false,
	//     "error": "my_error",
	//     "args": {
	//         "error": "my_error"
	//     }
	// }
	ApiTestResponse struct {
		OK    bool              `json:"ok"`
		Error string            `json:"error"`
		Args  map[string]string `json:"args"`
	}

	// https://api.slack.com/methods/auth.test
	// {
	//     "ok": true,
	//     "url": "https:\/\/myteam.slack.com\/",
	//     "team": "My Team",
	//     "user": "cal",
	//     "team_id": "T12345",
	//     "user_id": "U12345"
	// }
	AuthTestResponse struct {
		OK     bool   `json:"ok"`
		Error  string `json:"error"`
		URL    string `json:"url"`
		Team   string `json:"team"`
		User   string `json:"user"`
		TeamID string `json:"team_id"`
		UserID string `json:"user_id"`
	}

	// https://api.slack.com/methods/channels.archive
	// {
	//     "ok": true
	// }
	ChannelsArchiveResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}

	// https://api.slack.com/methods/channels.create
	// {
	//     "ok": true,
	//     "channel": {
	//         "id": "C024BE91L",
	//         "name": "fun",
	//         "created": 1360782804,
	//         "creator": "U024BE7LH",
	//         "is_archived": false,
	//         "is_member": true,
	//         "is_general": false,
	//         "last_read": "0000000000.000000",
	//         "latest": null,
	//         "unread_count": 0,
	//         "unread_count_display": 0,
	//         "members": [ ],
	//         "topic": { },
	//         "purpose": { }
	//     }
	// }
	ChannelsCreateResponse struct {
		OK      bool     `json:"ok"`
		Error   string   `json:"error"`
		Channel *Channel `json:"channel"`
	}

	// https://api.slack.com/methods/channels.history
	// {
	//     "ok": true,
	//     "latest": "1358547726.000003",
	//     "messages": [
	//         {
	//             "type": "message",
	//             "ts": "1358546515.000008",
	//             "user": "U2147483896",
	//             "text": "Hello"
	//         },
	//         {
	//             "type": "message",
	//             "ts": "1358546515.000007",
	//             "user": "U2147483896",
	//             "text": "World",
	//             "is_starred": true,
	//             "reactions": [
	//                 {
	//                     "name": "space_invader",
	//                     "count": 3,
	//                     "users": [ "U1", "U2", "U3" ]
	//                 },
	//                 {
	//                     "name": "sweet_potato",
	//                     "count": 5,
	//                     "users": [ "U1", "U2", "U3", "U4", "U5" ]
	//                 }
	//             ]
	//         },
	//         {
	//             "type": "something_else",
	//             "ts": "1358546515.000007",
	//             "wibblr": true
	//         }
	//     ],
	//     "has_more": false
	// }
	ChannelsHistoryResponse struct {
		OK       bool       `json:"ok"`
		Error    string     `json:"error"`
		Latest   Time       `json:"latest"`
		Messages []*Message `json:"messages"`
		HasMore  bool       `json:"has_more"`
	}

	// https://api.slack.com/methods/channels.info
	// {
	//     "ok": true,
	//     "channel": {
	//         "id": "C024BE91L",
	//         "name": "fun",
	//         "created": 1360782804,
	//         "creator": "U024BE7LH",
	//         "is_archived": false,
	//         "is_general": false,
	//         "is_member": true,
	//         "members": [ ],
	//         "topic": { },
	//         "purpose": { },
	//         "last_read": "1401383885.000061",
	//         "latest": { },
	//         "unread_count": 0,
	//         "unread_count_display": 0
	//     }
	// }
	ChannelsInfoResponse struct {
		OK      bool     `json:"ok"`
		Error   string   `json:"error"`
		Channel *Channel `json:"channel"`
	}

	// https://api.slack.com/methods/channels.invite
	// {
	//     "ok": true,
	//     "channel": {
	//         "id": "C024BE91L",
	//         "name": "fun",
	//         "created": 1360782804,
	//         "creator": "U024BE7LH",
	//         "is_archived": false,
	//         "is_member": true,
	//         "is_general": false,
	//         "last_read": "1401383885.000061",
	//         "latest": { },
	//         "unread_count": 0,
	//         "unread_count_display": 0,
	//         "members": [ ],
	//         "topic": {
	//             "value": "Fun times",
	//             "creator": "U024BE7LV",
	//             "last_set": 1369677212
	//         },
	//         "purpose": {
	//             "value": "This channel is for fun",
	//             "creator": "U024BE7LH",
	//             "last_set": 1360782804
	//         }
	//     }
	// }
	ChannelsInviteResponse struct {
		OK      bool     `json:"ok"`
		Error   string   `json:"error"`
		Channel *Channel `json:"channel"`
	}

	// https://api.slack.com/methods/channels.join
	// {
	//     "ok": true,
	//     "channel": {
	//         "id": "C024BE91L",
	//         "name": "fun",
	//         "created": 1360782804,
	//         "creator": "U024BE7LH",
	//         "is_archived": false,
	//         "is_member": true,
	//         "is_general": false,
	//         "last_read": "1401383885.000061",
	//         "latest": { },
	//         "unread_count": 0,
	//         "unread_count_display": 0,
	//         "members": [ ],
	//         "topic": {
	//             "value": "Fun times",
	//             "creator": "U024BE7LV",
	//             "last_set": 1369677212
	//         },
	//         "purpose": {
	//             "value": "This channel is for fun",
	//             "creator": "U024BE7LH",
	//             "last_set": 1360782804
	//         }
	//     }
	// }
	// {
	//     "ok": true,
	//     "already_in_channel": true,
	//     "channel": {
	//         "id": "C024BE91L",
	//         "name": "fun",
	//         "created": 1360782804,
	//         "creator": "U024BE7LH",
	//         "is_archived": false,
	//         "is_general": false
	//     }
	// }
	ChannelsJoinResponse struct {
		OK               bool     `json:"ok"`
		Error            string   `json:"error"`
		AlreadyInChannel bool     `json:"already_in_channel"`
		Channel          *Channel `json:"channel"`
	}

	// https://api.slack.com/methods/channels.kick
	// {
	//     "ok": true
	// }
	ChannelsKickResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}

	// https://api.slack.com/methods/channels.leave
	// {
	//     "ok": true
	// }
	// {
	//     "ok": true,
	//     "not_in_channel": true
	// }
	ChannelsLeaveResponse struct {
		OK           bool   `json:"ok"`
		Error        string `json:"error"`
		NotInChannel bool   `json:"not_in_channel"`
	}

	// https://api.slack.com/methods/channels.list
	// {
	//     "ok": true,
	//     "channels": [
	//         {
	//             "id": "C024BE91L",
	//             "name": "fun",
	//             "created": 1360782804,
	//             "creator": "U024BE7LH",
	//             "is_archived": false,
	//             "is_member": false,
	//             "num_members": 6,
	//             "topic": {
	//                 "value": "Fun times",
	//                 "creator": "U024BE7LV",
	//                 "last_set": 1369677212
	//             },
	//             "purpose": {
	//                 "value": "This channel is for fun",
	//                 "creator": "U024BE7LH",
	//                 "last_set": 1360782804
	//             }
	//         },
	//         ....
	//     ]
	// }
	ChannelsListResponse struct {
		OK       bool       `json:"ok"`
		Error    string     `json:"error"`
		Channels []*Channel `json:"channels"`
	}

	// https://api.slack.com/methods/channels.mark
	// {
	//     "ok": true
	// }
	ChannelsMarkResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}

	// https://api.slack.com/methods/channels.rename
	// {
	//     "ok": true,
	//     "channel": {
	//         "id": "C024BE91L",
	//         "is_channel": true,
	//         "name": "new_name",
	//         "created": 1360782804
	//     }
	// }
	ChannelsRenameResponse struct {
		OK      bool     `json:"ok"`
		Error   string   `json:"error"`
		Channel *Channel `json:"channel"`
	}

	// https://api.slack.com/methods/channels.setPurpose
	// {
	//     "ok": true,
	//     "purpose": "This is the new purpose!"
	// }
	ChannelsSetPurposeResponse struct {
		OK      bool   `json:"ok"`
		Error   string `json:"error"`
		Purpose string `json:"purpose"`
	}

	// https://api.slack.com/methods/channels.setTopic
	// {
	//     "ok": true,
	//     "topic": "This is the new topic!"
	// }
	ChannelsSetTopicResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
		Topic string `json:"topic"`
	}

	// https://api.slack.com/methods/channels.unarchive
	// {
	//     "ok": true
	// }
	ChannelsUnarchiveResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}

	// https://api.slack.com/methods/chat.delete
	// {
	//     "ok": true,
	//     "channel": "C024BE91L",
	//     "ts": "1401383885.000061"
	// }
	ChatDeleteResponse struct {
		OK        bool   `json:"ok"`
		Error     string `json:"error"`
		Channel   string `json:"channel"`
		Timestamp Time   `json:"ts"`
	}

	// https://api.slack.com/methods/chat.postMessage
	// {
	//     "ok": true,
	//     "ts": "1405895017.000506",
	//     "channel": "C024BE91L",
	//     "message": {}
	// }
	ChatPostMessageResponse struct {
		OK        bool     `json:"ok"`
		Error     string   `json:"error"`
		Timestamp Time     `json:"ts"`
		Channel   string   `json:"channel"`
		Message   *Message `json:"message"`
	}

	// https://api.slack.com/methods/chat.update
	// {
	//     "ok": true,
	//     "channel": "C024BE91L",
	//     "ts": "1401383885.000061",
	//     "text": "Updated Text"
	// }
	ChatUpdateResponse struct {
		OK        bool   `json:"ok"`
		Error     string `json:"error"`
		Channel   string `json:"channel"`
		Timestamp Time   `json:"ts"`
		Text      string `json:"text"`
	}

	// https://api.slack.com/methods/dnd.endDnd
	// {
	//     "ok": true
	// }
	DndEndDndResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}

	// https://api.slack.com/methods/dnd.endSnooze
	// {
	//     "ok": true,
	//     "dnd_enabled": true,
	//     "next_dnd_start_ts": 1450418400,
	//     "next_dnd_end_ts": 1450454400,
	//     "snooze_enabled": false
	// }
	DndEndSnoozeResponse struct {
		OK                    bool   `json:"ok"`
		Error                 string `json:"error"`
		DndEnabled            bool   `json:"dnd_enabled"`
		NextDndStartTimestamp Time   `json:"next_dnd_start_ts"`
		NextDndEndTimestamp   Time   `json:"next_dnd_end_ts"`
		SnoozeEnabled         bool   `json:"snooze_enabled"`
	}

	// https://api.slack.com/methods/dnd.info
	// {
	//     "ok": true,
	//     "dnd_enabled": true,
	//     "next_dnd_start_ts": 1450416600,
	//     "next_dnd_end_ts": 1450452600,
	//     "snooze_enabled": true,
	//     "snooze_endtime": 1450416600,
	//     "snooze_remaining": 1196
	// }
	DndInfoResponse struct {
		OK                    bool     `json:"ok"`
		Error                 string   `json:"error"`
		DndEnabled            bool     `json:"dnd_enabled"`
		NextDndStartTimestamp Time     `json:"next_dnd_start_ts"`
		NextDndEndTimestamp   Time     `json:"next_dnd_end_ts"`
		SnoozeEnabled         bool     `json:"snooze_enabled"`
		SnoozeEndTime         Time     `json:"snooze_endtime"`
		SnoozeRemaining       Duration `json:"snooze_remaining"`
	}

	// https://api.slack.com/methods/dnd.setSnooze
	// {
	//     "ok": true,
	//     "snooze_enabled": true,
	//     "snooze_endtime": 1450373897,
	//     "snooze_remaining": 60,
	// }
	DndSetSnoozeResponse struct {
		OK              bool     `json:"ok"`
		Error           string   `json:"error"`
		SnoozeEnabled   bool     `json:"snooze_enabled"`
		SnoozeEndTime   Time     `json:"snooze_endtime"`
		SnoozeRemaining Duration `json:"snooze_remaining"`
	}

	// https://api.slack.com/methods/dnd.teamInfo
	// {
	//     "ok": true,
	//     "users": {
	//         "U023BECGF": {
	//             "dnd_enabled": true,
	//             "next_dnd_start_ts": 1450387800,
	//             "next_dnd_end_ts": 1450423800
	//         },
	//         "U058CJVAA": {
	//             "dnd_enabled": false,
	//             "next_dnd_start_ts": 1,
	//             "next_dnd_end_ts": 1
	//         }
	//     }
	// }
	DndTeamInfoResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
		Users map[string]struct {
			DndEnabled            bool `json:"dnd_enabled"`
			NextDndEndTimestamp   Time `json:"next_dnd_end_ts"`
			NextDndStartTimestamp Time `json:"next_dnd_start_ts"`
		} `json:"users"`
	}

	// https://api.slack.com/methods/emoji.list
	// {
	//     "ok": true,
	//     "emoji": {
	//         "bowtie": "https:\/\/my.slack.com\/emoji\/bowtie\/46ec6f2bb0.png",
	//         "squirrel": "https:\/\/my.slack.com\/emoji\/squirrel\/f35f40c0e0.png",
	//         "shipit": "alias:squirrel"
	//     }
	// }
	EmojiListResponse struct {
		OK    bool              `json:"ok"`
		Error string            `json:"error"`
		Emoji map[string]string `json:"emoji"`
	}

	// https://api.slack.com/methods/files.comments.add
	// {
	//     "ok": true,
	//     "comment": {
	//         "id": "Fc1234567890",
	//         "created": 1356032811,
	//         "timestamp": 1356032811,
	//         "user": "U1234567890",
	//         "comment": "Everyone should take a moment to read this file."
	//     }
	// }
	FilesCommentsAddResponse struct {
		OK      bool     `json:"ok"`
		Error   string   `json:"error"`
		Comment *Comment `json:"comment"`
	}

	// https://api.slack.com/methods/files.comments.delete
	// {
	//     "ok": true
	// }
	FilesCommentsDeleteResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}

	// https://api.slack.com/methods/files.comments.edit
	// {
	//     "ok": true,
	//     "comment": {
	//         "id": "Fc1234567890",
	//         "created": 1356032811,
	//         "timestamp": 1356032811,
	//         "user": "U1234567890",
	//         "comment": "Everyone should take a moment to read this file, seriously."
	//     }
	// }
	FilesCommentsEditResponse struct {
		OK      bool     `json:"ok"`
		Error   string   `json:"error"`
		Comment *Comment `json:"comment"`
	}

	// https://api.slack.com/methods/files.delete
	// {
	//     "ok": true
	// }
	FilesDeleteResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}

	// https://api.slack.com/methods/files.info
	// {
	//     "ok": true,
	//     "file": {
	//         "id" : "F2147483862",
	//         "timestamp" : 1356032811,
	//         "name" : "file.htm",
	//         "title" : "My HTML file",
	//         "mimetype" : "text\/plain",
	//         "filetype" : "text",
	//         "pretty_type": "Text",
	//         "user" : "U2147483697",
	//         "mode" : "hosted",
	//         "editable" : true,
	//         "is_external": false,
	//         "external_type": "",
	//         "size" : 12345,
	//         "url": "https:\/\/slack-files.com\/files-pub\/T024BE7LD-F024BERPE-09acb6\/1.png",
	//         "url_download": "https:\/\/slack-files.com\/files-pub\/T024BE7LD-F024BERPE-09acb6\/download\/1.png",
	//         "url_private": "https:\/\/slack.com\/files-pri\/T024BE7LD-F024BERPE\/1.png",
	//         "url_private_download": "https:\/\/slack.com\/files-pri\/T024BE7LD-F024BERPE\/download\/1.png",
	//         "thumb_64": "https:\/\/slack-files.com\/files-tmb\/T024BE7LD-F024BERPE-c66246\/1_64.png",
	//         "thumb_80": "https:\/\/slack-files.com\/files-tmb\/T024BE7LD-F024BERPE-c66246\/1_80.png",
	//         "thumb_360": "https:\/\/slack-files.com\/files-tmb\/T024BE7LD-F024BERPE-c66246\/1_360.png",
	//         "thumb_360_gif": "https:\/\/slack-files.com\/files-tmb\/T024BE7LD-F024BERPE-c66246\/1_360.gif",
	//         "thumb_360_w": 100,
	//         "thumb_360_h": 100,
	//         "permalink" : "https:\/\/tinyspeck.slack.com\/files\/cal\/F024BERPE\/1.png",
	//         "edit_link" : "https:\/\/tinyspeck.slack.com\/files\/cal\/F024BERPE\/1.png/edit",
	//         "preview" : "&lt;!DOCTYPE html&gt;\n&lt;html&gt;\n&lt;meta charset=&#39;utf-8&#39;&gt;",
	//         "preview_highlight" : &lt;div class=\"sssh-code\"&gt;&lt;div class=\"sssh-line\"&gt;&lt;pre&gt;&lt;!DOCTYPE html...",
	//         "lines" : 123,
	//         "lines_more": 118,
	//         "is_public": true,
	//         "public_url_shared": false,
	//         "channels": ["C024BE7LT", ...],
	//         "groups": ["G12345", ...],
	//         "initial_comment": {...},
	//         "num_stars": 7,
	//         "is_starred": true
	//     },
	//     "comments": [
	//         {
	//             "id": "Fc027BN9L9",
	//             "timestamp": 1356032811,
	//             "user": "U2147483697",
	//             "comment": "This is a comment"
	//         }
	//     ],
	//     "paging": {
	//         "count": 100,
	//         "total": 2,
	//         "page": 1,
	//         "pages": 0
	//     }
	// }
	FilesInfoResponse struct {
		OK       bool       `json:"ok"`
		Error    string     `json:"error"`
		File     *File      `json:"file"`
		Comments []*Comment `json:"comments"`
		Paging   struct {
			Count int `json:"count"`
			Page  int `json:"page"`
			Pages int `json:"pages"`
			Total int `json:"total"`
		} `json:"paging"`
	}

	// https://api.slack.com/methods/files.list
	// {
	//     "ok": true,
	//     "files": [
	//         {...},
	//         {...},
	//         {...},
	//         ...
	//     ],
	//     "paging": {
	//         "count": 100,
	//         "total": 295,
	//         "page": 1,
	//         "pages": 3
	//     }
	// }
	FilesListResponse struct {
		OK     bool    `json:"ok"`
		Error  string  `json:"error"`
		Files  []*File `json:"files"`
		Paging struct {
			Count int `json:"count"`
			Page  int `json:"page"`
			Pages int `json:"pages"`
			Total int `json:"total"`
		} `json:"paging"`
	}

	// https://api.slack.com/methods/files.upload
	// {
	//     "ok": true,
	//     "file": {...}
	// }
	FilesUploadResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
		File  *File  `json:"file"`
	}

	// https://api.slack.com/methods/groups.archive
	// {
	//     "ok": true
	// }
	GroupsArchiveResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}

	// https://api.slack.com/methods/groups.close
	// {
	//     "ok": true
	// }
	// {
	//     "ok": true,
	//     "no_op": true,
	//     "already_closed": true
	// }
	GroupsCloseResponse struct {
		OK            bool   `json:"ok"`
		Error         string `json:"error"`
		NOOP          bool   `json:"no_op"`
		AlreadyClosed bool   `json:"already_closed"`
	}

	// https://api.slack.com/methods/groups.create
	// {
	//     "ok": true,
	//     "group": {
	//         "id": "G024BE91L",
	//         "name": "secretplans",
	//         "is_group": "true",
	//         "created": 1360782804,
	//         "creator": "U024BE7LH",
	//         "is_archived": false,
	//         "is_open": true,
	//         "last_read": "0000000000.000000",
	//         "latest": null,
	//         "unread_count": 0,
	//         "unread_count_display": 0,
	//         "members": [
	//             "U024BE7LH"
	//         ],
	//         "topic": {
	//             "value": "Secret plans on hold",
	//             "creator": "U024BE7LV",
	//             "last_set": 1369677212
	//         },
	//         "purpose": {
	//             "value": "Discuss secret plans that no-one else should know",
	//             "creator": "U024BE7LH",
	//             "last_set": 1360782804
	//         }
	//     }
	// }
	GroupsCreateResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
		Group *Group `json:"group"`
	}

	// https://api.slack.com/methods/groups.createChild
	// {
	//     "ok": true,
	//     "group": {
	//         "id": "G024BE91L",
	//         "name": "secretplans",
	//         "is_group": "true",
	//         "created": 1360782804,
	//         "creator": "U024BE7LH",
	//         "is_archived": false,
	//         "members": [
	//             "U024BE7LH"
	//         ]
	//     }
	// }
	GroupsCreateChildResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
		Group *Group `json:"group"`
	}

	// https://api.slack.com/methods/groups.history
	// {
	//     "ok": true,
	//     "latest": "1358547726.000003",
	//     "messages": [
	//         {
	//             "type": "message",
	//             "ts": "1358546515.000008",
	//             "user": "U2147483896",
	//             "text": "Hello"
	//         },
	//         {
	//             "type": "message",
	//             "ts": "1358546515.000007",
	//             "user": "U2147483896",
	//             "text": "World",
	//             "is_starred": true,
	//         },
	//         {
	//             "type": "something_else",
	//             "ts": "1358546515.000007",
	//             "wibblr": true
	//         }
	//     ],
	//     "has_more": false
	// }
	GroupsHistoryResponse struct {
		OK       bool       `json:"ok"`
		Error    string     `json:"error"`
		Latest   Time       `json:"latest"`
		Messages []*Message `json:"group"`
		HasMore  bool       `json:"has_more"`
	}

	// https://api.slack.com/methods/groups.info
	// {
	//     "ok": true,
	//     "group": {
	//         "id": "G024BE91L",
	//         "name": "secretplans",
	//         "is_group": "true",
	//         "created": 1360782804,
	//         "creator": "U024BE7LH",
	//         "is_archived": false,
	//         "members": [
	//             "U024BE7LH"
	//         ],
	//         "topic": { },
	//         "purpose": { },
	//         "last_read": "1401383885.000061",
	//         "latest": { }
	//         "unread_count": 0,
	//         "unread_count_display": 0

	//     },
	// }
	GroupsInfoResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
		Group *Group `json:"group"`
	}

	// https://api.slack.com/methods/groups.invite
	// {
	//     "ok": true,
	//     "group": {
	//
	//     },
	// }
	// {
	//     "ok": true,
	//     "already_in_group": true,
	//     "group": {},
	// }
	GroupsInviteResponse struct {
		OK             bool   `json:"ok"`
		Error          string `json:"error"`
		AlreadyInGroup bool   `json:"already_in_group"`
		Group          *Group `json:"group"`
	}

	// https://api.slack.com/methods/groups.kick
	// {
	//     "ok": true
	// }
	GroupsKickResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}

	// https://api.slack.com/methods/groups.leave
	// {
	//     "ok": true
	// }
	GroupsLeaveResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}

	// https://api.slack.com/methods/groups.list
	// {
	//     "ok": true,
	//     "groups": [
	//         {
	//             "id": "G024BE91L",
	//             "name": "secretplans",
	//             "created": 1360782804,
	//             "creator": "U024BE7LH",
	//             "is_archived": false,
	//             "members": [
	//                 "U024BE7LH"
	//             ],
	//             "topic": {
	//                 "value": "Secret plans on hold",
	//                 "creator": "U024BE7LV",
	//                 "last_set": 1369677212
	//             },
	//             "purpose": {
	//                 "value": "Discuss secret plans that no-one else should know",
	//                 "creator": "U024BE7LH",
	//                 "last_set": 1360782804
	//             }
	//         }
	//     ]
	// }
	GroupsListResponse struct {
		OK     bool     `json:"ok"`
		Error  string   `json:"error"`
		Groups []*Group `json:"groups"`
	}

	// https://api.slack.com/methods/groups.mark
	// {
	//     "ok": true
	// }
	GroupsMarkResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}

	// https://api.slack.com/methods/groups.open
	// {
	//     "ok": true
	// }
	// {
	//     "ok": true,
	//     "no_op": true,
	//     "already_open": true
	// }
	GroupsOpenResponse struct {
		OK          bool   `json:"ok"`
		Error       string `json:"error"`
		NOOP        bool   `json:"no_op"`
		AlreadyOpen bool   `json:"already_open"`
	}

	// https://api.slack.com/methods/groups.rename
	// {
	//     "ok": true,
	//     "channel": {
	//         "id": "C024BE91L",
	//         "is_group": true,
	//         "name": "new_name",
	//         "created": 1360782804
	//     }
	// }
	GroupsRenameResponse struct {
		OK      bool     `json:"ok"`
		Error   string   `json:"error"`
		Channel *Channel `json:"channel"`
	}

	// https://api.slack.com/methods/groups.setPurpose
	// {
	//     "ok": true,
	//     "purpose": "This is the new purpose!"
	// }
	GroupsSetPurposeResponse struct {
		OK      bool   `json:"ok"`
		Error   string `json:"error"`
		Purpose string `json:"purpose"`
	}

	// https://api.slack.com/methods/groups.setTopic
	// {
	//     "ok": true,
	//     "topic": "This is the new topic!"
	// }
	GroupsSetTopicResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
		Topic string `json:"topic"`
	}

	// https://api.slack.com/methods/groups.unarchive
	// {
	//     "ok": true
	// }
	GroupsUnarchiveResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}

	// https://api.slack.com/methods/im.close
	// {
	//     "ok": true
	// }
	// {
	//     "ok": true,
	//     "no_op": true,
	//     "already_closed": true
	// }
	ImCloseResponse struct {
		OK            bool   `json:"ok"`
		Error         string `json:"error"`
		NOOP          bool   `json:"no_op"`
		AlreadyClosed bool   `json:"already_closed"`
	}

	// https://api.slack.com/methods/im.history
	// {
	//     "ok": true,
	//     "latest": "1358547726.000003",
	//     "messages": [
	//         {
	//             "type": "message",
	//             "ts": "1358546515.000008",
	//             "user": "U2147483896",
	//             "text": "Hello"
	//         },
	//         {
	//             "type": "message",
	//             "ts": "1358546515.000007",
	//             "user": "U2147483896",
	//             "text": "World",
	//             "is_starred": true,
	//         },
	//         {
	//             "type": "something_else",
	//             "ts": "1358546515.000007",
	//             "wibblr": true
	//         }
	//     ],
	//     "has_more": false
	// }
	ImHistoryResponse struct {
		OK       bool       `json:"ok"`
		Error    string     `json:"error"`
		Latest   Time       `json:"latest"`
		Messages []*Message `json:"group"`
		HasMore  bool       `json:"has_more"`
	}

	// https://api.slack.com/methods/im.list
	// {
	//     "ok": true,
	//     "ims": [
	//         {
	//            "id": "D024BFF1M",
	//            "is_im": true,
	//            "user": "USLACKBOT",
	//            "created": 1372105335,
	//            "is_user_deleted": false
	//         },
	//         {
	//            "id": "D024BE7RE",
	//            "is_im": true,
	//            "user": "U024BE7LH",
	//            "created": 1356250715,
	//            "is_user_deleted": false
	//         },
	//
	//     ]
	// }
	ImListResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
		IMs   []*IM  `json:"ims"`
	}

	// https://api.slack.com/methods/im.mark
	// {
	//     "ok": true
	// }
	ImMarkResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}

	// https://api.slack.com/methods/im.open
	// {
	//     "ok": true,
	//     "channel": {
	//         "id": "D024BFF1M"
	//     }
	// }
	// {
	//     "ok": true,
	//     "no_op": true,
	//     "already_open": true,
	//     "channel": {
	//         "id": "D024BFF1M"
	//     }
	// }
	ImOpenResponse struct {
		OK          bool     `json:"ok"`
		Error       string   `json:"error"`
		NOOP        bool     `json:"no_op"`
		AlreadyOpen bool     `json:"already_open"`
		Channel     *Channel `json:"channel"`
	}

	// https://api.slack.com/methods/mpim.close
	// {
	//     "ok": true
	// }
	// {
	//     "ok": true,
	//     "no_op": true,
	//     "already_closed": true
	// }
	MpimCloseResponse struct {
		OK            bool   `json:"ok"`
		Error         string `json:"error"`
		NOOP          bool   `json:"no_op"`
		AlreadyClosed bool   `json:"already_closed"`
	}

	// https://api.slack.com/methods/mpim.history
	// {
	//     "ok": true,
	//     "latest": "1358547726.000003",
	//     "messages": [
	//         {
	//             "type": "message",
	//             "ts": "1358546515.000008",
	//             "user": "U2147483896",
	//             "text": "Hello"
	//         },
	//         {
	//             "type": "message",
	//             "ts": "1358546515.000007",
	//             "user": "U2147483896",
	//             "text": "World",
	//             "is_starred": true,
	//         },
	//         {
	//             "type": "something_else",
	//             "ts": "1358546515.000007",
	//             "wibblr": true
	//         }
	//     ],
	//     "has_more": false
	// }
	MpimHistoryResponse struct {
		OK       bool       `json:"ok"`
		Error    string     `json:"error"`
		Latest   Time       `json:"latest"`
		Messages []*Message `json:"group"`
		HasMore  bool       `json:"has_more"`
	}

	// https://api.slack.com/methods/mpim.list
	// {
	//     "ok": true,
	//     "groups": [
	//         {
	//             "id": "G024BE91L",
	//             "name": "dm-messaging-user-1",
	//             "created": 1360782804,
	//             "creator": "U024BE7LH",
	//             "is_archived": false,
	//             "is_mpim": true
	//             "members": [
	//                 "U024BE7LH",
	//                 "U1234567890",
	//                 "U2345678901",
	//                 "U3456789012"
	//             ],
	//             "topic": {
	//                 "value": "Group messaging.",
	//                 "creator": "U024BE7LH",
	//                 "last_set": 1360782804
	//             },
	//             "purpose": {
	//                 "value": "Group messaging with: @user @user_a @user_b @user_c",
	//                 "creator": "U024BE7LH",
	//                 "last_set": 1360782804
	//             }
	//         }
	//     ]
	// }
	MpimListResponse struct {
		OK     bool     `json:"ok"`
		Error  string   `json:"error"`
		Groups []*Group `json:"groups"`
	}

	// https://api.slack.com/methods/mpim.mark
	// {
	//     "ok": true
	// }
	MpimMarkResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}

	// https://api.slack.com/methods/mpim.open
	// {
	//     "ok": true,
	//     "group": {
	//         "id": "G024BE91L",
	//         "name": "mpdm-user--user_a--user_b--user_c-1",
	//         "is_group": "true",
	//         "created": 1360782804,
	//         "creator": "U024BE7LH",
	//         "is_archived": false,
	//         "is_open": false,
	//         "is_mpim": true,
	//         "last_read": "0000000000.000000",
	//         "latest": null,
	//         "unread_count": 0,
	//         "unread_count_display": 0,
	//         "members": [
	//             "U024BE7LH",
	//             "U1234567890",
	//             "U2345678901",
	//             "U3456789012"
	//         ],
	//         "topic": {
	//             "value": "Group messaging.",
	//             "creator": "U024BE7LH",
	//             "last_set": 1360782804
	//         },
	//         "purpose": {
	//             "value": "Group messaging with: @user @user_a @user_b @user_c",
	//             "creator": "U024BE7LH",
	//             "last_set": 1360782804
	//         }
	//     }
	// }
	MpimOpenResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
		Group *Group `json:"group"`
	}

	// https://api.slack.com/methods/oauth.access
	// {
	//     "access_token": "xoxt-23984754863-2348975623103",
	//     "scope": "read"
	// }
	OauthAccessResponse struct {
		AccessToken string `json:"access_token"`
		Scope       string `json:"scope"`
	}

	// https://api.slack.com/methods/pins.add
	// {
	//     "ok": true
	// }
	PinsAddResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}

	// https://api.slack.com/methods/pins.list
	// {
	//     "ok": true,
	//     "items": [
	//         {
	//             "type": "message",
	//             "channel": "C2147483705",
	//             "message": {...},
	//         },
	//         {
	//             "type": "file",
	//             "file": { ... },
	//         }
	//         {
	//             "type": "file_comment",
	//             "file": { ... },
	//             "comment": { ... },
	//         }
	//     ]
	// }
	PinsListResponse struct {
		OK    bool    `json:"ok"`
		Error string  `json:"error"`
		Items []*Item `json:"items"`
	}

	// https://api.slack.com/methods/pins.remove
	// {
	//     "ok": true
	// }
	PinsRemoveResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}

	// https://api.slack.com/methods/reactions.add
	// {
	//     "ok": true
	// }
	ReactionsAddResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}

	// https://api.slack.com/methods/reactions.get
	// {
	//     "type": "message",
	//     "channel": "C2147483705",
	//     "message": {
	//         ...
	//         "reactions": [
	//             {
	//                 "name": "astonished",
	//                 "count": 3,
	//                 "users": [ "U1", "U2", "U3" ]
	//             },
	//             {
	//                 "name": "clock1"
	//                 "count": 2,
	//                 "users": [ "U1", "U2", "U3" ]
	//             }
	//         ]
	//     },
	// },
	ReactionsGetResponse struct {
		OK      bool     `json:"ok"`
		Error   string   `json:"error"`
		Type    string   `json:"type"`
		Channel string   `json:"channel"`
		Message *Message `json:"message"`
	}

	// https://api.slack.com/methods/reactions.list
	// {
	//     "ok": true,
	//     "items": [
	//         {
	//             "type": "message",
	//             "channel": "C2147483705",
	//             "message": {
	//                 ...
	//                 "reactions": [
	//                     {
	//                         "name": "astonished",
	//                         "count": 3,
	//                         "users": [ "U1", "U2", "U3" ]
	//                     },
	//                     {
	//                         "name": "clock1"
	//                         "count": 2,
	//                         "users": [ "U1", "U2", "U3" ]
	//                     }
	//                 ]
	//             },
	//         },
	//         {
	//             "type": "file",
	//             "file": { ... },
	//             "reactions": [
	//                 {
	//                     "name": "thumbsup",
	//                     "count": 1,
	//                     "users": [ "U1" ]
	//                 }
	//             ]
	//         }
	//         {
	//             "type": "file_comment",
	//             "file": { ... },
	//             "comment": { ... },
	//             "reactions": [
	//                 {
	//                     "name": "facepalm",
	//                     "count": 1034,
	//                     "users": [ "U1", "U2", "U3", "U4", "U5" ]
	//                 }
	//             ]
	//         }
	//     ],
	//     "paging": {
	//         "count": 100,
	//         "total": 4,
	//         "page": 1,
	//         "pages": 1
	//     }
	// }
	ReactionsListResponse struct {
		OK     bool    `json:"ok"`
		Error  string  `json:"error"`
		Items  []*Item `json:"items"`
		Paging struct {
			Count int `json:"count"`
			Page  int `json:"page"`
			Pages int `json:"pages"`
			Total int `json:"total"`
		} `json:"paging"`
	}

	// https://api.slack.com/methods/reactions.remove
	// {
	//     "ok": true
	// }
	ReactionsRemoveResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}

	// https://api.slack.com/methods/rtm.start
	// {
	//     "ok": true,
	//     "url": "wss:\/\/ms9.slack-msgs.com\/websocket\/7I5yBpcvk",
	//     "self": {
	//         "id": "U023BECGF",
	//         "name": "bobby",
	//         "prefs": {},
	//         "created": 1402463766,
	//         "manual_presence": "active"
	//     },
	//     "team": {
	//         "id": "T024BE7LD",
	//         "name": "Example Team",
	//         "email_domain": "",
	//         "domain": "example",
	//         "icon": {},
	//         "msg_edit_window_mins": -1,
	//         "over_storage_limit": false
	//         "prefs": {},
	//         "plan": "std"
	//     },
	//     "users": [ ],
	//     "channels": [ ],
	//     "groups": [ ],
	//     "mpims": [ ],
	//     "ims": [ ],
	//     "bots": [ ],
	// }
	RTMStartResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
		URL   string `json:"url"`
		// Self  struct {
		// 	ID             string                 `json:"id"`
		// 	Name           string                 `json:"name"`
		// 	Prefs          map[string]interface{} `json:"prefs"`
		// 	Created        Time                   `json:"created"`
		// 	ManualPresence string                 `json:"manual_presence"`
		// } `json:"self"`
		// Team     *Team      `json:"team"`
		// Users    []string   `json:"users"`
		// Channels []*Channel `json:"channels"`
		// Groups   []*Group   `json:"groups"`
		// Mpims    []*MPIM    `json:"mpims"`
		// Ims      []*IM      `json:"ims"`
		// Bots     []string   `json:"bots"`
	}

	// https://api.slack.com/methods/search.all
	// {
	//     "ok": true,
	//     "query": "Best Pickles",
	//     "messages": {...},
	//     "files": {...}
	// }
	// {
	//     "matches": [],
	//     "paging": {
	//         "count": 100,  - number of records per page
	//         "total": 15,   - total records matching query
	//         "page": 1,     - page of records returned
	//         "pages": 1     - total pages matching query
	//     }
	// }
	SearchAllResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
		Query string `json:"query"`
		Files struct {
			Matches []*File `json:"matches"`
			Paging  struct {
				Count int `json:"count"`
				Page  int `json:"page"`
				Pages int `json:"pages"`
				Total int `json:"total"`
			} `json:"paging"`
		} `json:"files"`
		Messages struct {
			Matches []*Message `json:"matches"`
			Paging  struct {
				Count int `json:"count"`
				Page  int `json:"page"`
				Pages int `json:"pages"`
				Total int `json:"total"`
			} `json:"paging"`
		} `json:"messages"`
	}

	// https://api.slack.com/methods/search.files
	// {
	//     "ok": true,
	//     "query": "test",
	//     "files": {
	//         "total": 829,
	//         "paging": {
	//             "count": 20,
	//             "total": 829,
	//             "page": 1,
	//             "pages": 42
	//         },
	//         "matches": [
	//             {...},
	//             {...},
	//             {...}
	//         ]
	//     }
	// }
	SearchFilesResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
		Query string `json:"query"`
		Files struct {
			Total  int `json:"total"`
			Paging struct {
				Count int `json:"count"`
				Page  int `json:"page"`
				Pages int `json:"pages"`
				Total int `json:"total"`
			} `json:"paging"`
			Matches []*File `json:"matches"`
		} `json:"files"`
	}

	// https://api.slack.com/methods/search.messages
	// {
	//     "ok": true,
	//     "query": "test",
	//     "messages": {
	//         "total": 829,
	//         "paging": {
	//             "count": 20,
	//             "total": 829,
	//             "page": 1,
	//             "pages": 42
	//         },
	//         "matches": [
	//             {...},
	//             {...},
	//             {...}
	//         ]
	//     }
	// }
	// {
	//     "type": "message",
	//     "channel": {
	//         "id": "C2147483753",
	//         "name": "foo"
	//     },
	//     "user": "U2147483709",
	//     "username": "johnnytest",
	//     "ts": "1359414002.000003",
	//     "text": "mention test: johnnyrodgers".
	//     "permalink": "https:\/\/example.slack.com\/channels\/foo\/p1359414002000003",
	//     "previous_2": {
	//         "user": "U2147483709",
	//         "username": "johnnytest",
	//         "text": "This was said before before",
	//         "ts": "1359413987.000000",
	//         "type": "message"
	//     },
	//     "previous": {
	//         "user": "U2147483709",
	//         "username": "johnnytest",
	//         "text": "This was said before",
	//         "ts": "1359414001.000000",
	//         "type": "message"
	//     },
	//     "next": {
	//         "user": "U2147483709",
	//         "username": "johnnytest",
	//         "text": "This was said after",
	//         "ts": "1359414020.000000",
	//         "type": "message"
	//     },
	//     "next_2": {
	//         "user": "U2147483709",
	//         "username": "johnnytest",
	//         "text": "This was said after after",
	//         "ts": "1359414021.000000",
	//         "type": "message"
	//     }
	// }
	SearchMessagesResponse struct {
		OK       bool   `json:"ok"`
		Error    string `json:"error"`
		Query    string `json:"query"`
		Messages struct {
			Total  int `json:"total"`
			Paging struct {
				Count int `json:"count"`
				Page  int `json:"page"`
				Pages int `json:"pages"`
				Total int `json:"total"`
			} `json:"paging"`
			Matches []*ResultMessage `json:"matches"`
		} `json:"messages"`
	}

	// https://api.slack.com/methods/stars.add
	// {
	//     "ok": true
	// }
	StarsAddResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}

	// https://api.slack.com/methods/stars.list
	// {
	//     "ok": true,
	//     "items": [
	//         {
	//             "type": "message",
	//             "channel": "C2147483705",
	//             "message": {...}
	//         },
	//         {
	//             "type": "file",
	//             "file": { ... }
	//         }
	//         {
	//             "type": "file_comment",
	//             "file": { ... },
	//             "comment": { ... }
	//         }
	//         {
	//             "type": "channel",
	//             "channel": "C2147483705"
	//         }
	//     ],
	//     "paging": {
	//         "count": 100,
	//         "total": 4,
	//         "page": 1,
	//         "pages": 1
	//     }
	// }
	StarsListResponse struct {
		OK     bool    `json:"ok"`
		Error  string  `json:"error"`
		Items  []*Item `json:"items"`
		Paging struct {
			Count int `json:"count"`
			Page  int `json:"page"`
			Pages int `json:"pages"`
			Total int `json:"total"`
		} `json:"paging"`
	}

	// https://api.slack.com/methods/stars.remove
	// {
	//     "ok": true
	// }
	StarsRemoveResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}

	// https://api.slack.com/methods/team.accessLogs
	// {
	//     "ok": true,
	//     "logins": [
	//         {
	//             "user_id": "U12345",
	//             "username": "bob",
	//             "date_first": 1422922864,
	//             "date_last": 1422922864,
	//             "count": 1,
	//             "ip": "127.0.0.1",
	//             "user_agent": "SlackWeb Mozilla\/5.0 (Macintosh; Intel Mac OS X 10_10_2) AppleWebKit\/537.36 (KHTML, like Gecko) Chrome\/41.0.2272.35 Safari\/537.36",
	//             "isp": "BigCo ISP",
	//             "country": "US",
	//             "region": "CA"
	//         },
	//         {
	//             "user_id": "U45678",
	//             "username": "alice",
	//             "date_first": 1422922493,
	//             "date_last": 1422922493,
	//             "count": 1,
	//             "ip": "127.0.0.1",
	//             "user_agent": "SlackWeb Mozilla\/5.0 (iPhone; CPU iPhone OS 8_1_3 like Mac OS X) AppleWebKit\/600.1.4 (KHTML, like Gecko) Version\/8.0 Mobile\/12B466 Safari\/600.1.4",
	//             "isp": "BigCo ISP",
	//             "country": "US",
	//             "region": "CA"
	//         },
	//     ],
	//     "paging": {
	//         "count": 100,
	//         "total": 2,
	//         "page": 1,
	//         "pages": 1
	//     }
	// }
	TeamAccessLogsResponse struct {
		OK     bool     `json:"ok"`
		Error  string   `json:"error"`
		Logins []*Login `json:"logins"`
		Paging struct {
			Count int `json:"count"`
			Page  int `json:"page"`
			Pages int `json:"pages"`
			Total int `json:"total"`
		} `json:"paging"`
	}

	// https://api.slack.com/methods/team.info
	//  {
	//      "ok": true,
	//      "team": {
	//          "id": "T12345",
	//          "name": "My Team",
	//          "domain": "example",
	//          "email_domain": "",
	//          "icon": {
	//              "image_34": "https:\/\/...",
	//              "image_44": "https:\/\/...",
	//              "image_68": "https:\/\/...",
	//              "image_88": "https:\/\/...",
	//              "image_102": "https:\/\/...",
	//              "image_132": "https:\/\/...",
	//              "image_default": true
	//          }
	//      }
	//  }
	TeamInfoResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
		Team  *Team  `json:"team"`
	}

	// https://api.slack.com/methods/team.integrationLogs
	// {
	//     "ok": true,
	//     "logs": [
	//         {
	//             "service_id": "1234567890",
	//             "service_type": "Google Calendar",
	//             "user_id": "U1234ABCD",
	//             "user_name": "Johnny",
	//             "channel": "C1234567890",
	//             "date": "1392163200",
	//             "change_type": "enabled",
	//             "scope": "incoming-webhook"
	//         },
	//         {
	//             "app_id": "2345678901",
	//             "app_type": "Johnny App"
	//             "user_id": "U2345BCDE",
	//             "user_name": "Billy",
	//             "date": "1392163201",
	//             "change_type": "added"
	//             "scope": "chat:write:user,channels:read"
	//         },
	//         {
	//             "service_id": "3456789012",
	//             "service_type": "Airbrake",
	//             "user_id": "U3456CDEF",
	//             "user_name": "Joey",
	//             "channel": "C1234567890",
	//             "date": "1392163202",
	//             "change_type": "disabled",
	//             "reason": "user",
	//             "scope": "incoming-webhook"
	//         },
	//         "paging": {
	//             "count": 3,
	//             "total": 3,
	//             "page": 1,
	//             "pages": 1
	//         }
	//     ]
	// }
	TeamIntegrationLogsResponse struct {
		OK     bool        `json:"ok"`
		Error  string      `json:"error"`
		Logs   []*LogEntry `json:"logins"`
		Paging struct {
			Count int `json:"count"`
			Page  int `json:"page"`
			Pages int `json:"pages"`
			Total int `json:"total"`
		} `json:"paging"`
	}

	// https://api.slack.com/methods/usergroups.create
	// {
	//     "ok": true,
	//     "usergroup": {
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
	//             "channels": [],
	//             "groups": []
	//         },
	//         "user_count": "0"
	//     }
	// }
	UsergroupsCreateResponse struct {
		OK        bool       `json:"ok"`
		Error     string     `json:"error"`
		UserGroup *UserGroup `json:"user_group"`
	}

	// https://api.slack.com/methods/usergroups.disable
	// {
	//     "ok": true,
	//     "usergroup": {
	//         "id": "S0615G0KT",
	//         "team_id": "T060RNRCH",
	//         "is_usergroup": true,
	//         "name": "Marketing Team",
	//         "description": "Marketing gurus, PR experts and product advocates.",
	//         "handle": "marketing-team",
	//         "is_external": false,
	//         "date_create": 1446746793,
	//         "date_update": 1446747568,
	//         "date_delete": 1446747568,
	//         "auto_type": null,
	//         "created_by": "U060RNRCZ",
	//         "updated_by": "U060RNRCZ",
	//         "deleted_by": "U060RNRCZ",
	//         "prefs": {
	//             "channels": [],
	//             "groups": []
	//         },
	//         "user_count": "0"
	//     }
	// }
	UsergroupsDisableResponse struct {
		OK        bool       `json:"ok"`
		Error     string     `json:"error"`
		UserGroup *UserGroup `json:"user_group"`
	}

	// https://api.slack.com/methods/usergroups.enable
	// {
	//     "ok": true,
	//     "usergroup": {
	//         "id": "S0615G0KT",
	//         "team_id": "T060RNRCH",
	//         "is_usergroup": true,
	//         "name": "Marketing Team",
	//         "description": "Marketing gurus, PR experts and product advocates.",
	//         "handle": "marketing-team",
	//         "is_external": false,
	//         "date_create": 1446746793,
	//         "date_update": 1446747767,
	//         "date_delete": 0,
	//         "auto_type": null,
	//         "created_by": "U060RNRCZ",
	//         "updated_by": "U060RNRCZ",
	//         "deleted_by": null,
	//         "prefs": {
	//             "channels": [],
	//             "groups": []
	//         },
	//         "user_count": "0"
	//     }
	// }
	UsergroupsEnableResponse struct {
		OK        bool       `json:"ok"`
		Error     string     `json:"error"`
		UserGroup *UserGroup `json:"user_group"`
	}

	// https://api.slack.com/methods/usergroups.list
	// {
	//     "ok": true,
	//     "usergroups": [
	//         {
	//             "id": "S0614TZR7",
	//             "team_id": "T060RNRCH",
	//             "is_usergroup": true,
	//             "name": "Team Admins",
	//             "description": "A group of all Administrators on your team.",
	//             "handle": "admins",
	//             "is_external": false,
	//             "date_create": 1446598059,
	//             "date_update": 1446670362,
	//             "date_delete": 0,
	//             "auto_type": "admin",
	//             "created_by": "USLACKBOT",
	//             "updated_by": "U060RNRCZ",
	//             "deleted_by": null,
	//             "prefs": {
	//                 "channels": [],
	//                 "groups": []
	//             },
	//             "user_count": "2"
	//         },
	//         {
	//             "id": "S06158AV7",
	//             "team_id": "T060RNRCH",
	//             "is_usergroup": true,
	//             "name": "Team Owners",
	//             "description": "A group of all Owners on your team.",
	//             "handle": "owners",
	//             "is_external": false,
	//             "date_create": 1446678371,
	//             "date_update": 1446678371,
	//             "date_delete": 0,
	//             "auto_type": "owner",
	//             "created_by": "USLACKBOT",
	//             "updated_by": "USLACKBOT",
	//             "deleted_by": null,
	//             "prefs": {
	//                 "channels": [],
	//                 "groups": []
	//             },
	//             "user_count": "1"
	//         },
	//         {
	//             "id": "S0615G0KT",
	//             "team_id": "T060RNRCH",
	//             "is_usergroup": true,
	//             "name": "Marketing Team",
	//             "description": "Marketing gurus, PR experts and product advocates.",
	//             "handle": "marketing-team",
	//             "is_external": false,
	//             "date_create": 1446746793,
	//             "date_update": 1446747767,
	//             "date_delete": 1446748865,
	//             "auto_type": null,
	//             "created_by": "U060RNRCZ",
	//             "updated_by": "U060RNRCZ",
	//             "deleted_by": null,
	//             "prefs": {
	//                 "channels": [],
	//                 "groups": []
	//             },
	//             "user_count": "0"
	//         }
	//     ]
	// }
	UsergroupsListResponse struct {
		OK         bool         `json:"ok"`
		Error      string       `json:"error"`
		UserGroups []*UserGroup `json:"user_groups"`
	}

	// https://api.slack.com/methods/usergroups.update
	// {
	//     "ok": true,
	//     "usergroup": {
	//         "id": "S0615G0KT",
	//         "team_id": "T060RNRCH",
	//         "is_usergroup": true,
	//         "name": "Marketing Gurus",
	//         "description": "Marketing gurus, PR experts and product advocates.",
	//         "handle": "marketing-team",
	//         "is_external": false,
	//         "date_create": 1446746793,
	//         "date_update": 1446748574,
	//         "date_delete": 0,
	//         "auto_type": null,
	//         "created_by": "U060RNRCZ",
	//         "updated_by": "U060RNRCZ",
	//         "deleted_by": null,
	//         "prefs": {
	//             "channels": [],
	//             "groups": []
	//         },
	//         "user_count": "0"
	//     }
	// }
	UsergroupsUpdateResponse struct {
		OK        bool       `json:"ok"`
		Error     string     `json:"error"`
		UserGroup *UserGroup `json:"user_group"`
	}

	// https://api.slack.com/methods/usergroups.users.list
	// {
	//     "ok": true,
	//     "users": [
	//         "U060R4BJ4"
	//     ]
	// }
	UsergroupsUsersListResponse struct {
		OK    bool     `json:"ok"`
		Users []string `json:"users"`
	}

	// https://api.slack.com/methods/usergroups.users.update
	// {
	//     "ok": true,
	//     "usergroup": {
	//         "id": "S0616NG6M",
	//         "team_id": "T060R4BHN",
	//         "is_usergroup": true,
	//         "name": "Marketing Team",
	//         "description": "Marketing gurus, PR experts and product advocates.",
	//         "handle": "marketing-team",
	//         "is_external": false,
	//         "date_create": 1447096577,
	//         "date_update": 1447102109,
	//         "date_delete": 0,
	//         "auto_type": null,
	//         "created_by": "U060R4BJ4",
	//         "updated_by": "U060R4BJ4",
	//         "deleted_by": null,
	//         "prefs": {
	//             "channels": [],
	//             "groups": []
	//         },
	//         "users": [
	//             "U060R4BJ4",
	//             "U060RNRCZ"
	//         ],
	//         "user_count": 1
	//     }
	// }
	UsergroupsUsersUpdateResponse struct {
		OK        bool       `json:"ok"`
		Error     string     `json:"error"`
		UserGroup *UserGroup `json:"user_group"`
	}

	// https://api.slack.com/methods/users.getPresence
	// {
	//     "ok": true,
	//     "presence": "active",
	//     "online": true,
	//     "auto_away": false,
	//     "manual_away": false,
	//     "connection_count": 1,
	//     "last_activity": 1419027078
	// }
	UsersGetPresenceResponse struct {
		OK              bool     `json:"ok"`
		Error           string   `json:"error"`
		Presence        Presence `json:"presence"`
		Online          bool     `json:"online"`
		AutoAway        bool     `json:"auto_away"`
		ManualAway      bool     `json:"manual_away"`
		ConnectionCount int      `json:"connection_count"`
		LastActivity    Time     `json:"last_activity"`
	}

	// https://api.slack.com/methods/users.info
	// {
	//     "ok": true,
	//     "user": {
	//         "id": "U023BECGF",
	//         "name": "bobby",
	//         "deleted": false,
	//         "color": "9f69e7",
	//         "profile": {
	//             "first_name": "Bobby",
	//             "last_name": "Tables",
	//             "real_name": "Bobby Tables",
	//             "email": "bobby@slack.com",
	//             "skype": "my-skype-name",
	//             "phone": "+1 (123) 456 7890",
	//             "image_24": "https:\/\/...",
	//             "image_32": "https:\/\/...",
	//             "image_48": "https:\/\/...",
	//             "image_72": "https:\/\/...",
	//             "image_192": "https:\/\/..."
	//         },
	//         "is_admin": true,
	//         "is_owner": true,
	//         "has_2fa": true,
	//         "has_files": true
	//     }
	// }
	UsersInfoResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
		User  *User  `json:"user"`
	}

	// https://api.slack.com/methods/users.list
	// {
	//     "ok": true,
	//     "members": [
	//         {
	//             "id": "U023BECGF",
	//             "name": "bobby",
	//             "deleted": false,
	//             "color": "9f69e7",
	//             "profile": {
	//                 "first_name": "Bobby",
	//                 "last_name": "Tables",
	//                 "real_name": "Bobby Tables",
	//                 "email": "bobby@slack.com",
	//                 "skype": "my-skype-name",
	//                 "phone": "+1 (123) 456 7890",
	//                 "image_24": "https:\/\/...",
	//                 "image_32": "https:\/\/...",
	//                 "image_48": "https:\/\/...",
	//                 "image_72": "https:\/\/...",
	//                 "image_192": "https:\/\/..."
	//             },
	//             "is_admin": true,
	//             "is_owner": true,
	//             "has_2fa": false,
	//             "has_files": true
	//         }
	//     ]
	// }
	UsersListResponse struct {
		OK      bool    `json:"ok"`
		Error   string  `json:"error"`
		Members []*User `json:"members"`
	}

	// https://api.slack.com/methods/users.setActive
	// {
	//     "ok": true
	// }
	UsersSetActiveResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}

	// https://api.slack.com/methods/users.setPresence
	// {
	//     "ok": true
	// }
	UsersSetPresenceResponse struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}
)
