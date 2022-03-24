package controller

import (
	"admin_panel/models"
	"admin_panel/pkg/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetAllRBByContractorBIN  contract godoc
// @Summary Get Report RB
// @Description получение отчета по РБ
// @Accept  json
// @Produce  json
// @Tags reports
// @Param  contract  body models.RBRequest true "forming report"
// @Success 200 {array}  models.RbDTO
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /reports/rb [post]
func GetAllRBByContractorBIN(c *gin.Context) {
	var request models.RBRequest
	if err := c.BindJSON(&request); err != nil {
		log.Println("[controller][GetAllRBByContractorBIN] error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.SaveDoubtedDiscounts(request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	RbDTOs, err := service.GetAllRBByContractorBIN(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	for i := range RbDTOs {
		RbDTOs[i].Status = "Выполнено"
	}

	//SortedContracts := []models.RbDTO{}
	//for _, contract := range RbDTOs {
	//	if contract.ID != 0 {
	//		SortedContracts = append(SortedContracts, contract)
	//	}
	//}

	c.JSON(http.StatusOK, RbDTOs)
}

func FormExcelForRB(c *gin.Context) {
	var request models.RBRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.SaveDoubtedDiscounts(request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	if err := service.FormExcelForRBReport(request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.File("files/reports/rb/rb_report.xlsx")
}

func FormExcelForRBBrand(c *gin.Context) {
	//var request models.RBRequest
	//if err := c.BindJSON(&request); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
	//	return
	//}
	//
	//if err := service.FormExcelForRBReport(request); err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
	//	return
	//}

	c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.File("files/reports/rb/rb_report_template_brand.xlsx")
}

// GetDoubtedDiscounts  doubted_discounts godoc
// @Summary doubted_discounts
// @Description получение списка скидок для утверждения условия
// @Accept  json
// @Produce  json
// @Tags reports
// @Param  contract  body models.RBRequest true "getting doubted discounts"
// @Success 200 {array}  models.Discount
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /reports/doubted_discounts [post]
func GetDoubtedDiscounts(c *gin.Context) {

	var request models.RBRequest
	if err := c.BindJSON(&request); err != nil {
		log.Println("[controller][c.BindJSON] error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	discounts, err := service.GetDoubtedDiscounts(request)
	if err != nil {
		log.Println("[controller][service.GetDoubtedDiscounts] error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, discounts)
}

// SaveDoubtedDiscountsResults  doubted_discounts godoc
// @Summary doubted_discounts
// @Description сохранение списка скидок для утверждения условия
// @Accept  json
// @Produce  json
// @Tags reports
// @Param  contract  body models.DoubtedDiscountResponse true "saving doubted discounts"
// @Success 200 {object}  map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /reports/doubted_discounts [put]
func SaveDoubtedDiscountsResults(c *gin.Context) {

	var request models.DoubtedDiscountResponse
	if err := c.BindJSON(&request); err != nil {
		log.Println("[controller][c.BindJSON] error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	//err := service.SaveDoubtedDiscountsResults(request)
	//if err != nil {
	//	log.Println("[controller][service.GetDoubtedDiscounts] error is: ", err.Error())
	//	c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
	//	return
	//}

	c.JSON(http.StatusOK, gin.H{"reason": "данные успешно сохранены"})
}

func Check1CGetData(c *gin.Context) {
	var request models.RBRequest
	if err := c.BindJSON(&request); err != nil {
		log.Println("[controller][c.BindJSON] error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	requestFor1C := models.GetData1CRequest{
		ClientBin: request.BIN,
		DateStart: request.PeriodFrom,
		DateEnd:   request.PeriodTo,
		Type:      "sales_brand_only",
	}

	data, err := service.GetDataFrom1C(requestFor1C)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}
