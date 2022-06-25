package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
	config_path := "./config.yml"
	config, err := ParseConfig(config_path)
	if err != nil {
		t.Fatalf("unable to parse config: %v", err)
	}
	assert.Equal(t, "dummyId", config.Client_id)
	assert.Equal(t, "dummyToken", config.Api_token)
}
