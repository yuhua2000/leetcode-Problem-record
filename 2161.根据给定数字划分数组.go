package leetcode

func pivotArray(nums []int, pivot int) []int {
	var less, greater, equal []int

	for _, num := range nums {
		if num < pivot {
			less = append(less, num)
		} else if num > pivot {
			greater = append(greater, num)
		} else {
			equal = append(equal, num)
		}
	}

	result := append(less, equal...)
	return append(result, greater...)
}
