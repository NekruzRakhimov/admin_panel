package controller

import (
	"admin_panel/model"
	"admin_panel/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// CreateContract
//Creating contract godoc
// @Summary Creating contract
// @Description Creating contract
// @Accept  json
// @Produce  json
// @Tags contracts
// @Param  contract  body model.Contract true "creating contract"
// @Param  type  query string true "type of contract"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /contract/ [post]
func CreateContract(c *gin.Context) {
	var contract model.Contract

	contract.Type = c.Param("type")

	if err := c.BindJSON(&contract); err != nil {
		log.Println("[controller.CreateContract]|[c.BindJSO]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.CreateContract(contract); err != nil {
		log.Println("[controller.CreateContract]|[service.CreateContract]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "новый договор был успешно создан!"})
}

func GetAllContracts(c *gin.Context) {
	contractsMiniInfo, err := service.GetAllContracts()
	if err != nil {
		log.Println("[controller.GetAllContracts]|[service.GetAllContracts]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contractsMiniInfo)
}

func CreateMarketingContract(c *gin.Context) {
	var input model.MarketingServicesContract
	err := c.BindJSON(&input)
	fmt.Println("======================================================____", input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return

	}
	err = service.CreateMarketingContract(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "договор успешно создан"})
}

//func AddNewRight(c *gin.Context) {
//	var right model.Right
//	if err := c.BindJSON(&right); err != nil {
//		log.Println("[controller.AddNewRight]|[binding json]| error is: ", err.Error())
//		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
//		return
//	}
//
//	if err := service.AddNewRight(right); err != nil {
//		log.Println("[controller.AddNewRight]|[service.AddNewRight]| error is: ", err.Error())
//		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"reason": "новое право было успешно создано!"})
//}

func GetAllCurrency(c *gin.Context) {
	currency, err := service.GetAllCurrency()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, currency)

}
