package userservice

import (
  "context"
  "errors"
  "log/slog"
  "time"

  "github.com/wiliamvj/api-users-golang/config/env"
  "github.com/wiliamvj/api-users-golang/internal/dto"
  "github.com/wiliamvj/api-users-golang/internal/handler/response"
  "golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, u dto.LoginDTO) (*response.UserAuthToken, error) {
  user, err := s.repo.FindUserByEmail(ctx, u.Email)
  if err != nil {
    slog.Error("error to search user by email", "err", err, slog.String("package", "userservice"))
    return nil, errors.New("error to search user password")
  }
  if user == nil {
    slog.Error("user not found", slog.String("package", "userservice"))
    return nil, errors.New("user not found")
  }
  userPass, err := s.repo.GetUserPassword(ctx, user.ID)
  if err != nil {
    slog.Error("error to search user password", "err", err, slog.String("package", "userservice"))
    return nil, errors.New("error to search user password")
  }
  // compare password with password in database
  err = bcrypt.CompareHashAndPassword([]byte(userPass.Password), []byte(u.Password))
  if err != nil {
    slog.Error("invalid password", slog.String("package", "service_user"))
    return nil, errors.New("invalid password")
  }
  _, token, _ := env.Env.TokenAuth.Encode(map[string]interface{}{
    "id":    user.ID,
    "email": u.Email,
    "name":  user.Name,
    "exp":   time.Now().Add(time.Second * time.Duration(env.Env.JwtExpiresIn)).Unix(),
  })
  userAuthToken := response.UserAuthToken{
    AccessToken: token,
  }
  return &userAuthToken, nil
}
