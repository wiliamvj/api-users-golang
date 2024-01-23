package transaction

import (
  "context"
  "database/sql"
  "fmt"

  "github.com/wiliamvj/api-users-golang/internal/database/sqlc"
)

func Run(ctx context.Context, c *sql.DB, fn func(*sqlc.Queries) error) error {
  tx, err := c.BeginTx(ctx, nil)
  if err != nil {
    return err
  }
  q := sqlc.New(tx)
  err = fn(q)
  if err != nil {
    if errRb := tx.Rollback(); errRb != nil {
      return fmt.Errorf("error on rollback: %v, original error: %w", errRb, err)
    }
    return err
  }
  return tx.Commit()
}
