/*
Create a stack data structure with pop and push methods
*/
package main

import (
	"errors"
	"fmt"
	"sync"
)

var ErrEmptyStack = errors.New("stack is empty")

type Stack[T any] struct {
	elements []T
	mu       *sync.RWMutex
}

// Push adds an element to the stack
func (s *Stack[T]) Push(i ...T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.elements = append(s.elements, i...)
}

// Pop removes the last element of the stack and returns it
func (s *Stack[T]) Pop() (res T, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.elements) == 0 {
		return res, ErrEmptyStack
	}

	res = s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return res, nil
}

// Peek gets the last element from the stack without removing it
func (s *Stack[T]) Peek() (res T, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if len(s.elements) == 0 {
		return res, ErrEmptyStack
	}

	return s.elements[len(s.elements)-1], nil
}

// NewStack creates a new stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		mu: new(sync.RWMutex),
	}
}

func main() {
	s := NewStack[int]()
	s.Push(1, 2, 3, 4)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	s2 := NewStack[any]()
	s2.Push(1, 2, 3.03, "4")
	fmt.Println(s2.Pop())
	fmt.Println(s2.Pop())
	fmt.Println(s2.Pop())
	fmt.Println(s2.Pop())
	fmt.Println(s2.Pop())
}
