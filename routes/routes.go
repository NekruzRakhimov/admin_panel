package routes

import (
	"admin_panel/pkg/controller"
	"admin_panel/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
	//_ "github.com/rizalgowandy/go-swag-sample/docs/ginsimple" // you need to update github.com/rizalgowandy/go-swag-sample with your own project path
	_ "admin_panel/docs"
)

func RunAllRoutes() {
	r := gin.Default()

	// Исползование CORS

	r.Use(controller.CORSMiddleware())

	//r.Use(controller.CORSMiddleware())

	// Установка Logger-а
	utils.SetLogger()

	// Форматирование логов
	utils.FormatLogs(r)

	// Статус код 500, при любых panic()
	r.Use(gin.Recovery())

	// Исползование CORS
	r.Use(controller.CORSMiddleware())

	// Запуск роутов
	runAllRoutes(r)

	// Запуск сервера
	//_ = r.Run(fmt.Sprintf("%s:%s", "0.0.0.0", os.Getenv("PORT")))

	//_ = r.Run(":3000")

}

func runAllRoutes(r *gin.Engine) {
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", HealthCheck)
	r.POST("/contract/:type", controller.CreateContract)
	r.GET("/contract", controller.GetAllContracts)
	r.GET("/contract/:id/details", controller.GetContractDetails)

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

	//url := ginSwagger.URL("http://localhost:3000/swagger/doc.json") // The url pointing to API definition
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	//	use ginSwagger middleware to serve the API docs
	//	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//	Start server

	_ = r.Run(fmt.Sprintf("%s:%s", "0.0.0.0", os.Getenv("PORT")))
	//if err := r.Run(":3000"); err != nil {
	//	log.Fatal(err)
	//}

	//_ = r.Run(fmt.Sprintf("%s:%s", "0.0.0.0", os.Getenv("PORT")))
	//if err := r.Run(":3000"); err != nil {
	//	log.Fatal(err)
	//	//}

}

//func Init()  {
//	r := gin.New()
//
//	// Routes
//	r.GET("/ping", Ping)
//	r.GET("/", HealthCheck)
//
//	url := ginSwagger.URL("http://localhost:3000/swagger/doc.json") // The url pointing to API definition
//	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
//	// use ginSwagger middleware to serve the API docs
//	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
//
//	// Start server
//	if err := r.Run(":3000"); err != nil {
//		log.Fatal(err)
//	}
//}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *gin.Context) {
	//res := map[string]interface{}{
	//	"data": "Server is up and running",
	//}

	c.JSON(http.StatusOK, gin.H{"data": "Server is up and running"})
}

// Ping godoc
// @Summary Ping pong
// @Description Ping.
// @Tags root
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /ping [get]
func Ping(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Ping Pong"})

}
