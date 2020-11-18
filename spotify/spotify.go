package spotify

import (
	"encoding/json"
	"fmt"
	"github.com/zmb3/spotify"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"strings"
	"time"
)

func CreateSpotifyModel(accessToken string, logger *zap.Logger) *spotifyModel {
	token := &oauth2.Token{
		AccessToken:  accessToken,
		TokenType:    "Bearer",
		RefreshToken: "",
		Expiry:       time.Time{},
	}

	return &spotifyModel{
		logger: logger,
		client: spotify.Authenticator{}.NewClient(token),
	}
}

type spotifyModel struct {
	logger *zap.Logger
	client spotify.Client
}

func (s *spotifyModel) SearchTrack(trackName string, extended bool) *spotify.FullTrack {
	results, err := s.client.Search(trackName, spotify.SearchTypeTrack)
	if err != nil {
		s.logger.Error("can't search track", zap.Error(err), zap.String("track", trackName))
		return nil
	}
	const extendedSubstr = "extended"
	var result *spotify.FullTrack
	for i, track := range results.Tracks.Tracks {
		if i == 0 {
			result = &track
			if !extended {
				break
			}
		}
		if strings.Contains(strings.ToLower(track.Name), extendedSubstr) {
			result = &track
			break
		}
	}
	return result
}

func (s *spotifyModel) CreatePlaylist(userId, playlistName string) *spotify.FullPlaylist {
	playlist, err := s.client.CreatePlaylistForUser(userId, playlistName, "", true)
	if err != nil {
		s.logger.Error("can't create playlist", zap.Error(err), zap.String("userId", userId),
			zap.String("playlistName", playlistName))
		return nil
	}
	marshal, _ := json.Marshal(playlist)
	fmt.Println(string(marshal))
	return playlist
}

func (s *spotifyModel) AddTrackToPlaylist(playlistId, trackId spotify.ID) string {
	snapshotId, err := s.client.AddTracksToPlaylist(playlistId, trackId)
	if err != nil {
		s.logger.Error("can't add trackId to playlist", zap.Error(err), zap.String("playlistId", playlistId.String()),
			zap.String("trackId", trackId.String()))
		return ""
	}
	return snapshotId
}
