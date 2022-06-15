package mongodb

import "encoding/json"

type Tokens struct {
	Status  int    `json:"status"`
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
	Guid    string `json:"guid"`
}

func TokenEncodingJson(token Tokens) ([]byte, error) {
	return json.Marshal(token)
}

func DecodingJsonToken(jsonToken []byte) (Tokens, error) {
	var token Tokens
	return token, json.Unmarshal(jsonToken, &token)
}
