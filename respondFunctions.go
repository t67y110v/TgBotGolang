package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const res = "https://libretranslate.com/translate"

/*
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

*/
func translateFunction(text string) string {
	fmt.Println("зашел в перевод")
	var trnslt Translate
	trnslt.Tex = text
	trnslt.Src = "eng"
	trnslt.Trgt = "ru"

	buf, err := json.Marshal(trnslt)
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
	return translated.Tx
}
