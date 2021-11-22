package service

import (
	"admin_panel/model"
	"admin_panel/pkg/repository"
)

func GetAllRights() (rights []model.Right, err error) {
	return repository.GetAllRights()
}

func AddNewRight(right model.Right) error {
	return repository.AddNewRight(right)
}

func EditRight(right model.Right) error {
	return repository.EditRight(right)
}

func DeleteRight(id int) error {
	return repository.DeleteRight(id)
}
