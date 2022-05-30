package boot

import (
	"github.com/go-co-op/gocron"
	"os"
	"time"
	"tuble/src/utils"
)

// StartScheduler generates a new map in a scheduled manner.
// If the schedule is not configured, it will be disabled for this process.
func StartScheduler() {
	schedule := os.Getenv("CRON_MAP_GENERATE")
	if schedule != "" { //Only enable cron if schedule set.
		clock := gocron.NewScheduler(time.UTC)
		_, err := clock.Cron(schedule).Do(utils.MapToJson)
		if err != nil {
			panic(err)
		}
		clock.StartAsync()
	}
}
