package UsersRepositories

import (
	"github.com/HenCor2019/task-go/api/models"
	"github.com/HenCor2019/task-go/api/users/dtos"
	"gorm.io/gorm"
)

// Repository is the interface that wraps the basic CRUD operations.
//
//go:generate mockery --name=UserRepository --output=repository --inpackage
type UserRepository interface {
	Find() ([]models.User,error)
	CreateUser(createUserDto dtos.CreateUserDto) (models.User,error)
	FindById(userId string) (models.User,error)
	DeleteById(userId string) (models.User,error)
}

type Repository struct {
	db *gorm.DB
}


func New(repo *gorm.DB) UserRepository {
	return &Repository{
		db: repo,
	}
}
