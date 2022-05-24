package main

type RestResponse2 struct {
	Items []Item `json:"items"`
}
type Item struct {
	Id ItemInfo `json:"id"`
}
type ItemInfo struct {
	VideoId string `json:"videoId"`
}
type User struct {
	userID   int64
	userName string
}
