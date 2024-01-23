package userhandler

import (
  "encoding/json"
  "fmt"
  "log/slog"
  "net/http"

  "github.com/wiliamvj/api-users-golang/internal/common/utils"
  "github.com/wiliamvj/api-users-golang/internal/dto"
  "github.com/wiliamvj/api-users-golang/internal/handler/httperr"
  "github.com/wiliamvj/api-users-golang/internal/handler/validation"
)

// Create user
//	@Summary		Create new user
//	@Description	Endpoint for create user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			body	body	dto.CreateUserDto	true	"Create user dto"	true
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/user [post]
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
  var req dto.CreateUserDto

  if r.Body == http.NoBody {
    slog.Error("body is empty", slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("body is required")
    json.NewEncoder(w).Encode(msg)
    return
  }
  err := json.NewDecoder(r.Body).Decode(&req)
  if err != nil {
    slog.Error("error to decode body", "err", err, slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("error to decode body")
    json.NewEncoder(w).Encode(msg)
    return
  }
  httpErr := validation.ValidateHttpData(req)
  if httpErr != nil {
    slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "userhandler"))
    w.WriteHeader(httpErr.Code)
    json.NewEncoder(w).Encode(httpErr)
    return
  }
  err = h.service.CreateUser(r.Context(), req)
  if err != nil {
    slog.Error(fmt.Sprintf("error to create user: %v", err), slog.String("package", "userhandler"))
    if err.Error() == "cep not found" {
      w.WriteHeader(http.StatusNotFound)
      msg := httperr.NewNotFoundError("cep not found")
      json.NewEncoder(w).Encode(msg)
      return
    }
    if err.Error() == "user already exists" {
      w.WriteHeader(http.StatusBadRequest)
      msg := httperr.NewBadRequestError("user already exists")
      json.NewEncoder(w).Encode(msg)
      return
    }
    w.WriteHeader(http.StatusInternalServerError)
    msg := httperr.NewBadRequestError("error to create user")
    json.NewEncoder(w).Encode(msg)
    return
  }
}

// Update user
//	@Summary		Update user
//	@Description	Endpoint for update user
//	@Tags			user
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			body	body	dto.UpdateUserDto	false	"Update user dto"	true
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/user [patch]
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
  var req dto.UpdateUserDto

  user, err := utils.DecodeJwt(r)
  if err != nil {
    slog.Error("error to decode jwt", slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("error to decode jwt")
    json.NewEncoder(w).Encode(msg)
    return
  }
  if r.Body == http.NoBody {
    slog.Error("body is empty", slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("body is required")
    json.NewEncoder(w).Encode(msg)
    return
  }
  err = json.NewDecoder(r.Body).Decode(&req)
  if err != nil {
    slog.Error("error to decode body", "err", err, slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("error to decode body")
    json.NewEncoder(w).Encode(msg)
    return
  }
  httpErr := validation.ValidateHttpData(req)
  if httpErr != nil {
    slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "userhandler"))
    w.WriteHeader(httpErr.Code)
    json.NewEncoder(w).Encode(httpErr)
    return
  }
  err = h.service.UpdateUser(r.Context(), req, user.ID)
  if err != nil {
    slog.Error(fmt.Sprintf("error to update user: %v", err), slog.String("package", "userhandler"))
    if err.Error() == "user not found" {
      w.WriteHeader(http.StatusNotFound)
      msg := httperr.NewNotFoundError("user not found")
      json.NewEncoder(w).Encode(msg)
      return
    }
    if err.Error() == "cep not found" {
      w.WriteHeader(http.StatusNotFound)
      msg := httperr.NewNotFoundError("cep not found")
      json.NewEncoder(w).Encode(msg)
      return
    }
    if err.Error() == "user already exists" {
      w.WriteHeader(http.StatusBadRequest)
      msg := httperr.NewBadRequestError("user already exists with this email")
      json.NewEncoder(w).Encode(msg)
      return
    }
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(err)
    return
  }
}

// User details
//	@Summary		User details
//	@Description	Get user by id
//	@Tags			user
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"user id"
//	@Success		200	{object}	response.UserResponse
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/user [get]
func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
  user, err := utils.DecodeJwt(r)
  if err != nil {
    slog.Error("error to decode jwt", slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("error to decode jwt")
    json.NewEncoder(w).Encode(msg)
    return
  }
  res, err := h.service.GetUserByID(r.Context(), user.ID)
  if err != nil {
    slog.Error(fmt.Sprintf("error to get user: %v", err), slog.String("package", "userhandler"))
    if err.Error() == "user not found" {
      w.WriteHeader(http.StatusNotFound)
      msg := httperr.NewNotFoundError("user not found")
      json.NewEncoder(w).Encode(msg)
      return
    }
    w.WriteHeader(http.StatusInternalServerError)
    msg := httperr.NewBadRequestError("error to get user")
    json.NewEncoder(w).Encode(msg)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(res)
}

// Delete user
//	@Summary		Delete user
//	@Description	delete user by id
//	@Tags			user
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"user id"
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/user [delete]
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
  user, err := utils.DecodeJwt(r)
  if err != nil {
    slog.Error("error to decode jwt", slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("error to decode jwt")
    json.NewEncoder(w).Encode(msg)
    return
  }
  err = h.service.DeleteUser(r.Context(), user.ID)
  if err != nil {
    slog.Error(fmt.Sprintf("error to delete user: %v", err), slog.String("package", "userhandler"))
    if err.Error() == "user not found" {
      w.WriteHeader(http.StatusNotFound)
      msg := httperr.NewNotFoundError("user not found")
      json.NewEncoder(w).Encode(msg)
      return
    }
    w.WriteHeader(http.StatusInternalServerError)
    msg := httperr.NewBadRequestError("error to delete user")
    json.NewEncoder(w).Encode(msg)
    return
  }
  w.WriteHeader(http.StatusNoContent)
}

// Get many user
//	@Summary		Get many users
//	@Description	Get many users
//	@Tags			user
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.ManyUsersResponse
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/user [get]
func (h *handler) FindManyUsers(w http.ResponseWriter, r *http.Request) {
  res, err := h.service.FindManyUsers(r.Context())
  if err != nil {
    slog.Error(fmt.Sprintf("error to find many users: %v", err), slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusInternalServerError)
    msg := httperr.NewBadRequestError("error to find many users")
    json.NewEncoder(w).Encode(msg)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(res)
}

// Update user password
//	@Summary		Update user password
//	@Description	Endpoint for Update user password
//	@Tags			user
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			id		path	string						true	"user id"
//	@Param			body	body	dto.UpdateUserPasswordDto	true	"Update user password dto"	true
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/user/password [get]
func (h *handler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
  var req dto.UpdateUserPasswordDto

  user, err := utils.DecodeJwt(r)
  if err != nil {
    slog.Error("error to decode jwt", slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("error to decode jwt")
    json.NewEncoder(w).Encode(msg)
    return
  }
  if r.Body == http.NoBody {
    slog.Error("body is empty", slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("body is required")
    json.NewEncoder(w).Encode(msg)
    return
  }
  err = json.NewDecoder(r.Body).Decode(&req)
  if err != nil {
    slog.Error("error to decode body", "err", err, slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("error to decode body")
    json.NewEncoder(w).Encode(msg)
    return
  }
  httpErr := validation.ValidateHttpData(req)
  if httpErr != nil {
    slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "userhandler"))
    w.WriteHeader(httpErr.Code)
    json.NewEncoder(w).Encode(httpErr)
    return
  }
  err = h.service.UpdateUserPassword(r.Context(), &req, user.ID)
  if err != nil {
    slog.Error(fmt.Sprintf("error to update user password: %v", err), slog.String("package", "userhandler"))
    if err.Error() == "user not found" {
      w.WriteHeader(http.StatusNotFound)
      msg := httperr.NewNotFoundError("user not found")
      json.NewEncoder(w).Encode(msg)
      return
    }
    w.WriteHeader(http.StatusInternalServerError)
    msg := httperr.NewBadRequestError("error to update user password")
    json.NewEncoder(w).Encode(msg)
    return
  }
}
