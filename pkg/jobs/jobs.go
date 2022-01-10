package jobs

import (
	"admin_panel/pkg/service"
	"github.com/jasonlvhit/gocron"
)

func RunJobs()  {

	// вызов сервис
	gocron.Every(2).Seconds().Do(service.Notification)
	//

}
