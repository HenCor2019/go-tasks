package UsersRepositories

import (
	"github.com/HenCor2019/task-go/config/db"
	"github.com/HenCor2019/task-go/models"
	"github.com/HenCor2019/task-go/users/dtos"
)

func CreateUser(userToCreate dtos.CreateUserDto) (models.User, error) {
  user := CreateUserFromDto(userToCreate)
  savedUser := db.DB.Create(&user)

  if savedUser.Error != nil {
    return user,savedUser.Error
  }

  return user,nil
}

func Find() ([]models.User, error){
  users := [] models.User {}
  result := db.DB.Preload("Tasks").Find(&users)
  if result.Error != nil {
    return users,nil
  }

  return users,nil
}

func FindById(userId string) (models.User, error) {
  user := models.User {}
  result := db.DB.Preload("Tasks").First(&user, userId)
  if result.Error != nil {
    return user,nil
  }

  return user,nil
}

func DeleteById(userIdToDelete string) (models.User,error) {
  user := models.User {}
  result := db.DB.Delete(&user, userIdToDelete)
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

