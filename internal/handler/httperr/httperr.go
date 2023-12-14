package httperr

import "net/http"

type RestErr struct {
  Message string   `json:"message"`
  Err     string   `json:"error,omitempty"`
  Code    int      `json:"code"`
  Fields  []Fields `json:"fields,omitempty"`
}

type Fields struct {
  Field   string      `json:"field"`
  Value   interface{} `json:"value,omitempty"`
  Message string      `json:"message"`
}

func (r *RestErr) Error() string {
  return r.Message
}

func NewRestErr(m, e string, c int, f []Fields) *RestErr {
  return &RestErr{
    Message: m,
    Err:     e,
    Code:    c,
    Fields:  f,
  }
}

func NewBadRequestError(message string) *RestErr {
  return &RestErr{
    Message: message,
    Err:     "bad_request",
    Code:    http.StatusBadRequest,
  }
}

func NewUnauthorizedRequestError(message string) *RestErr {
  return &RestErr{
    Message: message,
    Err:     "unauthorized",
    Code:    http.StatusUnauthorized,
  }
}

func NewBadRequestValidationError(m string, c []Fields) *RestErr {
  return &RestErr{
    Message: m,
    Err:     "bad_request",
    Code:    http.StatusBadRequest,
    Fields:  c,
  }
}

func NewInternalServerError(message string) *RestErr {
  return &RestErr{
    Message: message,
    Err:     "internal_server_error",
    Code:    http.StatusInternalServerError,
  }
}

func NewNotFoundError(message string) *RestErr {
  return &RestErr{
    Message: message,
    Err:     "not_found",
    Code:    http.StatusNotFound,
  }
}

func NewForbiddenError(message string) *RestErr {
  return &RestErr{
    Message: message,
    Err:     "forbidden",
    Code:    http.StatusForbidden,
  }
}
