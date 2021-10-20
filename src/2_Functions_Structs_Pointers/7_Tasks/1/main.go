package main

import (
    "fmt"
    "math"
)

// https://stepik.org/lesson/229321/step/2?unit=201907

func main(){
    var a, b float64
    fmt.Scan(&a, &b)
    fmt.Println(math.Hypot(a, b))
}

