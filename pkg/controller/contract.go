package controller

import (
	"admin_panel/model"
	"admin_panel/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


// Create Marketing contract godoc
// @Summary Create marketing contract
//@Description Add by json MarketingServicesContract
//@Tags contract
// @Accept  json
// @Produce  json
// @Param user body model.MarketingServicesContract true "Add user"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /contract/ [post]
func CreateMarketingContract(c *gin.Context)  {
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

	c.JSON(http.StatusOK, gin.H{"reason": "договор успешно создано"})
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











func GetAllCurrency(c *gin.Context)  {
	currency, err := service.GetAllCurrency()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, currency)

}