package leetcode

func isZeroArray(nums []int, queries [][]int) bool {
	deltaArray := make([]int, len(nums)+1)
	for i := range len(queries) {
		l, r := queries[i][0], queries[i][1]
		deltaArray[l]++
		deltaArray[r+1]--
	}
	operationCounts := make([]int, len(nums)+1)
	currentOperations := 0
	for i := range len(deltaArray) {
		currentOperations += deltaArray[i]
		operationCounts[i] = currentOperations
	}

	for i := range len(nums) {
		if nums[i] > operationCounts[i] {
			return false
		}
	}
	return true
}
