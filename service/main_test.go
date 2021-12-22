package service

import (
	"log"
	"os"
	"testing"

	pb "github.com/FarrukhibnAkbar/ToDo/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client pb.TaskServiceClient

func TestMain(m *testing.M) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}
	client = pb.NewTaskServiceClient(conn)

	os.Exit(m.Run())
}
