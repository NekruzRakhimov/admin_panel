package repository

import (
	"admin_panel/db"
	"admin_panel/models"
)

func CreateSegment(segment models.Segment) error {

	err := db.GetDBConn().Create(&segment).Error
	if err != nil {
		return err
	}

	return nil

}

func GetSegmentByID(id int) (models.Segment, error) {
	var segment models.Segment
	err := db.GetDBConn().Raw("SELECT *FROM segment WHERE id = $1", id).Scan(&segment).Error
	if err != nil {
		return segment, err
	}

	return segment, nil

}
