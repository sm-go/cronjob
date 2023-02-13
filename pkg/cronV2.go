package pkg

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func world(name string) {
	message := fmt.Sprintf("Hi, %v", name)
	fmt.Println(message)
}

func RunCronJobsV2() {
	// 2
	s := cron.New()

	// 3
	s.AddFunc("@every 2s", func() {
		world("This is a world")
	})
	s.AddFunc("@reboot", func() {
		world("Run once after reboot")
	})
	s.AddFunc("@weekly", func() {
		world("Run once a week 0 0 * * 0")
	})
	s.AddFunc("@monthly", func() {
		world("Run once a month")
	})
	s.AddFunc("@yearly", func() {
		world("Run once a year")
	})

	s.AddFunc("@every 3s", func() {
		fmt.Println("Hello world")
	})
	s.AddFunc("* * * * *", func() {
		fmt.Println("Every one minus")
	})
	s.AddFunc("*/2 * * * *", func() {
		fmt.Println("Every second minus or even minus")
	})
	s.AddFunc("1-59/2 * * * *", func() {
		fmt.Println("Every second odd minus")
	})

	// 4
	s.Start()
}
