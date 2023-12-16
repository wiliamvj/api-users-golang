package main

import (
  "fmt"
  "log/slog"
  "net/http"

  "github.com/go-chi/chi"
  "github.com/wiliamvj/api-users-golang/config/env"
  "github.com/wiliamvj/api-users-golang/config/logger"
  _ "github.com/wiliamvj/api-users-golang/docs"
  "github.com/wiliamvj/api-users-golang/internal/database"
  "github.com/wiliamvj/api-users-golang/internal/database/sqlc"
  "github.com/wiliamvj/api-users-golang/internal/handler/routes"
  "github.com/wiliamvj/api-users-golang/internal/handler/userhandler"
  "github.com/wiliamvj/api-users-golang/internal/repository/userrepository"
  "github.com/wiliamvj/api-users-golang/internal/service/userservice"
)

func main() {
  logger.InitLogger()
  slog.Info("starting api")

  _, err := env.LoadingConfig(".")
  if err != nil {
    slog.Error("failed to load environment variables", err, slog.String("package", "main"))
    return
  }
  dbConnection, err := database.NewDBConnection()
  if err != nil {
    slog.Error("error to connect to database", "err", err, slog.String("package", "main"))
    return
  }

  queries := sqlc.New(dbConnection)

  // user
  userRepo := userrepository.NewUserRepository(dbConnection, queries)
  newUserService := userservice.NewUserService(userRepo)
  newUserHandler := userhandler.NewUserHandler(newUserService)

  // init routes
  router := chi.NewRouter()
  routes.InitUserRoutes(router, newUserHandler)
  routes.InitDocsRoutes(router)

  port := fmt.Sprintf(":%s", env.Env.GoPort)
  slog.Info(fmt.Sprintf("server running on port %s", port))
  err = http.ListenAndServe(port, router)
  if err != nil {
    slog.Error("error to start server", err, slog.String("package", "main"))
  }
}
