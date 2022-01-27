package controller

import (
	"admin_panel/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Noti struct {
	Contract_number string `json:"contract_number"`
}

func GetIdNotification(c *gin.Context) {
	n := Noti

	// он даст номер, после чего мы отправляем запрос
	c.ShouldBind(&n)
	fmt.Println("ContractNumber", n.Contract_number)
	id := service.GetContractNot(n.Contract_number)
	c.JSONP(http.StatusOK, gin.H{"message": id})

}

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
