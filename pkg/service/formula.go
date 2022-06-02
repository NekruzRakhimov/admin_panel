package service

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"encoding/json"
)

func CreateFormula(formula models.Formula) error {
	scheduleStr, err := json.Marshal(formula.Schedule)
	if err != nil {
		return err
	}

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

func GetAllFormulas() (formulas []models.Formula, err error) {
	formulas, err = repository.GetAllFormulas()
	if err != nil {
		return nil, err
	}

	for i, formula := range formulas {
		var schedule models.Schedule
		if err = json.Unmarshal([]byte(formula.ScheduleStr), &schedule); err != nil {
			return nil, err
		}

		formulas[i].Schedule = schedule
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

	return formula, nil
}
