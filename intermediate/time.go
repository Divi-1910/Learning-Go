package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	fmt.Println(current)

	specificTime := time.Date(2025, time.June, 17, 4, 9, 0, 0, time.UTC)
	fmt.Println(specificTime)

	paresedTime, _ := time.Parse("2006-01-02", "2020-05-01") // layout value : Mon Jan 2 15:04:05 MST 2006
	fmt.Println(paresedTime)

	t1 := time.Now()
	fmt.Println(t1.Format("Monday"))

	afterOneDay := t1.Add(time.Hour * 24)
	fmt.Println(afterOneDay.Weekday())

	loc, _ := time.LoadLocation("America/Los_Angeles")
	fmt.Println(loc)
	tInNewYork := time.Now().In(loc)
	fmt.Println(tInNewYork)

}
