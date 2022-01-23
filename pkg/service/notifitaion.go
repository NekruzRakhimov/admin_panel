package service

import (
	"admin_panel/db"
	"admin_panel/model"
	"admin_panel/pkg/repository"
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"time"
)

//Notification - берет те договора у которых истек срок
func Notification() {
	endDateContract := time.Now().Add((24 * 60) * time.Hour)
	var notifications []model.Notification
	var notification model.Notification
	//db.GetDBConn().Raw("SELECT cars_info -> 'brand' AS brand  FROM cars").Scan(&cars)
	db.GetDBConn().Raw("SELECT requisites ->> 'bin' AS bin, contract_parameters ->> 'contract_date' AS contract_date, contract_parameters ->> 'contract_number'  AS   contract_number, type, supplier_company_manager ->> 'email'  AS email FROM contracts WHERE status = 'в работе'").Scan(&notifications)
	//log.Println(" Массив Данных которые получили с уведомлений", notifications)

	for _, value := range notifications {
		var data struct {
			ID int `json:"id"`
		}

		log.Println("Contract Numb", value.ContractNumber)

		if value.Type == "supply" {
			value.Type = "Договор поставок"
		} else if value.Type == "marketing_services" {
			value.Type = "Договор маркетинговых услуг"
		}

		layoutISO := "02.1.2006"

		t, err := time.Parse(layoutISO, value.ContractDate)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(t, "ПРИМЕР ДАТЫ")
		res := endDateContract.After(t)
		log.Println("Проверка времени", res)
		if endDateContract.After(t) {
			//TODO: сделаем выборку дог номеров, если они не сущ, потом только добавить в бд
			db.GetDBConn().Raw("SELECT id, bin FROM notification where contract_number = $1", value.ContractNumber).Scan(&data)

			log.Println("Результат ID", data.ID)
			if data.ID == 0 {
				db.GetDBConn().Exec("INSERT into notification (bin, contract_date, contract_number, type, email) VALUES ($1, $2, $3, $4, $5)",
					value.Bin, value.ContractDate, value.ContractNumber, value.Type, value.Email).Scan(&notification)
				fmt.Println(notification, "данные о уведомлении")
				message := fmt.Sprintf("Номер договора: %s, Тип Договора: %s, Срок истечении договора: %s, БИН: %s, почта: %s ",
					value.ContractNumber, value.Type, value.ContractDate, value.Bin, value.Email)

				//TODO: после того отправилось ообщения, и если ошибки не возникли при этом, надо статус поменять
				// на доставлено
				SendNotification(value.Email, message)

				// если данные не записываются, то вызови их

			}

		}
	}

}

func GetContractNot(contractNum string) int {
	var data struct {
		ID  int    `json:"id"`
		Bin string `json:"bin"`
	}
	var notifications []model.Notification

	//db.GetDBConn().Raw("SELECT requisites -> 'bin' AS bin, contract_parameters -> 'contract_date' AS contract_date, contract_parameters -> 'contract_number'  AS   contract_number, type, supplier_company_manager -> 'email'  AS email FROM contracts WHERE status = 'в работе'").Scan(&notifications)
	fmt.Println(notifications)

	db.GetDBConn().Raw("SELECT id, bin FROM notification where contract_number = $1", contractNum).Scan(&data)
	fmt.Println(data)

	return data.ID

}

func GetNotifications() []model.Notification {
	return repository.GetNotification()

}

func SendNotification(email string, message string) string {

	fmt.Println(email, " EMAIL На почту которую ты отправил")
	fmt.Println(message, "MESSAGE На почту которую ты отправил")

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "aziz.rahimov0001@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", email)
	//m.SetAddressHeader("Cc", "aziz.rahimov0001@gmail.com", "Aziz")

	// Set E-Mail subject
	m.SetHeader("Subject", "Notification of Contract")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", message)

	// Settings for SMTP server
	//d := gomail.NewDialer("smtp.gmail.com", 587, "thief65mk@gmail.com", "Aziz65mk")
	d := gomail.NewDialer("smtp.gmail.com", 587, "aziz.rahimov0001@gmail.com", "Aziz65mk$")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		//panic(err)
	}
	fmt.Println("successfully sent email!")
	return "successfully sent email!"
}
