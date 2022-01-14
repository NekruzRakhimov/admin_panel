package service

import (
	"admin_panel/db"
	"admin_panel/model"
	"fmt"
)

//TODO:
// суть в логики в том, что он каждый день будет проверять, не истек ли срок договора
// если договора истек, добавляем в бд
// потом надо сделать проверку, если такой договор уже есть, то не добавлять в бд
// вопрос: как все это сделать?
// я думаю сделать select -> если договор найден, пропускай ее

func Notification() {

	//	//timeUpContracts := time.Now().Add(60 * time.Hour)
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

	var cars []model.Cars
	db.GetDBConn().Raw("SELECT cars_info -> 'brand' AS brand  FROM cars").Scan(&cars)

	fmt.Println(cars)

}
