package service

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
)

func CreateFormula(formula models.Formula) error {
	return repository.CreateFormula(formula)
}

func EditFormula(formula models.Formula) error {
	return repository.EditFormula(formula)
}
