package main

import (
	"admin_panel/routes"
	"admin_panel/utils"
)

// @title Gin Swagger Admin-Panel Api
// @version 1.0
// @description Админка, чтобы проверить роуты и права пользователей.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email aziz.rahimov0001@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {

	//host localhost:3000
	utils.ReadSettings()

	//db.StartDbConnection()

	routes.RunAllRoutes()
}
