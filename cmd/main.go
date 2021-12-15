package main

import (
	"net"

	"google.golang.org/grpc"

	"github.com/FarrukhibnAkbar/ToDo/config"
	pb "github.com/FarrukhibnAkbar/ToDo/genproto"
	"github.com/FarrukhibnAkbar/ToDo/pkg/db"
	"github.com/FarrukhibnAkbar/ToDo/pkg/logger"
	"github.com/FarrukhibnAkbar/ToDo/service"
	"github.com/FarrukhibnAkbar/ToDo/storage"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "ToDo")
	defer func(l logger.Logger) {
		err := logger.Cleanup(l)
		if err != nil {
			log.Fatal("failed cleanup logger", logger.Error(err))
		}
	}(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase),
	)

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connetion to postgres error", logger.Error(err))
	}

	pgStorage := storage.NewStoragePg(connDB)

	taskService := service.NewTaskService(pgStorage, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterTaskServiceServer(s, taskService)
	log.Info("main: server running", logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}