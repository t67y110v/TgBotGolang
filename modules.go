package main

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	ChatId int `json:"id"`
}

type RestResponse struct {
	Result []Update `json:"result"`
}
type BotMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}
type TranslatedText struct {
	Tx string `json:"translatedText"`
}

type Translate struct {
	Tex  string `json:"q"`
	Src  string `json:"source"`
	Trgt string `json:"target"`
}

/////////////////////////////////////////
/*type RestResponse2 struct {
	 Items []Item `json:"items"`
}
type Item struct {
	Id []ItemInfo `json:"id"`
}
type ItemInfo struct{
	VideoId string `"json:videoId"`
}*/
