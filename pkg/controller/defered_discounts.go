package controller

import (
	"admin_panel/model"
	"admin_panel/pkg/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetAllDeferredDiscounts(c *gin.Context) {
	var request model.RBRequest
	if err := c.BindJSON(&request); err != nil {
		log.Println("[controller][GetAllRBByContractorBIN] error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	contracts, err := service.GetAllDeferredDiscounts(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"discounts":    contracts,
		"total_amount": GetTotalFromRbDTO(contracts),
	})
}

func FormExcelForDeferredDiscounts(c *gin.Context) {
	var request model.RBRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.FormExcelForDeferredDiscounts(request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.File("files/reports/dd/reportDD.xlsx")

}

func GetTotalFromRbDTO(contracts []model.RbDTO) (totalAmount float32) {
	for _, contract := range contracts {
		totalAmount += contract.DiscountAmount
	}

	return totalAmount
}
