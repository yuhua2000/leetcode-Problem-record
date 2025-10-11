package leetcode

func maximumTripletValue1(nums []int) int64 {
	n := len(nums)
	var res int64
	for k := 2; k < n; k++ {
		m := nums[0]
		for j := 1; j < k; j++ {
			res = max(res, int64(m-nums[j])*int64(nums[k]))
			m = max(m, nums[j])
		}
	}

	return res
}

func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	leftMax := make([]int, n)
	rightMax := make([]int, n)

	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], nums[i-1])
	}

	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], nums[i+1])
	}

	var res int64

	for j := 1; j < n-1; j++ {
		res = max(res, int64(leftMax[j]-nums[j])*int64(rightMax[j]))
	}
	return res
}
