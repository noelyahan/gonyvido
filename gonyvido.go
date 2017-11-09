package gonyvido

import (
	"fmt"
	"github.com/gosuri/uiprogress"
	s "strings"
)

func Find(url string) []Video {
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

func (v Video) ShowProgress() {
	fmt.Println("Downloading: " + v.GetTitle() + " " + v.GetQuality())
	uiprogress.Start()
	bar := uiprogress.AddBar(100)
	bar.AppendCompleted()
	bar.PrependElapsed()
	pVal := 0
	for p := range v.DownloadP {
		if pVal != p {
			diff := p - pVal
			for i := 0; i < diff; i++ {
				bar.Incr()
			}
			pVal = p
		}
	}
	fmt.Println("Finished: " + v.GetTitle() + " quality: " + v.GetQuality())
}

func getFilteredVideo(url, vtype string) *Video {
	videos := Find(url)
	for _, video := range videos {
		if s.Contains(video.GetQuality(), vtype) {
			return &video
		}
	}
	return &Video{"", "", "", "", "", 0, 0, nil, "./"}
}
