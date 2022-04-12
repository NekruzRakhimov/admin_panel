package controller

import (
	"admin_panel/models"
	"admin_panel/pkg/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetAllDeferredDiscounts(c *gin.Context) {
	var request models.RBRequest
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

	if err := service.StoreDdReports(contracts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"discounts":    contracts,
		"total_amount": GetTotalFromRbDTO(contracts),
	})
}

func FormExcelForDeferredDiscounts(c *gin.Context) {
	var request models.RBRequest
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

func GetTotalFromRbDTO(contracts []models.RbDTO) (totalAmount float32) {
	for _, contract := range contracts {
		totalAmount += contract.DiscountAmount
	}

	return totalAmount
}

//GetAllDdStoredReports reports godoc
// @Summary Get All reports
// @Description Gel All reports
// @Accept  json
// @Produce  json
// @Tags contracts
// @Success 200 {array}  models.StoredReport
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /reports/dd/stored [get]
func GetAllDdStoredReports(c *gin.Context) {
	reports, err := service.GetAllDdStoredReports()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reports)
}

//GetStoredDdReportDetails reports godoc
// @Summary Get report Details
// @Description Gel report Details
// @Accept  json
// @Produce  json
// @Tags contracts
// @Param  id  path string true "id of report"
// @Success 200 {array}  models.RbDTO
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /reports/dd/stored/{id}/details [get]
func GetStoredDdReportDetails(c *gin.Context) {
	storedReportID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	rbDTO, err := service.GetStoredDdReportDetails(storedReportID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rbDTO)
}

func GetExcelForDdStoredExcelReport(c *gin.Context) {
	storedReportID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err = service.GetExcelForDdStoredExcelReport(storedReportID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.File("files/reports/dd/dd_stored_reports.xlsx")
}
