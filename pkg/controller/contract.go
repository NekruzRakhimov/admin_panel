package controller

import (
	"admin_panel/model"
	"admin_panel/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//CreateContract contract godoc
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
// @Router /contract/{type} [post]
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

//GetAllContracts contract godoc
// @Summary Get All Contracts
// @Description Gel All Contract
// @Accept  json
// @Produce  json
// @Tags contracts
// @Param  type  query string true "type of contract"
// @Success 200 {array}  model.ContractMiniInfo
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /contract/ [get]
func GetAllContracts(c *gin.Context) {
	contractType := c.Query("type")
	contractsMiniInfo, err := service.GetAllContracts(contractType)
	if err != nil {
		log.Println("[controller.GetAllContracts]|[service.GetAllContracts]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contractsMiniInfo)
}

func GetContractDetails(c *gin.Context) {
	contractIdStr := c.Query("id")
	contractId, err := strconv.Atoi(contractIdStr)
	if err != nil {
		log.Println("[controller.GetContractDetails]|[strconv.Atoi]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	contract, err := service.GetContractDetails(contractId)
	if err != nil {
		log.Println("[controller.GetContractDetails]|[service.GetContractDetails]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contract)
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
