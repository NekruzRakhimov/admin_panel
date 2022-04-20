package controller

import (
	"admin_panel/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetRegions regions godoc
// @Summary Get All regions
// @Description Gel All regions
// @Accept  json
// @Produce  json
// @Tags contracts
// @Success 200 {array}  models.Regions
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /regions_from_1c [get]
func GetRegions(c *gin.Context) {
	regions, err := service.GetRegionsFrom1C()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, regions)
}
