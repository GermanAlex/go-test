package main

import (
	"reflect"
	"statistics/pkg/storage"
	"statistics/pkg/storage/memdb"
	"testing"
)

func Test_task(t *testing.T) {
	db := &memdb.DB{}
	item := storage.Task{
		ID:    1,
		Title: "Test Task",
	}
	db.NewTask(item)
	got := task(db, item.ID)
	if !reflect.DeepEqual(got, item) {
		t.Errorf("got %v, want %v", got, item)
	}
}
