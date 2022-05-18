package controller

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"admin_panel/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func DDOne(c *gin.Context) {
	var request models.RBRequest
	c.ShouldBind(&request)

	contractsWithJson, err := repository.GetAllContractDetailByBIN(request.ClientCode, request.PeriodFrom, request.PeriodTo)
	if err != nil {
		return
	}

	contracts, err := service.BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return
	}

	fmt.Println("contracts", contracts)

	timeP, err := service.GetDD1st(request, contracts)
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, timeP)

}
