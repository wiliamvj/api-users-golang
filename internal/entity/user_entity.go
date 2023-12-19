package entity

import "time"

type UserEntity struct {
  ID        string    `json:"id"`
  Name      string    `json:"name"`
  Email     string    `json:"email"`
  Password  string    `json:"password,omitempty"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}
