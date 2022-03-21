package service

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
)

func GetAllRights() (rights []models.Right, err error) {
	return repository.GetAllRights()
}

func AddNewRight(right models.Right) error {
	return repository.AddNewRight(right)
}

func EditRight(right models.Right) error {
	return repository.EditRight(right)
}

func DeleteRight(id int) error {
	return repository.DeleteRight(id)
}

func GetRightByID(rightId int) (models.Right, error) {
	return repository.GetRightByID(rightId)
}
