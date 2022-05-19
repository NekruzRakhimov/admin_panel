package controller

import (
	"admin_panel/models"
	"admin_panel/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDefectsByPharmacyPF(c *gin.Context) {
	date := struct {
		Date string `json:"date"`
	}{}

	if err := c.BindJSON(&date); err != nil {
		c.JSON(http.StatusOK, gin.H{"reason": err.Error()})
		return
	}

	req := models.DefectsRequest{
		Startdate: fmt.Sprintf("%s 00:00:00", date.Date),
		Enddate:   fmt.Sprintf("%s 23:59:59", date.Date),
	}

	res, err := service.GetDefectsExt(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)

	//c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	//c.File("./files/defects/defects_pharmacy.xlsx")
}
