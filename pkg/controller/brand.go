package controller

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"admin_panel/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetBrands godoc
// @Summary     Получаем список брендов
// @Description  Получаем список брендов
// @Tags         brands
// @Accept       json
// @Produce      json
// @Success      200      {object}  models.Brand
// @Failure      400      {object}  map[string]interface{}
// @Failure      404      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /brands/ [get]
func GetBrands(c *gin.Context) {
	// тут получаем список всех наименований (товаров, таблеток и т.д)
	brands, err := service.GetBrands()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err})
		return
	}
	c.JSON(http.StatusOK, brands)
}

// GetSales godoc
// @Summary     Получаем список продаж и кол-во
// @Description  Получаем список продаж и кол-во
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        payload  body      models.DateSales  true  "Add BrandName"
// @Success      200      {object}  models.Sales
// @Failure      400      {object}  map[string]interface{}
// @Failure      404      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /sales/ [post]
func GetSales(c *gin.Context) {
	//по факту должны отправить период С по какое число получить продажи
	//TODO: на данный момент даты С и ПО вшиты в код (потом надо убрать их)
	//{
	//	"datestart":"01.01.2022 0:02:09",
	//	"dateend":"01.01.2022 0:02:09"
	//}
	//TODO:
	// и в ответ получаем 4 поля
	// "product_name":
	//  "product_code":
	//  "total": - сумма продаж
	//   "qnt_total": - кол-во
	payload := models.RBRequest{}
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err})
		return
	}

	//payload.PeriodFrom += service.TempDateCompleter
	//payload.PeriodTo += service.TempDateEnd

	//var brandInfo []models.BrandInfo
	req := models.ReqBrand{
		ClientBin:   payload.BIN,
		Beneficiary: payload.ContractorName,
		DateStart:   payload.PeriodFrom,
		DateEnd:     payload.PeriodTo,
		Type:        "sales",
	}

	brandInfo := []models.BrandInfo{}
	sales, err := service.GetSalesBrand(req, brandInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err})
		return
	}
	c.JSON(http.StatusOK, sales)

}

// AddBrand godoc
// @Summary     создает новый бренд
// @Description  создает новый бренд
// @Tags         brands
// @Accept       json
// @Produce      json
// @Param        brand_name   query     string  true  "brand_name"
// @Success      200      {object}  models.Sales
// @Failure      400      {object}  map[string]interface{}
// @Failure      404      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /add_brand/ [get]
func AddBrand(c *gin.Context) {
	brandName := c.Query("brand_name")
	brand, err := service.AddBrand(brandName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err})
		return
	}

	c.JSON(http.StatusOK, brand)

}

// GetBrandInfo godoc
// @Summary     получаем данные о брендах
// @Description  получаем данные о брендах
// @Tags         report
// @Accept       json
// @Produce      json
// @Param        brand_name   query     string  true  "brand_name"
// @Success      200      {object}  models.Sales
// @Failure      400      {object}  map[string]interface{}
// @Failure      404      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /add_brand/ [get]
func GetBrandInfo(c *gin.Context) {

	var req models.ReqBrand

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": req})
		return
	}
	brandInfo, _ := service.GetBrandInfo(req.ClientBin)

	c.JSON(http.StatusOK, brandInfo)

}

func GenerateReportBrand(c *gin.Context) {
	var req models.ReqBrand

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": req})
		return
	}
	brandInfo, _ := service.GetBrandInfo(req.ClientBin)

	brand, _ := service.GetSalesBrand(req, brandInfo)

	//TODO: след момент:
	// надо взять кол-во
	// кол-во умножаешь на сумму
	// находим
	// сумму скидки
	// и все данные запихать в эксель

	c.JSON(http.StatusOK, brand)

}

func GetExcelBrand(c *gin.Context) {
	//var req models.RBRequest
	////var req models.ReqBrand
	//
	////c.ShouldBind(&req)
	//c.BindJSON(&req)
	//
	//log.Println("запрос->>>: ", req)
	//discount, _ := service.GetRB2ndType(req)
	////discount, _ := service.GetSales(req)
	//
	//c.JSON(200, gin.H{"discount": discount})

}

func GetExcelGrowth(c *gin.Context) {
	var req models.RBRequest
	//var req models.ReqBrand

	//c.ShouldBind(&req)
	c.BindJSON(&req)

	contractsWithJson, err := repository.GetAllContractDetailByBIN(req.BIN, req.PeriodFrom, req.PeriodTo)
	if err != nil {
		return
	}

	contracts, err := service.BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return
	}

	fmt.Println("запрос для прироста ->>>: ", req)
	growth, _ := service.GetRB13thType(req, contracts)
	//discount := service.GetRB2ndType(req)
	//discount, _ := service.GetSales(req)

	c.JSON(200, gin.H{"growth": growth})

}

func GetRb1(c *gin.Context) {
	var req models.RBRequest
	//var req models.ReqBrand

	//c.ShouldBind(&req)
	c.BindJSON(&req)

	contractsWithJson, err := repository.GetAllContractDetailByBIN(req.BIN, req.PeriodFrom, req.PeriodTo)
	if err != nil {
		return
	}

	contracts, err := service.BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return
	}

	fmt.Println("запрос для прироста ->>>: ", req)
	rb, _ := service.GetRB1stType(req, contracts)
	//discount := service.GetRB2ndType(req)
	//discount, _ := service.GetSales(req)

	c.JSON(200, rb)
}

func GetRb3(c *gin.Context) {
	var req models.RBRequest
	//var req models.ReqBrand

	//c.ShouldBind(&req)
	c.BindJSON(&req)

	contractsWithJson, err := repository.GetAllContractDetailByBIN(req.BIN, req.PeriodFrom, req.PeriodTo)
	if err != nil {
		return
	}

	contracts, err := service.BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return
	}

	fmt.Println("запрос для прироста ->>>: ", req)
	rb, _ := service.GetRB3rdType(req, contracts)
	//discount := service.GetRB2ndType(req)
	//discount, _ := service.GetSales(req)

	c.JSON(200, rb)
}

func GetRb5(c *gin.Context) {
	var req models.RBRequest
	//var req models.ReqBrand

	//c.ShouldBind(&req)
	c.BindJSON(&req)

	contractsWithJson, err := repository.GetAllContractDetailByBIN(req.BIN, req.PeriodFrom, req.PeriodTo)
	if err != nil {
		return
	}

	contracts, err := service.BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return
	}

	fmt.Println("запрос для прироста ->>>: ", req)
	rb, _ := service.GetRB5thType(req, contracts)
	//discount := service.GetRB2ndType(req)
	//discount, _ := service.GetSales(req)

	c.JSON(200, rb)
}

func SaveDataFrom1C(c *gin.Context) {
	//log.Println("###1C BEGIN##")
	//var body []byte
	//c.Request.Body.Read()
	//log.Printf("%s"))
	//log.Println("###1C END##")
	var block models.Block
	if err := c.BindJSON(&block); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.SaveDataFrom1C(block); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "данные сохранены"})
}
