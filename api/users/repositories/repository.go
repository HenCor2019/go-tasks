package UsersRepositories

import (
	"github.com/HenCor2019/task-go/api/models"
	"github.com/HenCor2019/task-go/api/users/dtos"
	"gorm.io/gorm"
)

type UserRepository interface {
	Find() ([]models.User,error)
	CreateUser(createUserDto dtos.CreateUserDto) (models.User,error)
	FindById(userId string) (models.User,error)
	DeleteById(userId string) (models.User,error)
}

type Repository struct {
	db *gorm.DB
}

func New(repo *gorm.DB) *Repository {
	return &Repository{
		db: repo,
	}
}
