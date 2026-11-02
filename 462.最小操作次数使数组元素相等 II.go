package leetcode

import "slices"

func minMoves2(nums []int) int {
	slices.Sort(nums)

	result := 0
	mid := nums[len(nums)/2]
	for _, num := range nums {
		if num < mid {
			result += mid - num
		} else {
			result += num - mid
		}
	}

	return result
}
