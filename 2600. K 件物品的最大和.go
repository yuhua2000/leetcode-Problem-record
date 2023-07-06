func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	return min(k, numOnes) - max(k-numOnes-numZeros, 0)
}




