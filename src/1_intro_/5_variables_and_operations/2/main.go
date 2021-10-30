package main

import "fmt"

// https://stepik.org/lesson/228261/step/12?unit=200794

func main(){

  var a,b int
  fmt.Scan(&a) // считаем переменную 'a' с консоли
  fmt.Scan(&b) // считаем переменную 'b' с консоли

  fmt.Println((a*a) + (b*b))
}