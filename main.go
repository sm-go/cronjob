package main

import (
	"fmt"

	"github.com/smith-golang/cronjob/pkg"
)

func main() {
	// version one
	pkg.RunCronJobsV1()

	// cron job main function should be only one function

	// version two
	pkg.RunCronJobsV2()
	fmt.Scanln()
}
