package entity

import "time"

type ProductEntity struct {
  ID          string    `json:"id"`
  Title       string    `json:"title"`
  Price       int32     `json:"price"`
  Categories  []string  `json:"categories"`
  Description string    `json:"description"`
  CreatedAt   time.Time `json:"created_at"`
  UpdatedAt   time.Time `json:"updated_at"`
}

type ProductCategoryEntity struct {
  ID         string    `json:"id"`
  ProductID  string    `json:"product_id"`
  CategoryID string    `json:"category_id"`
  CreatedAt  time.Time `json:"created_at"`
  UpdatedAt  time.Time `json:"updated_at"`
}

type ProductWithCategoryEntity struct {
  ID          string           `json:"id"`
  Title       string           `json:"title"`
  Price       int32            `json:"price"`
  Description string           `json:"description"`
  Categories  []CategoryEntity `json:"categories"`
  CreatedAt   time.Time        `json:"created_at"`
}
