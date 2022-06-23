package main

import (
	"errors"
	"testing"
)

func TestBuildGetVideoReq(t *testing.T) {
	type test struct {
		input string
		exp   string
		err   error
	}
	tests := []test{
		{
			input: "1234567890",
			exp:   "https://api.twitch.tv/helix/videos?id=1234567890",
			err:   nil,
		},
		{
			input: "https://www.twitch.tv/videos/1234567890",
			exp:   "https://api.twitch.tv/helix/videos?id=1234567890",
			err:   nil,
		},
		{
			input: "https://www.twitch.tv/videos/1234567890/",
			exp:   "https://api.twitch.tv/helix/videos?id=1234567890",
			err:   nil,
		},
		{
			input: "https://www.twitch.tv/videos/1234567890?t=123",
			exp:   "https://api.twitch.tv/helix/videos?id=1234567890",
			err:   nil,
		},
		{
			input: "https://twitch.tv/videos/1234567890",
			exp:   "https://api.twitch.tv/helix/videos?id=1234567890",
			err:   nil,
		},
		{
			input: "https://twitch.tv/videos/1234567890/",
			exp:   "https://api.twitch.tv/helix/videos?id=1234567890",
			err:   nil,
		},
		{
			input: "https://twitch.tv/videos/1234567890?t=123",
			exp:   "https://api.twitch.tv/helix/videos?id=1234567890",
			err:   nil,
		},
		{
			input: "twitch.tv/videos/1234567890",
			exp:   "https://api.twitch.tv/helix/videos?id=1234567890",
			err:   nil,
		},
		{
			input: "twitch.tv/videos/1234567890?t=123",
			exp:   "https://api.twitch.tv/helix/videos?id=1234567890",
			err:   nil,
		},
		{
			input: "12345abcde",
			exp:   "",
			err:   ErrInvalidVodUrl,
		},
		{
			input: "https://www.twitch.tv/videos/12345abcde",
			exp:   "",
			err:   ErrInvalidVodUrl,
		},
		{
			input: "https://www.twitch.tv/videos/abcdejkgijk/",
			exp:   "",
			err:   ErrInvalidVodUrl,
		},
		{
			input: "https://www.twitch.tv/1234567890",
			exp:   "",
			err:   ErrInvalidVodUrl,
		},
		{
			input: "https://www.twitch.tv/videos",
			exp:   "",
			err:   ErrInvalidVodUrl,
		},
		{
			input: "https://www.twitch.com/videos/1234567890",
			exp:   "",
			err:   ErrInvalidVodUrl,
		},
		{
			input: "https://www.twitch.com/videos/1234567890/12345",
			exp:   "",
			err:   ErrInvalidVodUrl,
		},
		{
			input: "",
			exp:   "",
			err:   ErrInvalidVodUrl,
		},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got, err := buildGetVideoReq(tc.input)
			if err != nil && !errors.Is(err, tc.err) {
				t.Fatalf("got %v, exp %v", err, tc.err)
			}
			if got != tc.exp {
				t.Fatalf("got: %v, exp: %v", got, tc.exp)
			}
		})
	}
}

func TestIsTwitchVod(t *testing.T) {
	type test struct {
		input string
		exp   bool
	}
	good_urls := []test{
		{"https://www.twitch.tv/videos/1234567890", true},
		{"https://www.twitch.tv/videos/1234567890/", true},
		{"https://www.twitch.tv/videos/1234567890?t=123", true},
		{"https://twitch.tv/videos/1234567890", true},
		{"https://twitch.tv/videos/1234567890/", true},
		{"https://twitch.tv/videos/1234567890?t=123", true},
		{"www.twitch.tv/videos/1234567890", true},
		{"www.twitch.tv/videos/1234567890/", true},
		{"www.twitch.tv/videos/1234567890?t=123", true},
		{"twitch.tv/videos/1234567890", true},
		{"twitch.tv/videos/1234567890/", true},
		{"twitch.tv/videos/1234567890?t=123", true},
	}
	bad_urls := []test{
		{"https://www.twitch.com/videos/1234567890", false},
		{"https://www.twitch.com/videos/1234567890/", false},
		{"https://www.twitch.com/videos/1234567890?t=123", false},
		{"https://twitch.com/videos/1234567890", false},
		{"https://twitch.com/videos/1234567890/", false},
		{"https://twitch.com/videos/1234567890?t=123", false},
		{"www.twitch.com/videos/1234567890", false},
		{"www.twitch.com/videos/1234567890/", false},
		{"www.twitch.com/videos/1234567890?t=123", false},
		{"twitch.com/videos/1234567890", false},
		{"twitch.com/videos/1234567890/", false},
		{"twitch.com/videos/1234567890?t=123", false},
		{"https://www.twitch.tv/videos", false},
		{"https://www.twitch.tv", false},
		{"https://www.twitch.tv/1234567890", false},
	}
	for _, tc := range good_urls {
		if got := isTwitchVod(tc.input); got != tc.exp {
			t.Errorf("%s: got: %t, exp: %t", tc.input, got, tc.exp)
		}
	}
	for _, tc := range bad_urls {
		if got := isTwitchVod(tc.input); got != tc.exp {
			t.Errorf("%s: got: %t, exp[%t", tc.input, got, tc.exp)
		}
	}

}
