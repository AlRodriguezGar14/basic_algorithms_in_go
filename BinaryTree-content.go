package main

import "fmt"

type BinaryNode[T any] struct {
	Value T
	Left  *BinaryNode[T]
	Right *BinaryNode[T]
}

func TreeContentPre(curr *BinaryNode[int], path *[]int) {
	if curr == nil {
		return
	}
	*path = append(*path, curr.Value)
	TreeContentPre(curr.Left, path)
	TreeContentPre(curr.Right, path)
}

func TreeContentMid(curr *BinaryNode[int], path *[]int) {
	if curr == nil {
		return
	}
	TreeContentPre(curr.Left, path)
	*path = append(*path, curr.Value)
	TreeContentPre(curr.Right, path)
}
func TreeContentEnd(curr *BinaryNode[int], path *[]int) {
	if curr == nil {
		return
	}
	TreeContentPre(curr.Left, path)
	TreeContentPre(curr.Right, path)
	*path = append(*path, curr.Value)
}

func testBinaryTreeContent() {
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

	var path_pre []int
	var path_mid []int
	var path_end []int

	TreeContentPre(&tree, &path_pre)
	TreeContentMid(&tree, &path_mid)
	TreeContentEnd(&tree, &path_end)

	fmt.Println("Content of the tree when the content is added before the recurssion", path_pre)
	fmt.Println("Content of the tree when the content is added during the recurssion", path_mid)
	fmt.Println("Content of the tree when the content is added after the recurssion", path_end)
}
