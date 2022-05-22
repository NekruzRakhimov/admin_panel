package controller

import (
	"admin_panel/db"
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"admin_panel/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCarsBrand(c *gin.Context) {
	var cars []models.Cars

	db.GetDBConn().Raw("SELECT cars_info -> 'brand' AS brand  FROM cars").Scan(&cars)

	fmt.Println(cars)

	c.JSON(http.StatusOK, gin.H{"data2": cars})

}

func GetDisPer(c *gin.Context) {
	var bin models.ClientBin
	c.ShouldBind(&bin)

	period, err := repository.GetDiscountPeriod(bin.Bin)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, period)

}

func GetPurchase(c *gin.Context) {
	var brand models.ReqBrand
	c.ShouldBind(&brand)

	period, err := service.GetPurchase(brand)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, period)

}

func DiscountRBPeriodTime(c *gin.Context) {
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

	timeP, err := service.GetRB12thType(request, contracts)
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, timeP)

}

func DiscountRB7(c *gin.Context) {
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

	timeP, err := service.GetRB7thType(request, contracts)
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, timeP)

}

func DiscountRB4(c *gin.Context) {
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

	timeP, err := service.GetRB4thType(request, contracts)
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, timeP)

}

func DiscountRB14(c *gin.Context) {
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

	timeP, err := service.GetRB14ThType(request, contracts)
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, timeP)

}

func DiscountRB17(c *gin.Context) {
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

	timeP, err := service.GetRB17ThType(request, contracts)
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, timeP)

}

func DiscountRB16(c *gin.Context) {
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

	timeP, err := service.GetRB16ThType(request, contracts)
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, timeP)

}

func DiscountRB15(c *gin.Context) {
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

	timeP, err := service.GetRB15ThType(request, contracts)
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, timeP)

}

func GetContractCode(c *gin.Context) {
	var request models.RBRequest
	err := c.ShouldBind(&request)
	fmt.Println("REQUEST", request)
	if err != nil {
		c.JSON(400, err.Error())
	}
	code := service.GetExternalCode(request.BIN)
	c.JSON(200, code)

}

func GetAllContractDetailByBIN(c *gin.Context) {
	var req models.ReqBrand

	c.ShouldBind(&req)
	fmt.Println("start_date", req.DateStart)
	fmt.Println("end_date", req.DateEnd)

	bin, _ := repository.GetAllContractDetailByBIN(req.ClientCode, req.DateStart, req.DateEnd)
	c.JSON(200, bin)

}
