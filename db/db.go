package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var Client *mongo.Client

func ConnectToDB() {
	var uri = os.Getenv("MONGO_URI")
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	_Client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	// Copy the Variable
	Client = _Client

	if err != nil {
		fmt.Printf("%v", err)
		panic(err)
	}
}
