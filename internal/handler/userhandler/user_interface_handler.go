package userhandler

import (
  "net/http"

  "github.com/wiliamvj/api-users-golang/internal/service/userservice"
)

func NewUserHandler(service userservice.UserService) UserHandler {
  return &handler{
    service,
  }
}

type handler struct {
  service userservice.UserService
}

type UserHandler interface {
  CreateUser(w http.ResponseWriter, r *http.Request)
}
