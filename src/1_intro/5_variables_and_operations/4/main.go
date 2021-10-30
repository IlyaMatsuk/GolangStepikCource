package main

import "fmt"

// https://stepik.org/lesson/228261/step/14?thread=solutions&unit=200794

func main() {
	var a int
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	str := fmt.Sprint(a)
	fmt.Println(string(str[len(str)-1]))
}