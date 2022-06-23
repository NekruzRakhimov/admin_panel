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

func CreateDefectsOrder(order models.DefectOrder) (models.DefectOrder, error) {
	if err := db.GetDBConn().Table("formed_defects").Omit("created_at").Create(&order).Error; err != nil {
		return models.DefectOrder{}, err
	}

	return order, nil
}

func SaveFormedDefect(order models.DefectOrder) error {
	if err := db.GetDBConn().Table("formed_defects").Omit("created_at").Save(&order).Error; err != nil {
		return err
	}

	return nil
}

func GetAllFormedDefects() (orders []models.DefectOrder, err error) {
	sqlQuery := `SELECT id,
					   "date",
					   file_name,
					   status,
					   to_char(created_at, 'DD.MM.YYYY hh:mi:ss') as created_at
				FROM formed_defects`
	if err = db.GetDBConn().Raw(sqlQuery).Scan(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func GetFormedDefectByID(id int) (order models.DefectOrder, err error) {
	sqlQuery := `SELECT id,
					   "date",
					   file_name,
					   status,
					   to_char(created_at, 'DD.MM.YYYY hh:mi:ss') as created_at
				FROM formed_defects WHERE id = ?`
	if err = db.GetDBConn().Raw(sqlQuery, id).Scan(&order).Error; err != nil {
		return models.DefectOrder{}, err
	}

	return order, nil
}
