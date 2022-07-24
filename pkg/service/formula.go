package service

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"encoding/json"
	"fmt"
)

func CreateFormula(formula models.Formula) error {
	scheduleStr, err := json.Marshal(formula.Schedule)
	if err != nil {
		return err
	}
	pharStr, err := json.Marshal(formula.Pharmacies)
	if err != nil {
		return err
	}

	formula.PharmaciesStr = string(pharStr)
	formula.ScheduleStr = string(scheduleStr)
	return repository.CreateFormula(formula)
}

func EditFormula(formula models.Formula) error {
	scheduleStr, err := json.Marshal(formula.Schedule)
	if err != nil {
		return err
	}

	formula.ScheduleStr = string(scheduleStr)
	return repository.EditFormula(formula)
}

func DeleteFormula(id int) error {
	return repository.DeleteFormula(id)
}

func GetAllFormulas() (formulas []models.Formula, err error) {
	formulas, err = repository.GetAllFormulas()
	if err != nil {
		return nil, err
	}

	for i, formula := range formulas {
		var schedule models.Schedule
		var pharmacy []models.Pharmacy
		if err = json.Unmarshal([]byte(formula.ScheduleStr), &schedule); err != nil {
			return nil, err
		}

		if err = json.Unmarshal([]byte(formula.PharmaciesStr), &pharmacy); err != nil {
			return nil, err
		}

		formulas[i].Schedule = schedule
		formulas[i].Pharmacies = pharmacy
	}

	return formulas, nil
}

func GetFormulaByID(id int) (formula models.Formula, err error) {
	formula, err = repository.GetFormulaByID(id)
	if err != nil {
		return models.Formula{}, err
	}

	if err = json.Unmarshal([]byte(formula.ScheduleStr), &formula.Schedule); err != nil {
		return models.Formula{}, err
	}
	fmt.Println("FORMULA", formula.PharmaciesStr)
	if err = json.Unmarshal([]byte(formula.PharmaciesStr), &formula.Pharmacies); err != nil {
		return models.Formula{}, err
	}

	return formula, nil
}
