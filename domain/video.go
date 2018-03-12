package domain

import (
	s "strings"
	"fmt"
	"github.com/gosuri/uiprogress"
	"os"
	"net/http"
	"io"
)

type Video struct {
	title     string
	author    string
	quality   string
	videoType string
	url       string

	videoLength      int
	downloadedLength int
	downloadPro      chan float64
	savePath         string
}

func (v *Video) getUrl() string {
	return v.url
}

func (v Video) getExt() string {
	ext := s.Split(v.videoType, ";")[0]
	ext = s.Split(ext, "/")[1]
	return "." + ext
}

func (v Video) getFileName() string {
	fileName := s.Replace(v.title, "|", "", -1)
	return fileName
}

func (v Video) getSavePath() string {
	return v.savePath
}

func (v Video) GetType() string {
	return v.videoType
}

func (v Video) GetTitle() string {
	return v.title
}

func (v Video) GetAuthor() string {
	return v.author
}

func (v Video) GetQuality() string {
	return v.quality
}

func (v *Video) SetSavePath(savePath string) *Video {
	v.savePath = savePath
	if string(savePath[len(savePath)-1]) != "/" {
		v.savePath += "/"
	}
	return v
}

func NewVideo(t, a, q, vt, url string) Video {
	return Video{
		t,
		a,
		q,
		vt,
		url,
		0,
		0,
		make(chan float64),
		"./",
	}
}

func (v Video) showProgress() {
	fmt.Println("Downloading: " + v.GetTitle() + " " + v.GetQuality())
	uiprogress.Start()
	bar := uiprogress.AddBar(100)
	bar.AppendCompleted()
	bar.PrependElapsed()
	pVal := 0
	for p := range v.downloadPro {
		if pVal != int(p) {
			bar.Incr()
			pVal = int(p)
			if pVal >= 98 {
				for i := pVal; i < 100; i++ {
					bar.Incr()
				}
			}
		}
	}

	fmt.Println("Finished: " + v.GetTitle() + " quality: " + v.GetQuality())
}

func (v *Video) Write(b []byte) (n int, err error) {
	v.downloadedLength = v.downloadedLength + len(b)
	p := 100 / (float64(v.videoLength) / float64(v.downloadedLength))
	v.downloadPro <- p
	if p == 100 {
		close(v.downloadPro)
	}
	return len(b), nil
}

func (v *Video) Download() Video {
	// this has to have a public chan to notify the download is done
	go func() {
		resp, err := http.Get(v.getUrl())
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			fmt.Println(err)
		}
		v.videoLength = int(resp.ContentLength)

		out, err := os.Create(v.getSavePath() + v.getFileName() + v.getExt())
		if err != nil {
			fmt.Println(err)
		}
		mw := io.MultiWriter(out, v)
		_, err = io.Copy(mw, resp.Body)
		if err != nil {
			fmt.Println(err)
		}
	}()

	// show progress bar here
	v.showProgress()

	return Video{
		v.title,
		v.author,
		v.quality,
		v.videoType,
		v.url,
		v.videoLength,
		v.downloadedLength,
		v.downloadPro,
		"./",
	}
}
