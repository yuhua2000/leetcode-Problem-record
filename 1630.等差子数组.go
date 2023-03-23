/*
 * @lc app=leetcode.cn id=1630 lang=golang
 *
 * [1630] 等差子数组
 */

// @lc code=start
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	n := len(l)
	ans := make([]bool, n)
	for i := 0; i < n; i++ {
		ans[i] = isArithmetic(nums[l[i] : r[i]+1])
	}
	return ans
}

func isArithmetic(nums []int) bool {
	n := len(nums)
	max, min := nums[0], nums[0]
	for i := 0; i < n; i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	if max == min {
		return true
	}
	d := (max - min) / (n - 1)
	if min+d*(n-1) != max {
		return false
	}
	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}
	for i := 1; i < n; i++ {
		min += d
		if _, ok := m[min]; !ok {
			return false
		}
	}
	return true
}

// @lc code=end
