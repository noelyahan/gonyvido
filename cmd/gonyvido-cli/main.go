package main

import (
	"flag"
	"github.com/noelyahan/gonyvido"
)

const (
	defaultUrl      = "https://youtu.be/FBHEF1nxbi0"
	defaultSavePath = "./"
	defaultQuality  = "high" // medium, low
)

func main() {
	url := flag.String("url", defaultUrl, "Video download url")
	savePath := flag.String("path", defaultSavePath, "Video download path")
	quality := flag.String("quality", defaultQuality, "Video download quality <high, medium, low>")
	toMp3 := flag.Bool("mp3", false, "Enable this flag for mp3 file")

	flag.Parse()

	var video *gonyvido.Video

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
