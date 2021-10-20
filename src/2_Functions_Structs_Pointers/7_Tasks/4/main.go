package main

import "fmt"

// https://stepik.org/lesson/229321/step/5?unit=201907

func main() {
	var s string
	fmt.Scan(&s)
	for _, c := range s {
		c -= '0'
		fmt.Print(c * c)
	}
}