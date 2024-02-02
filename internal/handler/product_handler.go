package handler

import (
  "encoding/json"
  "fmt"
  "log/slog"
  "net/http"

  "github.com/go-chi/chi"
  "github.com/google/uuid"
  "github.com/wiliamvj/api-users-golang/internal/dto"
  "github.com/wiliamvj/api-users-golang/internal/handler/httperr"
  "github.com/wiliamvj/api-users-golang/internal/handler/validation"
)

func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
  var req dto.CreateProductDto

  if r.Body == http.NoBody {
    slog.Error("body is empty", slog.String("package", "producthandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("body is required")
    json.NewEncoder(w).Encode(msg)
    return
  }
  err := json.NewDecoder(r.Body).Decode(&req)
  if err != nil {
    slog.Error("error to decode body", "err", err, slog.String("package", "producthandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("error to decode body")
    json.NewEncoder(w).Encode(msg)
    return
  }
  httpErr := validation.ValidateHttpData(req)
  if httpErr != nil {
    slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "producthandler"))
    w.WriteHeader(httpErr.Code)
    json.NewEncoder(w).Encode(httpErr)
    return
  }

  err = h.productservice.CreateProduct(r.Context(), req)
  if err != nil {
    if err.Error() == "category not found" {
      w.WriteHeader(http.StatusNotFound)
      msg := httperr.NewNotFoundError("category not found")
      json.NewEncoder(w).Encode(msg)
      return
    }
    slog.Error(fmt.Sprintf("error to create category: %v", err), slog.String("package", "categoryhandler"))
    w.WriteHeader(http.StatusBadRequest)
  }
  w.WriteHeader(http.StatusCreated)
}

func (h *handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
  var req dto.UpdateProductDto

  productID := chi.URLParam(r, "id")
  if productID == "" {
    slog.Error("product id is required", slog.String("package", "producthandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("product id is required")
    json.NewEncoder(w).Encode(msg)
    return
  }
  _, err := uuid.Parse(productID)
  if err != nil {
    slog.Error(fmt.Sprintf("error to parse product id: %v", err), slog.String("package", "producthandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("invalid product id")
    json.NewEncoder(w).Encode(msg)
    return
  }
  if r.Body == http.NoBody {
    slog.Error("body is empty", slog.String("package", "producthandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("body is required")
    json.NewEncoder(w).Encode(msg)
    return
  }
  err = json.NewDecoder(r.Body).Decode(&req)
  if err != nil {
    slog.Error("error to decode body", "err", err, slog.String("package", "producthandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("error to decode body")
    json.NewEncoder(w).Encode(msg)
    return
  }
  httpErr := validation.ValidateHttpData(req)
  if httpErr != nil {
    slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "producthandler"))
    w.WriteHeader(httpErr.Code)
    json.NewEncoder(w).Encode(httpErr)
    return
  }
  err = h.productservice.UpdateProduct(r.Context(), productID, req)
  if err != nil {
    if err.Error() == "product not found" {
      w.WriteHeader(http.StatusNotFound)
      msg := httperr.NewNotFoundError("product not found")
      json.NewEncoder(w).Encode(msg)
      return
    }
    if err.Error() == "category not found" {
      w.WriteHeader(http.StatusNotFound)
      msg := httperr.NewNotFoundError("category not found")
      json.NewEncoder(w).Encode(msg)
      return
    }
    slog.Error(fmt.Sprintf("error to update category: %v", err), slog.String("package", "categoryhandler"))
    w.WriteHeader(http.StatusBadRequest)
  }
  w.WriteHeader(http.StatusOK)
}

func (h *handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
  productID := chi.URLParam(r, "id")
  if productID == "" {
    slog.Error("product id is required", slog.String("package", "producthandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("product id is required")
    json.NewEncoder(w).Encode(msg)
    return
  }
  _, err := uuid.Parse(productID)
  if err != nil {
    slog.Error(fmt.Sprintf("error to parse product id: %v", err), slog.String("package", "producthandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("invalid product id")
    json.NewEncoder(w).Encode(msg)
    return
  }
  err = h.productservice.DeleteProduct(r.Context(), productID)
  if err != nil {
    if err.Error() == "product not found" {
      w.WriteHeader(http.StatusNotFound)
      msg := httperr.NewNotFoundError("product not found")
      json.NewEncoder(w).Encode(msg)
      return
    }
    slog.Error(fmt.Sprintf("error to delete category: %v", err), slog.String("package", "categoryhandler"))
    w.WriteHeader(http.StatusBadRequest)
  }
  w.WriteHeader(http.StatusOK)
}

func (h *handler) FindManyProducts(w http.ResponseWriter, r *http.Request) {
  var req dto.FindProductDto

  err := json.NewDecoder(r.Body).Decode(&req)
  if err != nil {
    slog.Error("error to decode body", "err", err, slog.String("package", "producthandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("error to decode body")
    json.NewEncoder(w).Encode(msg)
    return
  }
  httpErr := validation.ValidateHttpData(req)
  if httpErr != nil {
    slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "producthandler"))
    w.WriteHeader(httpErr.Code)
    json.NewEncoder(w).Encode(httpErr)
    return
  }
  products, err := h.productservice.FindManyProducts(r.Context(), req)
  if err != nil {
    slog.Error(fmt.Sprintf("error to find many products: %v", err), slog.String("package", "producthandler"))
    w.WriteHeader(http.StatusBadRequest)
  }
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(products)
}
