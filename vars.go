package main

import "errors"

var Verbose bool
var Channel string
var Video_id string

var ErrInvalidVodUrl = errors.New("invalid vod url")
