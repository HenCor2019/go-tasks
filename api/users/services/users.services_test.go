package UsersServices

import (
	"errors"
	"testing"

	"github.com/HenCor2019/task-go/api/models"
	"github.com/HenCor2019/task-go/api/users/dtos"
	"github.com/gofiber/fiber/v2"
)

func TestCreateUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Result        models.User
		Dto           dtos.CreateUserDto
		ExpectedError error
	}{
		{
			Name:          "Should return a created user",
			ExpectedError: nil,
			Dto: dtos.CreateUserDto{
				Name:  "Henry Cortez",
				Email: "hcortez@gmail.com",
				Age:   18,
			},
			Result: models.User{
				ID:    1,
				Name:  "Henry Cortez",
				Email: "hcortez@gmail.com",
				Age:   18,
				Tasks: []models.Task{},
			},
		},

		{
			Name:          "Should throw an error if user cannor be created",
			ExpectedError: fiber.NewError(fiber.StatusBadRequest, "Cannot create the users"),
			Dto: dtos.CreateUserDto{
				Name:  "Katya Lopez",
				Email: "klopez@gmail.com",
				Age:   18,
			},
			Result: models.User{
				ID:    0,
				Name:  "Katya Cortez",
				Email: "hcortez@gmail.com",
				Age:   18,
				Tasks: []models.Task{},
			},
		},
	}
	repo.On("CreateUser", dtos.CreateUserDto{
		Name:  "Henry Cortez",
		Email: "hcortez@gmail.com",
		Age:   18,
	}).Return(models.User{
		ID:    1,
		Name:  "Henry Cortez",
		Email: "hcortez@gmail.com",
		Age:   18,
		Tasks: []models.Task{},
	}, nil)

	repo.On("CreateUser", dtos.CreateUserDto{
		Name:  "Katya Lopez",
		Email: "klopez@gmail.com",
		Age:   18,
	}).Return(models.User{
		ID:    0,
		Name:  "Katya Lopez",
		Email: "klopez@gmail.com",
		Age:   18,
		Tasks: []models.Task{},
	}, errors.New("Cannot create user in DB"))

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			repo.Mock.Test(t)
			t.Parallel()

			defer func() { recover() }()
			result := userServicesMock.CreateUser(tc.Dto)
			if result.ID != tc.Result.ID {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, result)
			}
		})
	}
}

func TestFindById(t *testing.T) {
	testCases := []struct {
		Name          string
		Result        models.User
		UserId        string
		ExpectedError error
	}{
		{
			Name:          "Should return a user",
			ExpectedError: nil,
			UserId:        "1",
			Result: models.User{
				ID:    1,
				Name:  "Henry Cortez",
				Email: "hcortez@gmail.com",
				Age:   1,
				Tasks: []models.Task{},
			},
		},

		{
			Name:          "Should throw an error if user doesn't exist",
			ExpectedError: fiber.NewError(fiber.StatusNotFound, "User not found"),
			UserId:        "2",
			Result:        models.User{},
		},
	}
	repo.On("FindById", "1").Return(models.User{
		ID:    1,
		Name:  "Henry Cortez",
		Email: "hcortez@gmail.com",
		Age:   18,
		Tasks: []models.Task{},
	}, nil)

	repo.On("FindById", "2").Return(models.User{}, errors.New("Cannot find in DB"))

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			repo.Mock.Test(t)
			t.Parallel()

			defer func() { recover() }()
			result := userServicesMock.FindById(tc.UserId)
			if result.ID != tc.Result.ID {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, result)
			}
		})
	}
}

func TestDeleteById(t *testing.T) {
	testCases := []struct {
		Name          string
		UserId        string
		ExpectedError error
		Result        models.User
	}{
		{
			Name:          "Should delete a user",
			ExpectedError: nil,
			UserId:        "1",
			Result: models.User{
				ID:    1,
				Name:  "Henry Cortez",
				Email: "hcortez@gmail.com",
				Age:   18,
				Tasks: []models.Task{},
			},
		},

		{
			Name:          "Should throw an error if user doesn't exist",
			ExpectedError: fiber.NewError(fiber.StatusNotFound, "User not found"),
			UserId:        "2",
		},

		{
			Name:          "Should throw an error if cannot delete the user",
			ExpectedError: fiber.NewError(fiber.StatusNotFound, "User not found"),
			UserId:        "3",
		},
	}
	repo.On("FindById", "1").Return(models.User{
		ID:    1,
		Name:  "Henry Cortez",
		Email: "hcortez@gmail.com",
		Age:   18,
		Tasks: []models.Task{},
	}, nil)

	repo.On("FindById", "1").Return(models.User{ID: 1}, nil).Once()
	repo.On("DeleteById", "1").Return(models.User{ID: 1}, nil).Once()

	repo.On("FindById", "2").Return(models.User{}, errors.New("Cannot find in DB")).Once()

	repo.On("FindById", "3").Return(models.User{ID: 1}, nil).Once()
	repo.On("DeleteById", "3").Return(models.User{}, errors.New("Cannot find in DB")).Once()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			repo.Mock.Test(t)
			t.Parallel()

			defer func() { recover() }()
			result := userServicesMock.DeleteById(tc.UserId)
			if result.ID != tc.Result.ID {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, result)
			}
		})
	}
}

func TestFind(t *testing.T) {
	testCases := []struct {
		Name          string
		ExpectedError error
		Result        []models.User
	}{
		{
			Name:          "Should return all users",
			ExpectedError: nil,
			Result: []models.User{
				{
					ID:    1,
					Name:  "Henry Cortez",
					Email: "hcortez@gmail.com",
					Age:   18,
					Tasks: []models.Task{},
				},
			},
		},

		{
			Name:          "Should throw an error if user doesn't exist",
			ExpectedError: fiber.NewError(fiber.StatusBadRequest, "Cannot find the users"),
		},
	}
	repo.On("Find").Return([]models.User{
		{
			ID:    1,
			Name:  "Henry Cortez",
			Email: "hcortez@gmail.com",
			Age:   18,
			Tasks: []models.Task{},
		},
	}, nil).Once()

	repo.On("Find").Return([]models.User{}, errors.New("Cannot find in DB")).Once()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			repo.Mock.Test(t)
			t.Parallel()

			defer func() { recover() }()
			result := userServicesMock.Find()
			if len(result) != len(tc.Result) {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, result)
			}
		})
	}
}
