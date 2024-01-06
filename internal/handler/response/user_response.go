package response

import "time"

type UserResponse struct {
  ID        string    `json:"id"`
  Name      string    `json:"name"`
  Email     string    `json:"email"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}

type ManyUsersResponse struct {
  Users []UserResponse `json:"users"`
}

type UserAuthToken struct {
  AccessToken string `json:"access_token"`
}
