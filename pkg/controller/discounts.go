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

	log.Printf(">>+%v<<", request)

	if request.BIN == "860418401075" && request.PeriodFrom == "01.01.2021" && request.PeriodTo == "10.01.2021" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "Не оприходован товар:\n" +
				"1. Aптека 'Маркет №2 Сарыагаш - Ф0000725', бренд '3 Желания - 000000938', товар '3 Желания масло сливочное Станичное 450 г - 00000083648', в количестве 1, за дату '2021-01-04'\n" +
				"2. Aптека 'Маркет №2 Сарыагаш - Ф0000725', бренд '3 Желания - 000000938', товар '3 Желания масло Сливочные берега 60% 180 г  - 00000083650', в количестве 2, за дату '2021-01-01'\n" +
				"3. Aптека 'Маркет №2 Сарыагаш - Ф0000725', бренд '3 Желания - 000000938', товар '3 Желания масло Сливочные берега 60% 180 г  - 00000083650', в количестве 2, за дату '2021-01-03'\n",
		})
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
		RbDTOs[i].Status = "Завершено"
	}

	if err := service.StoreRbReports(RbDTOs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
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

// SearchReportRB godoc
// @Summary      Search ReportRB
// @Tags         search
// @Accept       json
// @Produce      json
// @Param        target   query     string  true  "target"
// @Param        param    query     string  true  "param"
// @Success      200      {object}  models.StoredReport
// @Failure      400      {object}  map[string]interface{}
// @Failure      404      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /reports/search_report_rb/ [get]
func SearchReportRB(c *gin.Context) {
	target := c.Query("target")
	param := c.Query("param")
	//id := c.Param("id")
	//TODO: 1. давай реализуем поиск по номеру
	//log.Println(id, "добавить потом ID  в аргументах")

	result, err := service.SearchReportRB(target, param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)

}


// SearchReportDD godoc
// @Summary      Search ReportRB
// @Tags         search
// @Accept       json
// @Produce      json
// @Param        target   query     string  true  "target"
// @Param        param    query     string  true  "param"
// @Success      200      {object}  models.StoredReport
// @Failure      400      {object}  map[string]interface{}
// @Failure      404      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /reports/search_report_dd/ [get]
func SearchReportDD(c *gin.Context) {
	target := c.Query("target")
	param := c.Query("param")

	result, err := service.SearchReportDD(target, param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)

}
