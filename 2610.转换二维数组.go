package leetcode

import "slices"

func findMatrix(nums []int) [][]int {
	slices.Sort(nums)

	result := [][]int{[]int{nums[0]}}
	level := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			level++
		} else {
			level = 0
		}

		if level < len(result) {
			result[level] = append(result[level], nums[i])
		} else {
			result = append(result, []int{nums[i]})
		}
	}

	return result
}
