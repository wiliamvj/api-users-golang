package database

import (
  "database/sql"
  "log/slog"

  _ "github.com/lib/pq"
  "github.com/wiliamvj/api-users-golang/config/env"
)

func NewDBConnection() (*sql.DB, error) {
  postgresURI := env.Env.DatabaseURL
  db, err := sql.Open("postgres", postgresURI)
  if err != nil {
    return nil, err
  }
  err = db.Ping()
  if err != nil {
    db.Close()
    return nil, err
  }
  slog.Info("database connected", slog.String("package", "database"))

  return db, nil
}
