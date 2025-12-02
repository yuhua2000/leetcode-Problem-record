package leetcode

import "math"

func maxSubarraySum(nums []int, k int) int64 {
	preSum := int64(0)
	kSum := make([]int64, len(nums))
	for i := 0; i < len(kSum); i++ {
		kSum[i] = math.MaxInt64 / 2
	}
	kSum[k-1] = 0

	result := int64(math.MinInt64)
	for i := 0; i < len(nums); i++ {
		preSum += int64(nums[i])
		result = max(result, preSum-kSum[i%k])
		kSum[i%k] = min(kSum[i%k], preSum)
	}

	return result
}
