package repository

import (
	"admin_panel/db"
	"admin_panel/models"
	"fmt"
)

func CreateFormula(formula models.Formula) error {
	if err := db.GetDBConn().Table("formulas").Create(&formula).Error; err != nil {
		return err
	}

	return nil
}

func EditFormula(formula models.Formula) error {
	if err := db.GetDBConn().Table("formulas").Save(&formula).Error; err != nil {
		return err
	}
	return nil
}

func DeleteFormula(id int) error {
	fmt.Println("ID", id)
	update := "UPDATE formulas SET is_removed = ? WHERE id = ?"
	err := db.GetDBConn().Exec(update, true, id).Error
	if err != nil {
		return err
	}

	return nil
}

func GetAllFormulas() (formulas []models.Formula, err error) {
	sqlQuery := "SELECT * FROM formulas"
	if err = db.GetDBConn().Raw(sqlQuery).Scan(&formulas).Error; err != nil {
		return nil, err
	}
	return formulas, err
}

func GetFormulaByID(id int) (formula models.Formula, err error) {
	sqlQuery := "SELECT * FROM formulas WHERE id = ?"
	if err = db.GetDBConn().Raw(sqlQuery, id).Scan(&formula).Error; err != nil {
		return models.Formula{}, err
	}
	return formula, err
}
