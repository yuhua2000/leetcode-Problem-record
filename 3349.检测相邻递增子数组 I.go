package leetcode

func hasIncreasingSubarrays(nums []int, k int) bool {
	ans := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			ans[i] = ans[i+1] + 1
		}
	}
	for i := 0; i < len(ans)-k; i++ {
		if ans[i] >= k-1 && ans[i+k] >= k-1 {
			return true
		}
	}
	return false
}
