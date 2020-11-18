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
	config.SpotifyClient = os.Getenv("SPOTIFY_CLIENT")
	config.SpotifySecret = os.Getenv("SPOTIFY_SECRET")
	return config
}
