package main

import "fmt"

func main() {
    identityMatrix := [3][3]int{
        {1, 0, 0},
        {0, 1, 0},
        {0, 0, 1},
    }
    fmt.Println(identityMatrix)
    fmt.Println(cap(identityMatrix))
    fmt.Println()

    slice := []int{1, 2, 3}
    fmt.Printf("slice: %v %T\n", slice, slice)
    fmt.Printf("len:   %v %T\n", len(slice), len(slice))
    fmt.Printf("cap:   %v %T\n", cap(slice), cap(slice))

    newSlice := slice
    fmt.Printf("newSlice: %v %T\n", newSlice, newSlice)
    fmt.Printf("len:   %v %T\n", len(newSlice), len(newSlice))
    fmt.Printf("cap:   %v %T\n", cap(newSlice), cap(newSlice))
}
