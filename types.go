package slk

type (
	// https://api.slack.com/types/channel
	// {
	//     "id": "C024BE91L",
	//     "name": "fun",
	//     "is_channel": "true",
	//     "created": 1360782804,
	//     "creator": "U024BE7LH",
	//     "is_archived": false,
	//     "is_general": false,
	//     "members": [
	//         "U024BE7LH",
	//     ],
	//     "topic": {
	//         "value": "Fun times",
	//         "creator": "U024BE7LV",
	//         "last_set": 1369677212
	//     },
	//     "purpose": {
	//         "value": "This channel is for fun",
	//         "creator": "U024BE7LH",
	//         "last_set": 1360782804
	//     }
	//     "is_member": true,
	//     "last_read": "1401383885.000061",
	//     "latest": { … }
	//     "unread_count": 0,
	//     "unread_count_display": 0
	// }
	Channel struct {
		ID         string   `json:"id"`
		Name       string   `json:"name"`
		IsChannel  bool     `json:"is_channel"`
		Created    Time     `json:"created"`
		Creator    string   `json:"creator"`
		IsArchived bool     `json:"is_archived"`
		IsGeneral  bool     `json:"is_general"`
		Members    []string `json:"members"`
		Topic      struct {
			Creator string `json:"creator"`
			LastSet Time   `json:"last_set"`
			Value   string `json:"value"`
		} `json:"topic"`
		Purpose struct {
			Creator string `json:"creator"`
			LastSet Time   `json:"last_set"`
			Value   string `json:"value"`
		} `json:"purpose"`
		IsMember           bool     `json:"is_member"`
		LastRead           Time     `json:"last_read"`
		LatestMessage      *Message `json:"latest"`
		UnreadCount        int      `json:"unread_count"`
		UnreadCountDisplay int      `json:"unread_count_display"`
	}

	// https://api.slack.com/types/file
	// {
	//     "id" : "F2147483862",
	//     "created" : 1356032811,
	//     "timestamp" : 1356032811,
	//     "name" : "file.htm",
	//     "title" : "My HTML file",
	//     "mimetype" : "text\/plain",
	//     "filetype" : "text",
	//     "pretty_type": "Text",
	//     "user" : "U2147483697",
	//     "mode" : "hosted",
	//     "editable" : true,
	//     "is_external": false,
	//     "external_type": "",
	//     "username": "",
	//     "size" : 12345,
	//     "url_private": "https:\/\/slack.com\/files-pri\/T024BE7LD-F024BERPE\/1.png",
	//     "url_private_download": "https:\/\/slack.com\/files-pri\/T024BE7LD-F024BERPE\/download\/1.png",
	//     "thumb_64": "https:\/\/slack-files.com\/files-tmb\/T024BE7LD-F024BERPE-c66246\/1_64.png",
	//     "thumb_80": "https:\/\/slack-files.com\/files-tmb\/T024BE7LD-F024BERPE-c66246\/1_80.png",
	//     "thumb_360": "https:\/\/slack-files.com\/files-tmb\/T024BE7LD-F024BERPE-c66246\/1_360.png",
	//     "thumb_360_gif": "https:\/\/slack-files.com\/files-tmb\/T024BE7LD-F024BERPE-c66246\/1_360.gif",
	//     "thumb_360_w": 100,
	//     "thumb_360_h": 100,
	//     "thumb_480": "https:\/\/slack-files.com\/files-tmb\/T024BE7LD-F024BERPE-c66246\/1_480.png",
	//     "thumb_480_w": 480,
	//     "thumb_480_h": 480,
	//     "thumb_160": "https:\/\/slack-files.com\/files-tmb\/T024BE7LD-F024BERPE-c66246\/1_160.png",
	//     "permalink" : "https:\/\/tinyspeck.slack.com\/files\/cal\/F024BERPE\/1.png",
	//     "permalink_public" : "https:\/\/tinyspeck.slack.com\/T024BE7LD-F024BERPE-3f9216b62c",
	//     "edit_link" : "https:\/\/tinyspeck.slack.com\/files\/cal\/F024BERPE\/1.png/edit",
	//     "preview" : "&lt;!DOCTYPE html&gt;\n&lt;html&gt;\n&lt;meta charset='utf-8'&gt;",
	//     "preview_highlight" : &lt;div class=\"sssh-code\"&gt;&lt;div class=\"sssh-line\"&gt;&lt;pre&gt;&lt;!DOCTYPE html...",
	//     "lines" : 123,
	//     "lines_more": 118,
	//     "is_public": true,
	//     "public_url_shared": false,
	//     "display_as_bot" : false,
	//     "channels": ["C024BE7LT", ...],
	//     "groups": ["G12345", ...],
	//     "ims": ["D12345", ...],
	//     "initial_comment": {...},
	//     "num_stars": 7,
	//     "is_starred": true,
	//     "pinned_to": ["C024BE7LT", ...],
	//     "reactions": [
	//     {
	//         "name": "astonished",
	//         "count": 3,
	//         "users": [ "U1", "U2", "U3" ]
	//     },
	//     {
	//         "name": "facepalm",
	//         "count": 1034,
	//         "users": [ "U1", "U2", "U3", "U4", "U5" ]
	//     }
	//     ],
	//     "comments_count": 1
	// }
	File struct {
		ID                 string   `json:"id"`
		Created            Time     `json:"created"`
		Timestamp          Time     `json:"timestamp"`
		Name               string   `json:"name"`
		Title              string   `json:"title"`
		Mimetype           string   `json:"mimetype"`
		Filetype           string   `json:"filetype"`
		PrettyType         string   `json:"pretty_type"`
		User               string   `json:"user"`
		Mode               string   `json:"mode"`
		Editable           bool     `json:"editable"`
		IsExternal         bool     `json:"is_external"`
		ExternalType       string   `json:"external_type"`
		Username           string   `json:"username"`
		Size               int      `json:"size"`
		URLPrivate         string   `json:"url_private"`
		URLPrivateDownload string   `json:"url_private_download"`
		Thumb160           string   `json:"thumb_160"`
		Thumb360           string   `json:"thumb_360"`
		Thumb360Gif        string   `json:"thumb_360_gif"`
		Thumb360H          int      `json:"thumb_360_h"`
		Thumb360W          int      `json:"thumb_360_w"`
		Thumb480           string   `json:"thumb_480"`
		Thumb480H          int      `json:"thumb_480_h"`
		Thumb480W          int      `json:"thumb_480_w"`
		Thumb64            string   `json:"thumb_64"`
		Thumb80            string   `json:"thumb_80"`
		Permalink          string   `json:"permalink"`
		PermalinkPublic    string   `json:"permalink_public"`
		EditLink           string   `json:"edit_link"`
		Preview            string   `json:"preview"`
		PreviewHighlight   string   `json:"preview_highlight"`
		Lines              int      `json:"lines"`
		LinesMore          int      `json:"lines_more"`
		IsPublic           bool     `json:"is_public"`
		PublicURLShared    bool     `json:"public_url_shared"`
		DisplayAsBot       bool     `json:"display_as_bot"`
		Channels           []string `json:"channels"`
		Groups             []string `json:"groups"`
		Ims                []string `json:"ims"`
		InitialComment     *Comment `json:"initial_comment"`
		NumStars           int      `json:"num_stars"`
		IsStarred          bool     `json:"is_starred"`
		PinnedTo           []string `json:"pinned_to"`
		Reactions          []struct {
			Count int      `json:"count"`
			Name  string   `json:"name"`
			Users []string `json:"users"`
		} `json:"reactions"`
		CommentsCount int `json:"comments_count"`
	}

	// https://api.slack.com/types/group
	// {
	//     "id": "G024BE91L",
	//     "name": "secretplans",
	//     "is_group": "true",
	//     "created": 1360782804,
	//     "creator": "U024BE7LH",
	//     "is_archived": false,
	//     "is_mpim": false,
	//     "members": [
	//         "U024BE7LH"
	//     ],
	//     "topic": {
	//         "value": "Secret plans on hold",
	//         "creator": "U024BE7LV",
	//         "last_set": 1369677212
	//     },
	//     "purpose": {
	//         "value": "Discuss secret plans that no-one else should know",
	//         "creator": "U024BE7LH",
	//         "last_set": 1360782804
	//     },
	//     "last_read": "1401383885.000061",
	//     "latest": { … }
	//     "unread_count": 0,
	//     "unread_count_display": 0
	// }
	Group struct {
		ID         string   `json:"id"`
		Name       string   `json:"name"`
		IsGroup    bool     `json:"is_group"`
		Created    Time     `json:"created"`
		Creator    string   `json:"creator"`
		IsArchived bool     `json:"is_archived"`
		IsMpim     bool     `json:"is_mpim"`
		Members    []string `json:"members"`
		Topic      struct {
			Creator string `json:"creator"`
			LastSet Time   `json:"last_set"`
			Value   string `json:"value"`
		} `json:"topic"`
		Purpose struct {
			Creator string `json:"creator"`
			LastSet Time   `json:"last_set"`
			Value   string `json:"value"`
		} `json:"purpose"`
		LastRead           Time     `json:"last_read"`
		Latest             *Message `json:"latest"`
		UnreadCount        int      `json:"unread_count"`
		UnreadCountDisplay int      `json:"unread_count_display"`
	}

	// https://api.slack.com/types/im
	// {
	//     "id": "D024BFF1M",
	//     "is_im": true,
	//     "user": "U024BE7LH",
	//     "created": 1360782804,
	//     "is_user_deleted": false
	// }
	IM struct {
		ID            string `json:"id"`
		IsIm          bool   `json:"is_im"`
		User          string `json:"user"`
		Created       Time   `json:"created"`
		IsUserDeleted bool   `json:"is_user_deleted"`
	}

	// https://api.slack.com/types/mpim
	// {
	//     "id": "G024BE91L",
	//     "name": "mpdm-user1--user2--user3-1",
	//     "is_mpim": true,
	//     "is_group": false,
	//     "created": 1360782804,
	//     "creator": "U024BE7LH",
	//     "members": [
	//         "U024BE7LH"
	//     ],
	//     "last_read": "1401383885.000061",
	//     "latest": { … }
	//     "unread_count": 0,
	//     "unread_count_display": 0
	// }
	MPIM struct {
		ID                 string   `json:"id"`
		Name               string   `json:"name"`
		IsMpim             bool     `json:"is_mpim"`
		IsGroup            bool     `json:"is_group"`
		Created            Time     `json:"created"`
		Creator            string   `json:"creator"`
		Members            []string `json:"members"`
		LastRead           Time     `json:"last_read"`
		Latest             *Message `json:"latest"`
		UnreadCount        int      `json:"unread_count"`
		UnreadCountDisplay int      `json:"unread_count_display"`
	}

	// https://api.slack.com/types/user
	// {
	//     "id": "U023BECGF",
	//     "name": "bobby",
	//     "deleted": false,
	//     "color": "9f69e7",
	//     "profile": {
	//         "first_name": "Bobby",
	//         "last_name": "Tables",
	//         "real_name": "Bobby Tables",
	//         "email": "bobby@slack.com",
	//         "skype": "my-skype-name",
	//         "phone": "+1 (123) 456 7890",
	//         "image_24": "https:\/\/...",
	//         "image_32": "https:\/\/...",
	//         "image_48": "https:\/\/...",
	//         "image_72": "https:\/\/...",
	//         "image_192": "https:\/\/..."
	//     },
	//     "is_admin": true,
	//     "is_owner": true,
	//     "is_primary_owner": true,
	//     "is_restricted": false,
	//     "is_ultra_restricted": false,
	//     "has_2fa": false,
	//     "two_factor_type": 'sms',
	//     "has_files": true
	// }
	User struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Deleted bool   `json:"deleted"`
		Color   string `json:"color"`
		Profile struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			RealName  string `json:"real_name"`
			Email     string `json:"email"`
			Phone     string `json:"phone"`
			Skype     string `json:"skype"`
			Image24   string `json:"image_24"`
			Image32   string `json:"image_32"`
			Image48   string `json:"image_48"`
			Image72   string `json:"image_72"`
			Image192  string `json:"image_192"`
		} `json:"profile"`
		IsAdmin           bool   `json:"is_admin"`
		IsOwner           bool   `json:"is_owner"`
		IsPrimaryOwner    bool   `json:"is_primary_owner"`
		IsRestricted      bool   `json:"is_restricted"`
		IsUltraRestricted bool   `json:"is_ultra_restricted"`
		Has2fa            bool   `json:"has_2fa"`
		TwoFactorType     string `json:"two_factor_type"`
		HasFiles          bool   `json:"has_files"`
	}

	// https://api.slack.com/types/usergroup
	// {
	//     "id": "S0614TZR7",
	//     "team_id": "T060RNRCH",
	//     "is_usergroup": true,
	//     "name": "Team Admins",
	//     "description": "A group of all Administrators on your team.",
	//     "handle": "admins",
	//     "is_external": false,
	//     "date_create": 1446598059,
	//     "date_update": 1446670362,
	//     "date_delete": 0,
	//     "auto_type": "admin",
	//     "created_by": "USLACKBOT",
	//     "updated_by": "U060RNRCZ",
	//     "deleted_by": null,
	//     "prefs": {
	//         "channels": [],
	//         "groups": []
	//     },
	//     "users": [
	//         "U060RNRCZ",
	//         "U060ULRC0",
	//         "U06129G2V",
	//         "U061309JM"
	//     ],
	//     "user_count": "4"
	// }
	UserGroup struct {
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
		Users       []string               `json:"users"`
		UserCount   string                 `json:"user_count"`
	}

	// https://api.slack.com/docs/attachments
	// {
	//     "fallback": "Required plain-text summary of the attachment.",
	//     "color": "#36a64f",
	//     "pretext": "Optional text that appears above the attachment block",
	//     "author_name": "Bobby Tables",
	//     "author_link": "http://flickr.com/bobby/",
	//     "author_icon": "http://flickr.com/icons/bobby.jpg",
	//     "title": "Slack API Documentation",
	//     "title_link": "https://api.slack.com/",
	//     "text": "Optional text that appears within the attachment",
	//     "fields": [
	//         {
	//             "title": "Priority",
	//             "value": "High",
	//             "short": false
	//         }
	//     ],
	//     "image_url": "http://my-website.com/path/to/image.jpg",
	//     "thumb_url": "http://example.com/path/to/thumb.png",
	//     "footer": "Slack API",
	//     "footer_icon": "https://platform.slack-edge.com/img/default_application_icon.png",
	//     "ts": 123456789
	// }
	MessageAttachment struct {
		AuthorIcon     string                     `json:"author_icon,omitempty"`
		AuthorLink     string                     `json:"author_link,omitempty"`
		AuthorName     string                     `json:"author_name,omitempty"`
		Color          string                     `json:"color,omitempty"`
		Fallback       string                     `json:"fallback,omitempty"`
		Fields         []*MessageAttachmentField  `json:"fields,omitempty"`
		Footer         string                     `json:"footer,omitempty"`
		FooterIcon     string                     `json:"footer_icon,omitempty"`
		ImageURL       string                     `json:"image_url,omitempty"`
		MarkdownIn     []string                   `json:"mrkdwn_in,omitempty"`
		Pretext        string                     `json:"pretext,omitempty"`
		Text           string                     `json:"text,omitempty"`
		ThumbURL       string                     `json:"thumb_url,omitempty"`
		Title          string                     `json:"title,omitempty"`
		TitleLink      string                     `json:"title_link,omitempty"`
		Timestamp      int                        `json:"ts"`
		CallbackId     string                     `json:"callback_id"`
		AttachmentType string                     `json:"attachment_type"`
		Actions        []*MessageAttachmentAction `json:"actions"`
	}

	MessageAttachmentField struct {
		Short bool   `json:"short"`
		Title string `json:"title,omitempty"`
		Value string `json:"value,omitempty"`
	}

	// adding for button support
	MessageAttachmentAction struct {
		Name    string `json:"name,omitempty"`
		Text    string `json:"text,omitempty"`
		Style   string `json:"style,omitempty"`
		Type    string `json:"type,omitempty"`
		Value   string `json:"value,omitempty"`
		Confirm *MessageAttachmentActionConfirm
	}

	MessageAttachmentActionConfirm struct {
		Title       string `json:"title,omitempty"`
		Text        string `json:"text,omitempty"`
		OkText      string `json:"ok_text,omitempty"`
		DismissText string `json:"dismiss_text,omitempty"`
	}
)
