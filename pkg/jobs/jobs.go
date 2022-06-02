package jobs

import (
	"admin_panel/pkg/service"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"log"
)

func RunJobs() {
	fmt.Println("вызов Джоба")
	// вызов сервис
	//err := gocron.Every(24).Hours().Do(service.Notification)
	//if err != nil {
	//	log.Println("ошибка ")
	//	return
	//}
	<-gocron.Start()
	//

}

func RunJobsCheckEndContract() {
	fmt.Println("вызов Джоба")

	err := gocron.Every(24).Hours().Do(service.Notification)
	if err != nil {
		log.Println("ошибка ")
		return
	}
	<-gocron.Start()
	//

}
