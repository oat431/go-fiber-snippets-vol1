package batch

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/gofiber/fiber/v3/log"
)

func RegisterDailySummary(s gocron.Scheduler) error {

	_, err := s.NewJob(
		gocron.CronJob("* * * * *", false),
		gocron.NewTask(func() {
			log.Info("⏳ [CRON: DailySummary] Processing...")
		}),
	)

	return err
}
