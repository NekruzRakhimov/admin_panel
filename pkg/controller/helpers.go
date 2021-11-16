package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetImages(c *gin.Context) {
	imageName := c.Param("image_name")
	f1 := c.Param("f1") // first folder
	f2 := c.Param("f2") // second folder

	filePath := fmt.Sprintf("./files/%s/%s/%s", f1, f2, imageName)

	c.File(filePath)
}