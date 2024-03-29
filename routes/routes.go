package routes

import (
	"admin_panel/models"
	"admin_panel/pkg/controller"
	"admin_panel/pkg/service"
	"admin_panel/token"
	"admin_panel/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"math/rand"
	"net/http"
	"os"

	//_ "github.com/rizalgowandy/go-swag-sample/docs/ginsimple" // you need to update github.com/rizalgowandy/go-swag-sample with your own project path
	_ "admin_panel/docs"
)

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

func Check(c *gin.Context) {
	service.CreateNecessity()
	c.JSON(http.StatusOK, gin.H{"reason": "up and working"})
}

func GetData(c *gin.Context) {
	products, err := service.GetAllFormedGraphicsProducts(26)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	var data models.Data
	data.OrderId = "1"
	data.SupplierCode = "000000976"
	data.StoreCode = "Ф0001121 "

	for _, product := range products {
		data.Products = append(data.Products, models.DataProducts{
			ProductCode: product.ProductCode,
			SalesCount:  fmt.Sprintf("%2.2f", product.OrderQnt),
			Price:       fmt.Sprintf("%2.2f", rand.Float32()*100),
		})
	}

	c.JSON(http.StatusOK, data)
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

	cr.GET("/get_report", GetData)

	cr.GET("/graphic", controller.GetAllGraphics)
	cr.GET("/graphic/:id/details", controller.GetGraphicByID)
	cr.POST("/graphic", controller.CreateGraphic)
	cr.PUT("/graphic/:id", controller.EditGraphic)
	cr.PUT("/delete_graphic/:id", controller.DeleteGraphic)

	cr.POST("/auto_orders", controller.FormAutoOrder)
	cr.GET("/auto_orders", controller.GetAllAutoOrders)
	cr.POST("/auto_orders/:formula_id/send", controller.SendAutoOrdersTo1C)
	cr.DELETE("/auto_orders/:formula_id", controller.CancelFormedFormula)

	cr.GET("/auto_orders/:formula_id/graphics", controller.GetAllFormedGraphics)
	cr.DELETE("/auto_orders/:formula_id/graphics/:graphic_id", controller.CancelFormedGraphic)

	cr.GET("/auto_orders/:formula_id/graphics/:graphic_id/products", controller.GetAllFormedGraphicProducts)

	cr.GET("/auto_orders/types", controller.GetAllAutoOrderTypes)
	cr.GET("/auto_orders/analyses_period", controller.GetAllAutoOrderAnalysesPeriod)

	cr.POST("/formula", controller.CreateFormula)
	cr.GET("/formula", controller.GetAllFormulas)
	cr.GET("/formula/:id/details", controller.GetFormulaByID)
	cr.PUT("/formula/:id", controller.EditFormula)
	cr.PUT("/delete_formula/:id", controller.DeleteFormulaByID)

	cr.GET("/formula/parameters", controller.GetFormulaParameters)

	//cr.POST("/defects/pharmacy/PF", controller.GetDefectsByPharmacyPF)
	cr.POST("/defects/pharmacy/PF", controller.OrderDefectsPfReport)
	cr.GET("/defects/pharmacy/PF/list", controller.GetDefectsPfList)
	cr.GET("/defects/pharmacy/PF/excel/:id", controller.GetDefectExcel)

	//cr.POST("/defects/pharmacy/LS", controller.GetDefectsByPharmacyLS)
	cr.POST("/defects/pharmacy/LS", controller.OrderDefectsLsReport)
	cr.GET("/defects/pharmacy/LS/list", controller.GetDefectsLsList)
	cr.GET("/defects/pharmacy/LS/excel/:id", controller.GetDefectExcel)

	cr.POST("/check/sales_cgount", controller.GetSalesCount)

	cr.POST("/save_matrix", controller.SaveMatrix)

	//cr.GET("/letter", controller.GetSegmentByID)
	//cr.GET("/send_letter/:id", controller.SendLetter)

	hyperstockServ := service.NewHyperstocksService()
	defectServ := service.NewDefectsService()
	forecastServ := service.NewForecastService()
	controller.NewHyperstocksController(hyperstockServ).HandleRoutes(cr)
	controller.NewDefectsController(defectServ).HandleRoutes(cr)
	controller.NewForecastController(forecastServ).HandleRoutes(cr)
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
	r.GET("/check_enddate", controller.CheckEndContract)

	r.POST("/rb_brand/", controller.GetBrandInfo)
	r.POST("/rb_brand/excel/", controller.GenerateReportBrand)
	r.POST("/check_1c_get_data", controller.Check1CGetData)
	r.GET("/cars", controller.GetCarsBrand)
	//contract.POST("/products", controller.ConvertExcelToStruct)
	r.GET("/segments", controller.GetSegments)

	r.POST("/segment_product", controller.ConvertExcelToStructProductsAndRegion)
	r.POST("/segment", controller.CreateSegment)
	r.PUT("/segment/:id", controller.ChangeSegment)
	r.GET("/segment/:id", controller.GetSegmentByID)
	r.GET("/delete_segment/:id", controller.DeleteSegmentByID)
	r.GET("/letter", controller.SendLetter)

}

func Check1CRoutes(r *gin.RouterGroup) {
	r.GET("/contracts_from_1c/", controller.CheckContractIn1C)
	r.GET("/contracts_1c/", controller.CheckContractIn1C)
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
	r.GET("/pharmacies/", controller.ListOrganizations)
	r.GET("/organizations", controller.ListOrganizations)
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
	//reports.GET("/dd/stored", controller.GetAllDdStoredReports)
	reports.GET("/dd/stored", controller.GetAllStoredReports)
	reports.GET("/dd/stored/:id/details", controller.GetStoredDdReportDetails)
	reports.GET("/dd/stored/:id/details/excel", controller.GetExcelForDdStoredExcelReport)
	reports.GET("/search_report_dd/", controller.SearchReportDD)
}

func ContractRoutes(r *gin.RouterGroup) {
	r.GET("/segments_template", controller.GetSegmentsTemplate)

	r.GET("/search_history_ex/:id/", controller.SearchHistoryExecution)
	r.GET("/change_date_contract/", controller.ChangeDataContract)
	r.GET("/search_contract/", controller.SearchContractByNumber)
	r.GET("/search_history/:id", controller.SearchContractDC) // TODO: тут нам нужен ID договора (я тебе об этом говорил)
	r.GET("/suppliers/", controller.GetSuppliers)
	r.GET("/supplier_name/", controller.GetSupplierName)

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
