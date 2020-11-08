package youtube

import (
	"MusicHelper/webClient"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type youtubeModel struct {
	accessToken string
	endpoints map[string]string
}

func CreateYoutubeModel(token string) *youtubeModel {
	return &youtubeModel{accessToken:token, endpoints:endpoints}
}

func (y *youtubeModel) GetVideoInfo(videoId string) video {
	var newVideo video

	var params = map[string]string{
		"id": videoId,
		"part":"snippet",
		"key":y.accessToken,
	}

	respBody, err := webClient.MakeGetRequest(endpoints["video"], params)
	defer func() {
		err := respBody.Close()
		if err != nil {
			log.Printf("can't close resp body %v", err)
		}
	}()
	if err != nil {
		log.Printf("can't get video info %v", err)
		return newVideo
	}

	body, err := ioutil.ReadAll(respBody)
	if err != nil {
		log.Printf("can't parse resp body %v", err)
		return newVideo
	}
	fmt.Println(string(body))
	err = json.Unmarshal(body, &newVideo)
	if err != nil {
		log.Printf("can't parse resp body %v", err)
		return newVideo
	}
	return newVideo
}