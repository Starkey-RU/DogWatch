package main

import (
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"time"
)

func StartCronTask() {
	scheduler, err := gocron.NewScheduler(gocron.WithLocation(time.Local))
	if err != nil {
		panic(fmt.Errorf("scheduler creation error: %w", err))
	}
	
	_, err = scheduler.NewJob(
		gocron.CronJob("*/5 * * * *", false),
		gocron.NewTask(func() {
			fmt.Println("--- Cronjob at:", time.Now().Format("2006-01-02 15:04:05"))
			ch := make(chan string, 10)
			go func() {
				userseek(ch)
				close(ch)
			}()
			DogWatch(ch)
		}),
	)
	if err != nil {
		panic(fmt.Errorf("job creation error: %w", err))
	}
	scheduler.Start()
}
