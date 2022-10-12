package TasksRepositories

import (
	"github.com/HenCor2019/task-go/api/models"
	"github.com/HenCor2019/task-go/api/tasks/dtos"
)


func (r *Repository) Create(createTaskDto tasksDtos.CreateTaskDto, userId uint) (models.Task,error) {
  taskToSave := CreateTaskFromDto(createTaskDto)
  taskToSave.UserID = userId

  result := r.db.Model(&models.Task {}).Create(&taskToSave)
  if result.Error != nil {
    return taskToSave,result.Error
  }
  return taskToSave,nil
}

func (r *Repository) DeleteById(taskId string) (error) {
  taskToDelete := models.Task {}
  result := r.db.Model(&models.Task{}).Delete(&taskToDelete,taskId)

  if result.Error != nil {
    return result.Error
  }
  return nil
}

func (r *Repository) FindById(taskId string, userId uint) (models.Task,error) {
  task := models.Task {}

  result := r.db.Model(&models.Task{}).First(&task, models.Task {
    UserID: userId,
  })

  if result.Error != nil {
    return task,result.Error
  }

  return task,nil
}

func (r *Repository) FindUserById(userId string) (models.User,error) {
  user := models.User {}

  result := r.db.Model(&models.User{}).First(&user,userId)

  if result.Error != nil {
    return user,result.Error
  }

  return user,nil
}


func CreateTaskFromDto(createTaskDto tasksDtos.CreateTaskDto) models.Task {
  return models.Task {
    Title: createTaskDto.Title,
    Description: createTaskDto.Description,
  }
}
