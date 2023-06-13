package constant

import (
	"net/http"

	"github.com/bagastri07/custom-error/customerror"
)

var (
	ErrInternalServerError = customerror.CustomError{
		Message: "Internal Server Err",
		Code:    http.StatusInternalServerError,
	}
	ErrCompanyNotFound = customerror.CustomError{
		Message: "Company Not Found",
		Code:    http.StatusNotFound,
	}
)
