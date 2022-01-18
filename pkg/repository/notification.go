package repository

import (
	"admin_panel/db"
	"admin_panel/model"
)

func GetNotification() (notifications []model.Notification) {
	db.GetDBConn().Raw("SELECT bin, contract_number, contract_date, type, email, status FROM notification").Scan(&notifications)
	return notifications
}
