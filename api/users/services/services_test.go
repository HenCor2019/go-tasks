package UsersServices

import (
	"os"
	"testing"

	UsersRepositories "github.com/HenCor2019/task-go/api/users/repository"
)

var (
	repo             *UsersRepositories.MockUserRepository
	userServicesMock UserService
)

func TestMain(m *testing.M) {
	repo = &UsersRepositories.MockUserRepository{}

	userServicesMock = New(repo)
	code := m.Run()
	os.Exit(code)
}
