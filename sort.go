package tech_talk_table_driven_testing

// InsertionSort sorts the array using insertion sort.
func InsertionSort(arr []int) {
	current := 1
	for current < len(arr) {
		tmp := arr[current]
		innerPos := current - 1
		for innerPos >= 0 && arr[innerPos] > tmp {
			arr[innerPos+1] = arr[innerPos]
			innerPos--
		}
		arr[innerPos+1] = tmp
		current++
	}
}
