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

	//c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")

	if strings.HasSuffix(filePath, "pdf") {
		c.Writer.Header().Set("Content-Type", "application/pdf")
	} else if strings.HasSuffix(filePath, "docx") {
		c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	} else if strings.HasSuffix(filePath, "xlsx") {
		c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	} else if strings.HasSuffix(filePath, "xls") {
		c.Writer.Header().Set("Content-Type", "application/vnd.ms-excel")
	} else if strings.HasSuffix(filePath, "txt") {
		c.Writer.Header().Set("Content-Type", "text/plain")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "file extension not found"})
		return
	}

	c.File(filePath)
}
