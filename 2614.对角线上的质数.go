func diagonalPrime(nums [][]int) int {
	n, res := len(nums), 0
	for i := 0; i < len(nums); i++ {
		if nums[i][i] > res && isPrime(nums[i][i]) {
			res = nums[i][i]
		}
		if nums[i][n-i-1] > res && isPrime(nums[i][n-i-1]) {
			res = nums[i][n-i-1]
		}
	}

	return res
}

func isPrime(i int) bool {
	if i == 1 {
		return false
	}
	for fcator := 2; fcator*fcator <= i; fcator++ {
		if i%fcator == 0 {
			return false
		}
	}

	return true
}
