package main

import (
    "fmt"
    "strconv"
)

func main() {
    var i int32 = 888
    fmt.Printf("%v, %T\n", i, i)

    var j float32
    j = float32(i)
    fmt.Printf("%v, %T\n", j, j)

    var str string
    str = strconv.Itoa(int(i))
    fmt.Printf("%v, %T\n", str, str)

    var z complex128 = complex(float64(i), 0)
    fmt.Printf("%v, %T\n", z, z)

    a := 10
    b := 3
    fmt.Println(a | b)
    fmt.Println(a & b)
    fmt.Println(a ^ b)
    fmt.Println(a &^ b)
}
