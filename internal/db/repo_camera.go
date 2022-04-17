package db

import (
	"context"

	"github.com/evadcmd/mongo-example/internal/db/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Camera interface {
	FindOne(ctx context.Context, id string) (*model.Camera, error)
	FindMany(ctx context.Context, ids []string) ([]*model.Camera, error)
	// FindMany2(ctx context.Context, ids []string) ([]interface{}, error)
	FindAll(ctx context.Context) ([]*model.Camera, error)
	InsertOne(ctx context.Context, camera *model.Camera) error
	InsertMany(ctx context.Context, cmrs []*model.Camera) error
}

type camera struct {
	*mongo.Collection
}

func (cmr *camera) FindOne(ctx context.Context, id string) (*model.Camera, error) {
	res := cmr.Collection.FindOne(ctx, bson.M{"id": id})
	var camera model.Camera
	if err := res.Decode(&camera); err != nil {
		return nil, err
	}
	return &camera, nil
}

func (cmr *camera) FindMany(ctx context.Context, ids []string) ([]*model.Camera, error) {
	cur, err := cmr.Collection.Find(ctx, bson.M{"id": bson.M{"$in": ids}})
	if err != nil {
		return nil, err
	}
	cmrs := make([]*model.Camera, 0, len(ids))
	for cur.Next(ctx) {
		var cmr model.Camera
		cur.Decode(&cmr)
		cmrs = append(cmrs, &cmr)
	}
	return cmrs, nil
}

// not working
func (cmr *camera) FindMany2(ctx context.Context, ids []string) ([]interface{}, error) {
	filter := make([]bson.E, len(ids))
	for _, id := range ids {
		filter = append(filter, bson.E{Key: "id", Value: id})
	}
	cur, err := cmr.Collection.Find(ctx, bson.D(filter))
	if err != nil {
		return nil, err
	}
	res := make([]interface{}, 0, len(ids))
	for cur.Next(ctx) {
		var cmr model.Camera
		if err := cur.Decode(&cmr); err != nil {
			res = append(res, &cmr)
		}
	}
	return res, nil
}

func (cmr *camera) FindAll(ctx context.Context) ([]*model.Camera, error) {
	cur, err := cmr.Collection.Find(ctx, bson.M{})
	// using bson.D is also OK
	// cur, err := cmr.Collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var res []*model.Camera
	for cur.Next(ctx) {
		var c model.Camera
		if err := cur.Decode(&c); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	return res, nil
}

func (cmr *camera) InsertOne(ctx context.Context, camera *model.Camera) error {
	_, err := cmr.Collection.InsertOne(ctx, camera)
	if err != nil {
		return err
	}
	return nil
}

func (cmr *camera) InsertMany(ctx context.Context, cameras []*model.Camera) error {
	ctnr := make([]interface{}, len(cameras))
	for i, camera := range cameras {
		ctnr[i] = camera
	}
	if _, err := cmr.Collection.InsertMany(ctx, ctnr); err != nil {
		return err
	}
	return nil
}

var CameraRepo Camera

func initCameraCollection() error {
	ctx := context.TODO()
	database.CreateCollection(ctx, NameCollectionCamera)
	if _, err := database.Collection(NameCollectionCamera).Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"id": 1},
		Options: options.Index().SetUnique(true).SetBackground(false),
	}); err != nil {
		return err
	}
	return nil
}

func init() {
	CameraRepo = &camera{
		database.Collection(NameCollectionCamera),
	}
	initCameraCollection()
}
