package main

import (
	"fmt"
	"time"
)

// https://stepik.org/lesson/359395/step/3?unit=343626

func main() {
	var str string
	fmt.Scan(&str)
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		panic(err)
	}
	fmt.Println(t.Format(time.UnixDate))
}
