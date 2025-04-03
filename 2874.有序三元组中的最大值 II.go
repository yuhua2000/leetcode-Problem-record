func maximumTripletValue(nums []int) int64 {
	var (
		n        = len(nums)
		leftMax  = make([]int, n)
		rightMax = make([]int, n)
		ret      int64
	)

	for i := 1; i < n; i++ {
		leftMax[i] = max(nums[i-1], leftMax[i-1])
		rightMax[n-i-1] = max(nums[n-i], rightMax[n-i])
	}

	for i := 1; i < n-1; i++ {
		ret = max(ret, int64(leftMax[i]-nums[i])*int64(rightMax[i]))
	}

	return ret
}
