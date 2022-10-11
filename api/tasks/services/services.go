package TasksServices

import (
	"github.com/HenCor2019/task-go/api/models"
	"github.com/HenCor2019/task-go/api/users/dtos"
	"github.com/HenCor2019/task-go/api/tasks/repositories"
)

type TaskService interface {
	CreateUser(createUserDto dtos.CreateUserDto) models.User
	Find() []models.Task
	FindById(taskId string,userId string) models.Task
	DeleteById(taskId string,userId string) models.Task
}

type Service struct {
	repo *TasksRepositories.Repository
}

func New(repo *TasksRepositories.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
