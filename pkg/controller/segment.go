package controller

import (
	"admin_panel/models"
	"admin_panel/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateSegment(c *gin.Context) {
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

func GetSegmentByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "ERROR"})
		return
	}
	// по id берем поля из потребностей
	// потом из некоторых полей берем код сегмента и его данные
	// после чего должны соеденить потребности и сегменты в экселе и отправить на почту файл

	segment, err := service.GetSegmentByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, segment)
}

func SendLetter(c *gin.Context) {
	formedGraphicID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "ERROR"})
		return
	}

	formedGraphic, err := service.GetFormedGraphicByID(formedGraphicID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	formedGraphicProducts, err := service.GetAllFormedGraphicsProducts(formedGraphicID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

}
