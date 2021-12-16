package repository

import (
	"admin_panel/db"
	"admin_panel/model"
	"errors"
	"fmt"
)

func GetAllRights() (rights []model.Right, err error) {
	if err := db.GetDBConn().Table("rights").Where("is_removed = ?", false).Order("id").Scan(&rights).Error; err != nil {
		return nil, err
	}

	return rights, nil
}

func AddNewRight(right model.Right) error {
	if err := db.GetDBConn().Table("rights").Create(&right).Error; err != nil {
		return errors.New("право с таким кодом уже существоует")
	}

	return nil
}

func EditRight(right model.Right) error {
	if err := db.GetDBConn().Table("rights").Save(&right).Error; err != nil {
		return err
	}

	return nil
}

func DeleteRight(id int) error {
	if err := db.GetDBConn().Table("rights").Where("id = ?", id).Update("is_removed", true).Error; err != nil {
		return err
	}

	return nil
}

func GetAllRightsByRoleId(roleId int) (rights []model.RightDTO, err error) {
	sqlQuery := `SELECT rights.*,
				   case
					   when rights.id in (SELECT right_id FROM roles_rights WHERE role_id = ?)
						   then true
					   else false end is_attached
					FROM rights
					WHERE rights.is_removed = false`
	if err := db.GetDBConn().Raw(sqlQuery, roleId).Scan(&rights).Error; err != nil {
		return nil, err
	}

	fmt.Println(rights)
	return rights, nil
}

func AddRightsToRole(roleId int, rights []model.RightDTO) error {
	sqlQuery := "INSERT INTO roles_rights (role_id, right_id) VALUES(?, ?)"

	for _, right := range rights {
		if err := db.GetDBConn().Exec(sqlQuery, roleId, right.ID).Error; err != nil {
			return err
		}
	}

	return nil
}

func GetRightByID(rightId int) (right model.Right, err error) {
	right.ID = rightId
	if err := db.GetDBConn().Table("rights").Find(&right).Error; err != nil {
		return model.Right{}, err
	}

	return right, nil
}
