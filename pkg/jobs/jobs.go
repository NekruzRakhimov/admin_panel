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
	err := gocron.Every(30).Seconds().Do(service.Notification)
	if err != nil {
		log.Println("ошибка ")
		return
	}
	<-gocron.Start()
	//

}
