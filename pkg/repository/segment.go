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

func ChangeSegment(segment models.Segment) error {

	err := db.GetDBConn().Table("segment").Save(&segment).Error
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

func GetSegment(supplier string) (models.Segment, error) {
	var segment models.Segment
	err := db.GetDBConn().Raw("SELECT *FROM segment WHERE beneficiary = $1", supplier).Scan(&segment).Error
	if err != nil {
		return segment, err
	}

	return segment, nil

}

func GetSegments() ([]models.Segment, error) {
	var segment []models.Segment
	err := db.GetDBConn().Raw("SELECT *FROM segment").Scan(&segment).Error
	if err != nil {
		return segment, err
	}

	return segment, nil

}

func DeleteSegmentByID(id int) error {
	err := db.GetDBConn().Exec("DELETE FROM segment WHERE id = $1", id).Error
	if err != nil {
		return err
	}

	return err
}

func ChangeLetter(id int) error {
	err := db.GetDBConn().Exec("UPDATE formed_graphics SET is_letter = ? WHERE id = ?", true, id).Error
	if err != nil {
		return err
	}
	return nil
}
