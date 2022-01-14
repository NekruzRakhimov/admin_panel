package controller

import (
	"admin_panel/db"
	"admin_panel/model"
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
