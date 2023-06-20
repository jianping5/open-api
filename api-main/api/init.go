package api

import (
	"context"
	"api-main/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

// 连接数据库
func init() {
	connectMongo()
}

func connectMongo() {
	uri := config.Config.Mongo.Uri
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	DB = client
}