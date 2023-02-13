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

	s.AddFunc("@every 3s", func() {
		fmt.Println("Hello world")
	})

	// 4
	s.Start()
}
