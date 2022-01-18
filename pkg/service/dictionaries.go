package service

import (
	"admin_panel/model"
	"admin_panel/pkg/repository"
)

func GetAllCurrencies() ([]model.Currency, error) {
	return repository.GetAllCurrencies()
}

func GetAllPositions() ([]model.Position, error) {
	return repository.GetAllPositions()
}

func GetAllAddresses() ([]model.Address, error) {
	return repository.GetAllAddresses()
}

func GetAllFrequencyDeferredDiscounts() ([]model.FrequencyDeferredDiscount, error) {
	return repository.GetAllFrequencyDeferredDiscounts()
}

func GetAllDictionaries() (dictionaries []model.Dictionary, err error) {
	return repository.GetAllDictionaries()
}

func CreateDictionary(dictionary model.Dictionary) error {
	return repository.CreateDictionary(dictionary)
}

func EditDictionary(dictionary model.Dictionary) error {
	return repository.EditDictionary(dictionary)
}

func DeleteDictionary(dictionaryID int) error {
	return repository.DeleteDictionary(dictionaryID)
}

func GetAllDictionaryValues(dictionaryID int) (dictionaryValues []model.DictionaryValue, err error) {
	return repository.GetAllDictionaryValues(dictionaryID)
}

func CreateDictionaryValue(dictionaryValue model.DictionaryValue) error {
	return repository.CreateDictionaryValue(dictionaryValue)
}

func EditDictionaryValue(dictionaryValue model.DictionaryValue) error {
	return repository.EditDictionaryValue(dictionaryValue)
}

func DeleteDictionaryValue(dictionaryID, dictionaryValueID int) error {
	return repository.DeleteDictionaryValue(dictionaryID, dictionaryValueID)
}
