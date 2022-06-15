package boltdb

import (
	"errors"
	bDB "goBot/repository/DB"
	"strconv"

	"github.com/boltdb/bolt"
)

type InformRepos struct {
	db *bolt.DB
}

func InitBoltDB() (*bolt.DB, error) {
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bDB.Information))
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return db, nil
}

func NewInformRepos(db *bolt.DB) *InformRepos {
	return &InformRepos{db: db}
}

func (r *InformRepos) Save(chatID int64, info string, bucket bDB.Bucket) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put(intToBytes(chatID), []byte(info))
	})
}
func (r *InformRepos) Get(chatID int64, bucket bDB.Bucket) (string, error) {
	var info string
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		data := b.Get(intToBytes(chatID))
		info = string(data)
		return nil

	})

	if err != nil {
		return "", err
	}

	if info == "" {
		return "", errors.New("info not found")
	}

	return info, nil

}

func intToBytes(v int64) []byte {
	return []byte(strconv.FormatInt(v, 10))
}
