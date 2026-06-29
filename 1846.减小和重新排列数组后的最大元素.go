package leetcode

import (
	"slices"
)

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	slices.Sort(arr)
	arr[0] = 1

	for i := 1; i < len(arr); i++ {
		switch arr[i] - arr[i-1] {
		case 0, 1:
			continue
		default:
			arr[i] = arr[i-1] + 1
		}

	}

	return arr[len(arr)-1]
}
