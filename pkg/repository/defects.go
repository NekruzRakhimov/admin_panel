package repository

import (
	"admin_panel/db"
	"admin_panel/models"
	"fmt"
)

func SaveAllStoresFrom1C(stores []models.Store) error {
	fmt.Println(len(stores))
	for _, store := range stores {
		if err := db.GetDBConn().Table("stores").Create(&store).Error; err != nil {
			return nil
		}
	}

	return nil
}

func GetAllStores() (stores []models.Store, err error) {
	sqlQuery := "SELECT * FROM stores"
	if err = db.GetDBConn().Raw(sqlQuery).Scan(&stores).Error; err != nil {
		return nil, err
	}

	return stores, nil
}
