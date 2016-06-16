package slk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"strings"
)

type (
	Client struct {
		client *http.Client
		token  string
		Debug  bool
	}
)

const (
	SlackAPIURL = "https://slack.com/api"
	Version     = "slk 0.1"
)

func NewClient(token string) *Client {
	return &Client{
		client: &http.Client{},
		token:  token,
	}
}

func (c *Client) Do(method Method, args url.Values, v interface{}) error {
	var body io.Reader
	if args != nil {
		body = strings.NewReader(args.Encode())
	}
	return c.DoWithBody(method, "application/x-www-form-urlencoded;charset=UTF-8", body, v)
}

func (c *Client) DoWithBody(method Method, content_type string, body io.Reader, v interface{}) error {

	url_string := fmt.Sprintf("%s/%s?token=%s", SlackAPIURL, method, url.QueryEscape(c.token))

	req, err := http.NewRequest("POST", url_string, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", content_type)
	req.Header.Set("User-Agent", Version)

	if c.Debug {
		data, _ := httputil.DumpRequestOut(req, true)
		fmt.Println("REQUEST:", string(data)+"\n")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	body_bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if c.Debug {
		data, _ := httputil.DumpResponse(resp, false)
		fmt.Println("RESPONSE:", string(data))
		fmt.Println(string(body_bytes) + "\n")
	}

	if err := json.Unmarshal(body_bytes, v); err != nil {
		fmt.Println("ERROR:", method, string(body_bytes)+"\n")
		return err
	}

	// little magic to sniff out OK=false and error=something
	vv := reflect.ValueOf(v)

	// dereference if pointer (should be)
	if vv.Kind() == reflect.Ptr && !vv.IsNil() {
		vv = vv.Elem()
	}
	// if struct (should be)
	if vv.Kind() == reflect.Struct {
		if okv := vv.FieldByName("OK"); okv.Kind() == reflect.Bool && !okv.Bool() {
			if errv := vv.FieldByName("Error"); errv.Kind() == reflect.String {
				return ErrorForErrorCode(errv.String())
			}
			return errors.New("`OK` == false but missing field `Error`")
		}
	}

	return nil

}

var escaper = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	">", "&gt;",
)

func EscapeMessage(text string) string {
	return escaper.Replace(text)
}
