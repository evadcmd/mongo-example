package db

import (
	"context"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func TestMain(m *testing.M) {
	ctx := context.TODO()
	Connect(ctx)
	defer Disconnect(ctx)
	os.Exit(m.Run())
}

func TestPing(t *testing.T) {
	ctx := context.TODO()
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		t.Error(err)
	}
}

func TestCreateDB(t *testing.T) {
	ctx := context.TODO()
	if err := database.CreateCollection(ctx, "camera"); err != nil {
		t.Error(err)
	}
}

func TestGetDB(t *testing.T) {
	ctx := context.TODO()
	dbs, err := client.ListDatabaseNames(ctx, bson.D{})
	if err != nil {
		t.Error(err)
	}
	t.Log(dbs)
}
