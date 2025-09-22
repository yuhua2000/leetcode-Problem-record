package leetcode

func maxFrequencyElements(nums []int) int {
	numsMap := make(map[int]int)
	maxFrequency := 0
	result := 0
	for _, num := range nums {
		numsMap[num]++
		if numsMap[num] == maxFrequency {
			result += numsMap[num]
		}

		if numsMap[num] > maxFrequency {
			result = numsMap[num]
			maxFrequency = numsMap[num]
		}
	}

	return result
}
