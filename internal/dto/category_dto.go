package dto

type CreateCategoryDto struct {
  Title string `json:"title" validate:"required,min=3,max=30"`
}
