package main

import (
	"flag"
	gonyvido "github.com/noelyahan/gonyvido/api"
)

var url string
var savePath string
var quality string

func init() {
	const (
		defaultUrl      = "https://www.youtube.com/watch?v=PIyrV8UmNqg"
		defaultSavePath = "./"
		defaultQuality  = "high" // medium, low
	)
	// url flags
	flag.StringVar(&url, "url", defaultUrl, "Video download url")
	flag.StringVar(&url, "u", defaultUrl, "Video download url")

	// save path flags
	flag.StringVar(&savePath, "path", defaultSavePath, "Video save path")
	flag.StringVar(&savePath, "p", defaultSavePath, "Video save path")

	// quaity flags
	flag.StringVar(&quality, "quality", defaultQuality, "Video quality")
	flag.StringVar(&quality, "q", defaultQuality, "Video quality")
}

func main() {

	flag.Parse()

	switch quality {
	case "high":
		gonyvido.GetHQVideo(url).SetSavePath(savePath).Download().ShowProgress()
	case "medium":
		gonyvido.GetMQVideo(url).SetSavePath(savePath).Download().ShowProgress()
	case "low":
		gonyvido.GetLQVideo(url).SetSavePath(savePath).Download().ShowProgress()
	}

}
