package userrepository

import (
  "context"
  "database/sql"
  "time"

  "github.com/google/uuid"
  "github.com/wiliamvj/api-users-golang/internal/database/sqlc"
  "github.com/wiliamvj/api-users-golang/internal/entity"
)

func (r *repository) CreateUser(ctx context.Context, u *entity.UserEntity) error {
  err := r.queries.CreateUser(ctx, sqlc.CreateUserParams{
    ID:        u.ID,
    Name:      u.Name,
    Email:     u.Email,
    Password:  u.Password,
    CreatedAt: u.CreatedAt,
    UpdatedAt: u.UpdatedAt,
  })
  if err != nil {
    return err
  }
  err = r.queries.CreateUserAddress(ctx, sqlc.CreateUserAddressParams{
    ID:         uuid.New().String(),
    UserID:     u.ID,
    Cep:        u.Address.CEP,
    Ibge:       u.Address.IBGE,
    Uf:         u.Address.UF,
    City:       u.Address.City,
    Complement: sql.NullString{String: u.Address.Complement, Valid: u.Address.Complement != ""},
    Street:     u.Address.Street,
    CreatedAt:  time.Now(),
    UpdatedAt:  time.Now(),
  })
  if err != nil {
    return err
  }
  return nil
}

func (r *repository) FindUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
  user, err := r.queries.FindUserByEmail(ctx, email)
  if err != nil {
    return nil, err
  }
  userEntity := entity.UserEntity{
    ID:    user.ID,
    Name:  user.Name,
    Email: user.Email,
  }
  return &userEntity, nil
}

func (r *repository) FindUserByID(ctx context.Context, id string) (*entity.UserEntity, error) {
  user, err := r.queries.FindUserByID(ctx, id)
  if err != nil {
    return nil, err
  }
  userEntity := entity.UserEntity{
    ID:    user.ID,
    Name:  user.Name,
    Email: user.Email,
    Address: entity.UserAddress{
      CEP:        user.Cep,
      UF:         user.Uf,
      City:       user.City,
      Complement: user.Complement.String,
      Street:     user.Street,
    },
    CreatedAt: user.CreatedAt,
    UpdatedAt: user.UpdatedAt,
  }
  return &userEntity, nil
}

func (r *repository) UpdateUser(ctx context.Context, u *entity.UserEntity) error {
  err := r.queries.UpdateUser(ctx, sqlc.UpdateUserParams{
    ID:        u.ID,
    Name:      sql.NullString{String: u.Name, Valid: u.Name != ""},
    Email:     sql.NullString{String: u.Email, Valid: u.Email != ""},
    UpdatedAt: u.UpdatedAt,
  })
  if err != nil {
    return err
  }
  err = r.queries.UpdateUserAddress(ctx, sqlc.UpdateUserAddressParams{
    UserID:     u.ID,
    Cep:        sql.NullString{String: u.Address.CEP, Valid: u.Address.CEP != ""},
    Ibge:       sql.NullString{String: u.Address.IBGE, Valid: u.Address.IBGE != ""},
    Uf:         sql.NullString{String: u.Address.UF, Valid: u.Address.UF != ""},
    City:       sql.NullString{String: u.Address.City, Valid: u.Address.City != ""},
    Complement: sql.NullString{String: u.Address.Complement, Valid: u.Address.Complement != ""},
    Street:     sql.NullString{String: u.Address.Street, Valid: u.Address.Street != ""},
    UpdatedAt:  time.Now(),
  })
  if err != nil {
    return err
  }
  return nil
}

func (r *repository) DeleteUser(ctx context.Context, id string) error {
  err := r.queries.DeleteUser(ctx, id)
  if err != nil {
    return err
  }
  return nil
}

func (r *repository) FindManyUsers(ctx context.Context) ([]entity.UserEntity, error) {
  users, err := r.queries.FindManyUsers(ctx)
  if err != nil {
    return nil, err
  }
  var usersEntity []entity.UserEntity
  for _, user := range users {
    userEntity := entity.UserEntity{
      ID:    user.ID,
      Name:  user.Name,
      Email: user.Email,
      Address: entity.UserAddress{
        CEP:        user.Cep,
        UF:         user.Uf,
        City:       user.City,
        Street:     user.Street,
        Complement: user.Complement.String,
      },
      CreatedAt: user.CreatedAt,
      UpdatedAt: user.UpdatedAt,
    }
    usersEntity = append(usersEntity, userEntity)
  }
  return usersEntity, nil
}

func (r *repository) UpdatePassword(ctx context.Context, pass, id string) error {
  err := r.queries.UpdatePassword(ctx, sqlc.UpdatePasswordParams{
    ID:        id,
    Password:  pass,
    UpdatedAt: time.Now(),
  })
  if err != nil {
    return err
  }
  return nil
}

func (r *repository) GetUserPassword(ctx context.Context, id string) (string, error) {
  pass, err := r.queries.GetUserPassword(ctx, id)
  if err != nil {
    return "", err
  }
  return pass, nil
}
