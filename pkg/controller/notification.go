package controller

import (
	"admin_panel/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var SearchByNumber struct {
	ContractNumber string `json:"contract_number"`
}

//func SearchNotifications(c *gin.Context) {
//	number := SearchNotification
//
//	// он даст номер, после чего мы отправляем запрос
//	c.ShouldBind(&number)
//	fmt.Println("ContractNumber", number.ContractNumber)
//	id := service.GetContractNot(number.ContractNumber)
//	c.JSONP(http.StatusOK, gin.H{"message": id})
//
//}

// ListNotifications godoc
// @Summary      List Notifications
// @Description  get notifications
// @Tags         notifications
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.Notification
// @Router       /notifications [get]
func GetNotifications(c *gin.Context) {
	notifications := service.GetNotifications()

	c.JSON(http.StatusOK, notifications)

}

// SearchNotification godoc
// @Summary      Search Notification by Number
// @Description  add by json account
// @Tags         search
// @Accept       json
// @Produce      json
// @Param        contract_number	 query string  true  "contract_number"
// @Param        status	 query string  true  "status"
// @Success      200      {object}  model.Notification
// @Failure      400      {object}  map[string]interface{}
// @Failure      404      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
func SearchNotification(c *gin.Context) {
	contractNumber := c.Query("contract_number")

	notification, err := service.SearchNotification(contractNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err})
		return
	}
	c.JSON(200, notification)

}
