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

	//TODO:  вернуть ему данные получается
	contracts, err := service.GetAllRBByContractorBIN(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	rbSecondType, err := service.GetAllRBSecondTypeMock(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	contracts = append(contracts, rbSecondType...)

	rbThirdType, err := service.GetRBThirdType(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	contracts = append(contracts, rbThirdType...)

	rbFourthType := service.InfoPresentationDiscount(request)
	contracts = append(contracts, rbFourthType...)

	//var rewardAmount float32
	//for _, contract := range rbThirdType {
	//	rewardAmount += contract.DiscountAmount
	//}
	//fmt.Printf("@@@%+v###", rewardAmount)

	//if rbThirdType != nil {
	//	if len(contracts) > 0 {
	//		contracts[0].RewardAmount = rewardAmount
	//	}
	//}

	rbEighthType, err := service.GetRBEighthType(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	contracts = append(contracts, rbEighthType...)

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

	err := service.SaveDoubtedDiscountsResults(request)
	if err != nil {
		log.Println("[controller][service.GetDoubtedDiscounts] error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "данные успешно сохранены"})
}
