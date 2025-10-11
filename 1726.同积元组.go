package leetcode

/*
 * @lc app=leetcode.cn id=1726 lang=golang
 *
 * [1726] 同积元组
 */

// @lc code=start
func tupleSameProduct(nums []int) (result int) {
	n := len(nums)
	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			mp[nums[i]*nums[j]]++
		}
	}
	for _, v := range mp {
		result += v * (v - 1) * 4
	}
	return
}

// @lc code=end
