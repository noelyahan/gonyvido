package main

import (
	"flag"
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

	flag.Parse()

	switch *quality {
	case "high":
		gonyvido.GetHQVideo(*url).SetSavePath(*savePath).Download()
	case "medium":
		gonyvido.GetMQVideo(*url).SetSavePath(*savePath).Download()
	case "low":
		gonyvido.GetLQVideo(*url).SetSavePath(*savePath).Download()
	}

}
