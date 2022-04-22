package controller

import (
	"admin_panel/pkg/dto"
	"admin_panel/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DefectsController struct {
	s service.DefectsService
}

func NewDefectsController(s service.DefectsService) *DefectsController {
	return &DefectsController{s}
}

func (c *DefectsController) HandleRoutes(r *gin.RouterGroup) {
	r.GET("/defects/pharmacy", c.GetDefectsPharmacy)
	r.GET("/defects/stock", c.GetDefectsStock)
}

func (c *DefectsController) GetDefectsPharmacy(g *gin.Context) {
	err := g.Request.ParseForm()
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	searchParams, err := dto.ParseDefectsSearchParameters(g.Request.Form)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	res, err := c.s.GetDefectsPharmacy(searchParams)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	g.JSON(http.StatusOK, res)
}

func (c *DefectsController) GetDefectsStock(g *gin.Context) {
	err := g.Request.ParseForm()
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	searchParams, err := dto.ParseDefectsSearchParameters(g.Request.Form)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	res, err := c.s.GetDefectsStock(searchParams)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	g.JSON(http.StatusOK, res)
}
