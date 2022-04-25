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

	if err = c.s.GetHyperstocksPharmacy(searchParams); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	g.Writer.Header().Set("Content-Type", excelContentType)
	g.File(service.HyperstocksPathData + service.HyperstocksPharmacyFileName)
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

	if err = c.s.GetHyperstocksStock(searchParams); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	g.Writer.Header().Set("Content-Type", excelContentType)
	g.File(service.HyperstocksPathData + service.HyperstocksStockFileName)
}
