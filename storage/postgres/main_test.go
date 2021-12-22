package postgres

import (
	"log"
	"os"
	"testing"

	"github.com/FarrukhibnAkbar/ToDo/config"
	"github.com/FarrukhibnAkbar/ToDo/pkg/db"
	"github.com/FarrukhibnAkbar/ToDo/pkg/logger"
)

var pgRepo *taskRepo

func TestMain(m *testing.M) {
	cfg := config.Load()

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connetion to postgres error", logger.Error(err))
	}

	pgRepo = NewTaskRepo(connDB)

	os.Exit(m.Run())
}
