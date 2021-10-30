package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// https://stepik.org/lesson/359395/step/7?unit=343626

func main() {
	var str string
	str, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	strs := strings.Split(str, ",")

	var times [2]time.Time
	var err error
	template := "02.01.2006 15:04:05"
	for i := 0; i < 2; i++ {
		strs[i] = strings.TrimRight(strs[i], "\n")
		times[i], err = time.Parse(template, strs[i])
		if err != nil {
			panic(err)
		}
	}

	var diff time.Duration
	if times[0].Before(times[1]) {
		diff = times[1].Sub(times[0])
	} else {
		diff = times[0].Sub(times[1])
	}
	fmt.Println(diff)
}
