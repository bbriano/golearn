package main

import "fmt"

func main() {
    secret := 34
    guess := 0

    if guess == secret {
        fmt.Println("you got it!")
    } else if guess < secret {
        fmt.Println("higher")
    } else if guess > secret {
        fmt.Println("lower")
    }

    switch 1 {
    // case 1:
    case 1, 4:
        fmt.Println("One or Four")
        fallthrough
    case 2, 3:
        fmt.Println("Two or Three")
    case 5:
        fmt.Println("Five")
    default:
        fmt.Println("Something else")
    }
}
