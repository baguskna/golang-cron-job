package main

import (
	"fmt"
	"math/rand"
	"time"

	"golang-cron-job/app"

	"github.com/go-co-op/gocron"
	_ "github.com/go-sql-driver/mysql"
)

func runCronJobs() {
	app.NewDB()
	s := gocron.NewScheduler(time.UTC)
	defer app.GetDB().Close()

	s.Every(15).Seconds().Do(func() {
		water := rand.Intn(100)
		wind := rand.Intn(100)
		fmt.Println("=================")
		fmt.Println("water", water)
		if water < 5 {
			fmt.Println("status aman")
		} else if 6 <= water && water <= 8 {
			fmt.Println("status siaga")
		} else {
			fmt.Println("status bahaya")
		}

		fmt.Println("wind", wind)
		if wind < 6 {
			fmt.Println("status aman")
		} else if 7 <= wind && wind <= 15 {
			fmt.Println("status siaga")
		} else {
			fmt.Println("status bahaya")
		}
		fmt.Println("=================")

		SQL := "INSERT INTO water_wind(water, wind) value(?, ?)"
		_, err := app.GetDB().Exec(SQL, water, wind)
		if err != nil {
			panic(err)
		}
	})

	s.StartBlocking()
}

func main() {
	runCronJobs()
}
