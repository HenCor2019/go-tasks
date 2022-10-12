package UsersServices

import (
	"os"
	"testing"

	UsersRepositories "github.com/HenCor2019/task-go/api/users/repositories"
	mock "github.com/stretchr/testify/mock"
)

var repo *UsersRepositories.MockUserRepository
var userServicesMock *Service

func TestMain(m *testing.M) {
	repo = &UsersRepositories.MockUserRepository {}

	repo.On("CreateUser", mock.Anything, mock.Anything).Return(nil, nil)
	repo.On("Find", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil,nil)
	repo.On("FindById", mock.Anything, mock.Anything, mock.Anything).Return(nil,nil)
	repo.On("DeleteById", mock.Anything, mock.Anything, mock.Anything).Return(nil,nil)

 repo01 := &UsersRepositories.MockUserRepository {}
	userServicesMock = New(repo01)

	code := m.Run()
	os.Exit(code)
}
