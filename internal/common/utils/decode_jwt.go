package utils

import (
  "errors"
  "net/http"
  "strings"

  "github.com/golang-jwt/jwt/v4"
  "github.com/wiliamvj/api-users-golang/config/env"
)

type CurrentUser struct {
  ID    string `json:"id"`
  Email string `json:"email"`
  Name  string `json:"name"`
  Exp   int64  `json:"exp,omitempty"`
  jwt.RegisteredClaims
}

func DecodeJwt(r *http.Request) (*CurrentUser, error) {
  authHeader := r.Header.Get("Authorization")
  parts := strings.Split(authHeader, " ")
  if len(parts) != 2 || parts[0] != "Bearer" {
    return nil, errors.New("invalid authorization header")
  }

  tokenString := parts[1]
  key := &env.Env.JwtSecret
  var userClaim CurrentUser

  _, err := jwt.ParseWithClaims(tokenString, &userClaim, func(token *jwt.Token) (interface{}, error) {
    return []byte(*key), nil
  })
  if err != nil {
    return nil, err
  }
  return &userClaim, nil
}
