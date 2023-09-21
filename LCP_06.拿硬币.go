func minCount(coins []int) int {
	result := 0
	for _, coin := range coins {
		result += coin/2 + coin%2
	}
	return result
}
