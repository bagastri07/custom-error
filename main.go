package main

import (
	"errors"
	"fmt"

	"github.com/bagastri07/custom-error/constant"
	"github.com/bagastri07/custom-error/response"
	"github.com/gin-gonic/gin"
)

// MyService represents a service object
type MyService struct {
	CodeName string
}

func NewService() *MyService {
	return &MyService{
		CodeName: "03",
	}
}

// SomeMethod is a method of MyService that can return an error
func (s *MyService) SomeMethod() error {
	// Some logic that may result in an error
	err := someBusinessLogic()

	// Check if there was an error and return a custom error instance if necessary
	if err != nil {
		// return &customerror.CustomError{
		// 	Message: "meh",
		// 	Code:    http.StatusBadGateway,
		// 	ErrCode: s.CodeName,
		// }

		// doing log
		return constant.ErrCompanyNotFound.SetErrCode(s.CodeName)
	}

	// No error occurred
	return nil
}

// someBusinessLogic simulates some business logic that can return an error
func someBusinessLogic() error {
	// Simulate an error condition
	return errors.New("Simulated error")
}

func main() {
	router := gin.Default()
	service := NewService()

	router.GET("/endpoint", func(c *gin.Context) {
		err := service.SomeMethod()
		if err != nil {
			// Use the helper method to handle the error
			response.HandleError(c, err)
			return
		}

		// No error occurred
		c.JSON(200, gin.H{"message": "Success"})
	})

	err := router.Run(":4545")
	if err != nil {
		fmt.Println("Failed to start the server:", err)
	}
}
