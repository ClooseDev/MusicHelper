package main

import "os"

type config struct {
	SpotifyClient string
	SpotifySecret string

	YoutubeToken string
}

func NewConfig() *config {
	config := &config{}
	config.YoutubeToken = os.Getenv("YOUTUBE_ACCESS_TOKEN")
	return config
}
