package dto

type CreateUserDto struct {
  Name     string `json:"name" validate:"required,min=3,max=30"`
  Email    string `json:"email" validate:"required,email"`
  Password string `json:"password" validate:"required,min=8,max=30,containsany=!@#$%*"`
}
