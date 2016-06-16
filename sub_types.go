package slk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// types used throughout

type (
	Message struct {
		Type      string `json:"type"`
		IsStarred bool   `json:"is_starred"`
		Reactions []struct {
			Count int      `json:"count"`
			Name  string   `json:"name"`
			Users []string `json:"users"`
		} `json:"reactions"`
		Text      string `json:"text"`
		Timestamp Time   `json:"ts"`
		User      string `json:"user"`
	}

	ResultMessage struct {
		Type    string `json:"type"`
		Channel struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"channel"`
		Next struct {
			Text      string `json:"text"`
			Timestamp Time   `json:"ts"`
			Type      string `json:"type"`
			User      string `json:"user"`
			Username  string `json:"username"`
		} `json:"next"`
		Next2 struct {
			Text      string `json:"text"`
			Timestamp Time   `json:"ts"`
			Type      string `json:"type"`
			User      string `json:"user"`
			Username  string `json:"username"`
		} `json:"next_2"`
		Permalink string `json:"permalink"`
		Previous  struct {
			Text      string `json:"text"`
			Timestamp Time   `json:"ts"`
			Type      string `json:"type"`
			User      string `json:"user"`
			Username  string `json:"username"`
		} `json:"previous"`
		Previous2 struct {
			Text      string `json:"text"`
			Timestamp Time   `json:"ts"`
			Type      string `json:"type"`
			User      string `json:"user"`
			Username  string `json:"username"`
		} `json:"previous_2"`
		Text      string `json:"text"`
		Timestamp Time   `json:"ts"`
		User      string `json:"user"`
		Username  string `json:"username"`
	}

	Comment struct {
		ID        string `json:"id"`
		Created   Time   `json:"created"`
		Timestamp Time   `json:"timestamp"`
		User      string `json:"user"`
		Comment   string `json:"comment"`
	}

	// TODO: something
	Item struct {
		Type      string          `json:"type"`
		Channel   json.RawMessage `json:"channel"`
		Message   json.RawMessage `json:"message"`
		File      json.RawMessage `json:"file"`
		Comment   json.RawMessage `json:"comment"`
		Created   Time            `json:"created"`
		Reactions []struct {
			Name  string   `json:"name"`
			Count int      `json:"count"`
			Users []string `json:"users"`
		} `json:"reactions"`
		Timestamp Time `json:"ts"`
	}

	Team struct {
		ID          string `json:"id"`
		Domain      string `json:"domain"`
		EmailDomain string `json:"email_domain"`
		Icon        struct {
			Image102     string `json:"image_102"`
			Image132     string `json:"image_132"`
			Image34      string `json:"image_34"`
			Image44      string `json:"image_44"`
			Image68      string `json:"image_68"`
			Image88      string `json:"image_88"`
			ImageDefault bool   `json:"image_default"`
		} `json:"icon"`
		MsgEditWindow    Duration               `json:"msg_edit_window_mins"`
		Name             string                 `json:"name"`
		OverStorageLimit bool                   `json:"over_storage_limit"`
		Plan             string                 `json:"plan"`
		Prefs            map[string]interface{} `json:"prefs"`
	}

	Login struct {
		Count     int    `json:"count"`
		Country   string `json:"country"`
		DateFirst Time   `json:"date_first"`
		DateLast  Time   `json:"date_last"`
		IP        string `json:"ip"`
		Isp       string `json:"isp"`
		Region    string `json:"region"`
		UserAgent string `json:"user_agent"`
		UserID    string `json:"user_id"`
		Username  string `json:"username"`
	}

	LogEntry struct {
		ChangeType  string `json:"change_type"`
		Channel     string `json:"channel"`
		Date        Time   `json:"date"`
		Scope       string `json:"scope"`
		ServiceID   string `json:"service_id"`
		ServiceType string `json:"service_type"`
		UserID      string `json:"user_id"`
		UserName    string `json:"user_name"`
	}

	FileType string
	Presence string

	Time struct {
		time.Time
	}
	Duration struct {
		time.Duration
	}
)

const (
	FileTypeAll      FileType = `all`      // All files
	FileTypePosts    FileType = `posts`    // Posts
	FileTypeSnippets FileType = `snippets` // Snippets
	FileTypeImages   FileType = `images`   // Image files
	FileTypeGdocs    FileType = `gdocs`    // Google docs
	FileTypeZips     FileType = `zips`     // Zip files
	FileTypePdfs     FileType = `pdfs`     // PDF files

	PresenceAuto Presence = `auto`
	PresenceAway Presence = `away`

	ItemTypeChannel     = `channel`
	ItemTypeFile        = `file`
	ItemTypeFileComment = `file_comment`
	ItemTypeGroup       = `group`
	ItemTypeIM          = `im`
	ItemTypeMessage     = `message`
)

var unixNanoRegexp = regexp.MustCompile(`^\d+(?:\.\d+)?$`)

func (p *Presence) UnmarshallJSON(data []byte) error {
	data = bytes.Trim(data, `"`)
	*p = Presence(string(data))
	return nil
}

func NewTime(t time.Time) Time {
	return Time{t}
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.SlackString() + `"`), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {

	var err error

	data = bytes.Trim(data, `"`)

	if unixNanoRegexp.Match(data) {
		var sec, nsec int64
		if i := bytes.IndexByte(data, '.'); i > -1 {
			sec, _ = strconv.ParseInt(string(data[:i]), 10, 64)
			nsec, _ = strconv.ParseInt(string(data[i+1:]), 10, 64)
		} else {
			sec, _ = strconv.ParseInt(string(data), 10, 64)
		}
		t.Time = time.Unix(sec, nsec)
	} else {
		t.Time, err = time.Parse(time.RFC3339, string(data))
	}

	return err
}

func (t Time) SlackString() string {
	u := time.Duration(t.Time.UnixNano())
	return fmt.Sprintf("%d.%06d", u/time.Second, u%time.Second)
}

func NewDuration(t time.Duration) Duration {
	return Duration{t}
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return []byte(d.SlackString()), nil
}

func (d *Duration) UnmarshalJSON(data []byte) error {
	sec, _ := strconv.ParseInt(string(data), 10, 64)
	d.Duration = time.Duration(sec) * time.Second
	return nil
}

func (d Duration) SlackString() string {
	return strconv.Itoa(int(d.Duration / time.Second))
}
