package main

import "fmt"

func merge(l1 []int, l2 []int) []int {
	l3 := make([]int, len(l1)+len(l2))
	r1 := 0
	r2 := 0
	w := 0

	for r1 < len(l1) && r2 < len(l2) {
		if l1[r1] < l2[r2] {
			l3[w] = l1[r1]
			r1++
		} else {
			l3[w] = l2[r2]
			r2++
		}
		w++
	}

	if r1 >= len(l1) {
		copy(l3[w:], l2[r2:])
	}
	if r2 >= len(l2) {
		copy(l3[w:], l1[r1:])
	}

	return l3
}

func main() {
	l1 := []int{1, 3, 5, 6}
	l2 := []int{2, 4, 6, 8}
	fmt.Println(l1)
	fmt.Println(l2)
	l3 := merge(l1, l2)
	fmt.Println(l3)
}
