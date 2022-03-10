package controller

import (
	"admin_panel/model"
	"github.com/gin-gonic/gin"
)

func PresentationDiscount(c *gin.Context) {

	var rbReq model.RBRequest

	c.ShouldBind(&rbReq)

	//respDicsount, err := service.PresentationDiscount(rbReq)
	//respDicsount := service.InfoPresentationDiscount(rbReq)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"reason": err})
	//}

	//c.JSON(200, respDicsount)

}
