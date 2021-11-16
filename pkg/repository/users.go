package repository

import (
	"admin_panel/db"
	"admin_panel/models"
	"fmt"
)

func GetAllUsers() (users []models.User, err error) {
	if err := db.GetDBConn().Table("users").Where("is_removed = ?", false).Order("id").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func GetAllRolesByUserId(userId int) (roles []models.RoleDTO, err error) {
	sqlQuery := `SELECT roles.*,
				   case
					   when roles.id in (SELECT role_id FROM users_roles WHERE user_id = ?)
						   then true
					   else false end is_attached
					FROM roles
					WHERE roles.is_removed = false`
	if err := db.GetDBConn().Raw(sqlQuery, userId).Scan(&roles).Error; err != nil {
		return nil, err
	}

	fmt.Println()
	return roles, nil
}

func CreateNewUser(user models.User) (models.User, error) {
	if err := db.GetDBConn().Table("users").Save(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func AddRolesToUser(userId int, roles []models.RoleDTO) error {
	sqlQuery := "INSERT INTO users_roles (user_id, role_id) VALUES(?, ?)"

	for _, role := range roles {
		if err := db.GetDBConn().Exec(sqlQuery, userId, role.Id).Error; err != nil {
			return err
		}
	}

	return nil
}

func EditUser(user models.User) error {
	if err := db.GetDBConn().Table("users").Omit("roles").Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUser(roleId int) error {
	if err := db.GetDBConn().Table("users").Where("id = ?", roleId).Update("is_removed", true).Error; err != nil {
		return err
	}

	return nil
}

func AttachRoleToUser (userId, roleId int) error {
	sqlQuery := "INSERT INTO users_roles (user_id, role_id) VALUES(?, ?)"
	if err := db.GetDBConn().Exec(sqlQuery, userId, roleId).Error; err != nil {
		return err
	}

	return nil
}

func DetachRoleFromUser (userId, roleId int) error { // todo: затем заменить это удаление на soft_delete
	sqlQuery := "DELETE FROM users_roles WHERE user_id = ? AND role_id = ?"
	if err := db.GetDBConn().Exec(sqlQuery, userId, roleId).Error; err != nil {
		return err
	}

	return nil
}