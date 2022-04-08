package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("name")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	timeSign := fmt.Sprintf("%d", time.Now().UnixNano())

	filePath := fmt.Sprintf("files/layouts_%s_%s", timeSign, file.Filename)
	filePath = strings.Replace(filePath, " ", "", 111)

	if err = c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"file_path": filePath})
}

func DownloadFile(c *gin.Context) {
	filePath := c.Query("file_path")

	c.File(filePath)
}
