package main

import (
	"errors"
	"fmt"
)

type LNode struct {
	Data int
	Next *LNode
	Prev *LNode
}

type DoubleLinkedList struct {
	Length int
	Head   *LNode
}

func createSLL() *DoubleLinkedList {
	return &DoubleLinkedList{Length: 0, Head: nil}
}

func (dll *DoubleLinkedList) push(data int) {
	node := &LNode{Data: data, Next: nil}
	dll.Length++
	if dll.Head == nil {
		dll.Head = node
		return
	}
	prevHead := dll.Head
	dll.Head = node
	dll.Head.Next = prevHead
	dll.Head.Next.Prev = dll.Head
}

func (dll *DoubleLinkedList) remove(val int) (*LNode, error) {
	if dll.Head == nil {
		err := errors.New("No nodes linked")
		return &LNode{}, err
	}
	idx := 0
	curr := dll.Head
	var removed *LNode
	for idx < dll.Length && curr != nil {
		if curr.Data == val {
			removed = curr

			if curr.Prev != nil {
				curr.Prev.Next = curr.Next
			} else {
				dll.Head = curr.Next
			}
			if curr.Next != nil {
				curr.Next.Prev = curr.Prev
			}

			dll.Length--
			return removed, nil
		}
		if curr.Next != nil {
			curr = curr.Next
		}
		idx++
	}
	err := errors.New("No nodes with the provided value found")
	return &LNode{}, err
}

func test_ListNode() {
	test := createSLL()
	test.push(5)
	test.push(60)
	test.push(70)

	fmt.Println(test.Head.Next.Data)
	fmt.Println(test.Head.Data)
	removedNode, err := test.remove(60)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("removed:", removedNode, "with value", removedNode.Data)

	}
	fmt.Println(test.Head.Next.Data)
	fmt.Println(test.Head.Data)
}
