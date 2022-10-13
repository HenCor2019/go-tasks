package TasksServices

import (
	"github.com/HenCor2019/task-go/api/models"
	"github.com/HenCor2019/task-go/api/tasks/dtos"
	"github.com/HenCor2019/task-go/api/tasks/repositories"
)

type TaskService interface {
	CreateTask(createTaskDto tasksDtos.CreateTaskDto, userId string) models.Task
	DeleteTaskById(taskId string, userId string)
}

type Service struct {
	repo TasksRepositories.TaskRepository
}

func New(repo TasksRepositories.TaskRepository) TaskService {
	return &Service{repo}
}
