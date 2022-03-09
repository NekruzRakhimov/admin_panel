package controller

import (
	"admin_panel/db"
	"admin_panel/model"
	"admin_panel/pkg/repository"
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

	period, err := repository.GetDicsountPeriod(bin.Bin)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, period)

}
