package utils

import (
	"animaya/search-engine/search"
	"fmt"

	"github.com/robfig/cron"
)

func StartCronJobs() {
	c := cron.New()

	c.AddFunc("0 * * * *", search.RunEngine)
	c.AddFunc("15 * * * *", search.RunIndex)
	c.Start()
	cronCount := len(c.Entries())
	fmt.Printf("setup %d cron jobs \n", cronCount)
}

// func runEngine() {
// 	fmt.Println("Strting engine")
// }
