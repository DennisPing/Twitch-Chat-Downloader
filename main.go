package main

import (
	"flag"
	"fmt"
	"os"
)

func cliUsage() {
	fmt.Printf("Usage: %s [-v] URL\n", os.Args[0])
	fmt.Printf("Options:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = cliUsage
	input := parseArgs()
	// This input could either be a complete URL or a video ID
	get_video_req, err := buildGetVideoReq(input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(get_video_req)
	// https://www.reddit.com/r/PHP/comments/gtpm5r/what_are_the_benefits_of_using_env_over_envjson/
	LoadConfig()
}

// Parse the args and optional flags and return sanitized input
func parseArgs() string {
	flag.BoolVar(&Verbose, "v", false, "verbose output")
	flag.Parse()

	// Validate input args
	if flag.NArg() == 1 {
		return flag.Arg(0)
	} else if flag.NArg() < 1 {
		fmt.Println("Error: missing URL")
		cliUsage()
		os.Exit(1)
	} else {
		fmt.Printf("Error: too many arguments. Expected 1, got %d\n", flag.NArg())
		cliUsage()
		os.Exit(1)
	}
	return ""
}
