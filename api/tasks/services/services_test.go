package TasksServices

import (
	"os"
	"testing"

	TasksRepositories "github.com/HenCor2019/task-go/api/tasks/repositories"
)

var (
	repo             *TasksRepositories.MockTaskRepository
	taskServicesMock TaskService
)

func TestMain(m *testing.M) {
	repo = &TasksRepositories.MockTaskRepository{}
	taskServicesMock = New(repo)

	code := m.Run()
	os.Exit(code)
}
