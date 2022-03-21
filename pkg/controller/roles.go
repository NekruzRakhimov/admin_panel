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

// GetAllRoles Get All Roles by ADMIN godoc
// @Summary Get All Roles by Admin
// @Description Get All Roles by Admin
// @Accept  json
// @Produce  json
// @Tags roles
// @Success 200 {array} models.Role
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /roles/ [get]
func GetAllRoles(c *gin.Context) {
	roles, err := service.GetAllRolesFullInfo()
	if err != nil {
		log.Println("[controller.GetAllRoles]|[service.GetAllRolesFullInfo]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}

// EditRole Update Role by ADMIN godoc
// @Summary  Update Role by Admin
// @Description  Update Role by Admin
// @Accept  json
// @Produce  json
// @Tags roles
// @Param  id path string true "role ID"
// @Param  role body models.Role true "Update role"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /roles/{id} [put]
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

	role.ID = id
	if err := service.EditRole(role); err != nil {
		log.Println("[controller.EditRole]|[service.EditRole]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": fmt.Sprintf("роль c id = %d была успешно обновлена!", id)})
}

// AddNewRole Add Role by ADMIN godoc
// @Summary  Add Role by Admin
// @Description  Add Role by Admin
// @Accept  json
// @Produce  json
// @Tags roles
// @Param  role  body models.Role true "Add role"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /roles/ [post]
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

// DeleteRole Delete Role by ID  ADMIN godoc
// @Summary  Delete Role by ID Admin
// @Description  Delete Role by Admin
// @Accept  json
// @Produce  json
// @Tags roles
// @Param  id path string true "role ID"
// @Param  role  body models.Role true "Update role"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /roles/{id} [delete]
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

// AttachRightToRole Attach Right to Role godoc
// @Summary  Attach Right to Role
// @Description  Attach Right to Role
// @Accept  json
// @Produce  json
// @Tags roles
// @Param  role_id path string true "role ID"
// @Param  right_id path string true "right ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /attach_right/{role_id}/{right_id} [post]
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

// DetachRightFromRole Detach Right to Role godoc
// @Summary  Detach Right to Role
// @Description  Detach Right to Role
// @Accept  json
// @Produce  json
// @Tags roles
// @Param  role_id path string true "role ID"
// @Param  right_id path string true "right ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /attach_right/{role_id}/{right_id} [delete]
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

// GetRoleByID Get Role by ID godoc
// @Summary Get Role by ID
// @Description et Role by ID
// @Accept  json
// @Produce  json
// @Tags roles
// @Param  id path int true "role ID"
// @Success 200 {object} models.Role
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /roles/{id}/details [get]
func GetRoleByID(c *gin.Context) {
	roleIdStr := c.Param("id")
	roleId, err := strconv.Atoi(roleIdStr)
	if err != nil {
		log.Println("[controller.GetRoleByID]|[strconv.Atoi]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	role, err := service.GetRoleByID(roleId)
	if err != nil {
		log.Println("[controller.GetRoleByID]|[service.GetRoleByID]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, role)
}
