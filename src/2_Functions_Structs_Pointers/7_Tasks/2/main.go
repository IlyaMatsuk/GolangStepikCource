package main

import (
	"fmt"
	"strings"
)

// https://stepik.org/lesson/229321/step/3?unit=201907

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(strings.Join(strings.Split(str, ""), "*"))
}