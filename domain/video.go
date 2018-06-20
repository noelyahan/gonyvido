package domain

import (
	s "strings"
	"fmt"
	"github.com/gosuri/uiprogress"
	"os"
	"net/http"
	"io"
	"log"
	"os/exec"
)

type Video struct {
	title     string
	author    string
	quality   string
	videoType string
	url       string

	playerScript     string
	meta             map[string]interface{}
	videoLength      int
	downloadedLength int
	downloadPro      chan float64
	savePath         string
	done             chan struct{}
}

func (v *Video) getUrl() string {
	return v.url
}

func (v *Video) SetUrl(url string) {
	v.url = url
}

func (v *Video) GetMeta() map[string]interface{} {
	return v.meta
}

func (v *Video) GetPlayerScript() string {
	return v.playerScript
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

func NewVideo(t, a, q, vt, url, script string) Video {
	return Video{
		t,
		a,
		q,
		vt,
		url,
		script,
		make(map[string]interface{}),
		0,
		0,
		make(chan float64),
		"./",
		make(chan struct{}),
	}
}

func (v Video) showProgress() {
	fmt.Println(`Downloading:	` + v.GetTitle() + " " + v.GetQuality())
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
	uiprogress.Stop()

	fmt.Println(`Finished:	` + v.GetTitle() + ` quality: ` + v.GetQuality())
}

func (v *Video) Write(b []byte) (n int, err error) {
	v.downloadedLength = v.downloadedLength + len(b)
	p := 100 / (float64(v.videoLength) / float64(v.downloadedLength))
	v.downloadPro <- p
	if p == 100 {
		close(v.downloadPro)
		close(v.done)
	}
	return len(b), nil
}

func (v *Video) Download() *Video {
	//log.Panic(v.getUrl())

	// this has to have a public chan to notify the download is done
	go func() {
		resp, err := http.Get(v.getUrl())
		if err != nil {
			log.Fatalln("Please check your internet connection", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			log.Fatalln("Download failed")
		}
		v.videoLength = int(resp.ContentLength)

		if _, err := os.Stat(v.getSavePath()); err != nil {
			os.MkdirAll(v.getSavePath(), os.ModePerm)
		}

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

	return v
}

func (v *Video) ToMP3() {
	<-v.done
	mp3 := v.getFileName() + ".mp3"
	mp4 := v.getFileName() + v.getExt()
	//file := v.savePath + mp4
	ffmpeg, err := exec.LookPath("ffmpeg")
	if err != nil {
		log.Fatal("ffmpeg not found")
	}
	fmt.Println(`Converting:	` + v.GetTitle() + ` to mp3`)
	cmd := exec.Command(ffmpeg, "-y", "-loglevel", "quiet", "-i", v.getSavePath()+mp4, "-b:a", "320K", "-vn", v.getSavePath()+mp3)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to convert the mp3: ", err)
	}
	fmt.Println(`Finished:	` + v.GetTitle() + `.mp3`)
}

func removeFile(file string) {
	err := os.Remove(file)
	if err != nil {
		fmt.Println("Could not delete: ", file, err)
	}
}
