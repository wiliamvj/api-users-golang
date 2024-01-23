package response

import "time"

type UserResponse struct {
  ID        string      `json:"id"`
  Name      string      `json:"name"`
  Email     string      `json:"email"`
  Address   UserAddress `json:"address"`
  CreatedAt time.Time   `json:"created_at"`
  UpdatedAt time.Time   `json:"updated_at"`
}

type UserAddress struct {
  CEP        string `json:"cep"`
  UF         string `json:"uf"`
  City       string `json:"city"`
  Complement string `json:"complement,omitempty"`
  Street     string `json:"street"`
}

type ManyUsersResponse struct {
  Users []UserResponse `json:"users"`
}

type UserAuthToken struct {
  AccessToken string `json:"access_token"`
}
