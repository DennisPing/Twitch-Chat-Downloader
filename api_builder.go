package main

import (
	"fmt"
	"net/url"
	"path"
	"regexp"
	"strconv"
)

// Build the video api request given the user input.
// User input could be just the video_id or the complete URL.
func buildGetVideoReq(input string) (string, error) {
	video_id, err := strconv.Atoi(input)
	if err == nil { // This is probably just the video_id
		req := fmt.Sprintf("https://api.twitch.tv/helix/videos?id=%d", video_id)
		return req, nil
	} else {
		if isTwitchVod(input) { // This is probably a Twitch VOD URL
			u, _ := url.Parse(input)
			video_id, err := strconv.Atoi(path.Base(u.Path))
			if err != nil {
				return "", ErrInvalidVodUrl
			}
			req := fmt.Sprintf("https://api.twitch.tv/helix/videos?id=%d", video_id)
			return req, nil
		} else {
			return "", ErrInvalidVodUrl
		}
	}
}

// Check if an input URL is a Twitch VOD.
func isTwitchVod(input string) bool {
	re := regexp.MustCompile(`twitch\.tv/videos/[0-9]+`)
	return re.MatchString(input)
}
