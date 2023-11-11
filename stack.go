package main

import (
	"errors"
	"fmt"
)

type StackNode[T any] struct {
	Value T
	Prev  *StackNode[T]
}

type Stack[T any] struct {
	Length uint
	Head   *StackNode[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{Length: 0, Head: nil}
}

func (s *Stack[T]) PeekStack() (*T, error) {
	if s.Head == nil {
		var ErrEmptyStack = errors.New("The  stack is empty")
		return nil, ErrEmptyStack
	}
	return &s.Head.Value, nil
}

func (s *Stack[T]) PushStack(value T) {
	s.Length++
	node := &StackNode[T]{Value: value}
	if s.Head == nil {
		s.Head = node
		return
	}

	node.Prev = s.Head
	s.Head = node
}

func (s *Stack[T]) PopStack() (*T, error) {
	if s.Head == nil {
		var ErrEmptyStack = errors.New("Empty Stack, nothing to pop-out")
		return nil, ErrEmptyStack
	}
	s.Length--

	var head = s.Head
	s.Head = s.Head.Prev
	head.Prev = nil

	return &head.Value, nil
}

func testStack() {
	myStack := NewStack[int]()

	fmt.Println("Stack Contents before Pushing:")
	printStackContents(myStack)
	fmt.Printf("Stack Length: %d\n", myStack.Length)
	myStack.PushStack(1)
	myStack.PushStack(2)
	fmt.Println("Stack Contents after Pushing two values:")
	printStackContents(myStack)
	fmt.Printf("Stack Length: %d\n\n", myStack.Length)
	myStack.PushStack(3)

	fmt.Printf("Stack Length with three pushes: %d\n", myStack.Length)

	fmt.Println("Stack Contents before Popping:")
	printStackContents(myStack)

	peekedValue, peekErr := myStack.PeekStack()
	if peekErr != nil {
		fmt.Println("Peek Error:", peekErr)
	} else {
		fmt.Printf("Peeked Value (the head is always the last added value. This is different from the queue): %d\n", *peekedValue)
	}

	fmt.Println("\nPopping All Nodes:")
	for myStack.Length > 0 {
		poppedValue, popErr := myStack.PopStack()
		if popErr != nil {
			fmt.Println("Pop Error:", popErr)
		} else {
			fmt.Printf("Popped Value: %d\n", *poppedValue)
		}
	}

	fmt.Printf("\nStack Length after Pops: %d\n", myStack.Length)

	fmt.Println("Stack Contents after Popping:")
	printStackContents(myStack)
}

func printStackContents[T int](s *Stack[T]) {
	current := s.Head
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.Prev
	}
	fmt.Println()
}
