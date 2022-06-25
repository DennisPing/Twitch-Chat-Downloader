package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nicklaw5/helix"
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
	videoID, err := getVideoID(input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if err := LoadConfig(); err != nil {
		fmt.Printf("Error loading config file: %v\n", err)
		os.Exit(1)
	}
	opt := &helix.Options{
		ClientID:       os.Getenv("TCD_CLIENT_ID"),
		AppAccessToken: os.Getenv("TCD_API_TOKEN"),
	}
	client, err := helix.NewClient(opt)
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}
	params := &helix.VideosParams{
		IDs: []string{videoID},
	}
	resp, err := client.GetVideos(params)
	if resp.StatusCode != 200 {
		fmt.Printf("Error: %d - %v\n", resp.ErrorStatus, resp.ErrorMessage)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", resp.Data)
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
