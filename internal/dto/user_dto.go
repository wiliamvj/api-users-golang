package dto

type CreateUserDto struct {
  Name     string `json:"name" validate:"required,min=3,max=30"`
  Email    string `json:"email" validate:"required,email"`
  Password string `json:"password" validate:"required,min=8,max=30,containsany=!@#$%*"`
  CEP      string `json:"cep" validate:"required,min=8,max=8"`
}

type UpdateUserDto struct {
  Name  string `json:"name" validate:"omitempty,min=3,max=30"`
  Email string `json:"email" validate:"omitempty,email"`
  CEP   string `json:"cep" validate:"omitempty,min=8,max=8"`
}

type UpdateUserPasswordDto struct {
  Password    string `json:"password" validate:"required,min=8,max=30,containsany=!@#$%*"`
  OldPassword string `json:"old_password" validate:"required,min=8,max=30,containsany=!@#$%*"`
}

type LoginDTO struct {
  Email    string `json:"email" validate:"required,email"`
  Password string `json:"password" validate:"required,min=8,max=40"`
}
