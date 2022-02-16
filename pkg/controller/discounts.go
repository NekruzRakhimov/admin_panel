package controller

import (
	"admin_panel/model"
	"admin_panel/pkg/service"
	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	contracts, err := service.GetAllRBByContractorBIN(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	if contracts == nil {
		contracts = []model.RbDTO{}
	}

	c.JSON(http.StatusOK, contracts)

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

	c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.File("files/reports/rb/rb_report_template.xlsx")
}
