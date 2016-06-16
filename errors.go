package slk

import (
	"errors"
	"fmt"
)

const (
	AccountInactiveErrorCode                  = `account_inactive`
	AlreadyArchivedErrorCode                  = `already_archived`
	AlreadyInChannelErrorCode                 = `already_in_channel`
	AlreadyPinnedErrorCode                    = `already_pinned`
	AlreadyReactedErrorCode                   = `already_reacted`
	AlreadyStarredErrorCode                   = `already_starred`
	BadClientSecretErrorCode                  = `bad_client_secret`
	BadRedirectUriErrorCode                   = `bad_redirect_uri`
	BadTimestampErrorCode                     = `bad_timestamp`
	CantArchiveGeneralErrorCode               = `cant_archive_general`
	CantDeleteErrorCode                       = `cant_delete`
	CantDeleteFileErrorCode                   = `cant_delete_file`
	CantDeleteMessageErrorCode                = `cant_delete_message`
	CantEditErrorCode                         = `cant_edit`
	CantInviteErrorCode                       = `cant_invite`
	CantInviteSelfErrorCode                   = `cant_invite_self`
	CantKickFromGeneralErrorCode              = `cant_kick_from_general`
	CantKickFromLastChannelErrorCode          = `cant_kick_from_last_channel`
	CantKickSelfErrorCode                     = `cant_kick_self`
	CantLeaveGeneralErrorCode                 = `cant_leave_general`
	CantLeaveLastChannelErrorCode             = `cant_leave_last_channel`
	CantUpdateMessageErrorCode                = `cant_update_message`
	ChannelNotFoundErrorCode                  = `channel_not_found`
	ComplianceExportsPreventDeletionErrorCode = `compliance_exports_prevent_deletion`
	EditWindowClosedErrorCode                 = `edit_window_closed`
	FileCommentNotFoundErrorCode              = `file_comment_not_found`
	FileDeletedErrorCode                      = `file_deleted`
	FileNotFoundErrorCode                     = `file_not_found`
	FileNotSharedErrorCode                    = `file_not_shared`
	GroupContainsOthersErrorCode              = `group_contains_others`
	InvalidArrayArgErrorCode                  = `invalid_array_arg`
	InvalidAuthErrorCode                      = `invalid_auth`
	InvalidChannelErrorCode                   = `invalid_channel`
	InvalidCharsetErrorCode                   = `invalid_charset`
	InvalidClientIdErrorCode                  = `invalid_client_id`
	InvalidCodeErrorCode                      = `invalid_code`
	InvalidFormDataErrorCode                  = `invalid_form_data`
	InvalidNameErrorCode                      = `invalid_name`
	InvalidPostTypeErrorCode                  = `invalid_post_type`
	InvalidPresenceErrorCode                  = `invalid_presence`
	InvalidTimestampErrorCode                 = `invalid_timestamp`
	InvalidTsLatestErrorCode                  = `invalid_ts_latest`
	InvalidTsOldestErrorCode                  = `invalid_ts_oldest`
	IsArchivedErrorCode                       = `is_archived`
	LastMemberErrorCode                       = `last_member`
	LastRaChannelErrorCode                    = `last_ra_channel`
	MessageNotFoundErrorCode                  = `message_not_found`
	MigrationInProgressErrorCode              = `migration_in_progress`
	MissingDurationErrorCode                  = `missing_duration`
	MissingPostTypeErrorCode                  = `missing_post_type`
	MsgTooLongErrorCode                       = `msg_too_long`
	NameTakenErrorCode                        = `name_taken`
	NoChannelErrorCode                        = `no_channel`
	NoCommentErrorCode                        = `no_comment`
	NoItemSpecifiedErrorCode                  = `no_item_specified`
	NoReactionErrorCode                       = `no_reaction`
	NoTextErrorCode                           = `no_text`
	NotArchivedErrorCode                      = `not_archived`
	NotAuthedErrorCode                        = `not_authed`
	NotAuthorizedErrorCode                    = `not_authorized`
	NotEnoughUsersErrorCode                   = `not_enough_users`
	NotInChannelErrorCode                     = `not_in_channel`
	NotInGroupErrorCode                       = `not_in_group`
	NotPinnedErrorCode                        = `not_pinned`
	NotStarredErrorCode                       = `not_starred`
	OverPaginationLimitErrorCode              = `over_pagination_limit`
	PaidOnlyErrorCode                         = `paid_only`
	PermissionDeniedErrorCode                 = `permission_denied`
	PostingToGeneralChannelDeniedErrorCode    = `posting_to_general_channel_denied`
	RateLimitedErrorCode                      = `rate_limited`
	RequestTimeoutErrorCode                   = `request_timeout`
	RestrictedActionErrorCode                 = `restricted_action`
	SnoozeEndFailedErrorCode                  = `snooze_end_failed`
	SnoozeFailedErrorCode                     = `snooze_failed`
	SnoozeNotActiveErrorCode                  = `snooze_not_active`
	TooLongErrorCode                          = `too_long`
	TooManyEmojiErrorCode                     = `too_many_emoji`
	TooManyReactionsErrorCode                 = `too_many_reactions`
	TooManyUsersErrorCode                     = `too_many_users`
	UnknownErrorErrorCode                     = `unknown_error`
	UnknownTypeErrorCode                      = `unknown_type`
	UserDisabledErrorCode                     = `user_disabled`
	UserDoesNotOwnChannelErrorCode            = `user_does_not_own_channel`
	UserIsBotErrorCode                        = `user_is_bot`
	UserIsRestrictedErrorCode                 = `user_is_restricted`
	UserIsUltraRestrictedErrorCode            = `user_is_ultra_restricted`
	UserNotFoundErrorCode                     = `user_not_found`
	UserNotVisibleErrorCode                   = `user_not_visible`
	UsersListNotSuppliedErrorCode             = `users_list_not_supplied`
)

var (
	ErrAccountInactive                  = errors.New("account_inactive: Authentication token is for a deleted user or team.")
	ErrAlreadyArchived                  = errors.New("already_archived: Already been archived.")
	ErrAlreadyInChannel                 = errors.New("already_in_channel: Invited user is already in the channel.")
	ErrAlreadyPinned                    = errors.New("already_pinned: The specified item is already pinned to the channel.")
	ErrAlreadyReacted                   = errors.New("already_reacted: The specified item already has the user/reaction combination.")
	ErrAlreadyStarred                   = errors.New("already_starred: The specified item has already been starred by the authenticated user.")
	ErrBadClientSecret                  = errors.New("bad_client_secret: Value passed for client_secret was invalid.")
	ErrBadRedirectUri                   = errors.New("bad_redirect_uri: Value passed for redirect_uri did not match the redirect_uri in the original request.")
	ErrBadTimestamp                     = errors.New("bad_timestamp: Value passed for timestamp was invalid.")
	ErrCantArchiveGeneral               = errors.New("cant_archive_general: You cannot archive the general channel")
	ErrCantDelete                       = errors.New("cant_delete: The requested comment could not be deleted.")
	ErrCantDeleteFile                   = errors.New("cant_delete_file: Authenticated user does not have permission to delete this file.")
	ErrCantDeleteMessage                = errors.New("cant_delete_message: Authenticated user does not have permission to delete this message.")
	ErrCantEdit                         = errors.New("cant_edit: The requested file could not be found.")
	ErrCantInvite                       = errors.New("cant_invite: User cannot be invited to this channel/group.")
	ErrCantInviteSelf                   = errors.New("cant_invite_self: Authenticated user cannot invite themselves to a channel/group.")
	ErrCantKickFromGeneral              = errors.New("cant_kick_from_general: User cannot be removed from #general.")
	ErrCantKickFromLastChannel          = errors.New("cant_kick_from_last_channel: User cannot be removed from the last channel they're in.")
	ErrCantKickSelf                     = errors.New("cant_kick_self: You can't remove yourself from a group")
	ErrCantLeaveGeneral                 = errors.New("cant_leave_general: Authenticated user cannot leave the general channel")
	ErrCantLeaveLastChannel             = errors.New("cant_leave_last_channel: Authenticated user cannot leave the last channel they are in")
	ErrCantUpdateMessage                = errors.New("cant_update_message: Authenticated user does not have permission to update this message.")
	ErrChannelNotFound                  = errors.New("channel_not_found: Value passed for channel was invalid.")
	ErrComplianceExportsPreventDeletion = errors.New("compliance_exports_prevent_deletion: Compliance exports are on, messages can not be deleted")
	ErrEditWindowClosed                 = errors.New("edit_window_closed: The timeframe for editing the comment has expired.")
	ErrFileCommentNotFound              = errors.New("file_comment_not_found: File comment specified by file_comment does not exist.")
	ErrFileDeleted                      = errors.New("file_deleted: The requested file was previously deleted.")
	ErrFileNotFound                     = errors.New("file_not_found: The requested file could not be found.")
	ErrFileNotShared                    = errors.New("file_not_shared: File specified by file is not public nor shared to the channel.")
	ErrGroupContainsOthers              = errors.New("group_contains_others: Restricted accounts cannot archive groups containing others.")
	ErrInvalidArrayArg                  = errors.New("invalid_array_arg: The method was passed a PHP-style array argument (e.g. with a name like foo[7] ). These are never valid with the Slack API.")
	ErrInvalidAuth                      = errors.New("invalid_auth: Invalid authentication token.")
	ErrInvalidChannel                   = errors.New("invalid_channel: One or more channels supplied are invalid")
	ErrInvalidCharset                   = errors.New("invalid_charset: The method was called via a POST request, but the charset specified in the Content-Type header was invalid. Valid charset names are: utf-8 iso-8859-1 .")
	ErrInvalidClientId                  = errors.New("invalid_client_id: Value passed for client_id was invalid.")
	ErrInvalidCode                      = errors.New("invalid_code: Value passed for code was invalid.")
	ErrInvalidFormData                  = errors.New("invalid_form_data: The method was called via a POST request with Content-Type application/x-www-form-urlencoded or multipart/form-data, but the form data was either missing or syntactically invalid.")
	ErrInvalidName                      = errors.New("invalid_name: New name is invalid")
	ErrInvalidPostType                  = errors.New("invalid_post_type: The method was called via a POST request, but the specified Content-Type was invalid. Valid types are: application/json application/x-www-form-urlencoded multipart/form-data text/plain .")
	ErrInvalidPresence                  = errors.New("invalid_presence: Value passed for presence was invalid.")
	ErrInvalidTimestamp                 = errors.New("invalid_timestamp: Value passed for timestamp was invalid.")
	ErrInvalidTsLatest                  = errors.New("invalid_ts_latest: Value passed for latest was invalid")
	ErrInvalidTsOldest                  = errors.New("invalid_ts_oldest: Value passed for oldest was invalid")
	ErrIsArchived                       = errors.New("is_archived: Channel/Group has been archived.")
	ErrLastMember                       = errors.New("last_member: Authenticated user is the last member of a group and cannot leave it")
	ErrLastRaChannel                    = errors.New("last_ra_channel: You cannot archive the last channel for a restricted account.")
	ErrMessageNotFound                  = errors.New("message_not_found: Message specified by channel and/or timestamp does not exist.")
	ErrMigrationInProgress              = errors.New("migration_in_progress: Team is being migrated between servers. See the team_migration_started event documentation for details.")
	ErrMissingDuration                  = errors.New("missing_duration: No value provided for num_minutes")
	ErrMissingPostType                  = errors.New("missing_post_type: The method was called via a POST request and included a data payload, but the request did not include a Content-Type header.")
	ErrMsgTooLong                       = errors.New("msg_too_long: Message text is too long")
	ErrNameTaken                        = errors.New("name_taken: New channel name is taken")
	ErrNoChannel                        = errors.New("no_channel: No group name was passed.")
	ErrNoComment                        = errors.New("no_comment: The comment field was empty.")
	ErrNoItemSpecified                  = errors.New("no_item_specified: One of file, file_comment, or timestamp was not specified.")
	ErrNoReaction                       = errors.New("no_reaction: The specified item does not have the user/reaction combination.")
	ErrNoText                           = errors.New("no_text: No message text provided")
	ErrNotArchived                      = errors.New("not_archived: Channel/Group is not archived.")
	ErrNotAuthed                        = errors.New("not_authed: No authentication token provided.")
	ErrNotAuthorized                    = errors.New("not_authorized: Caller cannot rename this channel")
	ErrNotEnoughUsers                   = errors.New("not_enough_users: Needs at least 2 users to open")
	ErrNotInChannel                     = errors.New("not_in_channel: User was not in the channel.")
	ErrNotInGroup                       = errors.New("not_in_group: User or caller were are not in the group")
	ErrNotPinned                        = errors.New("not_pinned: The specified item is not pinned to the channel.")
	ErrNotStarred                       = errors.New("not_starred: The specified item is not currently starred by the authenticated user.")
	ErrOverPaginationLimit              = errors.New("over_pagination_limit: It is not possible to request more than 1000 items per page or more than 100 pages.")
	ErrPaidOnly                         = errors.New("paid_only: This is only available to paid teams.")
	ErrPermissionDenied                 = errors.New("permission_denied: The user does not have permission take that action.")
	ErrPostingToGeneralChannelDenied    = errors.New("posting_to_general_channel_denied: An admin has restricted posting to the #general channel.")
	ErrRateLimited                      = errors.New("rate_limited: Application has posted too many messages, read the Rate Limit documentation for more information")
	ErrRequestTimeout                   = errors.New("request_timeout: The method was called via a POST request, but the POST data was either missing or truncated.")
	ErrRestrictedAction                 = errors.New("restricted_action: A team preference prevents the authenticated user from that action.")
	ErrSnoozeEndFailed                  = errors.New("snooze_end_failed: There was a problem setting the user's Do Not Disturb status")
	ErrSnoozeFailed                     = errors.New("snooze_failed: There was a problem setting the user's Do Not Disturb status")
	ErrSnoozeNotActive                  = errors.New("snooze_not_active: Snooze is not active for this user and cannot be ended")
	ErrTooLong                          = errors.New("too_long: Text was longer than 250 characters.")
	ErrTooManyEmoji                     = errors.New("too_many_emoji: The limit for distinct reactions (i.e. emoji) on the item has been reached.")
	ErrTooManyReactions                 = errors.New("too_many_reactions: The limit for reactions a person may add to the item has been reached.")
	ErrTooManyUsers                     = errors.New("too_many_users: Needs at most 8 users to open")
	ErrUnknownError                     = errors.New("unknown_error: There was a mysterious problem ending the user's Do Not Disturb session")
	ErrUnknownType                      = errors.New("unknown_type: Value passed for types was invalid")
	ErrUserDisabled                     = errors.New("user_disabled: The user has been disabled.")
	ErrUserDoesNotOwnChannel            = errors.New("user_does_not_own_channel: Calling user does not own this DM channel.")
	ErrUserIsBot                        = errors.New("user_is_bot: This method cannot be called by a bot user.")
	ErrUserIsRestricted                 = errors.New("user_is_restricted: This method cannot be called by a restricted user or single channel guest.")
	ErrUserIsUltraRestricted            = errors.New("user_is_ultra_restricted: This method cannot be called by a single channel guest.")
	ErrUserNotFound                     = errors.New("user_not_found: Value passed for user was invalid.")
	ErrUserNotVisible                   = errors.New("user_not_visible: The requested user is not visible to the calling user")
	ErrUsersListNotSupplied             = errors.New("users_list_not_supplied: Missing users in request")

	errMap = map[string]error{
		AccountInactiveErrorCode:                  ErrAccountInactive,
		AlreadyArchivedErrorCode:                  ErrAlreadyArchived,
		AlreadyInChannelErrorCode:                 ErrAlreadyInChannel,
		AlreadyPinnedErrorCode:                    ErrAlreadyPinned,
		AlreadyReactedErrorCode:                   ErrAlreadyReacted,
		AlreadyStarredErrorCode:                   ErrAlreadyStarred,
		BadClientSecretErrorCode:                  ErrBadClientSecret,
		BadRedirectUriErrorCode:                   ErrBadRedirectUri,
		BadTimestampErrorCode:                     ErrBadTimestamp,
		CantArchiveGeneralErrorCode:               ErrCantArchiveGeneral,
		CantDeleteErrorCode:                       ErrCantDelete,
		CantDeleteFileErrorCode:                   ErrCantDeleteFile,
		CantDeleteMessageErrorCode:                ErrCantDeleteMessage,
		CantEditErrorCode:                         ErrCantEdit,
		CantInviteErrorCode:                       ErrCantInvite,
		CantInviteSelfErrorCode:                   ErrCantInviteSelf,
		CantKickFromGeneralErrorCode:              ErrCantKickFromGeneral,
		CantKickFromLastChannelErrorCode:          ErrCantKickFromLastChannel,
		CantKickSelfErrorCode:                     ErrCantKickSelf,
		CantLeaveGeneralErrorCode:                 ErrCantLeaveGeneral,
		CantLeaveLastChannelErrorCode:             ErrCantLeaveLastChannel,
		CantUpdateMessageErrorCode:                ErrCantUpdateMessage,
		ChannelNotFoundErrorCode:                  ErrChannelNotFound,
		ComplianceExportsPreventDeletionErrorCode: ErrComplianceExportsPreventDeletion,
		EditWindowClosedErrorCode:                 ErrEditWindowClosed,
		FileCommentNotFoundErrorCode:              ErrFileCommentNotFound,
		FileDeletedErrorCode:                      ErrFileDeleted,
		FileNotFoundErrorCode:                     ErrFileNotFound,
		FileNotSharedErrorCode:                    ErrFileNotShared,
		GroupContainsOthersErrorCode:              ErrGroupContainsOthers,
		InvalidArrayArgErrorCode:                  ErrInvalidArrayArg,
		InvalidAuthErrorCode:                      ErrInvalidAuth,
		InvalidChannelErrorCode:                   ErrInvalidChannel,
		InvalidCharsetErrorCode:                   ErrInvalidCharset,
		InvalidClientIdErrorCode:                  ErrInvalidClientId,
		InvalidCodeErrorCode:                      ErrInvalidCode,
		InvalidFormDataErrorCode:                  ErrInvalidFormData,
		InvalidNameErrorCode:                      ErrInvalidName,
		InvalidPostTypeErrorCode:                  ErrInvalidPostType,
		InvalidPresenceErrorCode:                  ErrInvalidPresence,
		InvalidTimestampErrorCode:                 ErrInvalidTimestamp,
		InvalidTsLatestErrorCode:                  ErrInvalidTsLatest,
		InvalidTsOldestErrorCode:                  ErrInvalidTsOldest,
		IsArchivedErrorCode:                       ErrIsArchived,
		LastMemberErrorCode:                       ErrLastMember,
		LastRaChannelErrorCode:                    ErrLastRaChannel,
		MessageNotFoundErrorCode:                  ErrMessageNotFound,
		MigrationInProgressErrorCode:              ErrMigrationInProgress,
		MissingDurationErrorCode:                  ErrMissingDuration,
		MissingPostTypeErrorCode:                  ErrMissingPostType,
		MsgTooLongErrorCode:                       ErrMsgTooLong,
		NameTakenErrorCode:                        ErrNameTaken,
		NoChannelErrorCode:                        ErrNoChannel,
		NoCommentErrorCode:                        ErrNoComment,
		NoItemSpecifiedErrorCode:                  ErrNoItemSpecified,
		NoReactionErrorCode:                       ErrNoReaction,
		NoTextErrorCode:                           ErrNoText,
		NotArchivedErrorCode:                      ErrNotArchived,
		NotAuthedErrorCode:                        ErrNotAuthed,
		NotAuthorizedErrorCode:                    ErrNotAuthorized,
		NotEnoughUsersErrorCode:                   ErrNotEnoughUsers,
		NotInChannelErrorCode:                     ErrNotInChannel,
		NotInGroupErrorCode:                       ErrNotInGroup,
		NotPinnedErrorCode:                        ErrNotPinned,
		NotStarredErrorCode:                       ErrNotStarred,
		OverPaginationLimitErrorCode:              ErrOverPaginationLimit,
		PaidOnlyErrorCode:                         ErrPaidOnly,
		PermissionDeniedErrorCode:                 ErrPermissionDenied,
		PostingToGeneralChannelDeniedErrorCode:    ErrPostingToGeneralChannelDenied,
		RateLimitedErrorCode:                      ErrRateLimited,
		RequestTimeoutErrorCode:                   ErrRequestTimeout,
		RestrictedActionErrorCode:                 ErrRestrictedAction,
		SnoozeEndFailedErrorCode:                  ErrSnoozeEndFailed,
		SnoozeFailedErrorCode:                     ErrSnoozeFailed,
		SnoozeNotActiveErrorCode:                  ErrSnoozeNotActive,
		TooLongErrorCode:                          ErrTooLong,
		TooManyEmojiErrorCode:                     ErrTooManyEmoji,
		TooManyReactionsErrorCode:                 ErrTooManyReactions,
		TooManyUsersErrorCode:                     ErrTooManyUsers,
		UnknownErrorErrorCode:                     ErrUnknownError,
		UnknownTypeErrorCode:                      ErrUnknownType,
		UserDisabledErrorCode:                     ErrUserDisabled,
		UserDoesNotOwnChannelErrorCode:            ErrUserDoesNotOwnChannel,
		UserIsBotErrorCode:                        ErrUserIsBot,
		UserIsRestrictedErrorCode:                 ErrUserIsRestricted,
		UserIsUltraRestrictedErrorCode:            ErrUserIsUltraRestricted,
		UserNotFoundErrorCode:                     ErrUserNotFound,
		UserNotVisibleErrorCode:                   ErrUserNotVisible,
		UsersListNotSuppliedErrorCode:             ErrUsersListNotSupplied,
	}
)

func ErrorForErrorCode(code string) error {
	err, exists := errMap[code]
	if exists {
		return err
	}
	return fmt.Errorf("Unrecognized error: %s", err)
}
