package dto

type CreateProductDto struct {
  Title       string   `json:"title" validate:"required,min=3,max=40"`
  Price       int32    `json:"price" validate:"required,min=1"`
  Categories  []string `json:"categories" validate:"required,min=1,dive,uuid4"`
  Description string   `json:"description" validate:"required,min=3,max=500"`
}

type UpdateProductDto struct {
  Title       string   `json:"title" validate:"omitempty,min=3,max=40"`
  Price       int32    `json:"price" validate:"omitempty,min=1"`
  Categories  []string `json:"categories" validate:"omitempty,min=1,dive,uuid4"`
  Description string   `json:"description" validate:"omitempty,min=3,max=500"`
}

type FindProductDto struct {
  Search     string   `json:"search" validate:"omitempty,min=2,max=40"`
  Categories []string `json:"categories" validate:"omitempty,min=1,dive,uuid4"`
}
