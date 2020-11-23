package main

import "fmt"

func main() {
    array := [6]int{1, 2, 3, 4, 5, 6}
    slice := array[1:len(array)-1]
    // slice := array[:]
    slice = append(slice, 1)
    slice[2] = 99
    fmt.Println(array)
    fmt.Println(slice)
}
