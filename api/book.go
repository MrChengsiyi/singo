package api

import (
"singo/service"

"github.com/gin-gonic/gin"
)

// CreateBook 书籍创建接口
func CreateBook(c *gin.Context) {
	var service service.CreateBookService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//ShowVideo 书籍详情接口
func ShowBook(c *gin.Context) {
	var service service.ShowBookService

	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

func ListBook(c *gin.Context) {
	var service service.ListBookService
	res := service.List()
	c.JSON(200, res)
	//if err := c.ShouldBind(&service); err == nil {
	//	res := service.List()
	//	c.JSON(200, res)
	//} else {
	//	c.JSON(200, ErrorResponse(err))
	//}
}

//书籍删除接口
func DeleteBook(c *gin.Context) {
	var service service.DeleteBookService
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}


//书籍信息更新接口
func UpdateBook(c *gin.Context) {
	var service service.UpdateBookService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

