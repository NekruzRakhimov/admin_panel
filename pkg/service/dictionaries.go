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
