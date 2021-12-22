package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/FarrukhibnAkbar/ToDo/config"
)

func ConnectToDB(cfg config.Config) (*sqlx.DB, error) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	connDB, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		return nil, err
	}

	// cleanUpFunc := func() {
	// 	connDB.Close()
	// }

	return connDB, nil
}

func ConnectDBForSuite(cfg config.Config) (*sqlx.DB, func()) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	connDB, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		panic(err)
	}

	CleanUpFunc := func() {
		connDB.Close()
	}

	return connDB, CleanUpFunc
}
