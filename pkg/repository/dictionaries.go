package repository

import (
	"admin_panel/db"
	"admin_panel/model"
)

func GetAllCurrencies() (currencies []model.Currency, err error) {
	sqlQuery := "SELECT * FROM currencies"
	if err := db.GetDBConn().Raw(sqlQuery).Scan(&currencies).Error; err != nil {
		return nil, err
	}

	return currencies, nil
}

func GetAllPositions() (positions []model.Position, err error) {
	sqlQuery := "SELECT * FROM positions"
	if err := db.GetDBConn().Raw(sqlQuery).Scan(&positions).Error; err != nil {
		return nil, err
	}

	return positions, nil
}

func GetAllAddresses() (addresses []model.Address, err error) {
	sqlQuery := "SELECT * FROM addresses"
	if err := db.GetDBConn().Raw(sqlQuery).Scan(&addresses).Error; err != nil {
		return nil, err
	}

	return addresses, nil
}

func GetAllFrequencyDeferredDiscounts() (frequencyDeferredDiscounts []model.FrequencyDeferredDiscount, err error) {
	sqlQuery := "SELECT * FROM frequency_deferred_discount"
	if err := db.GetDBConn().Raw(sqlQuery).Scan(&frequencyDeferredDiscounts).Error; err != nil {
		return nil, err
	}

	return frequencyDeferredDiscounts, nil
}

func GetAllDictionaries() (dictionaries []model.Dictionary, err error) {
	sqlQuery := "SELECT * FROM dictionaries WHERE is_removed = false ORDER BY id"
	err = db.GetDBConn().Raw(sqlQuery).Scan(&dictionaries).Error
	if err != nil {
		return nil, err
	}

	return dictionaries, nil
}

func CreateDictionary(dictionary model.Dictionary) error {
	if err := db.GetDBConn().Table("dictionaries").Omit("author").Create(&dictionary).Error; err != nil {
		return err
	}

	return nil
}

func EditDictionary(dictionary model.Dictionary) error {
	if err := db.GetDBConn().Table("dictionaries").Omit("author").Save(&dictionary).Error; err != nil {
		return err
	}

	return nil
}

func DeleteDictionary(dictionaryID int) error {
	sqlQuery := "UPDATE dictionaries SET is_removed = true, deleted_at = now() WHERE id = ?"
	if err := db.GetDBConn().Exec(sqlQuery, dictionaryID).Error; err != nil {
		return err
	}

	return nil
}

func GetAllDictionaryValues(dictionaryID int) (dictionaryValues []model.DictionaryValue, err error) {
	sqlQuery := "SELECT * FROM dictionary_values WHERE dictionary_id = ? ORDER BY id"
	err = db.GetDBConn().Raw(sqlQuery, dictionaryID).Scan(&dictionaryValues).Error
	if err != nil {
		return nil, err
	}

	return dictionaryValues, nil
}

func CreateDictionaryValue(dictionaryValue model.DictionaryValue) error {
	if err := db.GetDBConn().Table("dictionary_values").Create(&dictionaryValue).Error; err != nil {
		return err
	}

	return nil
}

func EditDictionaryValue(dictionaryValue model.DictionaryValue) error {
	if err := db.GetDBConn().Table("dictionary_values").Save(&dictionaryValue).Error; err != nil {
		return err
	}

	return nil
}

func DeleteDictionaryValue(dictionaryID, dictionaryValueID int) error {
	sqlQuery := "DELETE FROM dictionary_values WHERE id = ? AND dictionary_id = ?"
	if err := db.GetDBConn().Exec(sqlQuery, dictionaryValueID, dictionaryID).Error; err != nil {
		return err
	}

	return nil
}
