package service

import (
	"admin_panel/db"
	"admin_panel/model"
	"fmt"
	"log"
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
	scan := db.GetDBConn().Raw("SELECT requisites -> 'bin' AS bin, contract_parameters -> 'contract_date' AS contract_date, contract_parameters -> 'contract_number'  AS   contract_number, type, supplier_company_manager -> 'email'  AS email FROM contracts WHERE status = 'в работе'").Scan(&notifications)
	//log.Println(" Массив Данных которые получили с уведомлений", notifications)

	for _, value := range notifications {
		if value.Type == "supply" {
			value.Type = "Договор поставок"
		} else if value.Type == "marketing_services" {
			value.Type = "Договор маркетинговых услуг"
		}
		layout := "2006-01-02T15:04:05.000Z"
		//str := "2014-11-12T11:45:26.371Z"
		t, err := time.Parse(layout, value.ContractDate)
		if err != nil {
			fmt.Println(err)
		}
		res := endDateContract.After(t)
		log.Println("Проверка времени", res)
		if endDateContract.After(t) {

			if !db.GetDBConn().Raw("SELECT requisites -> 'bin' AS bin, contract_parameters -> 'contract_date' AS contract_date, contract_parameters -> 'contract_number'  AS   contract_number, type, supplier_company_manager -> 'email'  AS email FROM contracts WHERE status = 'в работе'").RecordNotFound() {
				log.Println("ДОговора не найдены")
			}

			//TODO: сделаем выборку дог номеров, если они не сущ, потом только добавить в бд
			if !db.GetDBConn().Raw("SELECT id FROM notification where contract_number = $", value.ContractNumber).RecordNotFound() {
				log.Println("ТАКС")
			}
			//log.Println(resultNotification.RecordNotFound(), "Проверка номера договора, если оно не найдено")
			log.Println("scan.RecordNotFound()", scan.RecordNotFound())
			//resultNotification.RecordNotFound(
			//if resultNotification.RecordNotFound() == true {
			//	db.GetDBConn().Exec("INSERT into notification (bin, contract_date, contract_number, type, email) VALUES ($1, $2, $3, $4, $5)",
			//		value.Bin, value.ContractDate, value.ContractNumber, value.Type, value.Email).Scan(&notification)
			//
			//}

			// наверное вот это не сработало

			//if scan.RecordNotFound() == false {
			// если запиши нет, то в этом случае добавлеяем данные в бд

			//TODO: после чего отправляем уведомлние
			// также тест, то что договор истекает и потом данные

			log.Println(notification.ContractNumber, "Номер контаркта")

			log.Println("Данные которые получили с уведомлений", notification)
			//}
			//TODO: но если все таки запись найдена, то можем обновить или ничего не делать
		}
	}

}
