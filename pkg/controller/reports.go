package controller

import (
	"admin_panel/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//GetAllStoredReports reports godoc
// @Summary Get All reports
// @Description Gel All reports
// @Accept  json
// @Produce  json
// @Tags contracts
// @Success 200 {array}  models.StoredReport
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /reports/rb/stored [get]
func GetAllStoredReports(c *gin.Context) {
	reports, err := service.GetAllStoredReports()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reports)
}

//GetStoredReportDetails reports godoc
// @Summary Get report Details
// @Description Gel report Details
// @Accept  json
// @Produce  json
// @Tags contracts
// @Param  id  path string true "id of report"
// @Success 200 {array}  models.RbDTO
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /reports/rb/stored/{id}/details [get]
func GetStoredReportDetails(c *gin.Context) {
	storedReportID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	rbDTO, err := service.GetStoredReportDetails(storedReportID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rbDTO)
}

func GetExcelForStoredExcelReport(c *gin.Context) {
	storedReportID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err = service.GetExcelForStoredExcelReport(storedReportID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.File("files/reports/rb/rb_stored_reports.xlsx")
}
