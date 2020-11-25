package main

import (
    "fmt"
    "sync"
)

var wg = sync.WaitGroup{}
var name = "begin"

func main() {
    wg.Add(1)
    go sayHello()
    name = "end"
    wg.Wait()
}

func sayHello() {
    fmt.Printf("hello, %v\n", name)
    wg.Done()
}
