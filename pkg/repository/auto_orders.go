package repository

import (
	"admin_panel/db"
	"admin_panel/models"
)

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
				FROM auto_order ORDER BY id DESC`
	if err = db.GetDBConn().Raw(sqlQuery).Scan(&autoOrders).Error; err != nil {
		return nil, err
	}

	return autoOrders, nil
}

func SaveFormedGraphics(formedGraphics []models.FormedGraphic) error {
	for _, graphic := range formedGraphics {
		if len(graphic.Products) == 0 {
			continue
		}

		if err := db.GetDBConn().Table("formed_graphics").Omit("formula_id", "graphic_name", "supplier", "store", "schedule", "products", "created_at").Create(&graphic).Error; err != nil {
			return err
		}

		if err := SaveFormedGraphicProducts(graphic.Products, graphic.ID); err != nil {
			return err
		}

	}

	return nil
}

func SaveFormedGraphicProducts(products []models.FormedGraphicProduct, formedGraphicID int) error {
	for _, product := range products {
		product.GraphicID = formedGraphicID
		if err := db.GetDBConn().Table("formed_graphic_products").Create(&product).Error; err != nil {
			return err
		}
	}

	return nil
}

func GetAllFormedGraphics(formulaID int) (graphics []models.FormedGraphic, err error) {
	sqlQuery := `SELECT fg.id,
					   g.number          as graphic_name,
					   g.supplier_name   as supplier,
					   g.store_name      as store,
					   fg.by_matrix,
					   g.application_day as schedule,
					   fg.product_availability_days,
					   fg.dister_days,
					   fg.store_days,
					   fg.status         as status
				FROM formed_graphics fg
						 JOIN graphics g ON fg.graphic_id = g.id
				WHERE fg.formula_id = ?`
	if err = db.GetDBConn().Raw(sqlQuery, formulaID).Scan(&graphics).Error; err != nil {
		return nil, err
	}

	return graphics, nil
}

func GetFormedGraphicByID(id int) (graphic models.FormedGraphic, err error) {
	sqlQuery := `SELECT fg.id,
					   fg.is_letter,
					   g.number          as graphic_name,
					   fg.graphic_id     as graphic_id,
					   g.supplier_name   as supplier,
					   g.store_name      as store,
					   fg.by_matrix,
					   g.application_day as schedule,
					   fg.product_availability_days,
					   fg.dister_days,
					   fg.store_days,
					   fg.status         as status,
					   fg.formula_id,
					   to_char(fg.created_at::date, 'DD.MM.YYYY') as created_at
				FROM formed_graphics fg
						 JOIN graphics g ON fg.graphic_id = g.id
				WHERE fg.id = ?`
	if err = db.GetDBConn().Raw(sqlQuery, id).Scan(&graphic).Error; err != nil {
		return models.FormedGraphic{}, err
	}

	return graphic, nil
}

func GetAllFormedGraphicsProducts(formedGraphicID int) (products []models.FormedGraphicProduct, err error) {
	sqlQuery := "SELECT * FROM formed_graphic_products WHERE graphic_id = ?"
	if err = db.GetDBConn().Raw(sqlQuery, formedGraphicID).Scan(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func CancelFormedFormula(formulaID int, comment string) error {
	sqlQuery := "UPDATE auto_order set status = ? where id = ?"
	if err := db.GetDBConn().Exec(sqlQuery, "отменен", formulaID).Error; err != nil {
		return err
	}

	return nil
}

func SendFormedFormula(formulaID int, comment string) error {
	sqlQuery := "UPDATE auto_order set status = ? where id = ?"
	if err := db.GetDBConn().Exec(sqlQuery, "отправлено", formulaID).Error; err != nil {
		return err
	}

	return nil
}

func CancelFormedGraphic(graphicID int, comment string) error {
	sqlQuery := "UPDATE formed_graphics set status = ? where id = ?"
	if err := db.GetDBConn().Exec(sqlQuery, "отменен", graphicID).Error; err != nil {
		return err
	}

	return nil
}

func CreateFormedFormula() (id int, err error) {
	var formula models.AutoOrder
	formula.Status = "в процессе"
	if err = db.GetDBConn().Table("auto_order").Omit("created_at").Create(&formula).Error; err != nil {
		return 0, err
	}

	return formula.ID, nil
}

func ChangeFormedFormulaStatus(id int) error {
	sqlQuery := "UPDATE auto_order set status = $1 WHERE id = $2"
	return db.GetDBConn().Exec(sqlQuery, "сформирован", id).Error
}
