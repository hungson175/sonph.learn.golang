package learngo

import (
	"reflect"
	"testing"
)

func TestFiboClose(t *testing.T) {
	cases := []struct {
		n        int
		expected []int
	}{
		{1, []int{0}},
		{2, []int{0, 1}},
		{3, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 3}},
		{7, []int{0, 1, 1, 2, 3, 5, 8}},
	}
	for _, v := range cases {
		result := RunFiboChannel(v.n)
		if !reflect.DeepEqual(v.expected, result) {
			t.Errorf("Expected %v But result %v", v.expected, result)
		}
	}
}

func TestFiboSelect(t *testing.T) {
	cases := []struct {
		n        int
		expected []int
	}{
		{1, []int{0}},
		{2, []int{0, 1}},
		{3, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 3}},
		{7, []int{0, 1, 1, 2, 3, 5, 8}},
	}
	for _, v := range cases {
		result := RunFiboSelect(v.n)
		if !reflect.DeepEqual(v.expected, result) {
			t.Errorf("Expected %v But result %v", v.expected, result)
		}
	}

}
