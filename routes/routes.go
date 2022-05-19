package routes

import (
	"admin_panel/pkg/controller"
	"admin_panel/pkg/service"
	"admin_panel/token"
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

func Check(c *gin.Context) {
	service.CreateNecessity()
}

func runAllRoutes(r *gin.Engine) {
	r.GET("/", Check)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/login", controller.LoginNew)

	cr := r.Group("/" /*, token.UserIdentity*/)
	// cr - closed routes
	//token.UserIdentity - это middleware для проверки токена

	cr.POST("/token", token.Token)
	//r.POST("/loginnew", controller.LoginNew)

	cr.POST("/file/upload", controller.UploadFile)
	cr.GET("/file/download", controller.DownloadFile)

	tempRoutes(cr)
	Check1CRoutes(cr)
	ContractRoutes(cr)
	DictionariesRoutes(cr)
	AdminRoutes(cr)
	ReportsRoutes(cr)
	NotificationsRoutes(cr)
	routesFor1C(cr)
	DDRoutes(cr)

	cr.GET("/graphic", controller.GetAllGraphics)
	cr.GET("/graphic/:id/details", controller.GetGraphicByID)
	cr.POST("/graphic", controller.CreateGraphic)
	cr.PUT("/graphic/:id", controller.EditGraphic)

	cr.GET("/auto_orders", controller.GetAllAutoOrders)
	cr.POST("/auto_orders", controller.FormAutoOrder)

	cr.POST("/defects/pharmacy/PF", controller.GetDefectsByPharmacyPF)

	hyperstockServ := service.NewHyperstocksService()
	defectServ := service.NewDefectsService()
	controller.NewHyperstocksController(hyperstockServ).HandleRoutes(cr)
	controller.NewDefectsController(defectServ).HandleRoutes(cr)
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

func routesFor1C(r *gin.RouterGroup) {
	r.POST("/1c/data", controller.SaveDataFrom1C)
	r.GET("/regions_from_1c", controller.GetRegions)
}

func DDRoutes(r *gin.RouterGroup) {
	r.POST("/getddone", controller.DDOne)

}

func tempRoutes(r *gin.RouterGroup) {
	r.POST("/getdisper", controller.GetDisPer)
	r.POST("/getpurchase", controller.GetPurchase)
	r.POST("/getdisp", controller.DiscountRBPeriodTime)
	r.POST("/getrbseven", controller.DiscountRB7)
	r.POST("/getrbfour", controller.DiscountRB4)
	r.POST("/getrbfourteen", controller.DiscountRB14)
	r.POST("/getcode", controller.GetContractCode)
	r.POST("/rbseventeen", controller.DiscountRB17)
	r.POST("/rbsixteen", controller.DiscountRB16)
	r.POST("/rbfifteen", controller.DiscountRB15)
	r.POST("/check_contract", controller.GetAllContractDetailByBIN)

	r.POST("/rb_brand/", controller.GetBrandInfo)
	r.POST("/rb_brand/excel/", controller.GenerateReportBrand)
	r.POST("/check_1c_get_data", controller.Check1CGetData)
	r.GET("/cars", controller.GetCarsBrand)
}

func Check1CRoutes(r *gin.RouterGroup) {
	r.GET("/contracts_from_1c/", controller.CheckContractIn1C)
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
	r.POST("/get_rb1", controller.GetRb1)
	r.POST("/get_rb3", controller.GetRb3)
	r.POST("/get_rb5", controller.GetRb5)
	r.GET("/store_regions", controller.GetStoreRegions)
	r.GET("/matrix/:store_code", controller.GetMatrix)

}

func ReportsRoutes(r *gin.RouterGroup) {
	reports := r.Group("/reports")
	reports.POST("/doubted_discounts", controller.GetDoubtedDiscounts)
	//reports.PUT("/doubted_discounts", controller.SaveDoubtedDiscountsResults)

	reports.POST("/rb", controller.GetAllRBByContractorBIN)
	reports.POST("/rb/update", controller.UpdateRbReports)
	reports.POST("/rb/excel", controller.FormExcelForRB)
	reports.GET("/rb/stored", controller.GetAllStoredReports)
	reports.GET("/rb/stored/:id/details", controller.GetStoredReportDetails)
	reports.GET("/rb/stored/:id/details/excel", controller.GetExcelForStoredExcelReport)
	reports.GET("/search_report_rb/", controller.SearchReportRB)
	//reports.GET("/rb_brand/excel", controller.FormExcelForRBBrand)

	reports.POST("/dd", controller.GetAllDeferredDiscounts)
	reports.POST("/dd/excel", controller.FormExcelForDeferredDiscounts)
	reports.GET("/dd/stored", controller.GetAllDdStoredReports)
	reports.GET("/dd/stored/:id/details", controller.GetStoredDdReportDetails)
	reports.GET("/dd/stored/:id/details/excel", controller.GetExcelForDdStoredExcelReport)
	reports.GET("/search_report_dd/", controller.SearchReportDD)
}

func ContractRoutes(r *gin.RouterGroup) {
	r.GET("/search_history_ex/:id/", controller.SearchHistoryExecution)
	r.GET("/change_date_contract/", controller.ChangeDataContract)
	r.GET("/search_contract/", controller.SearchContractByNumber)
	r.GET("/search_history/:id", controller.SearchContractDC) // TODO: тут нам нужен ID договора (я тебе об этом говорил)
	r.GET("/suppliers/", controller.GetSuppliers)

	r.POST("/suppliers", controller.SaveSuppliers)
	r.POST("/products", controller.GetProducts)

	contract := r.Group("/contract")
	contract.GET("", controller.GetAllContracts)
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
	contract.GET("/products_template", controller.GetProductsTemplate)
	contract.POST("/products", controller.ConvertExcelToStruct)
	contract.GET("/history/:id", controller.GetContractHistory)
	contract.GET("/status_history/:id", controller.GetContractStatusChangesHistory)

	contract.POST("/form/:contract_type/:with_temp_conditions", controller.FormContract)
}

func DictionariesRoutes(r *gin.RouterGroup) {
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

func AdminRoutes(r *gin.RouterGroup) {
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

func NotificationsRoutes(r *gin.RouterGroup) {
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
