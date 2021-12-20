package postgres

import (
	"testing"

	"github.com/FarrukhibnAkbar/ToDo/config"
	pb "github.com/FarrukhibnAkbar/ToDo/genproto"
	"github.com/FarrukhibnAkbar/ToDo/pkg/db"
	"github.com/FarrukhibnAkbar/ToDo/storage/repo"
	"github.com/stretchr/testify/suite"
)

type TaskRepositoryTestSuite struct {
	suite.Suite
	CleanupFunc func()
	Repository  repo.TaskStorageI
}

func (suite *TaskRepositoryTestSuite) SetupSuite() {
	pgPool, cleanup := db.ConnectDBForSuite(config.Load())

	suite.Repository = NewTaskRepo(pgPool)
	suite.CleanupFunc = cleanup
}

func (suite *TaskRepositoryTestSuite) TestTaskCRUD() {
	id := "0d512776-60ed-4980-b8a3-6904a2234fd4"
	assignee := "Farrux"

	task := pb.Task{
		Id:        id,
		Assignee:  "Farrux",
		Title:     "Gopher",
		Summary:   "Some ...",
		Deadline:  "2002-02-02",
		Status:    "active",
		CreatedAt: "",
		UpdatedAt: "",
	}

	_ = suite.Repository.Delete(id)

	taskTodo, err := suite.Repository.Create(task)
	suite.Nil(err)

	task = pb.Task{
		Id:       taskTodo.Id,
		Title:    taskTodo.Title,
		Assignee: taskTodo.Assignee,
		Summary:  taskTodo.Summary,
		Deadline: taskTodo.Deadline,
		Status:   taskTodo.Status,
	}

	getTask, err := suite.Repository.Get(task.Id)
	suite.Nil(err)
	suite.NotNil(getTask, "task must not be nil")
	suite.Equal(assignee, getTask.Assignee, "assignee must match")

	task.Title = "gofer"
	updateTask, err := suite.Repository.Update(task)
	suite.Nil(err)

	getTask, err = suite.Repository.Get(id)
	suite.Nil(err)
	suite.NotNil(getTask)
	suite.Equal(task.Title, updateTask.Title)

	listTasks, _, err := suite.Repository.List(1, 10)
	suite.Nil(err)
	suite.NotEmpty(listTasks)
	suite.Equal(task.Title, listTasks[0].Title)

	err = suite.Repository.Delete(id)
	suite.Nil(err)

}

func (suite *TaskRepositoryTestSuite) TearDownSuite() {
	suite.CleanupFunc()
}

func TestTaskRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryTestSuite))
}
