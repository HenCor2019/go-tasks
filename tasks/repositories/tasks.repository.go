package TasksRepositories

import (
	"github.com/HenCor2019/task-go/config/db"
	"github.com/HenCor2019/task-go/models"
	"github.com/HenCor2019/task-go/tasks/dtos"
)


func Create(createTaskDto tasksDtos.CreateTaskDto, userId uint) (models.Task,error) {
  taskToSave := CreateTaskFromDto(createTaskDto)
  taskToSave.UserID = userId

  result := db.DB.Model(&models.Task {}).Create(&taskToSave)
  if result.Error != nil {
    return taskToSave,result.Error
  }
  return taskToSave,nil
}

func DeleteById(idTaskToDelete string) (error) {
  taskToDelete := models.Task {}
  result := db.DB.Model(&models.Task{}).Delete(&taskToDelete,idTaskToDelete)

  if result.Error != nil {
    return result.Error
  }
  return nil
}

func FindById(taskId string, userId uint) (models.Task,error) {
  task := models.Task {}

  result := db.DB.Model(&models.Task{}).First(&task, models.Task {
    UserID: userId,
  })

  if result.Error != nil {
    return task,result.Error
  }

  return task,nil
}


func CreateTaskFromDto(createTaskDto tasksDtos.CreateTaskDto) models.Task {
  return models.Task {
    Title: createTaskDto.Title,
    Description: createTaskDto.Description,
  }
}
