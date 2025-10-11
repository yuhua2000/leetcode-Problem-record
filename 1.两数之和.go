package leetcode

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	prevNum := make(map[int]int)
	for i, num := range nums {
		if p, ok := prevNum[target-num]; ok {
			return []int{i, p}
		}
		prevNum[num] = i
	}
	return nil
}

// @lc code=end
