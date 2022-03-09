package controller

import (
	"admin_panel/db"
	"admin_panel/model"
	"admin_panel/pkg/repository"
	"admin_panel/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCarsBrand(c *gin.Context) {
	var cars []model.Cars

	db.GetDBConn().Raw("SELECT cars_info -> 'brand' AS brand  FROM cars").Scan(&cars)

	fmt.Println(cars)

	c.JSON(http.StatusOK, gin.H{"data2": cars})

}

func GetDisPer(c *gin.Context) {
	var bin model.ClientBin
	c.ShouldBind(&bin)

	period, err := repository.GetDiscountPeriod(bin.Bin)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, period)

}

func DiscountRBPeriodTime(c *gin.Context) {
	var request model.RBRequest
	c.ShouldBind(&request)
	timeP, err := service.DiscountRBPeriodTime(request)
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, timeP)

}

func GetContractCode(c *gin.Context) {
	var request model.RBRequest
	c.ShouldBind(&request)
	code := service.GetExternalCode(request.BIN)
	c.JSON(200, code)

}
