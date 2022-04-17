package main

import (
	"context"
	"log"

	"github.com/evadcmd/mongo-example/internal/db"
	"github.com/evadcmd/mongo-example/internal/db/model"
)

func main() {
	ctx := context.Background()
	db.Connect(ctx)
	defer db.Disconnect(ctx)
	err := db.CameraRepo.InsertOne(ctx, &model.Camera{ID: "a"})
	if err != nil {
		log.Fatal(err)
	}
	cmr, err := db.CameraRepo.FindOne(ctx, "a")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cmr)
}
