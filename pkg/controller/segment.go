package controller

import (
	"admin_panel/models"
	"admin_panel/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateSegment(c *gin.Context)  {
	var segment models.Segment

	err := c.ShouldBindJSON(&segment)
	if err != nil {
		c.JSON(400, gin.H{"reason": err.Error()})
		return
	}
	err = service.CreateSegment(segment)
	if err != nil {
		c.JSON(400, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"reason": "сегмент успешно создан"})

}
