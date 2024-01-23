package routes

import (
  "github.com/go-chi/chi"
  "github.com/go-chi/jwtauth"
  "github.com/wiliamvj/api-users-golang/config/env"

  "github.com/wiliamvj/api-users-golang/internal/handler/middleware"
  "github.com/wiliamvj/api-users-golang/internal/handler/userhandler"
)

func InitUserRoutes(router chi.Router, h userhandler.UserHandler) {
  router.Use(middleware.LoggerData)

  router.Post("/user", h.CreateUser)
  router.Route("/", func(r chi.Router) {
    r.Use(jwtauth.Verifier(env.Env.TokenAuth))
    r.Use(jwtauth.Authenticator)

    r.Patch("/user", h.UpdateUser)
    r.Get("/user", h.GetUserByID)
    r.Delete("/user", h.DeleteUser)
    r.Get("/user/list-all", h.FindManyUsers)
    r.Patch("/user/password", h.UpdateUserPassword)
  })
  router.Route("/auth", func(r chi.Router) {
    r.Post("/login", h.Login)
  })
}
