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

func GetAllRoles(c *gin.Context) {
	roles, err := service.GetAllRolesFullInfo()
	if err != nil {
		log.Println("[controller.GetAllRoles]|[service.GetAllRolesFullInfo]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}

func EditRole(c *gin.Context) {
	var role models.Role
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("[controller.EditRole]|[binding id param]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := c.BindJSON(&role); err != nil {
		log.Println("[controller.EditRole]|[binding json]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	role.Id = id
	if err := service.EditRole(role); err != nil {
		log.Println("[controller.EditRole]|[service.EditRole]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": fmt.Sprintf("роль c id = %d была успешно обновлена!", id)})
}

func AddNewRole(c *gin.Context) {
	var role models.Role
	if err := c.BindJSON(&role); err != nil {
		log.Println("[controller.AddNewRole]|[binding json]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.AddNewRole(role); err != nil {
		log.Println("[controller.AddNewRole]|[service.AddNewRole]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "новая роль была успешно создана!"})
}

func DeleteRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("[controller.DeleteRole]|[binding id param]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.DeleteRole(id); err != nil {
		log.Println("[controller.DeleteRole]|[service.DeleteRole]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": fmt.Sprintf("роль c id = %d было успешна удалена!", id)})
}

func AttachRightToRole(c *gin.Context) {
	roleId, err := strconv.Atoi(c.Param("role_id"))
	if err != nil {
		log.Println("[controller.AttachRightToRole]|[binding role_id param]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	rightId, err := strconv.Atoi(c.Param("right_id"))
	if err != nil {
		log.Println("[controller.AttachRightToRole]|[binding right_id param]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.AttachRightToRole(roleId, rightId); err != nil {
		log.Println("[controller.AttachRightToRole]|[service.AttachRightToRole]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": fmt.Sprintf("право c id = %d было успешна привязана к роли с id = %d", rightId, roleId)})
}

func DetachRightFromRole(c *gin.Context) {
	roleId, err := strconv.Atoi(c.Param("role_id"))
	if err != nil {
		log.Println("[controller.DetachRightFromRole]|[binding roleId param]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	rightId, err := strconv.Atoi(c.Param("right_id"))
	if err != nil {
		log.Println("[controller.DetachRightFromRole]|[binding rightId param]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.DetachRightFromRole(roleId, rightId); err != nil {
		log.Println("[controller.DetachRightFromRole]|[service.DetachRightFromRole]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": fmt.Sprintf("право c id = %d было успешна отнята от роли с id = %d", rightId, roleId)})
}