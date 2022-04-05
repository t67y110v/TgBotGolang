package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	YOUTUBE_SEARCH_URL = "https://www.googleapis.com/youtube/v3/search"
	YOUTUBE_API_TOKEN  = "YT token"
	YOUTUBE_VIDEO_URL  = "https://www.youtube.com/watch?v="
	botToken           = "token"
	botApi             = "https://api.telegram.org/bot"
	botUrl             = botApi + botToken
)

//точка входа программы
func main() {
	offset := 0
	for {
		updates, err := GetUpdates(botUrl, offset)
		if err != nil {
			log.Println("Smth went wrong", err.Error())
		}
		fmt.Println(updates)
		for _, update := range updates {
			err = respond(botUrl, update)
			offset = update.UpdateId + 1
		}
	}
}

// Запрос обновлений
func GetUpdates(botUrl string, offset int) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getupdates" + "?offset=" + strconv.Itoa(offset))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}

//отвечать на обновления
func respond(botUrl string, update Update) error {
	var botMessage BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
	botMessage.Text = update.Message.Text
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	if botMessage.Text == "/tasks" {

	}
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}

/*func getLastVideo(channelUrl string) (string, error){

}
func retrieveVideos(channelUrl string)([]Item, error){

}
func makeRequest(channelUrl string)(*http.Request,error){


} */
