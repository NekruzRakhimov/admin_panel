package controller

import (
	"admin_panel/models"
	"admin_panel/pkg/service"
	"fmt"
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
	fmt.Println(formedGraphicID, "ID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "ERROR"})
		return
	}

	//сформированый потребность граффика -
	formedGraphic, err := service.GetFormedGraphicByID(formedGraphicID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	fmt.Println("DATA", formedGraphic)

	formedGraphicProducts, err := service.GetAllFormedGraphicsProducts(formedGraphicID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	graphic, err := service.GetGraphicByID(formedGraphic.GraphicID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	fmt.Println("GRAP", graphic)

	c.JSON(200, gin.H{
		"formedGraphic":         formedGraphic,
		"formedGraphicProducts": formedGraphicProducts,
		"graphic":               graphic,
	})
	service.FillSegment(formedGraphic, formedGraphicProducts, graphic)
	segment, err := service.GetSegment(graphic.SupplierName)
	var email string
	if segment.Email != "" {
		email = segment.Email
		service.SendNotificationSegment("files/segments/segment.xlsx", email)
	} else {
		for _, value := range segment.Region {
			email = value.Email
			service.SendNotificationSegment("files/segments/segment.xlsx", email)
		}
	}
	//service.SendNotificationSegment("files/segments/segment.xlsx")
}
