package dtos

type CreateUserDto struct {
	Name  string `validate:"required,min=5,max=30"`
	Email string `validate:"required,email,min=5,max=30"`
	Age   uint8  `validate:"required,number,min=15,max=30"`
}
