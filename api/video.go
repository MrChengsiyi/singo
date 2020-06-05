package api

import (
	"singo/service"

	"github.com/gin-gonic/gin"
)

// CreateVideo 视频投稿接口
func CreateVideo(c *gin.Context) {
	var service service.CreateVideoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//ShowVideo 视频详情接口
func ShowVideo(c *gin.Context) {
	var service service.ShowVideoService
	res := service.Show(c.Param("ID"))
	c.JSON(200, res)
}

func ListVideo(c *gin.Context) {
	var service service.ListVideoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
