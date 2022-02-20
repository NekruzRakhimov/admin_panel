package controller

import (
	"admin_panel/model"
	"admin_panel/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetBrands godoc
// @Summary     Получаем список брендов
// @Description  Получаем список брендов
// @Tags         brands
// @Accept       json
// @Produce      json
// @Success      200      {object}  model.Brand
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
// @Param        payload  body      model.DateSales  true  "Add Brand"
// @Success      200      {object}  model.Sales
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
	payload := model.DateSales{}
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err})
		return
	}

	//sales, err := service.GetSales("01.01.2022"+service.TempDateCompleter, "01.01.2022"+service.TempDateCompleter)
	sales, err := service.GetSales(payload.Datestart+service.TempDateCompleter, payload.Dateend+service.TempDateEnd, payload.ClientBin)
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
// @Success      200      {object}  model.Sales
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

func GetBrandInfo(c *gin.Context) {

	var req model.ReqBrand

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": req})
		return
	}
	brandInfo, _ := service.GetBrandInfo(req.ClientBin)

	c.JSON(http.StatusOK, brandInfo)

}

func GenerateReportBrand(c *gin.Context) {
	var req model.ReqBrand

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": req})
		return
	}
	brandInfo, _ := service.GetBrandInfo(req.ClientBin)

	brand, _ := service.GetSalesBrand(req, brandInfo)

	c.JSON(http.StatusOK, brand)

}
