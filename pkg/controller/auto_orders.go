package controller

import (
	"admin_panel/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllAutoOrders TODO пока не сделаем формулы, не нужен
func GetAllAutoOrders(c *gin.Context) {
	autoOrders, err := service.GetAllAutoOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, autoOrders)
}

func SendAutoOrdersTo1C(c *gin.Context) {
	formulaID, err := strconv.Atoi(c.Param("formula_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "formula_id is invalid"})
		return
	}

	err = service.SendAutoOrderTo1C(formulaID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "отправка в 1С"})
}

func CancelFormedFormula(c *gin.Context) {
	comment := c.Query("comment")
	formulaID, err := strconv.Atoi(c.Param("formula_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "formula_id is invalid"})
		return
	}

	if err = service.CancelFormedFormula(formulaID, comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "успешно отменено"})
}

func CancelFormedGraphic(c *gin.Context) {
	comment := c.Query("comment")
	graphicID, err := strconv.Atoi(c.Param("graphic_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "graphic_id is invalid"})
	}

	if err = service.CancelFormedGraphic(graphicID, comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "успешно отменено"})
}

//FormAutoOrder auto_orders godoc
// @Summary Creating auto_orders
// @Description Creating auto_orders
// @Accept  json
// @Produce  json
// @Tags contracts
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auto_orders [post]
func FormAutoOrder(c *gin.Context) {
	if err := service.FormAutoOrders(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "потребность успешно сформирована"})
}

//GetAllFormedGraphics auto_orders godoc
// @Summary Get All auto_orders
// @Description Gel All auto_orders
// @Accept  json
// @Produce  json
// @Tags contracts
// @Success 200 {array}  models.FormedGraphic
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auto_orders/{formula_id}/graphics [get]
func GetAllFormedGraphics(c *gin.Context) {
	formulaID, err := strconv.Atoi(c.Param("formula_id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"reason": err.Error()})
		return
	}
	fmt.Println(formulaID)

	graphics, err := service.GetAllFormedGraphics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	//var notEmpty []models.FormedGraphic
	//for _, graphic := range graphics {
	//	if len(graphic.Products) > 0 {
	//		notEmpty = append(notEmpty, graphic)
	//	}
	//}

	c.JSON(http.StatusOK, graphics)
}

//GetAllFormedGraphicProducts auto_orders godoc
// @Summary Get All auto_ordered products
// @Description Gel All auto_ordered products
// @Accept  json
// @Produce  json
// @Tags contracts
// @Success 200 {array}  models.FormedGraphicProduct
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auto_orders/{formula_id}/graphics/{graphic_id}/products [get]
func GetAllFormedGraphicProducts(c *gin.Context) {
	formulaID, err := strconv.Atoi(c.Param("formula_id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"reason": err.Error()})
		return
	}
	fmt.Println(formulaID)

	formedGraphicID, err := strconv.Atoi(c.Param("graphic_id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"reason": "graphic_id param not found"})
		return
	}

	products, err := service.GetAllFormedGraphicsProducts(formedGraphicID)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
