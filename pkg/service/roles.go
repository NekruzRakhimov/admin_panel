package service

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"fmt"
)

func GetAllRolesFullInfo() (roles []models.Role, err error) {
	roles, err = repository.GetAllRoles()
	if err != nil {
		return nil, err
	}

	for i := range roles {
		rights, err := repository.GetAllRightsByRoleId(roles[i].Id)
		if err != nil {
			return nil, err
		}
		fmt.Println(rights)
		roles[i].Rights = rights
	}

	return roles, nil
}

func AddNewRole(role models.Role) error {
	role, err := repository.AddNewRole(role)
	if err != nil {
		return err
	}

	if err := repository.AddRightsToRole(role.Id, role.Rights); err != nil {
		return err
	}

	return nil
}

func EditRole(role models.Role) error {
	return repository.EditRole(role)
}

func DeleteRole(roleId int) error {
	return repository.DeleteRole(roleId)
}

func AttachRightToRole (roleId, rightId int) error {
	return repository.AttachRightToRole(roleId, rightId)
}

func DetachRightFromRole (roleId, rightId int) error {
	return repository.DetachRightFromRole(roleId, rightId)
}