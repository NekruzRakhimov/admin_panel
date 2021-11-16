package routes

import (
	"admin_panel/pkg/controller"
	"admin_panel/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func RunAllRoutes() {
	r := gin.Default()

	// Исползование CORS
	r.Use(controller.CORSMiddleware())

	// Установка Logger-а
	utils.SetLogger()

	// Форматирование логов
	utils.FormatLogs(r)

	// Статус код 500, при любых panic()
	r.Use(gin.Recovery())

	// Запуск роутов
	runAllRoutes(r)

	// Запуск сервера
	_ = r.Run(fmt.Sprintf("%s:%s", "0.0.0.0", os.Getenv("PORT")))
}

func runAllRoutes(r *gin.Engine) {
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", PingPong)

	users := r.Group("/users")
	users.GET("/", controller.GetAllUsers)
	users.POST("/", controller.CreateNewUser)
	users.PUT("/:id", controller.EditUser)
	users.DELETE("/:id", controller.DeleteUser)

	rights := r.Group("/rights")
	rights.GET("", controller.GetAllRights)
	rights.POST("", controller.AddNewRight)
	rights.PUT("/:id", controller.EditRight)
	rights.DELETE("/:id", controller.DeleteRight)

	roles := r.Group("/roles")
	roles.GET("", controller.GetAllRoles)
	roles.POST("", controller.AddNewRole)
	roles.PUT("/:id", controller.EditRole)
	roles.DELETE("/:id", controller.DeleteRole)

	r.POST("/attach_right/:role_id/:right_id", controller.AttachRightToRole)
	r.DELETE("/detach_right/:role_id/:right_id", controller.DetachRightFromRole)

	r.POST("/attach_role/:user_id/:role_id", controller.AttachRoleToUser)
	r.DELETE("/detach_role/:user_id/:role_id", controller.DetachRoleFromUser)
}

// PingPong Проверка
func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}
