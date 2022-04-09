package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//https://libretranslate.com/translate?q=%2Ftra&source=eng&target=ru
const res = "https://libretranslate.com/translate"

func translateFunction(text string) string {
	fmt.Println("зашел в перевод")
	req, err := makePostRequest(text)
	if err != nil {
		return ""
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	var translatedText TranslatedText
	err = json.Unmarshal(body, &translatedText)
	if err != nil {
		return ""
	}
	return ""

	/*buf, err := json.Marshal(trnslt)
	if err != nil {
		return "error1"
	}
	resp, err := http.Post(res, "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return "error2"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "error3"
	}
	var translated TranslatedText
	err = json.Unmarshal(body, &translated)
	if err != nil {
		return "error4"
	}
	return translated.Tx */
}

//https://libretranslate.com/translate?q=%2Ftra&source=eng&target=ru
func makePostRequest(text string) (*http.Request, error) {
	req, err := http.NewRequest("POST", res, nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Add("q", text)
	query.Add("source", "eng")
	query.Add("target", "ru")
	req.URL.RawQuery = query.Encode()
	fmt.Println(req.URL.String())
	return req, nil

}
