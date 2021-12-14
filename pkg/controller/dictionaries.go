package controller

import (
	"admin_panel/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetAllCurrencies dictionary godoc
// @Summary Get All Currencies
// @Description Gel All Currencies
// @Accept  json
// @Produce  json
// @Tags dictionary
// @Success 200 {array}  model.Currency
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /dictionary/currencies [get]
func GetAllCurrencies(c *gin.Context) {
	currencies, err := service.GetAllCurrencies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, currencies)
}

//GetAllPositions dictionary godoc
// @Summary Get All Positions
// @Description Gel All Positions
// @Accept  json
// @Produce  json
// @Tags dictionary
// @Success 200 {array}  model.Position
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /dictionary/positions [get]
func GetAllPositions(c *gin.Context) {
	positions, err := service.GetAllPositions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, positions)
}

//GetAllAddresses dictionary godoc
// @Summary Get All Addresses
// @Description Gel All Addresses
// @Accept  json
// @Produce  json
// @Tags dictionary
// @Success 200 {array}  model.Address
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /dictionary/addresses [get]
func GetAllAddresses(c *gin.Context) {
	addresses, err := service.GetAllAddresses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, addresses)
}

//GetAllFrequencyDeferredDiscounts dictionary godoc
// @Summary Get All FrequencyDeferredDiscounts
// @Description Gel All FrequencyDeferredDiscounts
// @Accept  json
// @Produce  json
// @Tags dictionary
// @Success 200 {array}  model.FrequencyDeferredDiscount
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /dictionary/frequency_deferred_discounts [get]
func GetAllFrequencyDeferredDiscounts(c *gin.Context) {
	frequencyDeferredDiscount, err := service.GetAllFrequencyDeferredDiscounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, frequencyDeferredDiscount)
}
