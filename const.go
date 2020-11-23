package main

import "fmt"

const a = 3

const (
    B  = 1 << (10 * iota)
    KB
    MB
    GB
    TB
    PB
)

func main() {
    var b int8 = 8
    var c float32 = 32
    fmt.Printf("%v, %T\n", a, a)
    fmt.Printf("%v, %T\n", a + b, a + b)
    fmt.Printf("%v, %T\n", a + c, a + c)
    fmt.Println()

    fileSizeBytes := PB
    fmt.Printf("%v PB\n", fileSizeBytes/PB)
    fmt.Printf("%v TB\n", fileSizeBytes/TB)
    fmt.Printf("%v GB\n", fileSizeBytes/GB)
    fmt.Printf("%v MB\n", fileSizeBytes/MB)
    fmt.Printf("%v KB\n", fileSizeBytes/KB)
    fmt.Printf("%v B\n",  fileSizeBytes/B)
}
