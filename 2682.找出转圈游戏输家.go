func circularGameLosers(n int, k int) []int {
	ops := make([]bool, n)
	index := 0
	for pace := k; !ops[index]; pace += k {
		ops[index] = true
		index = (index + pace) % n
	}
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		if !ops[i] {
			result = append(result, i+1)
		}
	}
	return result
}
