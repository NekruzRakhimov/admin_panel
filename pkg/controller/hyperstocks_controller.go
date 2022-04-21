package controller

import (
	"admin_panel/pkg/dto"
	"admin_panel/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HyperstocksController struct {
	s service.HyperstocksService
}

func NewHyperstocksController(s service.HyperstocksService) *HyperstocksController {
	return &HyperstocksController{s}
}

func (c *HyperstocksController) HandleRoutes(r *gin.RouterGroup) {
	r.GET("/hyperstocks/pharmacy", c.GetHyperstocksPharmacy)
	r.GET("/hyperstocks/stock", c.GetHyperstocksStock)
}

func (c *HyperstocksController) GetHyperstocksPharmacy(g *gin.Context) {
	err := g.Request.ParseForm()
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	searchParams, err := dto.ParseHyperstocksSearchParameters(g.Request.Form)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	res, err := c.s.GetHyperstocksPharmacy(searchParams)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	g.JSON(http.StatusOK, res)
}

func (c *HyperstocksController) GetHyperstocksStock(g *gin.Context) {
	err := g.Request.ParseForm()
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	searchParams, err := dto.ParseHyperstocksSearchParameters(g.Request.Form)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	res, err := c.s.GetHyperstocksStock(searchParams)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	g.JSON(http.StatusOK, res)
}
