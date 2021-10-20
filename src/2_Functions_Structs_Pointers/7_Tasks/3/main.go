package main

import "fmt"

// https://stepik.org/lesson/229321/step/4?unit=201907

func main(){
	var str string
	fmt.Scan(&str)
	var max_number byte
	
	for idx := range str {
		if str[idx] > max_number {
			max_number = str[idx]
		}
	}
	fmt.Println(string(max_number))
}