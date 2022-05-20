package controller

import (
	"admin_panel/pkg/dto"
	"admin_panel/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ForecastController struct {
	s service.ForecastService
}

func NewForecastController(s service.ForecastService) *ForecastController {
	return &ForecastController{s}
}

func (c *ForecastController) HandleRoutes(r *gin.RouterGroup) {
	r.GET("/forecast", c.GetForecastSales)
}

func (c *ForecastController) GetForecastSales(g *gin.Context) {
	err := g.Request.ParseForm()
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	searchParams, err := dto.ParseForecastSearchParameters(g.Request.Form)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	res, err := c.s.GetForecast(searchParams)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	g.JSON(http.StatusOK, res)
}
