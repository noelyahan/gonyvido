package gonyvido

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	s "strings"
)

const (
	YoutubeVideoInfoApi = "http://youtube.com/get_video_info?video_id="
	YoutubeWatchUrl     = "https://www.youtube.com/watch?v"
	YoutubeShortUrl     = "youtu.be"
	youtubeBaseURL      = "https://www.youtube.com/"
)

func GetYoutubeVideos(url string) ([]Video, error) {
	var videoId string
	splitted := s.Split(url, "=")
	if len(splitted) > 1 {
		videoId = splitted[1]
	}
	if s.Contains(url, YoutubeShortUrl) {
		splitted = s.Split(url, "/")
		if len(splitted) > 3 {
			videoId = splitted[3]
		}
	}
	if s.Contains(videoId, "&") {
		videoId = s.Split(videoId, "&")[0]
	}
	//infoApiUrl := YoutubeVideoInfoApi + videoId
	defaultApiUrl := YoutubeWatchUrl + "=" + videoId

	//videos, err := getVideoInfo(infoApiUrl)
	//if err != nil {
	videos, err := getVideoInfo(defaultApiUrl)
	if err != nil {
		fmt.Println("Sorry this video is cannot be download.", err)
		return nil, err
	}
	return videos, nil
	//}
	//return videos, nil
}

func getVideoInfo(url string) ([]Video, error) {
	// TODO add the url parser
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Stratergy pattern
	var title, author, streamMapStr, playerScript string
	if s.Contains(url, YoutubeVideoInfoApi) {
		title, author, streamMapStr, err = extractFromInfoApi(string(body))
	} else {
		title, author, streamMapStr, playerScript, err = extractFromDefault(string(body))
	}
	if err != nil {
		return nil, err
	}

	return constructVideos(title, author, streamMapStr, playerScript)
}

func extractFromDefault(strBody string) (string, string, string, string, error) {
	streamMapFilter := "url_encoded_fmt_stream_map\":"
	streamMapRegexp, err := regexp.Compile(streamMapFilter + "(.*?)\",")
	if err != nil {
		return "", "", "", "", err
	}
	titleRegexp, err := regexp.Compile("<title>(.*?)<\\/title>")
	if err != nil {
		return "", "", "", "", err
	}
	authorRegexp, err := regexp.Compile("author\":\"(.*?)\",")
	if err != nil {
		return "", "", "", "", err
	}
	playerScriptRegexp, err := regexp.Compile(`src="(.*?)base.js`)
	if err != nil {
		return "", "", "", "", err
	}
	title := titleRegexp.FindString(strBody)
	author := authorRegexp.FindString(strBody)
	title = title[len("<title>") : len(title)-len("</title>")]
	if len(author) > 0 {
		author = author[len("author:\"")+1 : len(author)-2]
	}
	playerScript := s.Replace(playerScriptRegexp.FindString(strBody), `src="`, "", -1)

	streamMapStr := streamMapRegexp.FindString(strBody)

	if len(streamMapStr) > 0 {
		streamMapStr = streamMapStr[len(streamMapFilter)+1 : len(streamMapStr)-2]
		streamMapStr = s.Replace(streamMapStr, "\\u0026", "&", -1)
	}

	return title, author, streamMapStr, playerScript, err
}

func extractFromInfoApi(strBody string) (string, string, string, error) {
	bodyMap, err := url.ParseQuery(strBody)
	if err != nil {
		return "", "", "", err
	}
	if bodyMap["status"][0] == "fail" {
		return "", "", "", errors.New(bodyMap["reason"][0])
	}
	var title, author string
	if _, exist := bodyMap["title"]; exist {
		title = bodyMap["title"][0]
	}
	if _, exist := bodyMap["author"]; exist {
		author = bodyMap["author"][0]
	}
	streamMapStr := bodyMap["url_encoded_fmt_stream_map"][0]
	return title, author, streamMapStr, nil
}

func constructVideos(title, author, streamMapStr, playerScript string) ([]Video, error) {
	streamList := s.Split(streamMapStr, ",")
	videos := make([]Video, 0)
	for _, streamItem := range streamList {
		stream, err := url.ParseQuery(streamItem)
		if err != nil {
			return nil, err
		}
		v := NewVideo(
			title,
			author,
			stream["quality"][0],
			stream["type"][0],
			stream["url"][0],
			playerScript,
		)
		for k, val := range stream {
			v.GetMeta()[k] = val[0]
		}
		url, _ := getDownloadURL(v.GetMeta(), v.GetPlayerScript())
		v.SetUrl(url.String())
		videos = append(videos, v)
	}
	return videos, nil
}
