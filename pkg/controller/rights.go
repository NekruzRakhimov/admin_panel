package controller

import (
	"admin_panel/model"
	"admin_panel/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)


// Get All Rights by ADMIN godoc
// @Summary Get All Rights by Admin
// @Description Get All Rights by Admin
// @Accept  json
// @Produce  json
// @Tags rights
// @Success 200 {array} model.Right
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /rights/ [get]
func GetAllRights(c *gin.Context) {
	rights, err := service.GetAllRights()
	if err != nil {
		log.Println("[controller.GetAllRights]|[service.GetAllRights]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rights)
}


// Add Right by ADMIN godoc
// @Summary Add Right by Admin
// @Description Add Right by Admin
// @Accept  json
// @Produce  json
// @Tags rights
// @Param  right  body model.Right true "add role"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /rights/ [post]
func AddNewRight(c *gin.Context) {
	var right model.Right
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



// Update Right by ADMIN godoc
// @Summary Update Right by Admin
// @Description Update Right by Admin
// @Accept  json
// @Produce  json
// @Tags rights
// @Param  id path string true "rigth ID"
// @Param  right  body model.Right true "update right"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /rights/{id} [put]
func EditRight(c *gin.Context) {
	var right model.Right
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

	right.ID = id
	if err := service.EditRight(right); err != nil {
		log.Println("[controller.EditRight]|[service.EditRight]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": fmt.Sprintf("право c id = %d было успешно обновлено!", id)})
}


// Delete Right by ADMIN godoc
// @Summary  Delete Right by Admin
// @Description  Delete Right by Admin
// @Accept  json
// @Produce  json
// @Tags rights
// @Param  id path string true "right ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /rights/{id} [delete]
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

