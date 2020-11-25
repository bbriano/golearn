package main

import "fmt"

func main() {
    s := NewIntArrayStack()
    fmt.Printf("%T\n", s)
    s.Push(1)
    s.Push(2)
    s.Push(3)
    fmt.Println("size:", s.Size())
    fmt.Println(s.Pop())
    fmt.Println(s.Pop())
    s.Push(5)
    fmt.Println(s.Pop())
    fmt.Println(s.Pop())
    fmt.Println("size:", s.Size())
}

type Stack interface {
    Push(int)
    Pop() int
    Size() int
}

type IntArrayStack struct {
    slice []int
    size int
}

func NewIntArrayStack() *IntArrayStack {
    return &IntArrayStack{}
}

func (s *IntArrayStack) Push(value int) {
    s.slice = append(s.slice, value)
    s.size++
}

func (s *IntArrayStack) Pop() int {
    top := s.slice[s.size-1]
    s.size--
    s.slice = s.slice[:s.size]
    return top
}

func (s *IntArrayStack) Size() int {
    return s.size
}
