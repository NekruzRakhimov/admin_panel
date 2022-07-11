package service

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"log"
	"math"
	"strconv"
)

func CreateNecessity() {
	log.Println("|#######################################|", "##########|", "###############################################|", "################|", "###########|")
	log.Println("|Наименование аптеки                    |", "код аптеки|", "Наименование номенклатуры                      |", "Код номенклатуры|", "Потребность|")
	log.Println("|#######################################|", "##########|", "###############################################|", "################|", "###########|")
	log.Println("|Аптека № 2, Шымкент, (Городской Акимат)|", "A0000120  |", "Kotex тампоны Active Normal № 8 шт             |", "00000064187     |", "100        |")
	log.Println("|---------------------------------------|", "----------|", "-----------------------------------------------|", "----------------|", "-----------|")
	log.Println("|Аптека № 2, Шымкент, (Городской Акимат)|", "A0000120  |", "Kotex тампоны Active Normal 24* № 16 шт        |", "00000065247     |", "120        |")
	log.Println("|---------------------------------------|", "----------|", "-----------------------------------------------|", "----------------|", "-----------|")
	log.Println("|Аптека № 2, Шымкент, (Городской Акимат)|", "A0000120  |", "Kotex тампоны Active Super № 16 шт             |", "00000065248     |", "175        |")
	log.Println("|---------------------------------------|", "----------|", "-----------------------------------------------|", "----------------|", "-----------|")
	log.Println("|Аптека № 2, Шымкент, (Городской Акимат)|", "A0000120  |", "Huggies трусики Elite Soft 3 (6-11 кг) № 25 шт |", "00000067262     |", "155        |")
	log.Println("|---------------------------------------|", "----------|", "-----------------------------------------------|", "----------------|", "-----------|")
	log.Println("|Аптека № 2, Шымкент, (Городской Акимат)|", "A0000120  |", "Huggies трусики Elite Soft 4 (9-14кг) № 21 шт  |", "00000067263     |", "190        |")
	log.Println("|---------------------------------------|", "----------|", "-----------------------------------------------|", "----------------|", "-----------|")
	log.Println("|Аптека № 2, Шымкент, (Городской Акимат)|", "A0000120  |", "Huggies трусики Elite Soft 4 (9-14кг) №42      |", "00000067264     |", "130        |")
	log.Println("|---------------------------------------|", "----------|", "-----------------------------------------------|", "----------------|", "-----------|")
	log.Println("|Аптека № 2, Шымкент, (Городской Акимат)|", "A0000120  |", "Huggies трусики Elite Soft 5 (12-17 кг) № 19 шт|", "00000067265     |", "140        |")
	log.Println("|---------------------------------------|", "----------|", "-----------------------------------------------|", "----------------|", "-----------|")
	log.Println("|Аптека № 2, Шымкент, (Городской Акимат)|", "A0000120  |", "Huggies трусики Elite Soft 5 (12-17 кг) №38    |", "00000067266     |", "200        |")
	log.Println("|---------------------------------------|", "----------|", "-----------------------------------------------|", "----------------|", "-----------|")

}

func GetAllAutoOrders() (autoOrders []models.AutoOrder, err error) {
	return repository.GetAllAutoOrders()
}

func FormAutoOrders() error {
	graphics, err := repository.GetAllGraphics()
	if err != nil {
		return err
	}

	var formedGraphics []models.FormedGraphic

	for _, graphic := range graphics {
		var formedGraphic models.FormedGraphic
		formedGraphic.GraphicID = graphic.ID
		formedGraphic.ByMatrix = true
		formedGraphic.ProductAvailabilityDays = 0
		formedGraphic.DisterDays = 0
		formedGraphic.StoreDays = 0
		formedGraphic.Status = "сформирован"

		storeCode := graphic.StoreCode
		supplierCode := graphic.SupplierCode

		matrix, err := GetMatrixExt(storeCode)
		if err != nil {
			return err
		}

		req := models.SalesCountRequest{
			Startdate: "01.01.2022 00:00:00",
			Enddate:   "27.05.2022 00:00:00",
			StoreCode: storeCode,
		}
		sales, err := GetSalesCountExt(req)
		if err != nil {
			return err
		}

		for _, product := range matrix {
			if product.SupplierCode == supplierCode && product.StoreCode == storeCode {
				min, _ := strconv.ParseFloat(product.Min, 2)
				max, _ := strconv.ParseFloat(product.Max, 2)
				if min != 0 {
					formedGraphic.FormulaID = 1
					for _, sale := range sales {
						if sale.ProductCode == product.ProductCode {
							salesCount, _ := strconv.ParseFloat(sale.SalesCount, 2)
							salesDayCount, _ := strconv.ParseFloat(sale.SalesDayCount, 2)
							totalStoreCount, _ := strconv.ParseFloat(sale.TotalStoreCount, 2)
							totalSalesDayCount, _ := strconv.ParseFloat(sale.TotalSalesDayCount, 2)

							koef := 15
							if graphic.OnceAMonth {
								koef = 45
							} else if graphic.TwiceAMonth {
								koef = 30
							}

							orderQnt := salesCount/salesDayCount*float64(koef) + min - totalStoreCount
							orderQnt = math.Ceil(orderQnt)

							formedGraphic.Products = append(formedGraphic.Products, models.FormedGraphicProduct{
								ProductCode:             product.ProductCode,
								ProductName:             product.ProductName,
								OrderQnt:                orderQnt,
								Days:                    int(salesDayCount),
								Remainder:               totalStoreCount,
								ProductAvailabilityDays: int(totalSalesDayCount),
								SalesCount:              salesCount,
								SalesDayCount:           salesDayCount,
								Koef:                    koef,
								TotalStoreCount:         totalStoreCount,
								Min:                     min,
								StoreCode:               storeCode,
							})
						}
					}
				} else if max != 0 {
					formedGraphic.FormulaID = 2
					for _, sale := range sales {
						if sale.ProductCode == product.ProductCode {
							//salesCount, _ := strconv.ParseFloat(sale.SalesCount, 2)
							salesDayCount, _ := strconv.ParseFloat(sale.SalesDayCount, 2)
							totalStoreCount, _ := strconv.ParseFloat(sale.TotalStoreCount, 2)
							totalSalesDayCount, _ := strconv.ParseFloat(sale.TotalSalesDayCount, 2)

							orderQnt := max - totalStoreCount
							orderQnt = math.Ceil(orderQnt)

							formedGraphic.Products = append(formedGraphic.Products, models.FormedGraphicProduct{
								ProductCode:             product.ProductCode,
								ProductName:             product.ProductName,
								OrderQnt:                orderQnt,
								Days:                    int(salesDayCount),
								Remainder:               totalStoreCount,
								ProductAvailabilityDays: int(totalSalesDayCount),
								Max:                     max,
								TotalStoreCount:         totalStoreCount,
							})
						}
					}
				}
			}
		}
		formedGraphics = append(formedGraphics, formedGraphic)
	}

	return repository.SaveFormedGraphics(formedGraphics)
}

func GetAllFormedGraphics() (graphics []models.FormedGraphic, err error) {
	return repository.GetAllFormedGraphics()
}

func GetFormedGraphicByID(id int) (graphic models.FormedGraphic, err error) {
	return repository.GetFormedGraphicByID(id)
}

func GetAllFormedGraphicsProducts(formedGraphicID int) (products []models.FormedGraphicProduct, err error) {
	return repository.GetAllFormedGraphicsProducts(formedGraphicID)
}

func CancelFormedFormula(formulaID int, comment string) error {
	return repository.CancelFormedFormula(formulaID, comment)
}

func CancelFormedGraphic(graphicID int, comment string) error {
	return repository.CancelFormedGraphic(graphicID, comment)
}
