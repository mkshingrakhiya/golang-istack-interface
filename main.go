package main

import (
	"errors"
	"fmt"
)

// IStack interface defines the methods for an integer stack
type IStack interface {
	InitStack(capacity int)
	Push(item int) error
	Pop() (int, error)
	IsStackEmpty() bool
}

// IntStack is a concrete implementation of the IStack interface
type IntStack struct {
	items    []int // Slice to hold stack items
	top      int   // Index of the top item
	capacity int   // Maximum capacity of the stack
}

// InitStack initializes the stack with the specified capacity
func (s *IntStack) InitStack(capacity int) {
	s.capacity = capacity
	s.items = make([]int, 0, capacity)
	s.top = -1
}

// Push adds an item to the top of the stack
func (s *IntStack) Push(item int) error {
	if s.top >= s.capacity-1 {
		return errors.New("stack overflow")
	}
	s.items = append(s.items, item)
	s.top++
	return nil
}

// Pop removes and returns the item at the top of the stack
func (s *IntStack) Pop() (int, error) {
	if s.IsStackEmpty() {
		return 0, errors.New("stack is empty")
	}
	item := s.items[s.top]
	s.items = s.items[:s.top] // Resize the slice
	s.top--
	return item, nil
}

// IsStackEmpty checks if the stack is empty
func (s *IntStack) IsStackEmpty() bool {
	return s.top == -1
}

func main() {
	var stack IStack = &IntStack{}
	stack.InitStack(5) // Initialize stack with a capacity of 5

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if topItem, err := stack.Pop(); err == nil {
		fmt.Println("Popped item:", topItem) // Should print 3
	} else {
		fmt.Println(err)
	}

	fmt.Println("Is stack empty?", stack.IsStackEmpty()) // Should print false

	// Pushing more items to demonstrate capacity
	for i := 4; i <= 6; i++ {
		if err := stack.Push(i); err != nil {
			fmt.Println(err) // Should print stack overflow after reaching max capacity
		}
	}

	// Print remaining items
	for !stack.IsStackEmpty() {
		if item, err := stack.Pop(); err == nil {
			fmt.Println("Popped item:", item) // Should print remaining items in LIFO order
		}
	}
}
