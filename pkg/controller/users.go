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

func GetAllUsers(c *gin.Context) {
	users, err := service.GetAllUsersFullInfo()
	if err != nil {
		log.Println("[controller.GetAllUsers]|[service.GetAllUsersFullInfo]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func CreateNewUser(c *gin.Context) {
	var role models.User
	if err := c.BindJSON(&role); err != nil {
		log.Println("[controller.CreateNewUser]|[binding json]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.CreateNewUser(role); err != nil {
		log.Println("[controller.CreateNewUser]|[service.CreateNewUser]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "пользователь успешно создан!"})
}

func EditUser(c *gin.Context) {
	var user models.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("[controller.EditUser]|[binding id param]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := c.BindJSON(&user); err != nil {
		log.Println("[controller.EditUser]|[binding json]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	user.Id = id
	if err := service.EditUser(user); err != nil {
		log.Println("[controller.EditRole]|[service.EditUser]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": fmt.Sprintf("данные о пользователе c id = %d была успешно обновлены!", id)})
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("[controller.DeleteUser]|[binding id param]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.DeleteUser(id); err != nil {
		log.Println("[controller.DeleteUser]|[service.DeleteUser]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": fmt.Sprintf("пользователь c id = %d был успешно удален!", id)})
}

func AttachRoleToUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		log.Println("[controller.AttachRoleToUser]|[binding user_id param]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	roleId, err := strconv.Atoi(c.Param("role_id"))
	if err != nil {
		log.Println("[controller.AttachRoleToUser]|[binding role_id param]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.AttachRoleToUser(userId, roleId); err != nil {
		log.Println("[controller.AttachRoleToUser]|[service.AttachRoleToUser]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": fmt.Sprintf("роль c id = %d было успешна привязана к пользоватлю с id = %d", roleId, userId)})
}

func DetachRoleFromUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		log.Println("[controller.DetachRoleFromUser]|[binding user_id param]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	roleId, err := strconv.Atoi(c.Param("role_id"))
	if err != nil {
		log.Println("[controller.DetachRoleFromUser]|[binding role_id param]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.DetachRoleFromUser(userId, roleId); err != nil {
		log.Println("[controller.DetachRoleFromUser]|[service.DetachRoleFromUser]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": fmt.Sprintf("роль c id = %d было успешна отнята от пользователя с id = %d", roleId, userId)})
}
