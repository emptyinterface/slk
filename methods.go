package slk

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type (
	Method string
)

const (
	ApiTestMethod               Method = `api.test`
	AuthTestMethod              Method = `auth.test`
	ChannelsArchiveMethod       Method = `channels.archive`
	ChannelsCreateMethod        Method = `channels.create`
	ChannelsHistoryMethod       Method = `channels.history`
	ChannelsInfoMethod          Method = `channels.info`
	ChannelsInviteMethod        Method = `channels.invite`
	ChannelsJoinMethod          Method = `channels.join`
	ChannelsKickMethod          Method = `channels.kick`
	ChannelsLeaveMethod         Method = `channels.leave`
	ChannelsListMethod          Method = `channels.list`
	ChannelsMarkMethod          Method = `channels.mark`
	ChannelsRenameMethod        Method = `channels.rename`
	ChannelsSetPurposeMethod    Method = `channels.setPurpose`
	ChannelsSetTopicMethod      Method = `channels.setTopic`
	ChannelsUnarchiveMethod     Method = `channels.unarchive`
	ChatDeleteMethod            Method = `chat.delete`
	ChatPostMessageMethod       Method = `chat.postMessage`
	ChatUpdateMethod            Method = `chat.update`
	DndEndDndMethod             Method = `dnd.endDnd`
	DndEndSnoozeMethod          Method = `dnd.endSnooze`
	DndInfoMethod               Method = `dnd.info`
	DndSetSnoozeMethod          Method = `dnd.setSnooze`
	DndTeamInfoMethod           Method = `dnd.teamInfo`
	EmojiListMethod             Method = `emoji.list`
	FilesCommentsAddMethod      Method = `files.comments.add`
	FilesCommentsDeleteMethod   Method = `files.comments.delete`
	FilesCommentsEditMethod     Method = `files.comments.edit`
	FilesDeleteMethod           Method = `files.delete`
	FilesInfoMethod             Method = `files.info`
	FilesListMethod             Method = `files.list`
	FilesUploadMethod           Method = `files.upload`
	GroupsArchiveMethod         Method = `groups.archive`
	GroupsCloseMethod           Method = `groups.close`
	GroupsCreateMethod          Method = `groups.create`
	GroupsCreateChildMethod     Method = `groups.createChild`
	GroupsHistoryMethod         Method = `groups.history`
	GroupsInfoMethod            Method = `groups.info`
	GroupsInviteMethod          Method = `groups.invite`
	GroupsKickMethod            Method = `groups.kick`
	GroupsLeaveMethod           Method = `groups.leave`
	GroupsListMethod            Method = `groups.list`
	GroupsMarkMethod            Method = `groups.mark`
	GroupsOpenMethod            Method = `groups.open`
	GroupsRenameMethod          Method = `groups.rename`
	GroupsSetPurposeMethod      Method = `groups.setPurpose`
	GroupsSetTopicMethod        Method = `groups.setTopic`
	GroupsUnarchiveMethod       Method = `groups.unarchive`
	ImCloseMethod               Method = `im.close`
	ImHistoryMethod             Method = `im.history`
	ImListMethod                Method = `im.list`
	ImMarkMethod                Method = `im.mark`
	ImOpenMethod                Method = `im.open`
	MpimCloseMethod             Method = `mpim.close`
	MpimHistoryMethod           Method = `mpim.history`
	MpimListMethod              Method = `mpim.list`
	MpimMarkMethod              Method = `mpim.mark`
	MpimOpenMethod              Method = `mpim.open`
	OauthAccessMethod           Method = `oauth.access`
	PinsAddMethod               Method = `pins.add`
	PinsListMethod              Method = `pins.list`
	PinsRemoveMethod            Method = `pins.remove`
	ReactionsAddMethod          Method = `reactions.add`
	ReactionsGetMethod          Method = `reactions.get`
	ReactionsListMethod         Method = `reactions.list`
	ReactionsRemoveMethod       Method = `reactions.remove`
	RTMStartMethod              Method = `rtm.start`
	SearchAllMethod             Method = `search.all`
	SearchFilesMethod           Method = `search.files`
	SearchMessagesMethod        Method = `search.messages`
	StarsAddMethod              Method = `stars.add`
	StarsListMethod             Method = `stars.list`
	StarsRemoveMethod           Method = `stars.remove`
	TeamAccessLogsMethod        Method = `team.accessLogs`
	TeamInfoMethod              Method = `team.info`
	TeamIntegrationLogsMethod   Method = `team.integrationLogs`
	UsergroupsCreateMethod      Method = `usergroups.create`
	UsergroupsDisableMethod     Method = `usergroups.disable`
	UsergroupsEnableMethod      Method = `usergroups.enable`
	UsergroupsListMethod        Method = `usergroups.list`
	UsergroupsUpdateMethod      Method = `usergroups.update`
	UsergroupsUsersListMethod   Method = `usergroups.users.list`
	UsergroupsUsersUpdateMethod Method = `usergroups.users.update`
	UsersGetPresenceMethod      Method = `users.getPresence`
	UsersInfoMethod             Method = `users.info`
	UsersListMethod             Method = `users.list`
	UsersSetActiveMethod        Method = `users.setActive`
	UsersSetPresenceMethod      Method = `users.setPresence`
)

// https://api.slack.com/methods/api.test
func (c *Client) ApiTest(optional url.Values) (*ApiTestResponse, error) {
	r := &ApiTestResponse{}
	err := c.Do(ApiTestMethod, optional, r)
	return r, err
}

// https://api.slack.com/methods/auth.test
func (c *Client) AuthTest() (*AuthTestResponse, error) {
	r := &AuthTestResponse{}
	err := c.Do(AuthTestMethod, nil, r)
	return r, err
}

// https://api.slack.com/methods/channels.archive
// req:channel // C1234567890 // Required // Channel to archive
func (c *Client) ChannelsArchive(channel string) (*ChannelsArchiveResponse, error) {
	r := &ChannelsArchiveResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	err := c.Do(ChannelsArchiveMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/channels.create
// req:name // mychannel // Required // Name of channel to create
func (c *Client) ChannelsCreate(name string) (*ChannelsCreateResponse, error) {
	r := &ChannelsCreateResponse{}
	args := url.Values{}
	args.Set("name", name)
	err := c.Do(ChannelsCreateMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/channels.history
// req:channel // C1234567890 // Required // Channel to fetch history for.
// opt:latest // 1234567890.123456 // Optional // , default=now End of time range of messages to include in results.
// opt:oldest // 1234567890.123456 // Optional // , default=0 Start of time range of messages to include in results.
// opt:inclusive // 1 // Optional // , default=0 Include messages with latest or oldest timestamp in results.
// opt:count // 100 // Optional // , default=100 Number of messages to return, between 1 and 1000.
// opt:unreads // 1 // Optional // , default=0 Include unread_count_display in the output?
func (c *Client) ChannelsHistory(channel string, optional url.Values) (*ChannelsHistoryResponse, error) {
	r := &ChannelsHistoryResponse{}
	var args url.Values
	if optional != nil {
		args = optional
	} else {
		args = url.Values{}
	}
	args.Set("channel", channel)
	err := c.Do(ChannelsHistoryMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/channels.info
// req:channel // C1234567890 // Required // Channel to get info on
func (c *Client) ChannelsInfo(channel string) (*ChannelsInfoResponse, error) {
	r := &ChannelsInfoResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	err := c.Do(ChannelsInfoMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/channels.invite
// req:channel // C1234567890 // Required // Channel to invite user to.
// req:user // U1234567890 // Required // User to invite to channel.
func (c *Client) ChannelsInvite(channel, user string) (*ChannelsInviteResponse, error) {
	r := &ChannelsInviteResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	args.Set("user", user)
	err := c.Do(ChannelsInviteMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/channels.join
// req:name // C024BE91L // Required // Name of channel to join
func (c *Client) ChannelsJoin(name string) (*ChannelsJoinResponse, error) {
	r := &ChannelsJoinResponse{}
	args := url.Values{}
	args.Set("name", name)
	err := c.Do(ChannelsJoinMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/channels.kick
// req:channel // C1234567890 // Required // Channel to remove user from.
// req:user // U1234567890 // Required // User to remove from channel.
func (c *Client) ChannelsKick(channel, user string) (*ChannelsKickResponse, error) {
	r := &ChannelsKickResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	args.Set("user", user)
	err := c.Do(ChannelsKickMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/channels.leave
// req:channel // C1234567890 // Required // Channel to leave
func (c *Client) ChannelsLeave(channel string) (*ChannelsLeaveResponse, error) {
	r := &ChannelsLeaveResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	err := c.Do(ChannelsLeaveMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/channels.list
// opt:exclude_archived // 1 // Optional // , default=0 Don't return archived channels.
func (c *Client) ChannelsList(optional url.Values) (*ChannelsListResponse, error) {
	r := &ChannelsListResponse{}
	err := c.Do(ChannelsListMethod, optional, r)
	return r, err
}

// https://api.slack.com/methods/channels.mark
// req:channel // C1234567890 // Required // Channel to set reading cursor in.
// req:ts // 1234567890.123456 // Required // Timestamp of the most recently seen message.
func (c *Client) ChannelsMark(channel string, ts Time) (*ChannelsMarkResponse, error) {
	r := &ChannelsMarkResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	args.Set("ts", ts.SlackString())
	err := c.Do(ChannelsMarkMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/channels.rename
// req:channel // C1234567890 // Required // Channel to rename
// req:name // new_name // Required // New name for channel.
func (c *Client) ChannelsRename(channel, name string) (*ChannelsRenameResponse, error) {
	r := &ChannelsRenameResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	args.Set("name", name)
	err := c.Do(ChannelsRenameMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/channels.setPurpose
// req:channel // C1234567890 // Required // Channel to set the purpose of
// req:purpose // My Purpose // Required // The new purpose
func (c *Client) ChannelsSetPurpose(channel, purpose string) (*ChannelsSetPurposeResponse, error) {
	r := &ChannelsSetPurposeResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	args.Set("purpose", purpose)
	err := c.Do(ChannelsSetPurposeMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/channels.setTopic
// req:channel // C1234567890 // Required // Channel to set the topic of
// req:topic // My Topic // Required // The new topic
func (c *Client) ChannelsSetTopic(channel, topic string) (*ChannelsSetTopicResponse, error) {
	r := &ChannelsSetTopicResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	args.Set("topic", topic)
	err := c.Do(ChannelsSetTopicMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/channels.unarchive
// req:channel // C1234567890 // Required // Channel to unarchive
func (c *Client) ChannelsUnarchive(channel string) (*ChannelsUnarchiveResponse, error) {
	r := &ChannelsUnarchiveResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	err := c.Do(ChannelsUnarchiveMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/chat.delete
// req:channel // C1234567890 // Required // Channel containing the message to be deleted.
// req:ts // 1405894322.002768 // Required // Timestamp of the message to be deleted.
func (c *Client) ChatDelete(channel string, ts Time) (*ChatDeleteResponse, error) {
	r := &ChatDeleteResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	args.Set("ts", ts.SlackString())
	err := c.Do(ChatDeleteMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/chat.postMessage
// req:channel // C1234567890 // Required // Channel, private group, or IM channel to send message to. Can be an encoded ID, or a name. See below for more details.
// req:text // Hello world // Required // Text of the message to send. See below for an explanation of formatting.
// opt:username // My Bot // Optional //  Name of bot.
// opt:as_user // true // Optional //  Pass true to post the message as the authed user, instead of as a bot
// opt:parse // full // Optional //  Change how messages are treated. See below.
// opt:link_names // 1 // Optional //  Find and link channel names and usernames.
// opt:attachments // [{"pretext": "pre-hello", "text": "text-world"}] // Optional //  Structured message attachments.
// opt:unfurl_links // true // Optional //  Pass true to enable unfurling of primarily text-based content.
// opt:unfurl_media // false // Optional //  Pass false to disable unfurling of media content.
// opt:icon_url // http://lorempixel.com/48/48 // Optional //  URL to an image to use as the icon for this message
// opt:icon_emoji // :chart_with_upwards_trend: // Optional //  emoji to use as the icon for this message. Overridesicon_url
func (c *Client) ChatPostMessage(channel, text string, optional url.Values) (*ChatPostMessageResponse, error) {
	r := &ChatPostMessageResponse{}
	var args url.Values
	if optional != nil {
		args = optional
	} else {
		args = url.Values{}
	}
	args.Set("channel", channel)
	args.Set("text", text)
	err := c.Do(ChatPostMessageMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/chat.update
// req:ts // 1405894322.002768 // Required // Timestamp of the message to be updated.
// req:channel // C1234567890 // Required // Channel containing the message to be updated.
// req:text // Hello world // Required // New text for the message, using the default formatting rules.
// opt:attachments // [{"pretext": "pre-hello", "text": "text-world"}] // Optional //  Structured message attachments.
// opt:parse // none // Optional //  Change how messages are treated. See below.
// opt:link_names // 1 // Optional //  Find and link channel names and usernames.
func (c *Client) ChatUpdate(channel, text string, ts Time, optional url.Values) (*ChatUpdateResponse, error) {
	r := &ChatUpdateResponse{}
	var args url.Values
	if optional != nil {
		args = optional
	} else {
		args = url.Values{}
	}
	args.Set("channel", channel)
	args.Set("text", text)
	args.Set("ts", ts.SlackString())
	err := c.Do(ChatUpdateMethod, args, r)
	return r, err

}

// https://api.slack.com/methods/dnd.endDnd
func (c *Client) DndEndDnd() (*DndEndDndResponse, error) {
	r := &DndEndDndResponse{}
	err := c.Do(DndEndDndMethod, nil, r)
	return r, err
}

// https://api.slack.com/methods/dnd.endSnooze
func (c *Client) DndEndSnooze() (*DndEndSnoozeResponse, error) {
	r := &DndEndSnoozeResponse{}
	err := c.Do(DndEndSnoozeMethod, nil, r)
	return r, err
}

// https://api.slack.com/methods/dnd.info
// opt:user // U1234 // Optional //  User to fetch status for (defaults to current user)
func (c *Client) DndInfo(optional url.Values) (*DndInfoResponse, error) {
	r := &DndInfoResponse{}
	err := c.Do(DndInfoMethod, optional, r)
	return r, err

}

// https://api.slack.com/methods/dnd.setSnooze
// req:num_minutes // 60 // Required // Number of minutes, from now, to snooze until.
func (c *Client) DndSetSnooze(minutes time.Duration) (*DndSetSnoozeResponse, error) {
	r := &DndSetSnoozeResponse{}
	args := url.Values{}
	args.Set("num_minutes", strconv.Itoa(int(minutes.Minutes())))
	err := c.Do(DndSetSnoozeMethod, args, r)
	return r, err

}

// https://api.slack.com/methods/dnd.teamInfo
// opt:users // U1234,U4567 // Optional //  Comma-separated list of users to fetch Do Not Disturb status for
func (c *Client) DndTeamInfo(optional url.Values) (*DndTeamInfoResponse, error) {
	r := &DndTeamInfoResponse{}
	err := c.Do(DndTeamInfoMethod, optional, r)
	return r, err
}

// https://api.slack.com/methods/emoji.list
func (c *Client) EmojiList() (*EmojiListResponse, error) {
	r := &EmojiListResponse{}
	err := c.Do(EmojiListMethod, nil, r)
	return r, err
}

// https://api.slack.com/methods/files.comments.add
// req:file // F1234567890 // Required // File to add a comment to.
// req:comment // Everyone should take a moment to read this file. // Required // Text of the comment to add.
func (c *Client) FilesCommentsAdd(file, comment string) (*FilesCommentsAddResponse, error) {
	r := &FilesCommentsAddResponse{}
	args := url.Values{}
	args.Set("file", file)
	args.Set("comment", comment)
	err := c.Do(FilesCommentsAddMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/files.comments.delete
// req:file // F1234567890 // Required // File to delete a comment from.
// req:id // Fc1234567890 // Required // The comment to delete.
func (c *Client) FilesCommentsDelete(file, comment_id string) (*FilesCommentsDeleteResponse, error) {
	r := &FilesCommentsDeleteResponse{}
	args := url.Values{}
	args.Set("file", file)
	args.Set("id", comment_id)
	err := c.Do(FilesCommentsDeleteMethod, args, r)
	return r, err

}

// https://api.slack.com/methods/files.comments.edit
// req:file // F1234567890 // Required // File containing the comment to edit.
// req:id // Fc1234567890 // Required // The comment to edit.
// req:comment // Everyone should take a moment to read this file, seriously. // Required // Text of the comment to edit.
func (c *Client) FilesCommentsEdit(file, comment_id, comment string) (*FilesCommentsEditResponse, error) {
	r := &FilesCommentsEditResponse{}
	args := url.Values{}
	args.Set("file", file)
	args.Set("id", comment_id)
	args.Set("comment", comment)
	err := c.Do(FilesCommentsEditMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/files.delete
// req:file // // Required // ID of file to delete.
func (c *Client) FilesDelete(file string) (*FilesDeleteResponse, error) {
	r := &FilesDeleteResponse{}
	args := url.Values{}
	args.Set("file", file)
	err := c.Do(FilesDeleteMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/files.info
// req:file // // Required // File to fetch info for
// opt:count // 20 // Optional // , default=100 Number of items to return per page.
// opt:page // 2 // Optional // , default=1 Page number of results to return.
func (c *Client) FilesInfo(file string, optional url.Values) (*FilesInfoResponse, error) {
	r := &FilesInfoResponse{}
	var args url.Values
	if optional != nil {
		args = optional
	} else {
		args = url.Values{}
	}
	args.Set("file", file)
	err := c.Do(FilesInfoMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/files.list
// opt:user // U1234567890 // Optional //  Filter files created by a single user.
// opt:channel // C1234567890 // Optional //  Filter files appearing in a specific channel, indicated by its ID.
// opt:ts_from // 123456789 // Optional // , default=0 Filter files created after this timestamp (inclusive).
// opt:ts_to // 123456789 // Optional // , default=now Filter files created before this timestamp (inclusive).
// opt:types // all,posts,snippets,images,gdocs,zips,pdfs // Optional (multiple) // , default=all Filter files by type:
// opt:count // 20 // Optional // , default=100 Number of items to return per page.
// opt:page // 2 // Optional // , default=1 Page number of results to return.
func (c *Client) FilesList(optional url.Values) (*FilesListResponse, error) {
	r := &FilesListResponse{}
	err := c.Do(FilesListMethod, optional, r)
	return r, err
}

// https://api.slack.com/methods/files.upload
// req:file // ... // Required // File contents via multipart/form-data.
// req:filename // foo.txt // Required // Filename of file.
// opt:content // ... // Optional //  File contents via a POST var.
// opt:filetype // php // Optional //  Slack-internal file type identifier.
// opt:title // My File // Optional //  Title of file.
// opt:initial_comment // Best! // Optional //  Initial comment to add to file.
// opt:channels // C1234567890 // Optional //  Comma-separated list of channel names or IDs where the file will be shared.
func (c *Client) FilesUpload(filename string, file io.Reader, optional url.Values) (*FilesUploadResponse, error) {

	r := &FilesUploadResponse{}

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	for key, _ := range optional {
		_ = writer.WriteField(key, optional.Get(key))
	}

	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(part, file); err != nil {
		return nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	// return object and error state even if error state
	// due to possible valuable response data
	return r, c.DoWithBody(FilesUploadMethod, writer.FormDataContentType(), &body, r)

}

// https://api.slack.com/methods/groups.archive
// req:channel // G1234567890 // Required // Private channel to archive
func (c *Client) GroupsArchive(channel string) (*GroupsArchiveResponse, error) {
	r := &GroupsArchiveResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	err := c.Do(GroupsArchiveMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/groups.close
// req:channel // G1234567890 // Required // Private channel to close.
func (c *Client) GroupsClose(channel string) (*GroupsCloseResponse, error) {
	r := &GroupsCloseResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	err := c.Do(GroupsCloseMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/groups.create
// req:name // // Required // Name of private channel to create
func (c *Client) GroupsCreate(name string) (*GroupsCreateResponse, error) {
	r := &GroupsCreateResponse{}
	args := url.Values{}
	args.Set("name", name)
	err := c.Do(GroupsCreateMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/groups.createChild
// req:channel // G1234567890 // Required // Private channel to clone and archive.
func (c *Client) GroupsCreateChild(channel string) (*GroupsCreateChildResponse, error) {
	r := &GroupsCreateChildResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	err := c.Do(GroupsCreateChildMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/groups.history
// req:channel // G1234567890 // Required // Private channel to fetch history for.
// opt:latest // 1234567890.123456 // Optional // , default=now End of time range of messages to include in results.
// opt:oldest // 1234567890.123456 // Optional // , default=0 Start of time range of messages to include in results.
// opt:inclusive // 1 // Optional // , default=0 Include messages with latest or oldest timestamp in results.
// opt:count // 100 // Optional // , default=100 Number of messages to return, between 1 and 1000.
// opt:unreads // 1 // Optional // , default=0 Include unread_count_display in the output?
func (c *Client) GroupsHistory(channel string, optional url.Values) (*GroupsHistoryResponse, error) {
	r := &GroupsHistoryResponse{}
	var args url.Values
	if optional != nil {
		args = optional
	} else {
		args = url.Values{}
	}
	args.Set("channel", channel)
	err := c.Do(GroupsHistoryMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/groups.info
// req:channel // C1234567890 // Required // Private channel to get info on
func (c *Client) GroupsInfo(channel string) (*GroupsInfoResponse, error) {
	r := &GroupsInfoResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	err := c.Do(GroupsInfoMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/groups.invite
// req:channel // G1234567890 // Required // Private channel to invite user to.
// req:user // U1234567890 // Required // User to invite.
func (c *Client) GroupsInvite(channel, user string) (*GroupsInviteResponse, error) {
	r := &GroupsInviteResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	args.Set("user", user)
	err := c.Do(GroupsInviteMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/groups.kick
// req:channel // G1234567890 // Required // Private channel to remove user from.
// req:user // U1234567890 // Required // User to remove from private channel.
func (c *Client) GroupsKick(channel, user string) (*GroupsKickResponse, error) {
	r := &GroupsKickResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	args.Set("user", user)
	err := c.Do(GroupsKickMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/groups.leave
// req:channel // G1234567890 // Required // Private channel to leave
func (c *Client) GroupsLeave(channel string) (*GroupsLeaveResponse, error) {
	r := &GroupsLeaveResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	err := c.Do(GroupsLeaveMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/groups.list
// opt:exclude_archived // 1 // Optional // , default=0 Don't return archived private channels.
func (c *Client) GroupsList(optional url.Values) (*GroupsListResponse, error) {
	r := &GroupsListResponse{}
	err := c.Do(GroupsListMethod, optional, r)
	return r, err
}

// https://api.slack.com/methods/groups.mark
// req:channel // G1234567890 // Required // Private channel to set reading cursor in.
// req:ts // 1234567890.123456 // Required // Timestamp of the most recently seen message.
func (c *Client) GroupsMark(channel string, ts Time) (*GroupsMarkResponse, error) {
	r := &GroupsMarkResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	args.Set("ts", ts.SlackString())
	err := c.Do(GroupsMarkMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/groups.open
// req:channel // G1234567890 // Required // Private channel to open.
func (c *Client) GroupsOpen(channel string) (*GroupsOpenResponse, error) {
	r := &GroupsOpenResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	err := c.Do(GroupsOpenMethod, args, r)
	return r, err

}

// https://api.slack.com/methods/groups.rename
// req:channel // C1234567890 // Required // Private channel to rename
// req:name // // Required // New name for private channel.
func (c *Client) GroupsRename(channel, name string) (*GroupsRenameResponse, error) {
	r := &GroupsRenameResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	args.Set("name", name)
	err := c.Do(GroupsRenameMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/groups.setPurpose
// req:channel // C1234567890 // Required // Private channel to set the purpose of
// req:purpose // My Purpose // Required // The new purpose
func (c *Client) GroupsSetPurpose(channel, purpose string) (*GroupsSetPurposeResponse, error) {
	r := &GroupsSetPurposeResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	args.Set("purpose", purpose)
	err := c.Do(GroupsSetPurposeMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/groups.setTopic
// req:channel // C1234567890 // Required // Private channel to set the topic of
// req:topic // My Topic // Required // The new topic
func (c *Client) GroupsSetTopic(channel, topic string) (*GroupsSetTopicResponse, error) {
	r := &GroupsSetTopicResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	args.Set("topic", topic)
	err := c.Do(GroupsSetTopicMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/groups.unarchive
// req:channel // G1234567890 // Required // Private channel to unarchive
func (c *Client) GroupsUnarchive(channel string) (*GroupsUnarchiveResponse, error) {
	r := &GroupsUnarchiveResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	err := c.Do(GroupsUnarchiveMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/im.close
// req:channel // D1234567890 // Required // Direct message channel to close.
func (c *Client) ImClose(channel string) (*ImCloseResponse, error) {
	r := &ImCloseResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	err := c.Do(ImCloseMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/im.history
// req:channel // D1234567890 // Required // Direct message channel to fetch history for.
// opt:latest // 1234567890.123456 // Optional // , default=now End of time range of messages to include in results.
// opt:oldest // 1234567890.123456 // Optional // , default=0 Start of time range of messages to include in results.
// opt:inclusive // 1 // Optional // , default=0 Include messages with latest or oldest timestamp in results.
// opt:count // 100 // Optional // , default=100 Number of messages to return, between 1 and 1000.
// opt:unreads // 1 // Optional // , default=0 Include unread_count_display in the output?
func (c *Client) ImHistory(channel string, optional url.Values) (*ImHistoryResponse, error) {
	r := &ImHistoryResponse{}
	var args url.Values
	if optional != nil {
		args = optional
	} else {
		args = url.Values{}
	}
	args.Set("channel", channel)
	err := c.Do(ImHistoryMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/im.list
func (c *Client) ImList() (*ImListResponse, error) {
	r := &ImListResponse{}
	err := c.Do(ImListMethod, nil, r)
	return r, err
}

// https://api.slack.com/methods/im.mark
// req:channel // D1234567890 // Required // Direct message channel to set reading cursor in.
// req:ts // 1234567890.123456 // Required // Timestamp of the most recently seen message.
func (c *Client) ImMark(channel string, ts Time) (*ImMarkResponse, error) {
	r := &ImMarkResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	args.Set("ts", ts.SlackString())
	err := c.Do(ImMarkMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/im.open
// req:user // U1234567890 // Required // User to open a direct message channel with.
func (c *Client) ImOpen(user string) (*ImOpenResponse, error) {
	r := &ImOpenResponse{}
	args := url.Values{}
	args.Set("user", user)
	err := c.Do(ImOpenMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/mpim.close
// req:channel // // Required // MPIM to close.
func (c *Client) MpimClose(channel string) (*MpimCloseResponse, error) {
	r := &MpimCloseResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	err := c.Do(MpimCloseMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/mpim.history
// req:channel // // Required // Multiparty direct message to fetch history for.
// opt:latest // 1234567890.123456 // Optional // , default=now End of time range of messages to include in results.
// opt:oldest // 1234567890.123456 // Optional // , default=0 Start of time range of messages to include in results.
// opt:inclusive // 1 // Optional // , default=0 Include messages with latest or oldest timestamp in results.
// opt:count // 100 // Optional // , default=100 Number of messages to return, between 1 and 1000.
// opt:unreads // 1 // Optional // , default=0 Include unread_count_display in the output?
func (c *Client) MpimHistory(channel string, optional url.Values) (*MpimHistoryResponse, error) {
	r := &MpimHistoryResponse{}
	var args url.Values
	if optional != nil {
		args = optional
	} else {
		args = url.Values{}
	}
	args.Set("channel", channel)
	err := c.Do(MpimHistoryMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/mpim.list
func (c *Client) MpimList() (*MpimListResponse, error) {
	r := &MpimListResponse{}
	err := c.Do(MpimListMethod, nil, r)
	return r, err
}

// https://api.slack.com/methods/mpim.mark
// req:channel // G1234567890 // Required // multiparty direct message channel to set reading cursor in.
// req:ts // 1234567890.123456 // Required // Timestamp of the most recently seen message.
func (c *Client) MpimMark(channel string, ts Time) (*MpimMarkResponse, error) {
	r := &MpimMarkResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	args.Set("ts", ts.SlackString())
	err := c.Do(MpimMarkMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/mpim.open
// req:users // U1234567890,U2345678901,U3456789012 // Required // Comma separated lists of users.  The ordering of the users is preserved whenever a MPIM group is returned.
func (c *Client) MpimOpen(users []string) (*MpimOpenResponse, error) {
	r := &MpimOpenResponse{}
	args := url.Values{}
	args.Set("users", strings.Join(users, ","))
	err := c.Do(MpimOpenMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/oauth.access
// req:client_id // 4b39e9-752c4 // Required // Issued when you created your application.
// req:client_secret // 33fea0113f5b1 // Required // Issued when you created your application.
// req:code // ccdaa72ad // Required // The code param returned via the OAuth callback.
// opt:redirect_uri // http://example.com // Optional //  This must match the originally submitted URI (if one was sent).
func (c *Client) OauthAccess(client_id, client_secret, code string, optional url.Values) (*OauthAccessResponse, error) {
	r := &OauthAccessResponse{}
	var args url.Values
	if optional != nil {
		args = optional
	} else {
		args = url.Values{}
	}
	args.Set("client_id", client_id)
	args.Set("client_secret", client_secret)
	args.Set("code", code)
	err := c.Do(OauthAccessMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/pins.add
// req:channel // C1234567890 // Required // Channel to pin the item in.
// opt:file // F1234567890 // Optional //  File to pin.
// opt:file_comment // Fc1234567890 // Optional //  File comment to pin.
// opt:timestamp // 1234567890.123456 // Optional //  Timestamp of the message to pin.
func (c *Client) PinsAdd(channel string, optional url.Values) (*PinsAddResponse, error) {
	r := &PinsAddResponse{}
	var args url.Values
	if optional != nil {
		args = optional
	} else {
		args = url.Values{}
	}
	args.Set("channel", channel)
	err := c.Do(PinsAddMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/pins.list
// req:channel // C1234567890 // Required // Channel to get pinned items for.
func (c *Client) PinsList(channel string) (*PinsListResponse, error) {
	r := &PinsListResponse{}
	args := url.Values{}
	args.Set("channel", channel)
	err := c.Do(PinsListMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/pins.remove
// req:channel // C1234567890 // Required // Channel where the item is pinned to.
// opt:file // F1234567890 // Optional //  File to un-pin.
// opt:file_comment // Fc1234567890 // Optional //  File comment to un-pin.
// opt:timestamp // 1234567890.123456 // Optional //  Timestamp of the message to un-pin.
func (c *Client) PinsRemove(channel string, optional url.Values) (*PinsRemoveResponse, error) {
	r := &PinsRemoveResponse{}
	var args url.Values
	if optional != nil {
		args = optional
	} else {
		args = url.Values{}
	}
	args.Set("channel", channel)
	err := c.Do(PinsRemoveMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/reactions.add
// req:name // thumbsup // Required // Reaction (emoji) name.
// opt:file // F1234567890 // Optional //  File to add reaction to.
// opt:file_comment // Fc1234567890 // Optional //  File comment to add reaction to.
// opt:channel // C1234567890 // Optional //  Channel where the message to add reaction to was posted.
// opt:timestamp // 1234567890.123456 // Optional //  Timestamp of the message to add reaction to.
func (c *Client) ReactionsAdd(name string, optional url.Values) (*ReactionsAddResponse, error) {
	r := &ReactionsAddResponse{}
	var args url.Values
	if optional != nil {
		args = optional
	} else {
		args = url.Values{}
	}
	args.Set("name", name)
	err := c.Do(ReactionsAddMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/reactions.get
// opt:file // F1234567890 // Optional //  File to get reactions for.
// opt:file_comment // Fc1234567890 // Optional //  File comment to get reactions for.
// opt:channel // C1234567890 // Optional //  Channel where the message to get reactions for was posted.
// opt:timestamp // 1234567890.123456 // Optional //  Timestamp of the message to get reactions for.
// opt:full // // Optional //  If true always return the complete reaction list.
func (c *Client) ReactionsGet(optional url.Values) (*ReactionsGetResponse, error) {
	r := &ReactionsGetResponse{}
	err := c.Do(ReactionsGetMethod, optional, r)
	return r, err
}

// https://api.slack.com/methods/reactions.list
// opt:user // U1234567890 // Optional //  Show reactions made by this user. Defaults to the authed user.
// opt:full // // Optional //  If true always return the complete reaction list.
// opt:count // 20 // Optional // , default=100 Number of items to return per page.
// opt:page // 2 // Optional // , default=1 Page number of results to return.
func (c *Client) ReactionsList(optional url.Values) (*ReactionsListResponse, error) {
	r := &ReactionsListResponse{}
	err := c.Do(ReactionsListMethod, optional, r)
	return r, err
}

// https://api.slack.com/methods/reactions.remove
// req:name // thumbsup // Required // Reaction (emoji) name.
// opt:file // F1234567890 // Optional //  File to remove reaction from.
// opt:file_comment // Fc1234567890 // Optional //  File comment to remove reaction from.
// opt:channel // C1234567890 // Optional //  Channel where the message to remove reaction from was posted.
// opt:timestamp // 1234567890.123456 // Optional //  Timestamp of the message to remove reaction from.
func (c *Client) ReactionsRemove(name string, optional url.Values) (*ReactionsRemoveResponse, error) {
	r := &ReactionsRemoveResponse{}
	var args url.Values
	if optional != nil {
		args = optional
	} else {
		args = url.Values{}
	}
	args.Set("name", name)
	err := c.Do(ReactionsRemoveMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/rtm.start
// opt:simple_latest // // Optional //  Return timestamp only for latest message object of each channel (improves performance).
// opt:no_unreads // // Optional //  Skip unread counts for each channel (improves performance).
// opt:mpim_aware // // Optional //  Returns MPIMs to the client in the API response.
func (c *Client) RTMStart(optional url.Values) (*RTMStartResponse, error) {
	r := &RTMStartResponse{}
	err := c.Do(RTMStartMethod, optional, r)
	return r, err
}

// https://api.slack.com/methods/search.all
// req:query // pickleface // Required // Search query. May contains booleans, etc.
// opt:sort // timestamp // Optional // , default=score Return matches sorted by either score or timestamp.
// opt:sort_dir // asc // Optional // , default=desc Change sort direction to ascending (asc) or descending (desc).
// opt:highlight // 1 // Optional //  Pass a value of 1 to enable query highlight markers (see below).
// opt:count // 20 // Optional // , default=20 Number of items to return per page.
// opt:page // 2 // Optional // , default=1 Page number of results to return.
func (c *Client) SearchAll(query string, optional url.Values) (*SearchAllResponse, error) {
	r := &SearchAllResponse{}
	var args url.Values
	if optional != nil {
		args = optional
	} else {
		args = url.Values{}
	}
	args.Set("query", query)
	err := c.Do(SearchAllMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/search.files
// req:query // pickleface // Required // Search query. May contain booleans, etc.
// opt:sort // timestamp // Optional // , default=score Return matches sorted by either score or timestamp.
// opt:sort_dir // asc // Optional // , default=desc Change sort direction to ascending (asc) or descending (desc).
// opt:highlight // 1 // Optional //  Pass a value of 1 to enable query highlight markers (see below).
// opt:count // 20 // Optional // , default=20 Number of items to return per page.
// opt:page // 2 // Optional // , default=1 Page number of results to return.
func (c *Client) SearchFiles(query string, optional url.Values) (*SearchFilesResponse, error) {
	r := &SearchFilesResponse{}
	var args url.Values
	if optional != nil {
		args = optional
	} else {
		args = url.Values{}
	}
	args.Set("query", query)
	err := c.Do(SearchFilesMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/search.messages
// req:query // pickleface // Required // Search query. May contains booleans, etc.
// opt:sort // timestamp // Optional // , default=score Return matches sorted by either score or timestamp.
// opt:sort_dir // asc // Optional // , default=desc Change sort direction to ascending (asc) or descending (desc).
// opt:highlight // 1 // Optional //  Pass a value of 1 to enable query highlight markers (see below).
// opt:count // 20 // Optional // , default=20 Number of items to return per page.
// opt:page // 2 // Optional // , default=1 Page number of results to return.
func (c *Client) SearchMessages(query string, optional url.Values) (*SearchMessagesResponse, error) {
	r := &SearchMessagesResponse{}
	var args url.Values
	if optional != nil {
		args = optional
	} else {
		args = url.Values{}
	}
	args.Set("query", query)
	err := c.Do(SearchMessagesMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/stars.add
// opt:file // F1234567890 // Optional //  File to add star to.
// opt:file_comment // Fc1234567890 // Optional //  File comment to add star to.
// opt:channel // C1234567890 // Optional //  Channel to add star to, or channel where the message to add star to was posted (used with timestamp).
// opt:timestamp // 1234567890.123456 // Optional //  Timestamp of the message to add star to.
func (c *Client) StarsAdd(optional url.Values) (*StarsAddResponse, error) {
	r := &StarsAddResponse{}
	err := c.Do(StarsAddMethod, optional, r)
	return r, err
}

// https://api.slack.com/methods/stars.list
// opt:user // U1234567890 // Optional //  Show stars by this user. Defaults to the authed user.
// opt:count // 20 // Optional // , default=100 Number of items to return per page.
// opt:page // 2 // Optional // , default=1 Page number of results to return.
func (c *Client) StarsList(optional url.Values) (*StarsListResponse, error) {
	r := &StarsListResponse{}
	err := c.Do(StarsListMethod, optional, r)
	return r, err
}

// https://api.slack.com/methods/stars.remove
// opt:file // F1234567890 // Optional //  File to remove star from.
// opt:file_comment // Fc1234567890 // Optional //  File comment to remove star from.
// opt:channel // C1234567890 // Optional //  Channel to remove star from, or channel where the message to remove star from was posted (used with timestamp).
// opt:timestamp // 1234567890.123456 // Optional //  Timestamp of the message to remove star from.
func (c *Client) StarsRemove(optional url.Values) (*StarsRemoveResponse, error) {
	r := &StarsRemoveResponse{}
	err := c.Do(StarsRemoveMethod, optional, r)
	return r, err
}

// https://api.slack.com/methods/team.accessLogs
// opt:count // 20 // Optional // , default=100 Number of items to return per page.
// opt:page // 2 // Optional // , default=1 Page number of results to return.
func (c *Client) TeamAccessLogs(optional url.Values) (*TeamAccessLogsResponse, error) {
	r := &TeamAccessLogsResponse{}
	err := c.Do(TeamAccessLogsMethod, optional, r)
	return r, err
}

// https://api.slack.com/methods/team.info
func (c *Client) TeamInfo() (*TeamInfoResponse, error) {
	r := &TeamInfoResponse{}
	err := c.Do(TeamInfoMethod, nil, r)
	return r, err
}

// https://api.slack.com/methods/team.integrationLogs
// opt:service_id // // Optional //  Filter logs to this service. Defaults to all logs.
// opt:app_id // // Optional //  Filter logs to this Slack app. Defaults to all logs.
// opt:user // U1234567890 // Optional //  Filter logs generated by this userâ€™s actions. Defaults to all logs.
// opt:change_type // added // Optional //  Filter logs with this change type. Defaults to all logs.
// opt:count // 20 // Optional // , default=100 Number of items to return per page.
// opt:page // 2 // Optional // , default=1 Page number of results to return.
func (c *Client) TeamIntegrationLogs(optional url.Values) (*TeamIntegrationLogsResponse, error) {
	r := &TeamIntegrationLogsResponse{}
	err := c.Do(TeamIntegrationLogsMethod, optional, r)
	return r, err
}

// https://api.slack.com/methods/usergroups.create
// req:name // My Test Team // Required // A name for the user group. Must be unique among user groups.
// opt:handle // // Optional //  A mention handle. Must be unique among channels, users and user groups.
// opt:description // // Optional //  A short description of the user group.
// opt:channels // // Optional //  A comma separated string of encoded channel IDs for which the user group uses as a default.
// opt:include_count // 1 // Optional //  Include the number of users in each user group.
func (c *Client) UsergroupsCreate(name string, optional url.Values) (*UsergroupsCreateResponse, error) {
	r := &UsergroupsCreateResponse{}
	var args url.Values
	if optional != nil {
		args = optional
	} else {
		args = url.Values{}
	}
	args.Set("name", name)
	err := c.Do(UsergroupsCreateMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/usergroups.disable
// req:usergroup // S0604QSJC // Required // The encoded ID of the user group to disable.
// opt:include_count // 1 // Optional //  Include the number of users in the user group.
func (c *Client) UsergroupsDisable(usergroup string, optional url.Values) (*UsergroupsDisableResponse, error) {
	r := &UsergroupsDisableResponse{}
	var args url.Values
	if optional != nil {
		args = optional
	} else {
		args = url.Values{}
	}
	args.Set("usergroup", usergroup)
	err := c.Do(UsergroupsDisableMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/usergroups.enable
// req:usergroup // S0604QSJC // Required // The encoded ID of the user group to enable.
// opt:include_count // 1 // Optional //  Include the number of users in the user group.
func (c *Client) UsergroupsEnable(usergroup string, optional url.Values) (*UsergroupsEnableResponse, error) {
	r := &UsergroupsEnableResponse{}
	var args url.Values
	if optional != nil {
		args = optional
	} else {
		args = url.Values{}
	}
	args.Set("usergroup", usergroup)
	err := c.Do(UsergroupsEnableMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/usergroups.list
// opt:include_disabled // 1 // Optional //  Include disabled user groups.
// opt:include_count // 1 // Optional //  Include the number of users in each user group.
// opt:include_users // 1 // Optional //  Include the list of users for each user group.
func (c *Client) UsergroupsList(optional url.Values) (*UsergroupsListResponse, error) {
	r := &UsergroupsListResponse{}
	err := c.Do(UsergroupsListMethod, optional, r)
	return r, err
}

// https://api.slack.com/methods/usergroups.update
// req:usergroup // S0604QSJC // Required // The encoded ID of the user group to update.
// opt:name // My Test Team // Optional //  A name for the user group. Must be unique among user groups.
// opt:handle // // Optional //  A mention handle. Must be unique among channels, users and user groups.
// opt:description // // Optional //  A short description of the user group.
// opt:channels // // Optional //  A comma separated string of encoded channel IDs for which the user group uses as a default.
// opt:include_count // 1 // Optional //  Include the number of users in the user group.
func (c *Client) UsergroupsUpdate(usergroup string, optional url.Values) (*UsergroupsUpdateResponse, error) {
	r := &UsergroupsUpdateResponse{}
	var args url.Values
	if optional != nil {
		args = optional
	} else {
		args = url.Values{}
	}
	args.Set("usergroup", usergroup)
	err := c.Do(UsergroupsUpdateMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/usergroups.users.list
// req:usergroup // S0604QSJC // Required // The encoded ID of the user group to update.
// opt:include_disabled // 1 // Optional //  Allow results that involve disabled user groups.
func (c *Client) UsergroupsUsersList(usergroup string, optional url.Values) (*UsergroupsUsersListResponse, error) {
	r := &UsergroupsUsersListResponse{}
	var args url.Values
	if optional != nil {
		args = optional
	} else {
		args = url.Values{}
	}
	args.Set("usergroup", usergroup)
	err := c.Do(UsergroupsUsersListMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/usergroups.users.update
// req:usergroup // S0604QSJC // Required // The encoded ID of the user group to update.
// req:users // U060R4BJ4,U060RNRCZ // Required // A comma separated string of encoded user IDs that represent the entire list of users for the user group.
// opt:include_count // 1 // Optional //  Include the number of users in the user group.
func (c *Client) UsergroupsUsersUpdate(usergroup string, users []string, optional url.Values) (*UsergroupsUsersUpdateResponse, error) {
	r := &UsergroupsUsersUpdateResponse{}
	var args url.Values
	if optional != nil {
		args = optional
	} else {
		args = url.Values{}
	}
	args.Set("usergroup", usergroup)
	args.Set("users", strings.Join(users, ","))
	err := c.Do(UsergroupsUsersUpdateMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/users.getPresence
// req:user // U1234567890 // Required // User to get presence info on. Defaults to the authed user.
func (c *Client) UsersGetPresence(user string) (*UsersGetPresenceResponse, error) {
	r := &UsersGetPresenceResponse{}
	args := url.Values{}
	args.Set("user", user)
	err := c.Do(UsersGetPresenceMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/users.info
// req:user // U1234567890 // Required // User to get info on
func (c *Client) UsersInfo(user string) (*UsersInfoResponse, error) {
	r := &UsersInfoResponse{}
	args := url.Values{}
	args.Set("user", user)
	err := c.Do(UsersInfoMethod, args, r)
	return r, err
}

// https://api.slack.com/methods/users.list
// opt:presence // 1 // Optional //  Whether to include presence data in the output
func (c *Client) UsersList(optional url.Values) (*UsersListResponse, error) {
	r := &UsersListResponse{}
	err := c.Do(UsersListMethod, optional, r)
	return r, err
}

// https://api.slack.com/methods/users.setActive
func (c *Client) UsersSetActive() (*UsersSetActiveResponse, error) {
	r := &UsersSetActiveResponse{}
	err := c.Do(UsersSetActiveMethod, nil, r)
	return r, err
}

// https://api.slack.com/methods/users.setPresence
// req:presence // (auto or away) // Required // Either
func (c *Client) UsersSetPresence(presence Presence) (*UsersSetPresenceResponse, error) {
	r := &UsersSetPresenceResponse{}
	args := url.Values{}
	args.Set("presense", string(presence))
	err := c.Do(UsersSetPresenceMethod, args, r)
	return r, err
}
