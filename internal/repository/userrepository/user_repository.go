package userrepository

import (
  "context"

  "github.com/wiliamvj/api-users-golang/internal/entity"
)

func (r *repository) CreateUser(ctx context.Context, u *entity.UserEntity) error {
  return nil
}

func (r *repository) FindUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
  return nil, nil
}

func (r *repository) FindUserByID(ctx context.Context, id string) (*entity.UserEntity, error) {
  return nil, nil
}

func (r *repository) UpdateUser(ctx context.Context, u *entity.UserEntity) error {
  return nil
}

func (r *repository) DeleteUser(ctx context.Context, id string) error {
  return nil
}

func (r *repository) FindManyUsers(ctx context.Context) ([]entity.UserEntity, error) {
  return nil, nil
}

func (r *repository) UpdatePassword(ctx context.Context, pass, id string) error {
  return nil
}
