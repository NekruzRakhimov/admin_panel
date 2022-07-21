package repository

import (
	"admin_panel/db"
	"admin_panel/models"
	"fmt"
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

	fmt.Println("вот так выглядить Сегмент", segment)

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
func GetSegmentName(beneficiary string) (segments []models.Segment, err error) {
	err = db.GetDBConn().Raw("SELECT id, segment_code, name_segment FROM  segment WHERE  beneficiary = $1", beneficiary).Scan(&segments).Error
	if err != nil {
		return nil, err
	}

	return segments, nil
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
