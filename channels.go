package main

import (
    "fmt"
    "sync"
)

var wg = sync.WaitGroup{}

func main() {
    channel := make(chan int)

    wg.Add(2)

    // Receiver
    go func(ch <-chan int) {
        for val := range ch {
            fmt.Println(val)
        }
        wg.Done()
    }(channel)

    // Sender
    go func(ch chan<- int) {
        ch <- 1
        ch <- 2
        ch <- 3
        close(ch)
        wg.Done()
    }(channel)

    wg.Wait()
}
