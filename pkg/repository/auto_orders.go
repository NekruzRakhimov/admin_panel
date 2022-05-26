package repository

import (
	"admin_panel/db"
	"admin_panel/models"
)

func CreateGraphic(graphic models.Graphic) error {
	if err := db.GetDBConn().Table("graphics").Create(&graphic).Error; err != nil {
		return err
	}

	return nil
}

func GetAllGraphics() (graphics []models.Graphic, err error) {
	sqlQuery := `SELECT id,
					   number,
					   author,
					   supplier_name,
					   supplier_code,
					   region_name,
					   region_code,
					   store_name,
					   store_code,
					   nomenclature_group,
					   execution_period,
					   once_a_month,
					   twice_a_month,
					   is_on,
					   to_char(auto_order_date::date, 'DD.MM.YYYY'),
					   created_at,
					   application_day
				from graphics`
	if err = db.GetDBConn().Raw(sqlQuery).Scan(&graphics).Error; err != nil {
		return nil, err
	}

	return graphics, nil
}

func GetGraphicByID(id int) (graphic models.Graphic, err error) {
	sqlQuery := `SELECT id,
					   number,
					   author,
					   supplier_name,
					   supplier_code,
					   region_name,
					   region_code,
					   store_name,
					   store_code,
					   nomenclature_group,
					   execution_period,
					   once_a_month,
					   twice_a_month,
					   is_on,
					   to_char(auto_order_date::date, 'DD.MM.YYYY'),
					   created_at,
					   application_day
				from graphics WHERE id = ?`
	if err = db.GetDBConn().Raw(sqlQuery, id).Scan(&graphic).Error; err != nil {
		return models.Graphic{}, err
	}

	return graphic, nil
}

func EditGraphic(graphic models.Graphic) error {
	if err := db.GetDBConn().Table("graphics").Create(&graphic).Error; err != nil {
		return err
	}

	return nil
}

func GetAllAutoOrders() (autoOrders []models.AutoOrder, err error) {
	sqlQuery := `SELECT id,
					   graphic_name,
					   formula,
					   by_matrix,
					   schedule,
					   formed_orders_count,
					   organization,
					   status,
					   store,
					   to_char(created_at::timestamptz, 'DD.MM.YYYY') as created_at
				FROM auto_order ORDER BY id`
	if err = db.GetDBConn().Raw(sqlQuery).Scan(&autoOrders).Error; err != nil {
		return nil, err
	}

	return autoOrders, nil
}
