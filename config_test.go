package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
	config_path := "./config.yml"
	config, err := parseConfig(config_path)
	if err != nil {
		t.Fatalf("unable to parse config: %v", err)
	}
	assert.Equal(t, "dummyId", config.ClientId)
	assert.Equal(t, "dummyToken", config.ApiToken)
	assert.Equal(t, "!", config.Format.Comments.Badges.Admin)
	assert.Equal(t, "$", config.Format.Comments.Badges.Bits)
	assert.Equal(t, "~", config.Format.Comments.Badges.Broadcaster)
	assert.Equal(t, "*", config.Format.Comments.Badges.GlobalMod)
	assert.Equal(t, "@", config.Format.Comments.Badges.Moderator)
	assert.Equal(t, "+", config.Format.Comments.Badges.Premium)
	assert.Equal(t, "&", config.Format.Comments.Badges.Staff)
	assert.Equal(t, "%", config.Format.Comments.Badges.Subscriber)
	assert.Equal(t, "+", config.Format.Comments.Badges.Turbo)
	assert.Equal(t, "[{timestamp[relative]}] <{commenter[badge]}{commenter[display_name]}> {message[body]}", config.Format.Comments.Format)
	assert.Equal(t, "%X", config.Format.Comments.Timestamp.Relative)
	assert.Equal(t, "{id}.txt", config.Format.Filename.Format)
	assert.Equal(t, "%x", config.Format.Filename.Timestamp.Absolute)
}
