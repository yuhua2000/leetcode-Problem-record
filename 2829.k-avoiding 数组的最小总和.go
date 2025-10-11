package leetcode

func minimumSum(n int, k int) int {
	result := 0
	for i := 1; i <= k/2 && i <= n; i++ {
		result += i
	}

	ans := k
	if ans == (k)/2 {
		ans++
	}
	for i := 0; i < n-k/2; i++ {
		result += ans
		ans++
	}
	return result
}
