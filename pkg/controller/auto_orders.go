package controller

import (
	"admin_panel/models"
	"admin_panel/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetStoreRegions(c *gin.Context) {
	storeRegions, err := service.GetStoreRegions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, storeRegions)
}

func GetMatrix(c *gin.Context) {
	//storeCode := "A0000120" //Аптека № 2, Шымкент, (Городской Акимат)
	storeCode := c.Param("store_code")
	matrix, err := service.GetMatrixExt(storeCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matrix)
}

//CreateGraphic graphic godoc
// @Summary Create Graphic
// @Description Create Graphic
// @Accept  json
// @Produce  json
// @Tags contracts
// @Param  contract  body models.Graphic true "creating Graphic"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /graphic [post]
func CreateGraphic(c *gin.Context) {
	var graphic models.Graphic
	if err := c.BindJSON(&graphic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.CreateGraphic(graphic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "график успешно создан"})
}

//GetAllGraphics Graphic godoc
// @Summary Get All Graphics
// @Description Gel All Graphics
// @Accept  json
// @Produce  json
// @Tags contracts
// @Param  q  query string false "param for searching"
// @Success 200 {array}  models.Graphic
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /graphic [get]
func GetAllGraphics(c *gin.Context) {
	graphics, err := service.GetAllGraphics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, graphics)
}

//GetGraphicByID Graphic godoc
// @Summary Get Graphic Details
// @Description Gel Graphic Details
// @Accept  json
// @Produce  json
// @Tags contracts
// @Param  id  path string true "id of Graphic"
// @Success 200 {object}  models.Graphic
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /graphic/{id}/details [get]
func GetGraphicByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "graphic_id не найден"})
		return
	}

	graphic, err := service.GetGraphicByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, graphic)
}

//EditGraphic Graphic godoc
// @Summary Editing Graphic
// @Description Editing Graphic
// @Accept  json
// @Produce  json
// @Tags contracts
// @Param  contract  body models.Graphic true "editing Graphic"
// @Param  id  path string true "id of Graphic"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /graphic/{id} [put]
func EditGraphic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "graphic_id не найден"})
		return
	}

	var graphic models.Graphic
	if err := c.BindJSON(&graphic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	graphic.ID = id
	if err := service.EditGraphic(graphic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "график успешно обновлен"})
}

func DeleteGraphic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "graphic_id не найден"})
		return
	}
	err = service.DeleteGraphic(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err})
		return

	}
	c.JSON(http.StatusOK, gin.H{"reason": "график успешно удален"})

}

// GetAllAutoOrders TODO пока не сделаем формулы, не нужен
func GetAllAutoOrders(c *gin.Context) {
	autoOrders, err := service.GetAllAutoOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, autoOrders)
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

func GetAllAutoOrderTypes(c *gin.Context) {
	types := []string{"Min", "Max"}
	c.JSON(http.StatusOK, types)
}

func GetAllAutoOrderAnalysesPeriod(c *gin.Context) {

	// вернуть массив строк
	//

	periods := []string{"1 неделя", "2 неделя", "1 месяц", "2 месяца", "3 месяца", "4 месяца",
		"5 месяцев", "6 месяцев", "7 месяцев", "8 месяцев", "9 месяцев", "10 месяцев", "11 месяцев", "12 месяцев",
		"13 месяцев", "14 месяцев", "15 месяцев", "16 месяцев", "17 месяцев", "18 месяцев", "19 месяцев", "20 месяцев", "21 месяца",
		"22 месяца", "23 месяца", "24 месяца", "за все время"}

	c.JSON(http.StatusOK, periods)
}
