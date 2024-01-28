package entity

import "time"

type ProductEntity struct {
  ID          string    `json:"id"`
  Title       string    `json:"title"`
  Price       int64     `json:"price"`
  Description string    `json:"description"`
  CreatedAt   time.Time `json:"created_at"`
  UpdatedAt   time.Time `json:"updated_at"`
}
