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
	//TODO: интеграция по договорам

	//TODO:  добавить функцию обработчика
	r.GET("/counterparty/:client", controller.CounterpartyContract)

	contract := r.Group("/contract")
	contract.GET("", controller.GetAllContracts)
	contract.GET("/products_template", controller.GetProductsTemplate)
	contract.POST("/:type", controller.CreateContract)
	contract.POST("/additional_agreement/:id", controller.AddAdditionalAgreement)
	contract.PUT("/:type/:id", controller.EditContract)
	contract.GET("/:id/details", controller.GetContractDetails)
	contract.PUT("/conform/:id", controller.ConformContract)
	contract.PUT("/cancel/:id", controller.CancelContract)
	contract.PUT("/finish/:id", controller.FinishContract)
	contract.PUT("/revision/:id", controller.RevisionContract)
	contract.POST("/products", controller.ConvertExcelToStruct)
	contract.GET("/history/:id", controller.GetContractHistory)

	contract.POST("/form/:with_temp_conditions", controller.FormContract)

	dictionary := r.Group("/dictionary")
	dictionary.GET("/currencies", controller.GetAllCurrencies)
	dictionary.GET("/positions", controller.GetAllPositions)
	dictionary.GET("/addresses", controller.GetAllAddresses)
	dictionary.GET("/frequency_deferred_discounts", controller.GetAllFrequencyDeferredDiscounts)

	users := r.Group("/users")
	users.GET("/", controller.GetAllUsers)
	users.GET("/:id/details", controller.GetUserById)
	users.POST("/", controller.CreateNewUser)
	users.PUT("/:id", controller.EditUser)
	users.DELETE("/:id", controller.DeleteUser)

	rights := r.Group("/rights")
	rights.GET("", controller.GetAllRights)
	rights.GET("/:id/details", controller.GetRightByID)
	rights.POST("", controller.AddNewRight)
	rights.PUT("/:id", controller.EditRight)
	rights.DELETE("/:id", controller.DeleteRight)

	roles := r.Group("/roles")
	roles.GET("", controller.GetAllRoles)
	roles.GET("/:id/details", controller.GetRoleByID)
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
	//_ = r.Run(fmt.Sprintf("%s:%s", "localhost", "3001"))
	//if err := r.Run(":3000"); err != nil {
	//	log.Fatal(err)
	//}

}

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
