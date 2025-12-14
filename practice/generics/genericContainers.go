package main

import "fmt"

// Real-life Use Case #2 - Generic Containers
// Instead of writing separate stack/queue/list for each type,
// you write one and reuse it:

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) GetLen() {
	fmt.Println(len(s.items))
}

func (s *Stack[T]) Pop() T {
	last := s.items[len(s.items) - 1]
	s.items = s.items[:len(s.items) - 1]
	return  last
}

func main() {
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	fmt.Println(intStack)

	stringStack := Stack[string]{}
	stringStack.Push("hello")
	fmt.Println(stringStack)

	// intStack.GetLen()

	// last := intStack.Pop()
	last := stringStack.Pop()

	fmt.Println("This is popped value:",last)
}