package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func UploadToken(c *gin.Context) {
	var service service.UploadTokenService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Post()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
