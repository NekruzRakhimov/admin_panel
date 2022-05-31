package service

import (
	"admin_panel/pkg/repository"
	"log"
	"time"
)

func CheckEndContract() {
	DateContracts, err := repository.CheckEndContract()
	if err != nil {
		log.Println(err)
		return
	}
	for _, contract := range DateContracts {
		t := time.Now()
		timeConctract, _ := ConvertStringTime(contract.EndDate)
		if !t.Before(timeConctract) {
			err = repository.ChangeStatusContract(contract.ID)
			log.Println(err)

		}

	}

}
