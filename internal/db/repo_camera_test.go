package db

import (
	"context"
	"testing"

	"github.com/evadcmd/mongo-example/internal/db/model"
)

func TestRepoFindOne(t *testing.T) {
	ctx := context.TODO()
	cmr, err := CameraRepo.FindOne(ctx, "a")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", cmr)
}

func TestRepoInsertMany(t *testing.T) {
	ctx := context.TODO()
	if err := CameraRepo.InsertMany(ctx, []*model.Camera{
		{ID: "a"},
		{ID: "b"},
		{ID: "c"},
		{ID: "d"},
	}); err != nil {
		t.Error(err)
	}
}

func TestRepoFindMany(t *testing.T) {
	ctx := context.TODO()
	cmrs, err := CameraRepo.FindMany(ctx, []string{"e", "d", "c", "b", "a"})
	if err != nil {
		t.Error(err)
	}
	for _, cmr := range cmrs {
		t.Logf("%+v", *&cmr.ID)
	}
}

/*
func TestRepoFindMany2(t *testing.T) {
	ctx := context.TODO()
	cmrs, err := CameraRepo.FindMany2(ctx, []string{"e", "d", "c", "b", "a"})
	if err != nil {
		t.Error(err)
	}
	for _, cmr := range cmrs {
		ptr, _ := cmr.(*model.Camera)
		t.Logf("%+v", *ptr)
	}
}
*/

func TestFindAll(t *testing.T) {
	ctx := context.TODO()
	cmrs, err := CameraRepo.FindAll(ctx)
	if err != nil {
		t.Log(err)
	}
	for _, cmr := range cmrs {
		t.Logf("%+v", *cmr)
	}
}
