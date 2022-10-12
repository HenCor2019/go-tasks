package UsersServices

import (
	"github.com/HenCor2019/task-go/api/models"
	"github.com/HenCor2019/task-go/api/users/dtos"
	"github.com/HenCor2019/task-go/api/users/repository"
)

//go:generate mockery --name=UserService --output=services --inpackage
type UserService interface {
	CreateUser(createUserDto dtos.CreateUserDto) models.User
	Find() []models.User
	FindById(userId string) models.User
	DeleteById(userId string) models.User
}

type Service struct {
	repo UsersRepositories.UserRepository
}

func New(repo UsersRepositories.UserRepository) UserService {
	return &Service{ repo }
}
