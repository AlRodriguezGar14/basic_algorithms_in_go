package main

import "fmt"

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

func test_ListNode() {
	test := createSLL()
	test.push(5)
	test.push(60)
	fmt.Println(test.Head.Data)
}
