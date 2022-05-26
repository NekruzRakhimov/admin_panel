package controller

import (
	"admin_panel/models"
	"admin_panel/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
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
		Startdate: date.Date,
		Enddate:   date.Date,
		IsPF:      true,
	}
	log.Println(time.Now(), " Started Defects - Main")
	fmt.Println(time.Now(), " Started Defects - Main")
	mainTime := time.Now()
	_, err := service.GetDefectsPF(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	//c.JSON(http.StatusOK, res)

	//c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	//c.File("./files/defects/defects_pharmacy.xlsx")

	log.Println(time.Now(), " Finished Defects - Main: durance[", time.Now().Sub(mainTime), "]")
	fmt.Println(time.Now(), " Finished Defects - Main: durance[", time.Now().Sub(mainTime), "]")
	c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.File("./files/defects/res.xlsx")
}

func GetDefectsByPharmacyLS(c *gin.Context) {
	date := struct {
		Date string `json:"date"`
	}{}

	if err := c.BindJSON(&date); err != nil {
		c.JSON(http.StatusOK, gin.H{"reason": err.Error()})
		return
	}

	req := models.DefectsRequest{
		Startdate: date.Date,
		Enddate:   date.Date,
	}

	_, err := service.GetDefectsLS(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	//c.JSON(http.StatusOK, res)

	//c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	//c.File("./files/defects/defects_pharmacy.xlsx")

	c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.File("./files/defects/res_ls.xlsx")
}

const TempData = " 00:00:00"

func GetSalesCount(c *gin.Context) {
	var req models.SalesCountRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	req.Startdate += TempData
	fmt.Println(req.Startdate)
	req.Enddate += TempData

	salesCount, err := service.GetSalesCountExt(req)
	if err != nil {
		fmt.Println(salesCount)
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, salesCount)
}
