package service

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
)

func GetAllCurrencies() ([]models.Currency, error) {
	return repository.GetAllCurrencies()
}

func GetAllPositions() ([]models.Position, error) {
	return repository.GetAllPositions()
}

func GetAllAddresses() ([]models.Address, error) {
	return repository.GetAllAddresses()
}

func GetAllFrequencyDeferredDiscounts() ([]models.FrequencyDeferredDiscount, error) {
	return repository.GetAllFrequencyDeferredDiscounts()
}

func GetAllDictionaries() (dictionaries []models.Dictionary, err error) {
	return repository.GetAllDictionaries()
}

func GetDictionaryByID(dictionaryID int) (models.Dictionary, error) {
	return repository.GetDictionaryByID(dictionaryID)
}

func CreateDictionary(dictionary models.Dictionary) error {
	return repository.CreateDictionary(dictionary)
}

func EditDictionary(dictionary models.Dictionary) error {
	return repository.EditDictionary(dictionary)
}

func DeleteDictionary(dictionaryID int) error {
	return repository.DeleteDictionary(dictionaryID)
}

func GetAllDictionaryValues(dictionaryID int) (dictionaryValues []models.DictionaryValue, err error) {
	return repository.GetAllDictionaryValues(dictionaryID)
}

func CreateDictionaryValue(dictionaryValue models.DictionaryValue) error {
	return repository.CreateDictionaryValue(dictionaryValue)
}

func EditDictionaryValue(dictionaryValue models.DictionaryValue) error {
	return repository.EditDictionaryValue(dictionaryValue)
}

func DeleteDictionaryValue(dictionaryID, dictionaryValueID int) error {
	return repository.DeleteDictionaryValue(dictionaryID, dictionaryValueID)
}

func GetSegments() []models.Segment {
	nomen1 := []models.ListsNomenclature{
		{
			ProductCode: "Аспирин С № 10 табл шипуч",
			ProductName: "00000001732",
		},
		{
			ProductCode: "Спринцовка-аспиратор САИ Б 1-3 со стекл. наконечн.",
			ProductName: "00000001639",
		},
		{
			ProductCode: "Фаспик 400 мг, №6, табл., покрытые оболочкой",
			ProductName: "00000002553",
		},
	}
	nomen2 := []models.ListsNomenclature{
		{
			ProductCode: "Head&Shoulders шампунь успокаивающий 400 мл ",
			ProductName: "00000008367",
		},
		{
			ProductCode: "Head&Shoulders шампунь успокаивающий 200 мл ",
			ProductName: "00000008835",
		},
		{
			ProductCode: "Спайдер мен мыло жидкое Web-Head 480 мл",
			ProductName: "00000017162",
		},
	}
	seg := []models.Segment{
		{
			SegmentCode:       "0000098",
			NameSegment:       "1",
			ListsNomenclature: nomen1,
			Counterparty:      "Юнифарма ООО",
			ForMarket:         true,
		},
		{
			SegmentCode:       "0000098",
			NameSegment:       "2",
			ListsNomenclature: nomen2,
			Counterparty:      "Farma Marketing (Фарма Маркетинг) ТОО",
			ForMarket:         true,
		},
	}

	return seg

}
