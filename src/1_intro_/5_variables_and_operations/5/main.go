package main

import "fmt"

// https://stepik.org/lesson/228261/step/15?unit=200794

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println((a % 100) / 10)
}