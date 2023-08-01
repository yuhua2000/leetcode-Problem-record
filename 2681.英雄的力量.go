func sumOfPower(nums []int) (res int) {
	n := len(nums)
	sort.Ints(nums)
	dp := 0
	preSum := 0

	mod := int(1e9 + 7)
	for i := 0; i < n; i++ {
		dp = (nums[i] + preSum) % mod
		preSum = (preSum + dp) % mod
		res = (res + (((nums[i]*nums[i])%mod)*dp)%mod) % mod
	}
	return res
}
