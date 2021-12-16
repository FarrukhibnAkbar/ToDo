package storage

import (
	"github.com/jmoiron/sqlx"

	"github.com/FarrukhibnAkbar/ToDo/storage/postgres"
	"github.com/FarrukhibnAkbar/ToDo/storage/repo"
)

type IStorage interface {
	Task() repo.TaskStorageI
}

type storagePg struct {
	db       *sqlx.DB
	taskRepo repo.TaskStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		taskRepo: postgres.NewTaskRepo(db),
	}
}

func (s storagePg) Task() repo.TaskStorageI {
	return s.taskRepo
}
