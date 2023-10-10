func sumDistance(nums []int, s string, d int) int {
	const mod = 1e9 + 7
	n := len(nums)
	for i := 0; i < n; i++ {
		if s[i] == 'R' {
			nums[i] += d
		} else {
			nums[i] -= d
		}
	}
	sort.Ints(nums)
	result := 0
	for i := 1; i < n; i++ {
		result += nums[i] - nums[i-1]*i%mod*(n-i)%mod
		result %= mod
	}
	return result
}
