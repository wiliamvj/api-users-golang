package entity

import "time"

type UserEntity struct {
  ID        string      `json:"id"`
  Name      string      `json:"name"`
  Email     string      `json:"email"`
  Password  string      `json:"password,omitempty"`
  Address   UserAddress `json:"address,omitempty"`
  CreatedAt time.Time   `json:"created_at"`
  UpdatedAt time.Time   `json:"updated_at"`
}

type UserAddress struct {
  CEP        string `json:"cep"`
  IBGE       string `json:"ibge"`
  UF         string `json:"uf"`
  City       string `json:"city"`
  Complement string `json:"complement,omitempty"`
  Street     string `json:"street"`
}
