func minOperations(nums []int, k int) int {
	numMap := make(map[int]bool, len(nums))
	minNum := nums[0]
	for _, num := range nums {
		numMap[num] = true
		minNum = min(minNum, num)
		if num < k {
			return -1
		}
	}

	result := len(numMap)
	if minNum == k {
		result--
	}

	return result
}
