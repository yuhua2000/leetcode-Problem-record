package leetcode

import "fmt"

func isGood(nums []int) bool {
	if len(nums) == 1 {
		return false
	}

	ans := make([]int, len(nums))

	for _, num := range nums {
		if num >= len(ans) {
			return false
		}

		ans[num]++
		if num < len(ans)-1 && ans[num] > 1 {
			return false
		}
	}

	return ans[len(ans)-1] == 2
}
