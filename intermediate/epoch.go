package main

import (
	"fmt"
	"time"
)

func main() {
	// 00:00:00 UTC on Jan 1 , 1970  --> unix epoch

	t := time.Since(time.Now())
	fmt.Println(t)

	now := time.Now().Unix()
	fmt.Println("Current Unix Time :", now)

	layout := "2006-01-02T15:04:05Z07:00"
	str := "2024-07-04T14:30:18Z"

	t1, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(t1)

}
