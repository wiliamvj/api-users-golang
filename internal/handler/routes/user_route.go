package routes

import (
  "github.com/go-chi/chi"
  "github.com/wiliamvj/api-users-golang/internal/handler/userhandler"
)

func InitUserRoutes(router chi.Router, h userhandler.UserHandler) {
  router.Route("/user", func(r chi.Router) {
    r.Post("/", h.CreateUser)
    r.Patch("/{id}", h.UpdateUser)
    r.Get("/{id}", h.GetUserByID)
    r.Delete("/{id}", h.DeleteUser)
    r.Get("/", h.FindManyUsers)
    r.Patch("/password/{id}", h.UpdateUserPassword)
  })
}
