package youtube

import (
	"strings"
)

type Video struct {
	BasicInfo
	Items []Item
}

type BasicInfo struct {
	Kind string
	Etag string
}

type Item struct {
	BasicInfo
	Snippet Snippet
}

type Snippet struct {
	Title       string
	Description string
	ChannelId   string
	Id          string
}

func (v *Video) GetTrackList() []string {
	res := make([]string, 0)
	for _, item := range v.Items {
		index := strings.Index(item.Snippet.Description, "Tracklist")
		split := strings.Split(item.Snippet.Description[index:], "\n")
		if len(split) > 1 {
			return split[1:]
		}
	}
	return res
}

func (v *Video) GetTitle() string {
	if len(v.Items) == 0 {
		return ""
	}
	return v.Items[0].Snippet.Title
}