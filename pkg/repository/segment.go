package repository

import (
	"admin_panel/db"
	"admin_panel/models"
)

func CreateSegment(segment models.Segment) error  {

	err := db.GetDBConn().Create(&segment).Error
	if err != nil {
		return err
	}


	return nil

}
