package slk

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestTime(t *testing.T) {

	vals := [][]string{
		{"1401383885.000061", "2014-05-29 13:18:05.000000061 -0400 EDT"},
		{`"1401383885.000061"`, "2014-05-29 13:18:05.000000061 -0400 EDT"},
		{"1401383885", "2014-05-29 13:18:05 -0400 EDT"},
		{`"1401383885"`, "2014-05-29 13:18:05 -0400 EDT"},
		{"2012-04-23T18:25:43.511Z", "2012-04-23 18:25:43.511 +0000 UTC"},
		{`"2012-04-23T18:25:43.511Z"`, "2012-04-23 18:25:43.511 +0000 UTC"},
	}

	for _, val := range vals {
		tv := Time{}
		if err := tv.UnmarshalJSON([]byte(val[0])); err != nil {
			t.Error(err)
		}
		if tv.String() != val[1] {
			t.Errorf("Expected %q, got %q", val[1], tv.String())
		}
	}

}

func TestSlackTime(t *testing.T) {

	vals := []string{
		`1401383884.000061`,
		`1401383885.000062`,
		`1401383885.000063`,
		`1401383886.111164`,
	}

	for _, val := range vals {
		tv := Time{}
		if err := tv.UnmarshalJSON([]byte(val)); err != nil {
			t.Error(err)
		}
		if tv.SlackString() != val {
			t.Errorf("Expected %q, got %q", val, tv.SlackString())
		}
	}

}

func TestUnmarshal(t *testing.T) {

	const valjson = `{"ok":true,"snooze_endtime":1450373897,"snooze_remaining":60}`

	type testStruct struct {
		OK              bool     `json:"ok"`
		SnoozeEndTime   Time     `json:"snooze_endtime"`
		SnoozeRemaining Duration `json:"snooze_remaining"`
	}

	var val testStruct

	if err := json.NewDecoder(strings.NewReader(valjson)).Decode(&val); err != nil {
		fmt.Println(err)
	}

	var expected = testStruct{
		OK:              true,
		SnoozeEndTime:   Time{time.Unix(1450373897, 0)},
		SnoozeRemaining: Duration{60 * time.Second},
	}

	if val.OK != expected.OK {
		t.Errorf("Expected: %#v, got: %#v\n", expected.OK, val.OK)
	}
	if val.SnoozeEndTime != expected.SnoozeEndTime {
		t.Errorf("Expected: %#v, got: %#v\n", expected.SnoozeEndTime, val.SnoozeEndTime)
	}
	if val.SnoozeRemaining != expected.SnoozeRemaining {
		t.Errorf("Expected: %#v, got: %#v\n", expected.SnoozeRemaining, val.SnoozeRemaining)
	}

}
