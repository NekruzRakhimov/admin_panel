package service

import (
	"admin_panel/db"
	"admin_panel/model"
	"fmt"
	"time"
)

//TODO:
// суть в логики в том, что он каждый день будет проверять, не истек ли срок договора
// если договора истек, добавляем в бд
// потом надо сделать проверку, если такой договор уже есть, то не добавлять в бд
// вопрос: как все это сделать?
// я думаю сделать select -> если договор найден, пропускай ее

func Notification() {

	endDateContract := time.Now().Add((24 * 60) * time.Hour)

	var notifications []model.Notification
	var notification model.Notification

	//db.GetDBConn().Raw("SELECT cars_info -> 'brand' AS brand  FROM cars").Scan(&cars)
	scan := db.GetDBConn().Raw("SELECT requisites -> 'bin' AS bin, contract_parameters -> 'contract_date' AS end_date, contract_parameters -> contract_number  AS   contract_number, type, supplier_company_manager -> email  AS email FROM contacts").Scan(&notifications)
	for _, value := range notifications {
		layout := "2006-01-02T15:04:05.000Z"
		//str := "2014-11-12T11:45:26.371Z"
		t, err := time.Parse(layout, value.ContractDate)
		if err != nil {
			fmt.Println(err)
		}
		if endDateContract.After(t) {
			if scan.RecordNotFound() == false {
				// если запиши нет, то в этом случае добавлеяем данные в бд
				db.GetDBConn().Raw("INSERT into notifications (bin, contract_date, contract_number, type, email) VALUES ($1, $2, $3, $4, $5)",
					value.Bin, value.ContractDate, value.ContractNumber, value.Type, value.Email).Scan(&notification)
			}
			//TODO: но если все таки запись найдена, то можем обновить или ничего не делать
		}
	}



}
