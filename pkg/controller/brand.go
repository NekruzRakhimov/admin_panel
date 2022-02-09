package controller

import (
	"admin_panel/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetBrands godoc
// @Summary     Получаем список брендов
// @Description  Получаем список брендов
// @Tags         brand
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
// @Success      200      {object}  model.Sales
// @Failure      400      {object}  map[string]interface{}
// @Failure      404      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /sales/ [get]
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

	sales, err := service.GetSales()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err})
		return
	}
	c.JSON(http.StatusOK, sales)

}
