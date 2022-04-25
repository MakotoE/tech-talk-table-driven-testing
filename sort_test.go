package tech_talk_table_driven_testing

import "testing"

func TestInsertionSortSimple(t *testing.T) {
	arr := []int{0}
	InsertionSort(arr)
	if arr[0] != 0 {
		t.Error("incorrect result")
	}

	arr = []int{1, 0}
	InsertionSort(arr)
	if arr[0] != 0 || arr[1] != 1 {
		t.Error("incorrect result")
	}

	arr = []int{0, 2, 1}
	InsertionSort(arr)
	if arr[0] != 0 || arr[1] != 1 || arr[2] != 2 {
		t.Error("incorrect result")
	}
}

// EqualArrays returns true if the arrays are equal.
func EqualArrays(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		arr      []int
		expected []int
	}{
		{
			nil,
			nil,
		},
		{
			[]int{0},
			[]int{0},
		},
		{
			[]int{0, 1},
			[]int{0, 1},
		},
		{
			[]int{1, 0},
			[]int{0, 1},
		},
		{
			[]int{0, 0},
			[]int{0, 0},
		},
		{
			[]int{-1, 0, 1, 2, 3},
			[]int{-1, 0, 1, 2, 3},
		},
		{
			[]int{3, 2, 1, 0, -1},
			[]int{-1, 0, 1, 2, 3},
		},
		{
			[]int{0, 1, 2, -1, 3, 1},
			[]int{-1, 0, 1, 1, 2, 3},
		},
	}

	for i, test := range tests {
		InsertionSort(test.arr)
		if !EqualArrays(test.arr, test.expected) {
			t.Errorf("actual: %v, expected: %v, test case %v", test.arr, test.expected, i)
		}
	}
}
