package errors

import (
	"net/http"
)

type Rest struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *Rest {
	return &Rest{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *Rest {
	return &Rest{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string) *Rest {
	return &Rest{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

func NewForbiddenAccessError(message string) *Rest {
	return &Rest{
		Message: message,
		Status:  http.StatusForbidden,
		Error:   "forbidden_access_error",
	}
}

func NewUnauthorizedError(message string) *Rest {
	return &Rest{
		Message: message,
		Status:  http.StatusUnauthorized,
		Error:   "unauthorized_error",
	}
}

func NewError(status int, message string) *Rest {
	return &Rest{
		Message: message,
		Status:  status,
		Error:   "error",
	}
}

