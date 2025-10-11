package leetcode

func minimumOperations(nums []int) int {
	result := 0
	for _, num := range nums {
		if num%3 != 0 {
			result++
		}
	}
	return result
}

// 给你一个整数数组 nums 。一次操作中，你可以将 nums 中的 任意 一个元素增加或者减少 1 。

// 请你返回将 nums 中所有元素都可以被 3 整除的 最少 操作次数。
