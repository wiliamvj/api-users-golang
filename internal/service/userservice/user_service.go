package userservice

import (
  "context"
  "fmt"
  "time"

  "github.com/wiliamvj/api-users-golang/internal/dto"
  "github.com/wiliamvj/api-users-golang/internal/handler/response"
)

func (s *service) CreateUser(ctx context.Context, u dto.CreateUserDto) error {
  return nil
}

func (s *service) UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error {
  return nil
}

func (s *service) GetUserByID(ctx context.Context, id string) (*response.UserResponse, error) {
  userFake := response.UserResponse{
    ID:        "123",
    Name:      "John Doe",
    Email:     "jonh.doe@email.com",
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }
  return &userFake, nil
}

func (s *service) DeleteUser(ctx context.Context, id string) error {
  return nil
}

func (s *service) FindManyUsers(ctx context.Context) (response.ManyUsersResponse, error) {
  // create fake users
  usersFake := response.ManyUsersResponse{}
  for i := 0; i < 5; i++ {
    userFake := response.UserResponse{
      ID:        "123",
      Name:      "John Doe",
      Email:     fmt.Sprintf("jonh.doe-%v@email.com", i),
      CreatedAt: time.Now(),
      UpdatedAt: time.Now(),
    }
    usersFake.Users = append(usersFake.Users, userFake)
  }
  return usersFake, nil
}

func (s *service) UpdateUserPassword(ctx context.Context, u *dto.UpdateUserPasswordDto, id string) error {
  fmt.Println("new password: ", u.Password)
  fmt.Println("old password: ", u.OldPassword)
  return nil
}
