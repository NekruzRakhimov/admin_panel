package controller

import (
	"admin_panel/model"
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
// @Param  contract  body model.RBRequest true "forming report"
// @Success 200 {array}  model.RbDTO
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /reports/rb [post]
func GetAllRBByContractorBIN(c *gin.Context) {
	var request model.RBRequest
	if err := c.BindJSON(&request); err != nil {
		log.Println("[controller][GetAllRBByContractorBIN] error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.SaveDoubtedDiscounts(request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	var contracts []model.RbDTO
	// #1
	contracts1stType, err := service.GetAllRBByContractorBIN(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	for i := range contracts1stType {
		contracts1stType[i].DiscountType = service.RB1Name
	}
	contracts = append(contracts, contracts1stType...)

	// #2
	rbSecondType, err := service.GetAllRBSecondType(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	for i := range rbSecondType {
		rbSecondType[i].DiscountType = service.RB2Name
	}
	contracts = append(contracts, rbSecondType...)

	// #3
	rbThirdType, err := service.GetRBThirdType(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	for i := range rbThirdType {
		rbThirdType[i].DiscountType = service.RB3Name
	}
	contracts = append(contracts, rbThirdType...)

	// #4
	rbFourthType, err := service.InfoPresentationDiscount(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	for i := range rbFourthType {
		rbFourthType[i].DiscountType = service.RB4Name
	}
	contracts = append(contracts, rbFourthType...)

	// #8
	rbEighthType, err := service.GetRBEighthType(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	for i := range rbEighthType {
		rbEighthType[i].DiscountType = service.RB8Name
	}
	contracts = append(contracts, rbEighthType...)

	// #10
	rbTenthType, err := service.GetRbTenthType(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	for i := range rbTenthType {
		rbTenthType[i].DiscountType = service.RB10Name
	}
	contracts = append(contracts, rbTenthType...)

	// #12
	rb12thType, err := service.RbDiscountForSalesGrowth(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	for i := range rb12thType {
		rb12thType[i].DiscountType = service.RB12Name
	}
	contracts = append(contracts, rb12thType...)

	// #13
	rb13thType, err := service.DiscountRBPeriodTime(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	for i := range rb13thType {
		rb13thType[i].DiscountType = service.RB13Name
	}

	contracts = append(contracts, rb13thType...)

	SortedContracts := []model.RbDTO{}
	for _, contract := range contracts {
		if contract.ID != 0 {
			SortedContracts = append(SortedContracts, contract)
		}
	}

	c.JSON(http.StatusOK, SortedContracts)
}

func FormExcelForRB(c *gin.Context) {
	var request model.RBRequest
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
	//if request.BIN == "010203040506" {
	//	c.File("files/reports/rb/report_brand_new.xlsx")
	//	return
	//}

	c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.File("files/reports/rb/rb_report.xlsx")
}

func FormExcelForRBBrand(c *gin.Context) {
	//var request model.RBRequest
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
// @Param  contract  body model.RBRequest true "getting doubted discounts"
// @Success 200 {array}  model.Discount
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /reports/doubted_discounts [post]
func GetDoubtedDiscounts(c *gin.Context) {

	var request model.RBRequest
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
// @Param  contract  body model.DoubtedDiscountResponse true "saving doubted discounts"
// @Success 200 {object}  map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /reports/doubted_discounts [put]
func SaveDoubtedDiscountsResults(c *gin.Context) {

	var request model.DoubtedDiscountResponse
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
	var request model.RBRequest
	if err := c.BindJSON(&request); err != nil {
		log.Println("[controller][c.BindJSON] error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	requestFor1C := model.GetData1CRequest{
		ClientBin:      request.BIN,
		Beneficiary:    "",
		DateStart:      request.PeriodFrom,
		DateEnd:        request.PeriodTo,
		Type:           "sales",
		TypeValue:      "",
		TypeParameters: nil,
		Contracts:      nil,
	}

	data, err := service.GetDataFrom1C(requestFor1C)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}
