# Setup Go Batch Process


## 1. install go cron package
```bash
go get github.com/go-co-op/gocron/v2
```

## 2. Create a Go file for batch processing job
```go
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

```

## 3. Register the batch job in the RegisteringCronjob function
```go
package bootstap

import (
	"go-fiber-snippets/job"

	"github.com/go-co-op/gocron/v2"
	"github.com/gofiber/fiber/v3/log"
)

func RegisterJobs() (gocron.Scheduler, error) {

	s, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	job1 := job.RegisterDailySummary(s)
	if job1 != nil {
		log.Error("❌ Failed to register DailySummary Cron Job: ", job1)
		return nil, job1
	}
	
    // job2
	// job3

	log.Info("✅ All Cron Jobs registered successfully")

	return s, nil
}
```

## 4. Start the scheduler in the main function
```go
func StartServer() {
	scheduler, batchFailed := bootstap.RegisterJobs()
	if batchFailed != nil {
		log.Fatal("Failed to start scheduler: ", batchFailed)
	}
	scheduler.Start()
}
```
