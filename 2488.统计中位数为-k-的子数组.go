package leetcode

/*
 * @lc app=leetcode.cn id=2488 lang=golang
 *
 * [2488] 统计中位数为 K 的子数组
 */

// @lc code=start
func countSubarrays(nums []int, k int) int {
	n := len(nums)
	index := -1
	for i := 0; i < n; i++ {
		if nums[i] == k {
			index = i
			break
		}
	}
	ans := 0
	sum := 0
	counts := map[int]int{}
	counts[0] = 1
	for i := 0; i < n; i++ {
		sum += sign(nums[i] - k)
		if i < index {
			counts[sum]++
		} else {
			prev0 := counts[sum]
			prev1 := counts[sum-1]
			ans += prev0 + prev1
		}
	}
	return ans
}

func sign(num int) int {
	if num == 0 {
		return 0
	}
	if num > 0 {
		return 1
	}
	return -1
}

// @lc code=end
