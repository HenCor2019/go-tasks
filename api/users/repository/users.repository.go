package UsersRepositories

import (
	"github.com/HenCor2019/task-go/api/models"
	"github.com/HenCor2019/task-go/api/users/dtos"
)

func (r *Repository) CreateUser(userToCreate dtos.CreateUserDto) (models.User, error) {
  user := CreateUserFromDto(userToCreate)
  savedUser := r.db.Create(&user)

  if savedUser.Error != nil {
    return user,savedUser.Error
  }

  return user,nil
}

func (r *Repository) Find() ([]models.User, error){
  users := [] models.User {}
  result := r.db.Preload("Tasks").Find(&users)
  if result.Error != nil {
    return users,nil
  }

  return users,nil
}

func (r *Repository) FindById(userId string) (models.User, error) {
  user := models.User {}
  result := r.db.Preload("Tasks").First(&user, userId)
  if result.Error != nil {
    return user,nil
  }

  return user,nil
}

func (r *Repository) DeleteById(userIdToDelete string) (models.User,error) {
  user := models.User {}
  result := r.db.Delete(&user, userIdToDelete)
  if result.Error != nil {
    return user,nil
  }

  return user,nil
}

func CreateUserFromDto(userDto dtos.CreateUserDto) models.User {
  return models.User {
    Name: userDto.Name,
    Email: userDto.Email,
    Age: userDto.Age,
  }
}

