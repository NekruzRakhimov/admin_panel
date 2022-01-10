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

type LoginData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// Login log in godoc
// @Summary login
//@Description login
//@Tags Auth
// @Accept  json
// @Produce  json
// @Param user body LoginData true "login"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /login [post]
func Login(c *gin.Context) {
	var request LoginData
	if err := c.BindJSON(&request); err != nil {
		log.Println("[controller.Login]|[controller.c.BindJSON(&request)]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	accessToken, err := service.GenerateToken(request.Login, request.Password)
	if err != nil {
		log.Println("[controller.Login]|[service.GenerateToken]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	refreshToken, err := service.GenerateToken(request.Login, accessToken)
	if err != nil {
		log.Println("[controller.Login]|[service.GenerateToken]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access":  accessToken,
		"refresh": refreshToken,
	})
}

// GetAllUsers Get All Users godoc
// @Summary Get All Users
// @Description Get All Users
// @Accept  json
// @Produce  json
// @Tags users
// @Success 200 {array} model.User
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/ [get]
func GetAllUsers(c *gin.Context) {
	users, err := service.GetAllUsersFullInfo()
	if err != nil {
		log.Println("[controller.GetAllUsers]|[service.GetAllUsersFullInfo]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// CreateNewUser Add User godoc
// @Summary Add an user
//@Description Add by json user
//@Tags users
// @Accept  json
// @Produce  json
// @Param user body model.User true "Add user"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/ [post]
func CreateNewUser(c *gin.Context) {
	var role model.User
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

// EditUser Update User godoc
// @Summary Update an user
// @Description Update by json user
// @Tags users
// @Accept  json
// @Produce  json
// @Param  id path int true "user ID"
// @Param  account body model.User true "Update account"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{id} [put]
func EditUser(c *gin.Context) {
	var user model.User
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

	user.ID = id
	if err := service.EditUser(user); err != nil {
		log.Println("[controller.EditRole]|[service.EditUser]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": fmt.Sprintf("данные о пользователе c id = %d была успешно обновлены!", id)})
}

//DeleteUser godoc
//@Summary Delete user by ID
//@Tags users
//@Produce json
//@Param id path string true "User ID"
//@Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
//@Router /users/{id} [delete]
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

// AttachRoleToUser Attach Role To User godoc
// @Summary Attach Role To User
// @Description Attach by json Role To User
// @Tags users
// @Accept  json
// @Produce  json
// @Param  id path string true "user ID"
// @Param  id path string true "role ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /attach_role/{user_id}/{role_id} [post]
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

// DetachRoleFromUser Detach Role To User godoc
// @Summary Detach Role To User
// @Description Detach by json Role To User
// @Tags users
// @Accept  json
// @Produce  json
// @Param  id path string true "user ID"
// @Param  id path string true "role ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /detach_role/{user_id}/{role_id} [delete]
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

// GetUserById Get User by ID godoc
// @Summary Get User by ID
// @Description Get User by ID
// @Accept  json
// @Produce  json
// @Tags users
// @Param  id path int true "user ID"
// @Success 200 {object} model.User
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{id}/details [get]
func GetUserById(c *gin.Context) {
	userIdStr := c.Param("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Println("[controller.GetUserById]|[strconv.Atoi]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	user, err := service.GetUserById(userId)
	if err != nil {
		log.Println("[controller.GetUserById]|[service.GetUserById]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
