package entity

import "time"

type CategoryEntity struct {
  ID        string    `json:"id"`
  Title     string    `json:"title"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}
