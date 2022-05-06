package service

import (
	"admin_panel/db"
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"strings"
)

func SaveDataFrom1C(block models.Block) error {
	return repository.SaveDataFrom1C(block)
}

func SaveDoubtedDiscounts(request models.RBRequest) error {
	for _, discount := range request.DoubtedDiscounts {
		if err := repository.SaveDoubtedDiscounts(request.BIN, request.PeriodFrom, request.PeriodTo, discount.ContractNumber, discount.Discounts); err != nil {
			return err
		}
	}

	return nil
}

func GetAllRBByContractorBIN(request models.RBRequest) (rbDTOs []models.RbDTO, err error) {
	contractsWithJson, err := repository.GetAllContractDetailByBIN(request.BIN, request.PeriodFrom, request.PeriodTo)
	if err != nil {
		return
	}

	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return
	}

	for i, contract := range contracts {
		if contract.AdditionalAgreementNumber != 0 {
			var contractType string
			//ДС №1 к Договору маркетинговых услуг №1111 ИП  “Adal Trade“
			//marketing_services
			//supply
			switch contract.Type {
			case "marketing_services":
				contractType = "маркетинговых услуг"
			case "supply":
				contractType = "поставок"
			}

			contracts[i].ContractParameters.ContractNumber = fmt.Sprintf("ДС №%d к Договору %s №%s %s",
				contract.AdditionalAgreementNumber, contractType,
				contract.ContractParameters.ContractNumber,
				contract.Requisites.Beneficiary)
		}
	}

	// #1
	RB1stType, err := GetRB1stType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, RB1stType...)

	// #2
	rb2ndType := GetRB2ndType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb2ndType...)

	// #3
	rb3rdType, err := GetRB3rdType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb3rdType...)

	// #4
	rbFourthType, err := GetRB4thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rbFourthType...)

	// #5
	rb5thType, err := GetRB5thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb5thType...)

	// #6
	rb6thType, err := GetRB6thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb6thType...)

	// #7
	rb7thType, err := GetRB7thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb7thType...)

	// #8
	rb8thType, err := GetRB8thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb8thType...)

	// #9
	rb9thType, err := GetRB9thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb9thType...)

	// #10
	rbTenthType, err := GetRb10thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rbTenthType...)

	// #11
	rb11thType, err := GetRB11thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb11thType...)

	// #12
	rb12thType, err := GetRB12thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb12thType...)

	// #13
	rb13thType, err := GetRB13thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb13thType...)

	// #14
	rb14ThType, err := GetRB14ThType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb14ThType...)

	// #15
	rb15ThType, err := GetRB15ThType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb15ThType...)

	// #16
	rb16ThType, err := GetRB16ThType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb16ThType...)

	// #17
	rb17ThType, err := GetRB17ThType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb17ThType...)

	return
}

func GetTotalSalesForSku(sales models.Sales, sku string) (totalSum float32) {
	for _, s := range sales.SalesArr {
		if s.ProductCode == sku {
			totalSum += s.Total * s.QntTotal
		}
	}

	return totalSum
}

func GetAllProductsSku(contracts []models.Contract) (SkuArr []string) {
	for _, contract := range contracts {
		for _, product := range contract.Products {
			SkuArr = append(SkuArr, product.Sku)
		}
	}

	return SkuArr
}

func GetTotalFromSalesByBrand(sales models.Sales, brand string) (totalAmount float32) {
	for _, s := range sales.SalesArr {
		if s.BrandName == brand {
			totalAmount += s.QntTotal * s.Total
		}
	}

	return totalAmount
}

func DefiningRBReport(contracts []models.Contract, totalAmount float64, request models.RBRequest) (contractsRB []models.RbDTO) {
	for _, contract := range contracts {
		var contractRB []models.RbDTO
		for _, discount := range contract.Discounts {
			if discount.Code == "TOTAL_AMOUNT_OF_SELLING" && discount.IsSelected {
				log.Printf("\n[CONTRACT_DISCOUNT][%s] %+v\n", contract.ContractParameters.ContractNumber, contract.Discounts)
				contractRB = DiscountToReportRB(discount, contract, totalAmount, request)
			}
		}
		contractsRB = append(contractsRB, contractRB...)
	}

	return contractsRB
}

func DiscountToReportRB(discount models.Discount, contract models.Contract, totalAmount float64, request models.RBRequest) (contractsRB []models.RbDTO) {
	fmt.Println("<begin>")
	var contractRB models.RbDTO

	if len(discount.Periods) > 0 {
		contractRB = models.RbDTO{
			ID:             contract.ID,
			ContractNumber: contract.ContractParameters.ContractNumber,
			DiscountType:   RB1Name,
			StartDate:      discount.Periods[0].PeriodFrom,
			EndDate:        discount.Periods[0].PeriodTo,
		}

		var (
			maxDiscountAmount float32 // сумма закупа
			maxRewardAmount   int     // Сумма вознаграждения
			maxLeasePlan      float32 // план закупа
			isCompleted       bool
		)

		for _, period := range discount.Periods {
			if period.PeriodFrom >= request.PeriodFrom && period.PeriodTo <= request.PeriodTo {
				if float64(period.TotalAmount) <= totalAmount {
					if period.TotalAmount >= maxLeasePlan {
						log.Printf("\n[CONTRACT_PERIODS][%s] %+v\n", contract.ContractParameters.ContractNumber, discount.Periods)
						maxDiscountAmount = float32(period.RewardAmount)
						maxRewardAmount = period.RewardAmount
						maxLeasePlan = period.TotalAmount
						isCompleted = true
					}
				} /*else {
					maxRewardAmount = period.RewardAmount
					maxLeasePlan = period.TotalAmount
				}*/
			}
		}

		if !isCompleted && len(discount.Periods) > 0 {
			maxRewardAmount = discount.Periods[0].RewardAmount
			maxLeasePlan = discount.Periods[0].TotalAmount
		}

		// Сумма скидки	| Сумма вознаграждения	| План закупа

		contractRB.RewardAmount = float32(maxRewardAmount)
		contractRB.LeasePlan = maxLeasePlan
		contractRB.DiscountAmount = maxDiscountAmount

		//if len(discount.Periods) > 1 && totalAmount >= discount.Periods[1].TotalAmount && discount.Periods[1].RewardAmount > discount.Periods[0].RewardAmount {
		//	fmt.Printf("worked [totalAmount = %d AND discount.Periods[0].TotalAmount = %d]\n", totalAmount, discount.Periods[0].TotalAmount)
		//	contractRB.DiscountAmount = float32(discount.Periods[0].RewardAmount)
		//} else if totalAmount >= discount.Periods[0].TotalAmount {
		//	fmt.Printf("worked [totalAmount = %d AND discount.Periods[0].TotalAmount = %d]\n", totalAmount, discount.Periods[0].TotalAmount)
		//	contractRB.DiscountAmount = float32(discount.Periods[0].RewardAmount)
		//}
	}
	contractsRB = append(contractsRB, contractRB)

	fmt.Println("<end>")
	return contractsRB
}

func GetTotalAmount(sales models.Sales) float32 {
	var amount float32
	for _, s := range sales.SalesArr {
		amount += s.Total
	}

	return amount
}

func GetTotalAmountPurchase(purchase models.Purchase) float32 {
	var amount float32
	for _, s := range purchase.PurchaseArr {
		amount += float32(s.Total)
	}

	return amount
}

func GetTotalAmountFrom1CDataSalesOrPurchases(data []models.GetData1CProducts) float32 {
	var amount float32
	for _, s := range data {
		amount += s.Total * s.QntTotal
	}

	return amount
}

func TrimDate(fullDate string) string {
	arr := strings.Split(fullDate, " ")
	if len(arr) > 0 {
		return arr[0]
	}
	return ""
}

func BulkConvertContractFromJsonB(contractsWithJson []models.ContractWithJsonB) (contracts []models.Contract, err error) {
	for i := range contractsWithJson {
		contract, err := ConvertContractFromJsonB(contractsWithJson[i])
		if err != nil {
			log.Println("Error: service.BulkConvertContractFromJsonB. Error is: ", err.Error())
			continue
		}
		contracts = append(contracts, contract)
	}

	return
}

func GetDoubtedDiscounts(request models.RBRequest) (doubtedDiscounts []models.DoubtedDiscount, err error) {
	contractsWithJson, err := repository.GetAllContractDetailByBIN(request.BIN, request.PeriodFrom, request.PeriodTo)
	if err != nil {
		return nil, err
	}

	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		return nil, err
	}

	var (
	//hasPresentation bool
	//hasMTZ bool
	)

	for _, contract := range contracts {
		var DoubtedDiscountDetails []models.DoubtedDiscountDetails
		for _, discount := range contract.Discounts {
			var DoubtedDiscountDetail models.DoubtedDiscountDetails
			if (discount.Code == RB4Code || discount.Code == RB11Code) && discount.IsSelected == true {
				DoubtedDiscountDetail.Name = discount.Name
				DoubtedDiscountDetail.Code = discount.Code
				DoubtedDiscountDetail.IsCompleted = repository.DoubtedDiscountExecutionCheck(request, contract.ContractParameters.ContractNumber, discount.Code)
				DoubtedDiscountDetails = append(DoubtedDiscountDetails, DoubtedDiscountDetail)
			}
		}
		if len(DoubtedDiscountDetails) > 0 {
			doubtedDiscounts = append(doubtedDiscounts, models.DoubtedDiscount{
				ContractNumber: contract.ContractParameters.ContractNumber,
				Discounts:      DoubtedDiscountDetails,
			})
		}
	}

	return doubtedDiscounts, nil
}

func GetTotalPurchasesForBrands(sales models.Purchase, brand string) (totalAmount float32) {
	for _, s := range sales.PurchaseArr {
		if s.BrandCode == brand || s.BrandName == brand {
			totalAmount += float32(s.Total * s.QntTotal)
		}
	}

	return totalAmount
}

func GeAllBrands(brandsDTO []models.BrandDTO) (brands []string) {
	for _, brand := range brandsDTO {
		brands = append(brands, brand.BrandName)
	}

	return brands
}

//TODO: необходимо реализовать
func SearchHistoryDiscount(field string, param string) ([]models.SearchContract, error) {
	var search []models.SearchContract
	if field == "author" {
		query := fmt.Sprintf("SELECT id, manager AS author, status," +
			"created_at, contract_parameters ->> 'end_date' AS end_date, comment FROM  contracts " +
			"WHERE  manager  like  $1")
		err := db.GetDBConn().Raw(query, "%"+param+"%").Scan(&search).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return search, err
		}
		return search, nil
	}
	//это чтобы понять из какого объекта будем доставать поля из JSONB
	var jsonBTable string
	if field == "contract_number" {
		jsonBTable = "contract_parameters"
	} else if field == "beneficiary" {
		jsonBTable = "requisites"
	}
	query := fmt.Sprintf("SELECT id, manager AS author, status,"+
		"created_at, contract_parameters ->> 'end_date' AS end_date, comment FROM  contracts"+
		"WHERE  %s ->> $1 like  $2", jsonBTable)

	err := db.GetDBConn().Raw(query, field, "%"+param+"%").Scan(&search).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return search, err
	}

	return search, nil

}
