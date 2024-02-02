package response

import (
  "time"
)

type ProductResponse struct {
  ID          string             `json:"id"`
  Title       string             `json:"title"`
  Price       int32              `json:"price"`
  Description string             `json:"description,omitempty"`
  Categories  []CategoryResponse `json:"categories"`
  CreatedAt   time.Time          `json:"created_at"`
}
