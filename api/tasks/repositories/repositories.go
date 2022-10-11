package TasksRepositories

import (
	"github.com/HenCor2019/task-go/api/models"
	"github.com/HenCor2019/task-go/api/tasks/dtos"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(createTaskDto tasksDtos.CreateTaskDto) (models.Task,error)
	FindById(taskId string) (models.Task,error)
 FindUserById(userId string) (models.User,error)
	DeleteById(taskId string) (models.Task,error)
}

type Repository struct {
	db *gorm.DB
}

func New(repo *gorm.DB) *Repository {
	return &Repository{
		db: repo,
	}
}
