package categoryrepository

import (
  "context"

  "github.com/wiliamvj/api-users-golang/internal/database/sqlc"
  "github.com/wiliamvj/api-users-golang/internal/entity"
)

func (r *repository) CreateCategory(ctx context.Context, c *entity.CategoryEntity) error {
  err := r.queries.CreateCategory(ctx, sqlc.CreateCategoryParams{
    ID:        c.ID,
    Title:     c.Title,
    CreatedAt: c.CreatedAt,
    UpdatedAt: c.UpdatedAt,
  })
  if err != nil {
    return err
  }
  return nil
}
