package productrepository

import (
  "context"
  "database/sql"
  "log/slog"

  "github.com/wiliamvj/api-users-golang/internal/database/sqlc"
  "github.com/wiliamvj/api-users-golang/internal/dto"
  "github.com/wiliamvj/api-users-golang/internal/entity"
  "github.com/wiliamvj/api-users-golang/internal/repository/transaction"
)

func (r *repository) CreateProduct(ctx context.Context, p *entity.ProductEntity, c []entity.ProductCategoryEntity) error {
  err := transaction.Run(ctx, r.db, func(q *sqlc.Queries) error {
    var err error
    err = q.CreateProduct(ctx, sqlc.CreateProductParams{
      ID:          p.ID,
      Title:       p.Title,
      Price:       p.Price,
      Description: sql.NullString{String: p.Description, Valid: p.Description != ""},
      CreatedAt:   p.CreatedAt,
      UpdatedAt:   p.UpdatedAt,
    })
    if err != nil {
      return err
    }
    for _, category := range c {
      err = q.CreateProductCategory(ctx, sqlc.CreateProductCategoryParams{
        ID:         category.ID,
        ProductID:  p.ID,
        CategoryID: category.CategoryID,
        CreatedAt:  category.CreatedAt,
        UpdatedAt:  category.UpdatedAt,
      })
      if err != nil {
        return err
      }
    }
    return nil
  })
  if err != nil {
    slog.Error("error to create product, roll back applied", "err", err)
    return err
  }
  return nil
}

func (r *repository) GetCategoryByID(ctx context.Context, id string) (bool, error) {
  exists, err := r.queries.GetCategoryByID(ctx, id)
  if err != nil || err == sql.ErrNoRows {
    return false, err
  }
  return exists, nil
}

func (r *repository) GetProductByID(ctx context.Context, id string) (bool, error) {
  exists, err := r.queries.GetProductByID(ctx, id)
  if err != nil || err == sql.ErrNoRows {
    return false, err
  }
  return exists, nil
}

func (r *repository) UpdateProduct(ctx context.Context, p *entity.ProductEntity, c []entity.ProductCategoryEntity) error {
  err := transaction.Run(ctx, r.db, func(q *sqlc.Queries) error {
    var err error
    err = q.UpdateProduct(ctx, sqlc.UpdateProductParams{
      ID:          p.ID,
      Title:       sql.NullString{String: p.Title, Valid: p.Title != ""},
      Price:       sql.NullInt32{Int32: p.Price, Valid: p.Price != 0},
      Description: sql.NullString{String: p.Description, Valid: p.Description != ""},
      UpdatedAt:   p.UpdatedAt,
    })
    if err != nil {
      return err
    }
    for _, category := range c {
      err = q.CreateProductCategory(ctx, sqlc.CreateProductCategoryParams{
        ID:         category.ID,
        ProductID:  p.ID,
        CategoryID: category.CategoryID,
        CreatedAt:  category.CreatedAt,
        UpdatedAt:  category.UpdatedAt,
      })
      if err != nil {
        return err
      }
    }
    return nil
  })
  if err != nil {
    slog.Error("error to update product, roll back applied", "err", err)
    return err
  }
  return nil
}

func (r *repository) GetCategoriesByProductID(ctx context.Context, id string) ([]string, error) {
  categories, err := r.queries.GetCategoriesByProductID(ctx, id)
  if err != nil {
    return nil, err
  }
  return categories, nil
}

func (r *repository) DeleteProductCategory(ctx context.Context, productID, categoryID string) error {
  err := r.queries.DeleteProductCategory(ctx, sqlc.DeleteProductCategoryParams{
    ProductID:  productID,
    CategoryID: categoryID,
  })
  if err != nil {
    return err
  }
  return nil
}

func (r *repository) DeleteProduct(ctx context.Context, id string) error {
  err := r.queries.DeleteProduct(ctx, id)
  if err != nil {
    return err
  }
  return nil
}

func (r *repository) FindManyProducts(ctx context.Context, d dto.FindProductDto) ([]entity.ProductWithCategoryEntity, error) {
  products, err := r.queries.FindManyProducts(ctx, sqlc.FindManyProductsParams{
    Categories: d.Categories,
    Search:     sql.NullString{String: d.Search, Valid: d.Search != ""},
  })
  if err != nil {
    return nil, err
  }
  var response []entity.ProductWithCategoryEntity
  for _, p := range products {
    var category []entity.CategoryEntity
    categories, err := r.queries.GetProductCategories(ctx, p.ID)
    if err != nil {
      return nil, err
    }
    for _, c := range categories {
      category = append(category, entity.CategoryEntity{
        ID:    c.ID,
        Title: c.Title,
      })
    }
    response = append(response, entity.ProductWithCategoryEntity{
      ID:          p.ID,
      Title:       p.Title,
      Description: p.Description.String,
      Price:       p.Price,
      Categories:  category,
      CreatedAt:   p.CreatedAt,
    })
  }
  return response, nil
}
