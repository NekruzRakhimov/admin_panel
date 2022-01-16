package service

import (
	"admin_panel/db"
	"admin_panel/model"
)

//TODO:
// суть в логики в том, что он каждый день будет проверять, не истек ли срок договора
// если договора истек, добавляем в бд
// потом надо сделать проверку, если такой договор уже есть, то не добавлять в бд
// вопрос: как все это сделать?
// я думаю сделать select -> если договор найден, пропускай ее

func Notification() {


	//	fmt.Println("ВЫЗОВ")
	//	//var contracts []model.Contract
	//	var data string
	//	//db.GetDBConn().Model(&contracts).Find()
	//	//	db.GetDBConn().Model(&contracts).Find("contracts")
	//	//var bDate []byte
	/////	db.GetDBConn().Raw("SELECT contract_parameters -> 'contract_date' AS data FROM contracts").Scan(&data)
	//	//for _, value := range contracts {
	//	//	//TODO: тут будет сравнение
	//	//
	//	//}
	//	//fmt.Println(string(bDate), "бинарные данные")
	//
	//	//json.Unmarshal(bDate, &date)
	//	fmt.Println(data, "ДАТА")
	var notification []model.Notification
	//db.GetDBConn().Raw("SELECT cars_info -> 'brand' AS brand  FROM cars").Scan(&cars)
	scan := db.GetDBConn().Raw("SELECT requisites -> 'bin' AS bin, contract_parameters -> 'contract_date' AS end_date, contract_parameters -> contract_number  AS   contract_number, type, supplier_company_manager -> email  AS email FROM contacts").Scan(&notification)
	if scan.RecordNotFound() == false{
		// добавить в бд
	}
	// если запись найдена то обновляем
	//timeUpContracts := time.Now().Add(1440 * time.Hour)


	//TODO:  запихнуть эти даннные в другую таблицу
	//fmt.Println(cars)

}
