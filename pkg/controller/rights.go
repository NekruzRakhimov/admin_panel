package controller

import (
	"admin_panel/models"
	"admin_panel/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetAllRights(c *gin.Context) {
	rights, err := service.GetAllRights()
	if err != nil {
		log.Println("[controller.GetAllRights]|[service.GetAllRights]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rights)
}

func AddNewRight(c *gin.Context) {
	var right models.Right
	if err := c.BindJSON(&right); err != nil {
		log.Println("[controller.AddNewRight]|[binding json]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.AddNewRight(right); err != nil {
		log.Println("[controller.AddNewRight]|[service.AddNewRight]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "новое право было успешно создано!"})
}

func EditRight(c *gin.Context) {
	var right models.Right
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("[controller.EditRight]|[binding id param]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := c.BindJSON(&right); err != nil {
		log.Println("[controller.EditRight]|[binding json]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	right.Id = id
	if err := service.EditRight(right); err != nil {
		log.Println("[controller.EditRight]|[service.EditRight]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": fmt.Sprintf("право c id = %d было успешно обновлено!", id)})
}

func DeleteRight(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("[controller.DeleteRight]|[binding id param]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.DeleteRight(id); err != nil {
		log.Println("[controller.DeleteRight]|[service.DeleteRight]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": fmt.Sprintf("право c id = %d было успешно удалено!", id)})
}

