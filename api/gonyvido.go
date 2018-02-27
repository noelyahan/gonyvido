package api

import (
	"fmt"
	s "strings"
	"github.com/noelyahan/gonyvido/domain"
	"github.com/noelyahan/gonyvido/youtube"
)

func Find(url string) []domain.Video {
	videos, err := youtube.GetYoutubeVideos(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return videos
}

func GetHQVideo(url string) *domain.Video {
	return getFilteredVideo(url, "hd")
}

func GetMQVideo(url string) *domain.Video {
	return getFilteredVideo(url, "medium")
}

func GetLQVideo(url string) *domain.Video {
	return getFilteredVideo(url, "small")
}


func getFilteredVideo(url, vtype string) *domain.Video {
	videos := Find(url)
	for _, video := range videos {
		if s.Contains(video.GetQuality(), vtype) {
			return &video
		}
	}
	v := domain.NewVideo("", "", "", "", "")
	return &v
}
