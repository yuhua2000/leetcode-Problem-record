package leetcode

func isPossibleToSplit(nums []int) bool {
	numFlags := [101]int{}
	for _, num := range nums {
		numFlags[num]++
		if numFlags[num] == 3 {
			return false
		}
	}
	return true
}
