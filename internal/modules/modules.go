package modules

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
	UserID   int64
	UserName string
}
