package learngo

import (
	"reflect"
	"testing"
)

func setUpTree1() *Tree {
	t := Tree{nil, 3, nil}
	t.Left = &Tree{nil, 1, nil}
	t.Left.Left = &Tree{nil, 1, nil}
	t.Left.Right = &Tree{nil, 2, nil}
	t.Right = &Tree{nil, 8, nil}
	t.Right.Left = &Tree{nil, 5, nil}
	t.Right.Right = &Tree{nil, 13, nil}
	return &t
}
func setUpTree2() *Tree {
	t := Tree{nil, 8, nil}
	t.Left = &Tree{nil, 3, nil}
	t.Left.Left = &Tree{nil, 1, nil}
	t.Left.Left.Left = &Tree{nil, 1, nil}
	t.Left.Left.Right = &Tree{nil, 2, nil}
	t.Left.Right = &Tree{nil, 5, nil}
	t.Right = &Tree{nil, 13, nil}
	return &t
}
func TestTreeSequence(testCase *testing.T) {
	t := setUpTree1()
	result := t.GetSequence()
	//TODO: Write a wrapper over the testing
	expected := []int{1, 1, 2, 3, 5, 8, 13}
	if !reflect.DeepEqual(result, expected) {
		testCase.Errorf("Expected %v But result %v", expected, result)
	}

	t2 := setUpTree2()
	result = t2.GetSequence()
	//TODO: Write a wrapper over the testing
	expected = []int{1, 1, 2, 3, 5, 8, 13}
	if !reflect.DeepEqual(result, expected) {
		testCase.Errorf("Expected %v But result %v", expected, result)
	}

	//Single tree
	t1 := Tree{nil, 1, nil}
	expected = []int{1}
	result = t1.GetSequence()
	if !reflect.DeepEqual(expected, result) {
		testCase.Errorf("Expected %v But result %v", expected, result)
	}
}

func TestTreeEquivalent(testCase *testing.T) {
	t1 := setUpTree1()
	t2 := setUpTree2()
	t3 := &Tree{nil, 1, nil}
	cases := []struct {
		ta, tb *Tree
		eq     bool
	}{
		{t1, t2, true},
		{t2, t1, true},
		{t1, t3, false},
		{t3, t1, false},
		{t3, t2, false},
		{t2, t3, false},
	}
	for _, test := range cases {
		result := test.ta.EqualSequence(test.tb)
		if test.eq != result {
			testCase.Errorf("Expected %v But result %v", test.eq, result)
		}
	}

}
