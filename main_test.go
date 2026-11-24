package main

import (
	"reflect"
	"testing"
)

func TestSortStack(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{3, 1, 4, 1, 5, 9, 2, 6}, []int{1, 1, 2, 3, 4, 5, 6, 9}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{-1, -3, -2}, []int{-3, -2, -1}},
	}
	for _, c := range cases {
		in := append([]int{}, c.in...)
		if got := SortStack(in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("SortStack(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
