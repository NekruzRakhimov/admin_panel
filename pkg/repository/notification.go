package repository

import (
	"admin_panel/db"
	"admin_panel/models"
	"errors"
	"github.com/jinzhu/gorm"
)

func GetNotification() (notifications []models.Notification) {
	db.GetDBConn().Raw("SELECT id, bin, contract_number, contract_date, type, email, status FROM notification").Scan(&notifications)
	return notifications
}

func SearchNotification(number string) ([]models.Notification, error) {
	var notifications []models.Notification

	err := db.GetDBConn().Raw("SELECT id, bin, contract_number, contract_date, type, email, status FROM notification WHERE contract_number like $1", "%"+number+"%").Scan(&notifications).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return notifications, nil

}
