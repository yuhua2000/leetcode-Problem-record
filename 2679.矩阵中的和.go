package leetcode

import "sort"

func matrixSum(nums [][]int) (result int) {
	ans := make([]int, len(nums[0]))
	for i := 0; i < len(nums); i++ {
		sort.Ints(nums[i])
		for j := 0; j < len(nums[i]); j++ {
			if nums[i][j] > ans[j] {
				result += (nums[i][j] - ans[j])
				ans[j] = nums[i][j]
			}
		}
	}
	return result
}
