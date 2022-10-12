package UsersServices

import (
	"os"
	"testing"

	"github.com/HenCor2019/task-go/api/models"
	UsersRepositories "github.com/HenCor2019/task-go/api/users/repository"
)

var repo *UsersRepositories.MockUserRepository
var userServicesMock UserService

func TestMain(m *testing.M) {
	repo = &UsersRepositories.MockUserRepository{}
	userServicesMock = New(repo)
	repo.On("Find").Return([]models.User{{
		Name:  "Henry Cortez",
		Email: "hcortez@gmail.com",
		Age:   18,
		Tasks: []models.Task{},
	}})

	code := m.Run()
	os.Exit(code)
}
