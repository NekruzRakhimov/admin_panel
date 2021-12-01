package service

import (
	"admin_panel/db"
	"admin_panel/model"
	"admin_panel/pkg/repository"
)

func GetAllCurrency() (rights []model.Currency, err error) {
	return repository.GetAllCurrency()
}

func CreateMarketingContract(contract model.MarketingServicesContract)  {
	var input model.MarketingServicesContract
	db.GetDBConn().Model(&input).Create(contract)

}
