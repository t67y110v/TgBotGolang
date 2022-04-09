package main

import (
	"bolt"
	"strconv"
)

type Bucket string

const (
	AccessTokens  Bucket = "access_tokens"
	RequestTokens Bucket = "request_tokens"
)

type TokenReprository struct {
	db *bolt.DB
}

func NewTokenRepository(db *bolt.DB) *TokenReprository {
	return &TokenReprository{db: db}
}
func (r *TokenReprository) Save(chatId int64, token string, bucket Bucket) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put(intToByte(chatId), []byte(token))
	})

}
func (r *TokenReprository) Get(chatId int64, bucket Bucket) (string, error) {
	var token string
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		data := b.Get(intToByte(chatId))
		token = string(data)
		return nil
	})
	if err != nil {
		return "", err
	}
	if token == "" {
		return "", err
	}
	return token, nil
}

func intToByte(v int64) []byte {
	return []byte(strconv.FormatInt(v, 10))
}
