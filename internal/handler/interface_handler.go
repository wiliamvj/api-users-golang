package handler

import (
  "net/http"

  "github.com/wiliamvj/api-users-golang/internal/service/categoryservice"
  "github.com/wiliamvj/api-users-golang/internal/service/productservice"
  "github.com/wiliamvj/api-users-golang/internal/service/userservice"
)

func NewHandler(userService userservice.UserService,
  categoryService categoryservice.CategoryService,
  productservice productservice.ProductService) Handler {
  return &handler{
    userService:     userService,
    categoryService: categoryService,
    productservice:  productservice,
  }
}

type handler struct {
  userService     userservice.UserService
  categoryService categoryservice.CategoryService
  productservice  productservice.ProductService
}

type Handler interface {
  CreateUser(w http.ResponseWriter, r *http.Request)
  UpdateUser(w http.ResponseWriter, r *http.Request)
  GetUserByID(w http.ResponseWriter, r *http.Request)
  DeleteUser(w http.ResponseWriter, r *http.Request)
  FindManyUsers(w http.ResponseWriter, r *http.Request)
  UpdateUserPassword(w http.ResponseWriter, r *http.Request)
  Login(w http.ResponseWriter, r *http.Request)

  CreateCategory(w http.ResponseWriter, r *http.Request)

  CreateProduct(w http.ResponseWriter, r *http.Request)
  UpdateProduct(w http.ResponseWriter, r *http.Request)
  DeleteProduct(w http.ResponseWriter, r *http.Request)
  FindManyProducts(w http.ResponseWriter, r *http.Request)
}
