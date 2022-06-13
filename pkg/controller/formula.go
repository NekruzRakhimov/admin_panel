package controller

import (
	"admin_panel/models"
	"admin_panel/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//CreateFormula
// @Summary Create Formula
// @Description Create Formula
// @Accept  json
// @Produce  json
// @Tags formula
// @Param  Formula  body models.Formula true "creating Formula"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /formula [post]
func CreateFormula(c *gin.Context) {
	var formula models.Formula
	if err := c.BindJSON(&formula); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.CreateFormula(formula); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "формула успешно создана"})
}

//EditFormula Formula godoc
// @Summary Editing Formula
// @Description Editing Formula
// @Accept  json
// @Produce  json
// @Tags formula
// @Param  contract  body models.Formula true "editing Formula"
// @Param  id  path string true "id of Formula"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /formula/{id} [put]
func EditFormula(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "id not found"})
		return
	}

	var formula models.Formula
	if err := c.BindJSON(&formula); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	formula.ID = id

	if err := service.EditFormula(formula); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "формула успешно создана"})
}

func DeleteFormulaByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "id not found"})
		return
	}
	err = service.DeleteFormula(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err})
		return

	}
	c.JSON(http.StatusOK, gin.H{"reason": "формула успешно удалена"})

}

//GetAllFormulas formula godoc
// @Summary Get All formulas
// @Description Gel All formulas
// @Accept  json
// @Produce  json
// @Tags formula
// @Success 200 {array}  models.Formula
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /formula [get]
func GetAllFormulas(c *gin.Context) {
	formulas, err := service.GetAllFormulas()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, formulas)
}

//GetFormulaByID formula godoc
// @Summary Get formula Details
// @Description Gel formula Details
// @Accept  json
// @Produce  json
// @Tags contracts
// @Param  id  path string true "id of formula"
// @Success 200 {object}  models.Formula
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /formula/{id}/details [get]
func GetFormulaByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "id not found"})
		return
	}

	formula, err := service.GetFormulaByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, formula)
}

//GetFormulaParameters formula godoc
// @Summary Get All formulas parameters
// @Description Gel All parameters
// @Accept  json
// @Produce  json
// @Tags formula
// @Success 200 {array}  models.FormulaParameters
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /formula/parameters [get]
func GetFormulaParameters(c *gin.Context) {
	parameters := []models.FormulaParameters{
		{
			NameRus: "Мин",
			Code:    "min",
		},
		{
			NameRus: "Макс",
			Code:    "max",
		},
		{
			NameRus: "Остатки",
			Code:    "remainder",
		},
		{
			NameRus: "Продажи",
			Code:    "sales",
		},
		{
			NameRus: "Дни продаж",
			Code:    "sales_days",
		},
		{
			NameRus: "Товар в пути",
			Code:    "good_on_road",
		},
	}

	c.JSON(http.StatusOK, parameters)
}
