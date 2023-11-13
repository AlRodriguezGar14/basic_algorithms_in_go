package main

import "fmt"

func compareBT(a *BinaryNode[int], b *BinaryNode[int]) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Value != b.Value {
		return false
	}
	return compareBT(a.Left, b.Left) && compareBT(a.Right, b.Right)
}

func testCompareBT() {

	a := BinaryNode[int]{
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

	b := BinaryNode[int]{
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
			Value: 5,
			Right: &BinaryNode[int]{
				Value: 15,
			},
			Left: &BinaryNode[int]{
				Value: 10,
				Right: &BinaryNode[int]{
					Value: 7,
				},
			},
		},
	}

	c := BinaryNode[int]{
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
			Value: 5,
			Right: &BinaryNode[int]{
				Value: 15,
			},
			Left: &BinaryNode[int]{
				Value: 10,
				Right: &BinaryNode[int]{
					Value: 7,
				},
			},
		},
	}
	fmt.Println("a and b are exactly the same?", compareBT(&a, &b))
	fmt.Println("b and c are exactly the same?", compareBT(&b, &c))
}
