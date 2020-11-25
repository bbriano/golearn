package main

import (
    "fmt"
)

func main() {
    sayHello("briano")
    fmt.Println("sum is:", sum(1, 2, 3, 4))

    fib, err := fibonacci(6)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("fib =", fib)

    func() {
        fmt.Println("anon")
    }()

    hello := func() {
        // fmt.Println("hello")
    }
    hello()
}

func sayHello(name string) {
    fmt.Printf("hello, %v\n", name)
}

func sum(numbers ...int) int {
    total := 1
    for _, n := range numbers {
        total += n
    }
    return total
}

func fibonacci (n int) (int, error) {
    if n < 0 {
        return 0, fmt.Errorf("undefined for n < 0")
    }
    current, next, temp := 0, 1, 0
    for i := 0; i < n; i++ {
        temp = next
        next = current + next
        current = temp
    }
    return current, nil
}
