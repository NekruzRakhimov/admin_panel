package repository

import (
	"admin_panel/db"
	"admin_panel/model"
)

func GetAllRoles() (roles []model.Role, err error) {
	if err := db.GetDBConn().Table("roles").Where("is_removed = ?", false).Order("id").Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}

func AddNewRole(role model.Role) (model.Role, error) {
	if err := db.GetDBConn().Table("roles").Save(&role).Error; err != nil {
		return model.Role{}, err
	}

	return role, nil
}

func EditRole(role model.Role) error {
	if err := db.GetDBConn().Table("roles").Omit("rights").Save(&role).Error; err != nil {
		return err
	}

	return nil
}

func DeleteRole(roleId int) error {
	if err := db.GetDBConn().Table("roles").Where("id = ?", roleId).Update("is_removed", true).Error; err != nil {
		return err
	}

	return nil
}

func AttachRightToRole(roleId, rightId int) error {
	sqlQuery := "INSERT INTO roles_rights (role_id, right_id) VALUES(?, ?)"
	if err := db.GetDBConn().Exec(sqlQuery, roleId, rightId).Error; err != nil {
		return err
	}

	return nil
}

func DetachRightFromRole(roleId, rightId int) error { // todo: затем заменить это удаление на soft_delete
	sqlQuery := "DELETE FROM roles_rights WHERE role_id = ? AND right_id = ?"
	if err := db.GetDBConn().Exec(sqlQuery, roleId, rightId).Error; err != nil {
		return err
	}

	return nil
}
