package main

import "fmt"

// SortStack sorts a stack (slice whose end is the top) so the largest value
// ends on top, using only stack operations and recursion — no arrays, no
// index access, just pop, recurse, and sorted re-insertion.
func SortStack(stack []int) []int {
	if len(stack) <= 1 {
		return stack
	}
	top := stack[len(stack)-1]
	rest := SortStack(stack[:len(stack)-1])
	return insertSorted(rest, top)
}

// insertSorted pushes v into an already-sorted stack, keeping it sorted:
// pop while the top exceeds v, place v, push the popped values back.
func insertSorted(stack []int, v int) []int {
	if len(stack) == 0 || stack[len(stack)-1] <= v {
		return append(stack, v)
	}
	top := stack[len(stack)-1]
	rest := insertSorted(stack[:len(stack)-1], v)
	return append(rest, top)
}

func main() {
	fmt.Println(SortStack([]int{3, 1, 4, 1, 5, 9, 2, 6}))
}
