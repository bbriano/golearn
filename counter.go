package main

import (
	"fmt"
)

func main() {
	var c Counter = &IntCounter{}

	c.Increment()
	fmt.Println(c)

	c.Increment()
	fmt.Println(c)

	c.Decrement()
	fmt.Println(c)
}

type Counter interface {
	Increment()
	Decrement()
}

type IntCounter struct {
	value int
}

func (c *IntCounter) String() string {
	return fmt.Sprint(c.value)
}

func (c *IntCounter) Increment() {
	c.value++
}

func (c *IntCounter) Decrement() {
	c.value--
}
