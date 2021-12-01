package service

import (
	"admin_panel/model"
	"admin_panel/pkg/repository"
)

func GetAllCurrency() (rights []model.Currency, err error) {
	return repository.GetAllCurrency()
}

func CreateMarketingContract(contract model.MarketingServicesContract)  error {

	return repository.CreateMarketingContract(contract)

}


//func AddNewRight(right model.Right) error {
//	return repository.AddNewRight(right)
//}