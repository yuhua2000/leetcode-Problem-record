package leetcode

import "math"

func sortPermutation(nums []int) int {
	res := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] == i {
			continue
		}
		res &= nums[i]
	}

	if res == math.MaxInt32 {
		return 0
	}
	return res
}
