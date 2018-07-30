package gonyvido

import (
	"fmt"
	s "strings"
)

func find(url string) []Video {
	videos, err := GetYoutubeVideos(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return videos
}

func GetHQVideo(url string) *Video {
	return getFilteredVideo(url, "hd")
}

func GetMQVideo(url string) *Video {
	return getFilteredVideo(url, "medium")
}

func GetLQVideo(url string) *Video {
	return getFilteredVideo(url, "small")
}

func getFilteredVideo(url, vtype string) *Video {
	videos := find(url)
	for _, video := range videos {
		if s.Contains(video.GetQuality(), vtype) {
			return &video
		}
	}
	v := Video{}
	return &v
}
