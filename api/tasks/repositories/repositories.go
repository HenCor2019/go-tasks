package TasksRepositories

import (
	"github.com/HenCor2019/task-go/api/models"
	"github.com/HenCor2019/task-go/api/tasks/dtos"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(createTaskDto tasksDtos.CreateTaskDto, userId uint) (models.Task, error)
	FindById(taskId string, userId uint) (models.Task, error)
	FindUserById(userId string) (models.User, error)
	DeleteById(taskId string) error
}

type Repository struct {
	db *gorm.DB
}

func New(repo *gorm.DB) TaskRepository {
	return &Repository{
		db: repo,
	}
}
