package tasksDtos

type CreateTaskDto struct {
  Title string `validate:"required,min=5,max=20"`
  Description string `validate:"required,min=5,max=40"`
}
