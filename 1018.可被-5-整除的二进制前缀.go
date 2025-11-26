package leetcode

/*
 * @lc app=leetcode.cn id=1018 lang=golang
 *
 * [1018] 可被 5 整除的二进制前缀
 */

// @lc code=start
func prefixesDivBy5(nums []int) []bool {
	num := 0
	result := make([]bool, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		num = (num<<1 + nums[i]) % 5
		result = append(result, (num%5 == 0))
	}
	return result
}

// @lc code=end
