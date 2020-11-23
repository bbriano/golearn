package main

import "fmt"

func main() {
    myString := "hello"
    myRune := 'a'
    myBytes := []byte(myString)

    fmt.Printf("%v, %T\n", myString, myString)
    fmt.Printf("%v, %T\n", myBytes, myBytes)
    fmt.Printf("%v, %T\n", myRune, myRune)
}
