package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// https://stepik.org/lesson/359395/step/4?unit=343626

func main() {
	var str string
	str, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.TrimRight(str, "\n")
	template := "2006-01-02 15:04:05"
	t, err := time.Parse(template, str)
	if err != nil {
		panic(err)
	}
	if t.Hour() >= 13 {
		t = t.Add(time.Hour * 24)
	}
	fmt.Println(t.Format(template))
}
