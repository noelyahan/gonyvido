package main

import (
	"flag"
	"github.com/noelyahan/gonyvido/domain"
	gonyvido "github.com/noelyahan/gonyvido/api"
)

const (
	defaultUrl      = "https://youtu.be/FBHEF1nxbi0"
	defaultSavePath = "./"
	defaultQuality  = "high" // medium, low
)

func main() {
	url := flag.String("url", defaultUrl, "Video download url")
	savePath := flag.String("path", defaultSavePath, "Video download path")
	quality := flag.String("quality", defaultQuality, "Video download quality")
	toMp3 := flag.Bool("mp3", false, "y/n to download mp3")

	flag.Parse()

	var video *domain.Video

	switch *quality {
	case "high":
		video = gonyvido.GetHQVideo(*url).SetSavePath(*savePath).Download()
	case "medium":
		video = gonyvido.GetMQVideo(*url).SetSavePath(*savePath).Download()
	case "low":
		video = gonyvido.GetLQVideo(*url).SetSavePath(*savePath).Download()
	}

	if *toMp3 {
		video.ToMP3()
	}

}
