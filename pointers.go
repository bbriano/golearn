package main

import (
    "fmt"
)

func main() {
    a := [3]int{1, 2, 3}
    b := &a[0]
    c := &a[1]
    fmt.Println(a, b, c)
    fmt.Printf("%v %v %v\n", a, b, c)
}
