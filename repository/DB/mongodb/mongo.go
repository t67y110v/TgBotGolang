package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RefreshToken struct {
	Guid    string    `bson:"_id"`
	Refresh string    `bson:"refresh"`
	Time    time.Time `bson:"time"`
}

var collection *mongo.Collection
var ctx = context.TODO()

func InitDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("auth").Collection("refresh-tokens")
}

func InsertRefreshToken(refresh, guid string) error {
	token := RefreshToken{Refresh: refresh, Guid: guid, Time: time.Now().Add(60 * 24 * time.Hour)}
	_, err := collection.InsertOne(context.TODO(), token)
	if err != nil {
		return err
	}
	return nil
}

func ReadRefreshToken(guid string) (*RefreshToken, error) {
	filter := bson.D{{"_id", guid}}
	var result *RefreshToken
	var err error
	if err = collection.FindOne(context.TODO(), filter).Decode(&result); err == nil {
		return result, err
	}
	return nil, err
}

func UpdateRefreshToken(refresh, guid string) error {
	filter := bson.D{{"_id", guid}}
	update := bson.D{
		{"$set", bson.D{
			{"refresh", refresh},
			{"time", time.Now().Add(60 * 24 * time.Hour)},
		}},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
