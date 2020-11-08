package main

import (
	"MusicHelper/youtube"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/url"
)

var videoUrl = "https://www.youtube.com/watch?v=5AHor0kj31M"

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	newConfig := NewConfig()

	parse, err := url.Parse(videoUrl)
	if err != nil {
		log.Fatalf("can't parse video url %v", err)
	}

	videoId := parse.Query().Get("v")

	youtubeModel := youtube.CreateYoutubeModel(newConfig.YoutubeToken)

	videoInfo := youtubeModel.GetVideoInfo(videoId)
	videoInfo.GetTrackList()

	fmt.Println(videoInfo)
}
