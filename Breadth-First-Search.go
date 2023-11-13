package main

import "fmt"

func bfs(head BinaryNode[int], needle int) bool {
	queue := NewQueue[BinaryNode[int]]()

	queue.Enqueue(head)

	for queue.Length > 0 {
		curr, err := queue.Dequeue()

		if err != nil {
			return false
		}

		if curr.Value == needle {
			return true
		}
		if curr.Left != nil {
			queue.Enqueue(*curr.Left)
		}
		if curr.Right != nil {
			queue.Enqueue(*curr.Right)
		}
	}
	return false
}

func testBreadthFirst() {
	tree := BinaryNode[int]{
		Value: 20,
		Right: &BinaryNode[int]{
			Value: 50,
			Right: &BinaryNode[int]{
				Value: 100,
			},
			Left: &BinaryNode[int]{
				Value: 30,
				Right: &BinaryNode[int]{
					Value: 45,
				},
				Left: &BinaryNode[int]{
					Value: 29,
				},
			},
		},
		Left: &BinaryNode[int]{
			Value: 10,
			Right: &BinaryNode[int]{
				Value: 15,
			},
			Left: &BinaryNode[int]{
				Value: 5,
				Right: &BinaryNode[int]{
					Value: 7,
				},
			},
		},
	}

	fmt.Println("The value 10 is in the tree?", bfs(tree, 10))
	fmt.Println("The value 11 is in the tree?", bfs(tree, 11))
}
