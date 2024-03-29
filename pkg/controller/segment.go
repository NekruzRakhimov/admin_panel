package controller

import (
	"admin_panel/models"
	"admin_panel/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	segment.SegmentCode = uuid.New().String()
	err = service.CreateSegment(segment)
	if err != nil {
		c.JSON(400, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"reason": "сегмент успешно создан"})

}

func ChangeSegment(c *gin.Context) {
	var segment models.Segment
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&segment)
	if err != nil {
		c.JSON(400, gin.H{"reason": err.Error()})
		return
	}
	segment.ID = id

	//segment.SegmentCode = uuid.New().String()
	err = service.ChangeSegment(segment)
	if err != nil {
		c.JSON(400, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"reason": "сегмент успешно обновлен"})

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
	//formedGraphicID, err := strconv.Atoi(c.Param("id"))
	//fmt.Println(formedGraphicID, "ID")
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"reason": "ERROR"})
	//	return
	//}

	formedGraphics, err := service.GetAllFormedGraphics(1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	for _, g := range formedGraphics {
		//сформированый потребность граффика -
		formedGraphic, err := service.GetFormedGraphicByID(g.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		fmt.Println("formedGraphic", formedGraphic.CreatedAt)

		formedGraphicProducts, err := service.GetAllFormedGraphicsProducts(formedGraphic.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}
		// от сюда взять дату  и номер
		graphic, err := service.GetGraphicByID(formedGraphic.GraphicID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		formula, err := service.GetFormulaByID(formedGraphic.FormulaID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}
		fmt.Println("formula", formula)

		fmt.Println("GRAP", graphic)
		fmt.Println("GRAP", graphic)

		c.JSON(200, gin.H{
			"formedGraphic":         formedGraphic,
			"formedGraphicProducts": formedGraphicProducts,
			"graphic":               graphic,
			"formula":               formula,
		})
		fmt.Println("STATUS", formedGraphic.Status)

		if formedGraphic.Status == "сформирован" {

			err := service.ChangeLetter(formedGraphic.ID)
			if err != nil {
				c.JSON(400, gin.H{"reason": err})
				return
			}
			//
			service.FillSegment(formedGraphic, formedGraphicProducts, graphic, formula)
			segment, _ := service.GetSegment(graphic.SupplierName)
			var email string
			if segment.Email != "" {
				email = segment.Email
				fmt.Println("почта", email)
				err := service.SendNotificationSegment("files/segments/segment.xlsx", email)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"reason": err})
					return
				}
			} else {
				for _, value := range segment.Region {
					email = value.Email
					err := service.SendNotificationSegment("files/segments/segment.xlsx", email)
					if err != nil {
						c.JSON(http.StatusBadRequest, gin.H{"reason": err})
						return
					}
				}
			}
			//		//service.SendNotificationSegment("files/segments/segment.xlsx", email)
		}
	}

}
