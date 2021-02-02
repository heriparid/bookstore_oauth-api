package errors

import "net/http"

// RestErr struct to inform error
type RestErr struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:error`
}

// NewBadRequestError create new object
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

// NewNotFoundRequestError object
func NewNotFoundRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

// NewInternalServerError func
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
