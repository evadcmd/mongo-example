package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

// the mongo.Client ref to opearate for convenience
// would be initialized in func init()
var client *mongo.Client

// the mongo.Database ref to opearate for convenience
// would be initialized in func init()
var database *mongo.Database

func Connect(ctx context.Context) {
	if err := client.Connect(ctx); err != nil {
		log.Fatalf("connect db failed: %v", err)
	}
	if err := initCameraCollection(); err != nil {
		log.Fatalf("connect camera collection failed: %v", err)
	}
}

func Disconnect(ctx context.Context) {
	if err := client.Disconnect(ctx); err != nil {
		log.Fatalf("disconnect db failed: %v", err)
	}
}

func init() {
	var err error
	client, err = mongo.NewClient(dbOpt)
	if err != nil {
		panic(err)
	}
	database = client.Database(NameDB)
}
