package res

import (
	"github.com/gin-gonic/gin"
)

func Res(c *gin.Context, code int, status bool, message string, data interface{}) {

	response := gin.H{
		"status":  status,
		"message": message,
	}

	if data == nil {
		data = nil
	} else {
		response["data"] = data
	}

	c.JSON(code, response)
}
