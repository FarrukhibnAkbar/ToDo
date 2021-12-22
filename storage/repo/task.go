package repo

import (
	"time"

	pb "github.com/FarrukhibnAkbar/ToDo/genproto"
)

type TaskStorageI interface {
	Create(pb.Task) (pb.Task, error)
	Get(id string) (pb.Task, error)
	List(page, limit int64) ([]*pb.Task, int64, error)
	Update(pb.Task) (pb.Task, error)
	Delete(id string) error
	ListOverdue(page, limit int64, time time.Time) ([]*pb.Task, int64, error)
}
