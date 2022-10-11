package UsersServices

import (
	"github.com/HenCor2019/task-go/api/models"
	"github.com/HenCor2019/task-go/api/users/dtos"
	"github.com/HenCor2019/task-go/api/users/repositories"
)

type UserService interface {
	CreateUser(createUserDto dtos.CreateUserDto) models.User
	Find() []models.User
	FindById(userId string) models.User
	DeleteById(userId string) models.User
}

type Service struct {
	repo *UsersRepositories.Repository
}

func New(repo *UsersRepositories.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
