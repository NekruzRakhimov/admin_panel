package controller

import (
	"admin_panel/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStoreRegions(c *gin.Context) {
	storeRegions, err := service.GetStoreRegions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, storeRegions)
}

func GetMatrix(c *gin.Context) {
	//storeCode := "A0000120" //Аптека № 2, Шымкент, (Городской Акимат)
	storeCode := c.Param("store_code")
	matrix, err := service.GetMatrix(storeCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matrix)
}
