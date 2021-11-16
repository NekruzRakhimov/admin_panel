package main

import (
	"admin_panel/db"
	"admin_panel/routes"
	"admin_panel/utils"
)

func main() {
	utils.ReadSettings()

	db.StartDbConnection()

	routes.RunAllRoutes()
}
