package routes

import (
	"admin_panel/pkg/controller"
	"admin_panel/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"

	//_ "github.com/rizalgowandy/go-swag-sample/docs/ginsimple" // you need to update github.com/rizalgowandy/go-swag-sample with your own project path
	_ "admin_panel/docs"
)

//

func RunAllRoutes() {
	r := gin.Default()

	// Использование CORS
	r.Use(controller.CORSMiddleware())

	// Установка Logger-а
	utils.SetLogger()

	// Форматирование логов
	utils.FormatLogs(r)

	// Статус код 500, при любых panic()
	r.Use(gin.Recovery())

	//r.Use(limits.RequestSizeLimiter(100))

	// Запуск end-point'ов
	runAllRoutes(r)

	// Запуск сервера
	runServer(r)

}

func runAllRoutes(r *gin.Engine) {

	r.GET("/", HealthCheck)
	//r.POST("/rbdiscountforsalesgrowth", controller.RbDiscountForSalesGrowth)
	r.POST("/login", controller.Login)
	r.POST("/getdisper", controller.GetDisPer)
	r.POST("/getdisp", controller.DiscountRBPeriodTime)
	r.POST("/getcode", controller.GetContractCode)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tempRoutes(r)
	Check1CRoutes(r)
	ContractRoutes(r)
	DictionariesRoutes(r)
	AdminRoutes(r)
	ReportsRoutes(r)
	NotificationsRoutes(r)
	routesFor1C(r)
}

func runServer(r *gin.Engine) {
	var (
		port string
		host string
	)
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "3000"
		host = "localhost"
	} else {
		host = "0.0.0.0"
	}
	err := r.Run(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Println(err)
	}
}

func routesFor1C(r *gin.Engine) {
	r.POST("/1c/data", controller.SaveDataFrom1C)
}

func tempRoutes(r *gin.Engine) {
	r.POST("/rb_brand/", controller.GetBrandInfo)
	r.POST("/rb_brand/excel/", controller.GenerateReportBrand)

	r.POST("/check_1c_get_data", controller.Check1CGetData)

	r.GET("/cars", controller.GetCarsBrand)
}

func Check1CRoutes(r *gin.Engine) {
	r.GET("/counterparty/:client", controller.CounterpartyContract)

	r.POST("/price_type", controller.GetPriceType)
	r.POST("/create_price_type", controller.CreatePriceType)
	r.GET("/currencies", controller.GetCurrencies) // add swagger

	r.GET("/brands/", controller.GetBrands)
	r.GET("/add_brand/", controller.AddBrand)

	r.GET("/country/", controller.GetCountries)
	r.POST("/sales/", controller.GetSales)

	r.POST("/presentationdiscount", controller.PresentationDiscount)
	r.POST("/get_excell_brand", controller.GetExcelBrand)
	r.POST("/get_excell_growth", controller.GetExcelGrowth)
}

func ReportsRoutes(r *gin.Engine) {
	reports := r.Group("/reports")
	reports.POST("/doubted_discounts", controller.GetDoubtedDiscounts)
	//reports.PUT("/doubted_discounts", controller.SaveDoubtedDiscountsResults)
	reports.POST("/rb", controller.GetAllRBByContractorBIN)
	reports.POST("/rb/excel", controller.FormExcelForRB)
	//reports.GET("/rb_brand/excel", controller.FormExcelForRBBrand)

	reports.POST("/dd", controller.GetAllDeferredDiscounts)
	reports.POST("/dd/excel", controller.FormExcelForDeferredDiscounts)
}

func ContractRoutes(r *gin.Engine) {
	r.GET("/search_history_ex/:id/", controller.SearchHistoryExecution)
	r.GET("/change_date_contract/", controller.ChangeDataContract)
	r.GET("/search_contract/", controller.SearchContractByNumber)
	r.GET("/search_history/:id", controller.SearchContractDC) // TODO: тут нам нужен ID договора (я тебе об этом говорил)

	contract := r.Group("/contract")
	contract.GET("", controller.GetAllContracts)
	contract.GET("/products_template", controller.GetProductsTemplate)
	contract.POST("/:type", controller.CreateContract)
	contract.POST("/additional_agreement/:id", controller.AddAdditionalAgreement)
	contract.POST("/individual_contract", controller.AddIndividualContract) //TODO: PutBack Id
	contract.GET("/individual_contract/:id", controller.GetIndividContract) //TODO: PutBack Id
	contract.PUT("/:type/:id", controller.EditContract)
	contract.GET("/:id/details", controller.GetContractDetails)
	contract.PUT("/conform/:id", controller.ConformContract)
	contract.PUT("/cancel/:id", controller.CancelContract)
	contract.PUT("/finish/:id", controller.FinishContract)
	contract.PUT("/revision/:id", controller.RevisionContract)
	contract.POST("/products", controller.ConvertExcelToStruct)
	contract.GET("/history/:id", controller.GetContractHistory)
	contract.GET("/status_history/:id", controller.GetContractStatusChangesHistory)

	contract.POST("/form/:contract_type/:with_temp_conditions", controller.FormContract)
}

func DictionariesRoutes(r *gin.Engine) {
	dictionary := r.Group("/dictionary")
	dictionary.GET("", controller.GetAllDictionaries)
	dictionary.GET("/:id", controller.GetAllDictionaryByID)
	dictionary.POST("", controller.CreatDictionary)
	dictionary.PUT("/:id", controller.EditDictionary)
	dictionary.DELETE("/:id", controller.DeleteDictionary)

	dictionaryValues := dictionary.Group("/:id/value")
	dictionaryValues.GET("", controller.GetAllDictionaryValues)
	dictionaryValues.POST("", controller.CreateDictionaryValue)
	dictionaryValues.PUT("/:value_id", controller.EditDictionaryValue)
	dictionaryValues.DELETE("/:value_id", controller.DeleteDictionaryValue)

	dictionary.GET("/currencies", controller.GetAllCurrencies)
	dictionary.GET("/positions", controller.GetAllPositions)
	dictionary.GET("/addresses", controller.GetAllAddresses)
	dictionary.GET("/frequency_deferred_discounts", controller.GetAllFrequencyDeferredDiscounts)

}

func AdminRoutes(r *gin.Engine) {
	r.POST("/client_search", controller.SearchBinClient)

	users := r.Group("/users")
	users.GET("/search/:user_number", controller.FindUserByTableName)
	users.GET("/", controller.GetAllUsers)
	users.GET("/:id/details", controller.GetUserById)
	users.POST("/", controller.CreateNewUser)
	users.PUT("/:id", controller.EditUser)
	users.DELETE("/:id", controller.DeleteUser)
	//	users.GET("/search/:user_number", controller.FindUserByTableName)

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
}

func NotificationsRoutes(r *gin.Engine) {
	//r.POST("/getcontractnumb", controller.SearchNotifications)
	r.GET("/notifications", controller.GetNotifications)
	r.GET("/search_notification/", controller.SearchNotification)
	r.GET("/notification", controller.Notification)
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description gets the status of server.
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
