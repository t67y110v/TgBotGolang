package tatsoft

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Res struct {
	XMLName      xml.Name `xml:"res"`
	Text         string   `xml:"chardata"`
	ResponseType string   `xml:"responseType"`
	Word         string   `xml:"word"`
	POS          string   `xml:"POS"`
	Translation  string   `xml:"translation"`
	Examples     string   `xml:"examples"`
	Mt           string   `xml:"mt"`
}

const res = "https://translate.tatar/translate_hack"

func tatSoftTranslate(tx string, lang int) string {
	req, err := makeGetRequest(tx, lang)
	if err != nil {
		return "Request error"
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "Creating response error"
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Reading from response error"
	}
	defer resp.Body.Close()
	var translatedText Res
	xml.Unmarshal(body, &translatedText)
	if err != nil {
		return "UNmarshelling error"
	}
	return translatedText.Translation
}
func makeGetRequest(text string, lang int) (*http.Request, error) {
	req, err := http.NewRequest("GET", res, nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Add("lang", strconv.Itoa(lang))
	query.Add("text", text)
	req.URL.RawQuery = query.Encode()
	fmt.Println(req.URL.String())
	return req, nil
}
