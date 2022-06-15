package db

type Bucket string

const (
	Information Bucket = "information"
)

type InformRepos interface {
	Save(chatID int64, info string, bucket Bucket) error
	Get(chatID int64, bucket Bucket) (string, error)
}
