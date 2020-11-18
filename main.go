package main

import (
	"MusicHelper/spotify"
	"MusicHelper/youtube"
	"fmt"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"net/url"
	"strings"
)

var videoUrl = "https://www.youtube.com/watch?v=5AHor0kj31M"
var spotifyToken = "BQBlCkKAm44VqhrXauMZDKrk4rp_l_JfVXlMIKO7WsdUxAEsYpbsNGL96zj5cH1x9QCgZeGAQZk2-ODCSKRlyb3KE3Doaf1IlDahu8KstJrC4QZGERPPcBpa6P3arKnLupTDg5Tjfe1IiSEmXUuLcrVLuJEKFJZqm2BDruY2zoonaht1ZJgsZ90FIfcOIj7B73UaNA"
var spotifyUserId = "r8gnh1qh6d2yp4uzcsbhvh65l"

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

type songS struct {
	artist string
	song string
}

func main() {
	newConfig := NewConfig()
	parse, err := url.Parse(videoUrl)
	if err != nil {
		log.Fatalf("can't parse video url %v", err)
	}
	videoId := parse.Query().Get("v")
	video := getVideo(videoId, newConfig.YoutubeToken)

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	spotifyModel := spotify.CreateSpotifyModel(spotifyToken, logger)

	playlist := spotifyModel.CreatePlaylist(spotifyUserId, video.GetTitle() + " extended")
	if playlist == nil {
		log.Fatal("playlist is empty")
	}
	for _, trackStr := range video.GetTrackList() {
		trackStr = sanitize(trackStr[2:])
		//splitted := strings.Split(trackStr, "-")
		//if len(splitted) < 2 {
		//	logger.Error("can't split track on artist and song name")
		//	continue
		//}
		//song := &songS{
		//	artist: splitted[0],
		//	song:   splitted[1],
		//}
		track := spotifyModel.SearchTrack(trackStr, true)
		if track == nil {
			logger.Error("track is empty", zap.String("trackStr", trackStr))
			continue
		}
		snapId := spotifyModel.AddTrackToPlaylist(playlist.ID, track.ID)
		if len(snapId) != 0 {
			logger.Info("track added", zap.String("name", track.Name), zap.String("playlist", playlist.Name))
		}
	}


}

func sanitize(song string) string {
	rpl := strings.NewReplacer("ft.", "",
		"&", "", "-", "")
	return rpl.Replace(song)
}

func searchSong(songName string) {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	spotifyModel := spotify.CreateSpotifyModel(spotifyToken, logger)
	track := spotifyModel.SearchTrack(songName, false)
	fmt.Println(track.Name)

}

//func createPlaylist(playlistName string) {
//	spotifyModel := spotify.CreateSpotifyModel(spotifyToken)
//
//	newPlaylist := spotifyModel.CreatePlaylist(spotifyUserId, playlistName)
//	if newPlaylist != nil {
//		fmt.Printf("%v", *newPlaylist)
//	}
//}

func getVideo(videoId, youtubeToken string) youtube.Video {
	youtubeModel := youtube.CreateYoutubeModel(youtubeToken)
	videoInfo := youtubeModel.GetVideoInfo(videoId)
	return videoInfo
}

