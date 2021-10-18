package main

import "fmt"

// https://stepik.org/lesson/228261/step/16?unit=200794

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println("It is", a/30, "hours", 2*(a%30), "minutes.")
}