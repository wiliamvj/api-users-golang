package productrepository

import (
  "context"
  "database/sql"

  "github.com/wiliamvj/api-users-golang/internal/database/sqlc"
  "github.com/wiliamvj/api-users-golang/internal/dto"
  "github.com/wiliamvj/api-users-golang/internal/entity"
)

func NewProductRepository(db *sql.DB, q *sqlc.Queries) ProductRepository {
  return &repository{
    db,
    q,
  }
}

type repository struct {
  db      *sql.DB
  queries *sqlc.Queries
}

type ProductRepository interface {
  CreateProduct(ctx context.Context, p *entity.ProductEntity, c []entity.ProductCategoryEntity) error
  GetCategoryByID(ctx context.Context, id string) (bool, error)
  GetProductByID(ctx context.Context, id string) (bool, error)
  UpdateProduct(ctx context.Context, p *entity.ProductEntity, c []entity.ProductCategoryEntity) error
  GetCategoriesByProductID(ctx context.Context, id string) ([]string, error)
  DeleteProductCategory(ctx context.Context, productID, categoryID string) error
  DeleteProduct(ctx context.Context, id string) error
  FindManyProducts(ctx context.Context, d dto.FindProductDto) ([]entity.ProductWithCategoryEntity, error)
}
