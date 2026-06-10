package leetcode

func maxTotalValue(nums []int, k int) int64 {
	minN, maxN := nums[0], nums[0]
	for _, num := range nums {
		minN = min(minN, num)
		maxN = max(maxN, num)
	}
	return int64((maxN - minN) * k)
}
