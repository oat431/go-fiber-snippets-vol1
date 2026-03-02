package job

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/gofiber/fiber/v3/log"
)

func RegisterDailySummary(s gocron.Scheduler) error {

	_, err := s.NewJob(
		gocron.CronJob("* * * * *", false),
		gocron.NewTask(ExampleBusinessLogic),
	)

	return err
}

func ExampleBusinessLogic() {
	log.Info("⏳ [CRON: DailySummary] Processing...")
}
