package cron

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/gofiber/fiber/v3/log"
)

// RegisterAll creates and returns a scheduler with all cron jobs.
func RegisterAll() (gocron.Scheduler, error) {
	s, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	// Example: run every minute
	_, err = s.NewJob(
		gocron.CronJob("* * * * *", false),
		gocron.NewTask(dailySummary),
	)
	if err != nil {
		return nil, err
	}

	// Add more jobs here:
	// _, err = s.NewJob(gocron.CronJob("0 9 * * *", false), gocron.NewTask(morningReport))

	log.Info("All cron jobs registered")
	return s, nil
}

// dailySummary is a placeholder for your business logic.
func dailySummary() {
	log.Info("[cron:daily-summary] processing…")
}
