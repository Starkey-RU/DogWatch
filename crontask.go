package main

import (
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"time"
)

// функция для cron job вызовов основных функций - uptime.go И dogwatch.go
func StartCronTask() {
	scheduler, err := gocron.NewScheduler(gocron.WithLocation(time.Local))
	if err != nil {
		panic(fmt.Sprintf("scheduler creation error: %v", err))
	}
	// cronjob (5 mins)
	_, err = scheduler.NewJob(
		gocron.CronJob("*/5 * * * *", false),
		gocron.NewTask(func() {
			fmt.Println("---Cronjob at:", time.Now().Format("2006-01-02 15:04:05"))
			ch := make(chan string, 10)
			go func() {
				userseek(ch)
				close(ch)
			}()
			DogWatch(ch)
		}),
	)
	if err != nil {
		panic(fmt.Sprintf("job creation error: %v", err))
	}
	scheduler.Start()
}
