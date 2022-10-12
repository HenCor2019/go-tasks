package TasksServices

import (
	"errors"
	"testing"

	"github.com/HenCor2019/task-go/api/models"
	tasksDtos "github.com/HenCor2019/task-go/api/tasks/dtos"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
	// "github.com/gofiber/fiber/v2"
)

func TestCreateTask(t *testing.T) {
	testCases := []struct {
		Name          string
		Result        models.Task
		Dto           tasksDtos.CreateTaskDto
		UserId        string
		ExpectedError error
	}{
		{
			Name: "Should create a task",
			Dto: tasksDtos.CreateTaskDto{
				Title:       "Test title to create",
				Description: "Test description to create",
			},
			Result: models.Task{
				ID:          1,
				Title:       "Test task",
				Description: "Test task description",
				UserID:      1,
			},
			UserId:        "1",
			ExpectedError: nil,
		},

		{
			Name: "Should throw an error if user doesn't exist",
			Dto: tasksDtos.CreateTaskDto{
				Title:       "Test title to create",
				Description: "Test description to create",
			},
			Result: models.Task{
				ID:          0,
				Title:       "Test task",
				Description: "Test task description",
				UserID:      1,
			},
			UserId:        "2",
			ExpectedError: fiber.NewError(fiber.StatusBadRequest, "User not found"),
		},

		{
			Name: "Should throw an error if task cannot be created",
			Dto: tasksDtos.CreateTaskDto{
				Title:       "Test title to create",
				Description: "Test description to create",
			},
			Result: models.Task{
				ID:          1,
				Title:       "Test task",
				Description: "Test task description",
				UserID:      2,
			},
			UserId:        "3",
			ExpectedError: fiber.NewError(fiber.StatusBadRequest, "Cannot save the task"),
		},
	}

	repo.On("FindUserById", "1").Return(models.User{
		ID:   1,
		Name: "Test name",
		Age:  18,
	}, nil)

	repo.On("Create", tasksDtos.CreateTaskDto{
		Title:       "Test title to create",
		Description: "Test description to create",
	}, cast.ToUint(1)).Return(models.Task{
		ID:          1,
		Title:       "Test task",
		Description: "Test task description",
		UserID:      1,
	}, nil)

	repo.On("FindUserById", "2").Return(models.User{
		ID:   0,
		Name: "",
		Age:  0,
	}, errors.New("User not found"))

	repo.On("FindUserById", "3").Return(models.User{
		ID:   1,
		Name: "Test name",
		Age:  18,
	}, nil)

	repo.On("Create", tasksDtos.CreateTaskDto{
		Title:       "Test title to create",
		Description: "Test description to create",
	}, cast.ToUint(3)).Return(models.Task{
		ID:          0,
		Title:       "Test task",
		Description: "Test task description",
	}, errors.New("Cannot create the task"))

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			repo.Mock.Test(t)
			t.Parallel()

			defer func() { recover() }()
			result := taskServicesMock.CreateTask(tc.Dto, tc.UserId)
			if result.ID != tc.Result.ID {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, result)
			}
		})
	}
}

func TestDeleteTaskById(t *testing.T) {
	testCases := []struct {
		Name          string
		Result        models.Task
		TaskId        string
		UserId        string
		ExpectedError error
	}{
		{
			Name:          "Should throw an error if user doesn't exist",
			Result:        models.Task{},
			TaskId:        "4",
			UserId:        "4",
			ExpectedError: fiber.NewError(fiber.StatusBadRequest, "User not found"),
		},

		{
			Name:          "Should throw an error if user is not the task owner",
			Result:        models.Task{},
			TaskId:        "5",
			UserId:        "5",
			ExpectedError: fiber.NewError(fiber.StatusForbidden, "You cannot delete this task"),
		},

		{
			Name:          "Should throw an error if task cannot be deleted",
			Result:        models.Task{},
			TaskId:        "6",
			UserId:        "6",
			ExpectedError: fiber.NewError(fiber.StatusForbidden, "You cannot delete this task"),
		},

		{
			Name:          "Should delete the task",
			Result:        models.Task{},
			TaskId:        "7",
			UserId:        "7",
			ExpectedError: nil,
		},
	}

	repo.On("FindUserById", "4").Return(models.User{}, errors.New("User not found"))

	repo.On("FindUserById", "5").Return(models.User{ID: 5}, nil)
	repo.On("FindById", "5", cast.ToUint(5)).Return(models.Task{ID: 0}, errors.New("Task doesn't exist"))

	repo.On("FindUserById", "6").Return(models.User{ID: 6}, nil)
	repo.On("FindById", "6", cast.ToUint(6)).Return(models.Task{ID: 6}, nil)
	repo.On("DeleteById", "6").Return(errors.New("Cannot delete the task"))

	repo.On("FindUserById", "7").Return(models.User{ID: 7}, nil)
	repo.On("FindById", "7", cast.ToUint(7)).Return(models.Task{ID: 7}, nil)
	repo.On("DeleteById", "7").Return(nil)

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			repo.Mock.Test(t)
			t.Parallel()

			defer func() { recover() }()
			taskServicesMock.DeleteTaskById(tc.TaskId, tc.UserId)
		})
	}
}
