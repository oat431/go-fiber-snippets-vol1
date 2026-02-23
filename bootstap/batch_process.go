package bootstap

import (
	"go-fiber-snippets/batch"

	"github.com/go-co-op/gocron/v2"
	"github.com/gofiber/fiber/v3/log"
)

func RegisterJobs() (gocron.Scheduler, error) {

	s, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	job1 := batch.RegisterDailySummary(s)
	if job1 != nil {
		log.Error("❌ Failed to register DailySummary Cron Job: ", job1)
		return nil, job1
	}

	log.Info("✅ All Cron Jobs registered successfully")

	return s, nil
}
