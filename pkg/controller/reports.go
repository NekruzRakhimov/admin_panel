package controller

import (
	"admin_panel/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
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
