package main

import (
    "fmt"
    "log"
)

func main() {
    fmt.Println("start")
    iPanick()
    fmt.Println("end")
}

func iPanick() {
    fmt.Println("i panick")
    defer func() {
        err := recover()
        if err != nil {
            log.Println("ERROR:", err)
        }
    }()
    panic("PANIK")
}
