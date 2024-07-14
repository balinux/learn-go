package main

import (
	"fmt"
	"time"
)

func main() {
	timenow := time.Now()
	fmt.Printf("time %v\n", timenow)

	time2 := time.Date(2011, time.June, 30, 12, 44, 33, 123123, time.UTC)
	fmt.Println("time2", time2)

	// method milik
	now := time.Now()
	fmt.Println("year:", now.Year(), "month:", now.Month(), now.Local(), now.Location().String())

	// parsing string ke time.Time
	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = "2015-09-02 08:04:00"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())
}
