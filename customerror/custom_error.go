package customerror

// CustomError is a custom error struct that implements the error interface
type CustomError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	ErrCode string `json:"errCode"`
}

// Error returns the error message
func (ce *CustomError) Error() string {
	return ce.Message
}

func (ce *CustomError) SetErrCode(errCode string) *CustomError {
	ce.ErrCode = errCode
	return ce
}
