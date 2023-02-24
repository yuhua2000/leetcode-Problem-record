/*
 * @lc app=leetcode.cn id=2357 lang=golang
 *
 * [2357] 使数组中所有元素都等于零
 */

// @lc code=start
func minimumOperations1(nums []int) int {
	sort.Ints(nums)
	ans := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]-sum > 0 {
			ans++
			sum += min(nums[i], nums[i]-sum)
		}
	}
	return ans
}

func minimumOperations(nums []int) int {
	n := len(nums)
	m := make(map[int]struct{}, n/2)
	nS := struct{}{}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			m[nums[i]] = nS
		}
	}
	return len(m)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

