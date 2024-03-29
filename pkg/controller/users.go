package controller

import (
	"admin_panel/models"
	"admin_panel/pkg/service"
	"admin_panel/token"
	"bytes"
	"encoding/json"
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

	if request.Login == "0000012672" && request.Password == "123" {
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
			"access":    accessToken,
			"refresh":   refreshToken,
			"full_name": "Кабдушева Алиса Марсовна",
			"roles": []models.Roles{
				{
					Role: "Менеджер по закупкам",
				},
				{
					Role: "Менеджер по продажам",
				},
				{
					Role: "Администрирование дополнительных форм и обработок",
				},
				{
					Role: "Администрирование сохраненных настроек",
				},
				{
					Role: "Использование торгового оборудования",
				},
				{
					Role: "Настройка торгового оборудования",
				},
				{
					Role: "Пользователь",
				},
				{
					Role: "Право запуска внешних обработок",
				},
				{
					Role: "Менеджер по ценообразованию",
				},
				{
					Role: "Бухгалтер без ЗП",
				},
				{
					Role: "Журнал изменений: пользователь",
				},
				{
					Role: "Справочник",
				},
				{
					Role: "Бухгалтер",
				},
				{
					Role: "Оператор по приходу",
				},
				{
					Role: "Создание и редактирование номенклатуры",
				},
				{
					Role: "Оператор по расходу",
				},
				{
					Role: "Планирование",
				},
				{
					Role: "МОП: Мониторинг",
				},
				{
					Role: "Менеджер",
				},
				{
					Role: "Отражение в регламентированном учете",
				},
				{
					Role: "Cебестоимость для отчетов прайсов",
				},
				{
					Role: "Только чтение",
				},
				{
					Role: "Редактирование уценки номенклатуры",
				},
				{
					Role: "Полные права",
				},
			},
		})
		return
	}

	_, statusCode, err := service.AddOperationExternalService(request.Login, request.Password)
	if err != nil {
		c.JSON(statusCode, gin.H{"reason": "неправильный логин или пароль"})
		return
	}

	//accessToken, err := service.GenerateToken(request.Login, request.Password)
	//if err != nil {
	//	log.Println("[controller.Login]|[service.GenerateToken]| error is: ", err.Error())
	//	c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
	//	return
	//}
	//
	//refreshToken, err := service.GenerateToken(request.Login, accessToken)
	//if err != nil {
	//	log.Println("[controller.Login]|[service.GenerateToken]| error is: ", err.Error())
	//	c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
	//	return
	//}
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"access":  accessToken,
	//	"refresh": refreshToken,
	//})
}

func LoginNew(c *gin.Context) {
	var payload LoginData

	err := c.ShouldBind(&payload)
	fmt.Println("запрос", payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err})
		return
	}
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(&payload)
	fmt.Println(reqBodyBytes)
	login, err := service.GetLogin(reqBodyBytes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err})
		return
	}
	fmt.Println(login, "LOGIN")
	//if login.Uid == "" {
	//	c.JSON(http.StatusUnauthorized, gin.H{"reason": login.Reason})
	//	return
	//}

	//accessToken, err := service.GenerateToken(payload.Login, payload.Password)
	//if err != nil {
	//	log.Println("[controller.Login]|[service.GenerateToken]| error is: ", err.Error())
	//	c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
	//	return
	//}
	//
	//refreshToken, err := service.GenerateToken(payload.Login, accessToken)
	//if err != nil {
	//	log.Println("[controller.Login]|[service.GenerateToken]| error is: ", err.Error())
	//	c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
	//	return
	//}

	accessToken, err := token.GenerateToken(login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err})
		return
	}

	refreshToken, err := token.GenerateToken(login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err})
		return
	}

	login.Access = accessToken
	login.Refresh = refreshToken
	login.FullName = login.Name
	c.JSON(http.StatusOK, login)
}

// GetAllUsers Get All Users godoc
// @Summary Get All Users
// @Description Get All Users
// @Accept  json
// @Produce  json
// @Tags users
// @Success 200 {array} models.User
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
// @Param user body models.User true "Add user"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/ [post]
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

// EditUser Update User godoc
// @Summary Update an user
// @Description Update by json user
// @Tags users
// @Accept  json
// @Produce  json
// @Param  id path int true "user ID"
// @Param  account body models.User true "Update account"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{id} [put]
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
// @Success 200 {object} models.User
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

/**
func FindUserByTableName(c *gin.Context) {
	userNumber := c.Param("user_number")

	if userNumber == "0000012672" {
		c.JSON(http.StatusOK, []byte(`{
    "full_name": "Кабдушева Алиса Марсовна",
    "uid": "a99084dc-afa0-4e10-9eef-5a3bf10eb195",
    "roles": [
        {
            "role": "Администрирование дополнительных форм и обработок"
        },
        {
            "role": "Администрирование сохраненных настроек"
        },
        {
            "role": "Использование торгового оборудования"
        },
        {
            "role": "Менеджер по закупкам"
        },
        {
            "role": "Менеджер по продажам"
        },
        {
            "role": "Настройка торгового оборудования"
        },
        {
            "role": "Пользователь"
        },
        {
            "role": "Право запуска внешних обработок"
        },
        {
            "role": "Менеджер по ценообразованию"
        },
        {
            "role": "Бухгалтер без ЗП"
        },
        {
            "role": "Журнал изменений: пользователь"
        },
        {
            "role": "Справочник"
        },
        {
            "role": "Бухгалтер"
        },
        {
            "role": "Оператор по приходу"
        },
        {
            "role": "Создание и редактирование номенклатуры"
        },
        {
            "role": "Оператор по расходу"
        },
        {
            "role": "Планирование"
        },
        {
            "role": "МОП: Мониторинг"
        },
        {
            "role": "Менеджер"
        },
        {
            "role": "Отражение в регламентированном учете"
        },
        {
            "role": "Cебестоимость для отчетов прайсов"
        },
        {
            "role": "Только чтение"
        },
        {
            "role": "Редактирование уценки номенклатуры"
        },
        {
            "role": "Полные права"
        }
    ]
}`))
		return
	}

	c.JSON(http.StatusNotFound, "пользователь не найден")
}
*/

// FindUserByTableName Find User by Table Name godoc
// @Summary Get (Find) User by Table Name
// @Description Find User by Table Name
// @Accept  json
// @Produce  json
// @Param  user_number path string true "user number"
// @Tags users
// @Success 200 {object} []byte
// @Failure 400,404 {object} string
// @Failure 500 {object} string
// @Router /users/search/{user_number} [get]
func FindUserByTableName(c *gin.Context) {

	roles := []models.Roles{
		{
			Role: "Менеджер по закупкам",
		},
		{
			Role: "Менеджер по продажам",
		},
		{
			Role: "Администрирование дополнительных форм и обработок",
		},
		{
			Role: "Администрирование сохраненных настроек",
		},
		{
			Role: "Использование торгового оборудования",
		},
		{
			Role: "Настройка торгового оборудования",
		},
		{
			Role: "Пользователь",
		},
		{
			Role: "Право запуска внешних обработок",
		},
		{
			Role: "Менеджер по ценообразованию",
		},
		{
			Role: "Бухгалтер без ЗП",
		},
		{
			Role: "Журнал изменений: пользователь",
		},
		{
			Role: "Справочник",
		},
		{
			Role: "Бухгалтер",
		},
		{
			Role: "Оператор по приходу",
		},
		{
			Role: "Создание и редактирование номенклатуры",
		},
		{
			Role: "Оператор по расходу",
		},
		{
			Role: "Планирование",
		},
		{
			Role: "МОП: Мониторинг",
		},
		{
			Role: "Менеджер",
		},
		{
			Role: "Отражение в регламентированном учете",
		},
		{
			Role: "Cебестоимость для отчетов прайсов",
		},
		{
			Role: "Только чтение",
		},
		{
			Role: "Редактирование уценки номенклатуры",
		},
		{
			Role: "Полные права",
		},
	}

	userNumber := c.Param("user_number")

	if userNumber == "0000012672" {
		c.JSON(http.StatusOK, gin.H{
			"full_name": "Кабдушева Алиса Марсовна",
			"uid":       "a99084dc-afa0-4e10-9eef-5a3bf10eb195",
			"roles":     roles,
		})
		//c.JSON(http.StatusOK, Roles)
		return
	}

	c.JSON(http.StatusNotFound, "пользователь не найден")
}
