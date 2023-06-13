package response

import (
	"github.com/bagastri07/custom-error/constant"
	cErr "github.com/bagastri07/custom-error/customerror"
	"github.com/gin-gonic/gin"
)

// HandleError handles the error and sends an appropriate JSON response
func HandleError(c *gin.Context, err error) {
	if customErr, ok := err.(*cErr.CustomError); ok {
		c.JSON(customErr.Code, customErr)
	} else {
		c.JSON(constant.ErrInternalServerError.Code, constant.ErrInternalServerError.SetErrCode("undefined"))
	}
}
