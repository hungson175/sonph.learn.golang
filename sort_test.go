package learngo

import (
	"reflect"
	"testing"
)

func TestSwap(t *testing.T) {
	cases := []struct {
		x, y, lx, ly int
	}{
		{0, 100, 100, 0},
		{0, 0, 0, 0},
		{1, -1, -1, 1},
	}
	for _, v := range cases {
		x, y, lx, ly := v.x, v.y, v.lx, v.ly
		Swap(&x, &y)
		if x != lx || y != ly {
			t.Errorf("Expected=(%v,%v) Result = (%v,%v)", lx, ly, x, y)
		}
	}
}

func TestSort(t *testing.T) {
	cases := []struct {
		original, sorted []int
	}{
		{[]int{3, 1, 2, 4}, []int{1, 2, 3, 4}},
		{[]int{}, []int{}},
		{[]int{0, 0, 2, 3}, []int{0, 0, 2, 3}},
		{[]int{0, 2, -1, 3}, []int{-1, 0, 2, 3}},
	}

	for _, v := range cases {
		original, sorted := v.original, v.sorted
		compare := func(a, b int) int {
			if a == b {
				return 0
			}
			if a < b {
				return -1
			}
			return 1
		}
		result := Sort(original, compare)

		if !reflect.DeepEqual(result, sorted) {
			t.Errorf("Expected %v But result %v", sorted, result)
		}
	}
}
