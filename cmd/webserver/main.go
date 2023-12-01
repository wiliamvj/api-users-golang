package main

import (
  "log/slog"

  "github.com/wiliamvj/api-users-golang/config/logger"
)

func main() {
  logger.InitLogger()

  slog.Info("starting api")
}
