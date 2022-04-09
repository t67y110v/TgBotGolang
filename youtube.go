package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	YOUTUBE_SEARCH_URL = "https://www.googleapis.com/youtube/v3/search"
	YOUTUBE_API_TOKEN  = ""
	YOUTUBE_VIDEO_URL  = "https://www.youtube.com/watch?v="
)

func GetLast(channelUrl string) (string, error) {
	items, err := retriveVideos(channelUrl)
	if err != nil {
		return "nil", err
	}
	if len(items) < 1 {
		return "", errors.New("no video found")
	}
	return YOUTUBE_VIDEO_URL + items[0].Id.VideoId, nil
}

func retriveVideos(channelUrl string) ([]Item, error) {
	req, err := makeRequest(channelUrl, 1)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var restResponse RestResponse2
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Items, nil
}

func makeRequest(channelUrl string, maxRes int) (*http.Request, error) {
	lastSlashIndex := strings.LastIndex(channelUrl, "/")
	channelId := channelUrl[lastSlashIndex+1:]
	req, err := http.NewRequest("GET", YOUTUBE_SEARCH_URL, nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Add("part", "id")
	query.Add("channelId", channelId)
	query.Add("maxResults", strconv.Itoa(maxRes))
	query.Add("order", "date")
	query.Add("key", YOUTUBE_API_TOKEN)
	req.URL.RawQuery = query.Encode()
	fmt.Println(req.URL.String())
	return req, nil
}
