package main

import (
	"errors"
	"fmt"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type Queue[T any] struct {
	Length uint
	Head   *Node[T]
	Tail   *Node[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{Head: nil, Tail: nil, Length: 0}
}

func (q *Queue[T]) Peek() (*T, error) {
	if q.Head != nil {
		return &q.Head.Value, nil
	}
	var ErrHeadEmpty = errors.New("Head is empty")
	return nil, ErrHeadEmpty
}

func (q *Queue[T]) Dequeue() (*T, error) {
	if q.Head == nil {
		var ErrNoHead = errors.New("No content in the head")
		return nil, ErrNoHead
	}
	q.Length--
	head := q.Head
	q.Head = q.Head.Next

	head.Next = nil

	return &head.Value, nil
}

// Add a new value to the queue
// If the queue is empty, both head and tail are the same
func (q *Queue[T]) Enqueue(value T) {
	q.Length++
	var node = &Node[T]{Value: value}
	if q.Head == nil {
		q.Head = node
		q.Tail = q.Head
		return
	}
	q.Tail.Next = node // The current tail points to the new node
	q.Tail = node      // Update the tail to be the new node

}

func testQueue() {
	myQueue := NewQueue[int]()

	myQueue.Enqueue(1)
	myQueue.Enqueue(2)
	myQueue.Enqueue(3)

	fmt.Printf("Queue Length: %d\n", myQueue.Length)

	fmt.Println("Queue Contents before Dequeuing:")
	printQueueContents(myQueue)

	peekedValue, peekErr := myQueue.Peek()
	if peekErr != nil {
		fmt.Println("Peek Error:", peekErr)
	} else {
		fmt.Printf("Peeked Value: %d\n", *peekedValue)
	}

	fmt.Println("\nDequeuing All Nodes:")
	for myQueue.Length > 0 {
		dequeuedValue, dequeueErr := myQueue.Dequeue()
		if dequeueErr != nil {
			fmt.Println("Dequeue Error:", dequeueErr)
		} else {
			fmt.Printf("Dequeued Value: %d\n", *dequeuedValue)
		}
	}

	fmt.Printf("\nQueue Length after Dequeues: %d\n", myQueue.Length)

	fmt.Println("Queue Contents after Dequeuing:")
	printQueueContents(myQueue)
}

func printQueueContents[T int](q *Queue[T]) {
	current := q.Head
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.Next
	}
	fmt.Println()
}
