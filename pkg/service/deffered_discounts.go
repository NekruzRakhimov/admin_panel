package service

import (
	"admin_panel/model"
	"admin_panel/pkg/repository"
)

const (
	DD1Name = "Отложенная скидка за своевременную оплату"
	DD2Name = "Отложенная скидка за своевременную оплату в виде Кредит Ноты"
	DD3Name = "Отложенная скидка за своевременную оплату в виде Оплаты на счет"
	DD4Name = "Отложенная скидка за своевременную оплату в виде Товара"
	DD5Name = "Отложенная скидка за выполнение квартального плана в виде кредит ноты"
	DD6Name = "Отложенная скидка за выполнение годового  плана в виде кредит ноты"
)
const (
	DD1Code = "deferred_discount_for_timely_payment"
	DD2Code = "deferred_discount_for_timely_payment_in_credit_note_form"
	DD3Code = "deferred_discount_for_timely_payment_in_pay_to_account_form"
	DD4Code = "deferred_discount_for_timely_payment_in_goods_form"
	DD5Code = "deferred_discount_for_implementation_of_quarterly_plan_in_credit_note_form"
	DD6Code = "deferred_discount_for_implementation_of_annual_plan_in_credit_note_form"
)

func GetAllDeferredDiscounts(request model.RBRequest) (RbDTO []model.RbDTO, err error) {
	contractsWithJson, err := repository.GetAllContractDetailByBIN(request.BIN, request.PeriodFrom, request.PeriodTo)
	if err != nil {
		return nil, err
	}

	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return nil, err
	}

	totalAmount := float32(5000_000.0)

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if (discount.Code == DD1Code || discount.Code == DD2Code || discount.Code == DD3Code ||
				discount.Code == DD4Code || discount.Code == DD5Code || discount.Code == DD6Code) && discount.IsSelected == true {
				RbDTO = append(RbDTO, model.RbDTO{
					ContractNumber:  contract.ContractParameters.ContractNumber,
					StartDate:       contract.ContractParameters.StartDate,
					EndDate:         contract.ContractParameters.EndDate,
					DiscountPercent: discount.DiscountPercent,
					DiscountAmount:  totalAmount * discount.DiscountPercent / 100,
					DiscountType:    discount.Name,
				})
			}
		}
	}

	return RbDTO, nil
}
