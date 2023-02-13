package pkg

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func RunCronJobsV1() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Second().Do(func() {
		hello("Smith")
	})
	s.StartBlocking()
}

func hello(name string) {
	message := fmt.Sprintf("Hi, %v", name)
	fmt.Println(message)
}
