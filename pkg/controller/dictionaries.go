package controller

import (
	"admin_panel/models"
	"admin_panel/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//GetAllDictionaries dictionary godoc
// @Summary Get All Dictionaries
// @Description Gel All Dictionaries
// @Accept  json
// @Produce  json
// @Tags dictionary
// @Success 200 {array}  models.Dictionary
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /dictionary [get]
func GetAllDictionaries(c *gin.Context) {
	dictionaries, err := service.GetAllDictionaries()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dictionaries)
}

//GetAllDictionaryByID dictionary godoc
// @Summary Get All Dictionaries
// @Description Gel All Dictionaries
// @Accept  json
// @Produce  json
// @Tags dictionary
// @Param  id path int true "Dictionary ID"
// @Success 200 {object}  models.Dictionary
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /dictionary/{id} [get]
func GetAllDictionaryByID(c *gin.Context) {
	dictionaryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	dictionary, err := service.GetDictionaryByID(dictionaryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dictionary)
}

// CreatDictionary Create Dictionary godoc
// @Summary Create Dictionary
// @Description Create Dictionary
// @Tags dictionary
// @Accept  json
// @Produce  json
// @Param Dictionary body models.Dictionary true "Create Dictionary"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /dictionary [post]
func CreatDictionary(c *gin.Context) {
	var dictionary models.Dictionary
	err := c.BindJSON(&dictionary)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.CreateDictionary(dictionary); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "новый справочник был успешно создан!"})
}

// EditDictionary Update Dictionary godoc
// @Summary Update Dictionary user
// @Description Update Dictionary
// @Tags dictionary
// @Accept  json
// @Produce  json
// @Param  id path int true "Dictionary ID"
// @Param  Dictionary body models.Dictionary true "Update account"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /dictionary/{id} [put]
func EditDictionary(c *gin.Context) {
	var dictionary models.Dictionary
	err := c.BindJSON(&dictionary)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	dictionary.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.EditDictionary(dictionary); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "справочник был успешно изменен!"})
}

//DeleteDictionary godoc
//@Summary Dictionary user by ID
//@Tags dictionary
//@Produce json
//@Param id path string true "dictionary ID"
//@Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
//@Router /dictionary/{id} [delete]
func DeleteDictionary(c *gin.Context) {
	dictionaryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.DeleteDictionary(dictionaryID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "справочник был успешно удален!"})
}

// CreateDictionaryValue Create Dictionary Value godoc
// @Summary Create Dictionary Value
// @Description Create Dictionary Value
// @Tags dictionary_values
// @Accept  json
// @Produce  json
// @Param id path string true "dictionary ID"
// @Param DictionaryValue body models.DictionaryValue true  "Dictionary Value"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /dictionary/{id}/value [post]
func CreateDictionaryValue(c *gin.Context) {
	var dictionaryValue models.DictionaryValue
	err := c.BindJSON(&dictionaryValue)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	dictionaryValue.DictionaryID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.CreateDictionaryValue(dictionaryValue); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "новое значение для справочника было успешно создано!"})
}

//GetAllDictionaryValues Get All Dictionary Values godoc
// @Summary Get All Dictionary Values
// @Description Gel All Dictionary Values
// @Accept  json
// @Produce  json
// @Tags dictionary_values
// @Param id path string true "dictionary ID"
// @Success 200 {array}  models.DictionaryValue
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /dictionary/{id}/value [get]
func GetAllDictionaryValues(c *gin.Context) {
	dictionaryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	dictionaryValues, err := service.GetAllDictionaryValues(dictionaryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dictionaryValues)
}

// EditDictionaryValue Dictionary Values godoc
// @Summary Update Dictionary Values
// @Description Update Dictionary Values
// @Tags dictionary_values
// @Accept  json
// @Produce  json
// @Param id path string true "dictionary ID"
// @Param DictionaryValue body models.DictionaryValue true  "Dictionary Value"
// @Param value_id path string true "value ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /dictionary/{id}/value/{value_id} [put]
func EditDictionaryValue(c *gin.Context) {
	dictionaryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	dictionaryValueID, err := strconv.Atoi(c.Param("value_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	var dictionaryValue models.DictionaryValue
	err = c.BindJSON(&dictionaryValue)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	dictionaryValue.ID = dictionaryValueID
	dictionaryValue.DictionaryID = dictionaryID

	if err := service.EditDictionaryValue(dictionaryValue); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "значение справочника было успешно изменено!"})

}

// DeleteDictionaryValue Delete Dictionary Values godoc
// @Summary Delete Dictionary Values by ID
// @Tags dictionary_values
// @Produce json
// @Param id path string true "dictionary ID"
// @Param value_id path string true "value ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /dictionary/{id}/value/{value_id} [delete]
func DeleteDictionaryValue(c *gin.Context) {
	dictionaryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	dictionaryValueID, err := strconv.Atoi(c.Param("value_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.DeleteDictionaryValue(dictionaryID, dictionaryValueID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "значение справочника было успешно удалено!"})
}

//GetAllCurrencies dictionary godoc
// @Summary Get All Currencies
// @Description Gel All Currencies
// @Accept  json
// @Produce  json
// @Tags dictionary
// @Success 200 {array}  models.Currency
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
// @Success 200 {array}  models.Position
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
// @Success 200 {array}  models.Address
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
// @Success 200 {array}  models.FrequencyDeferredDiscount
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
