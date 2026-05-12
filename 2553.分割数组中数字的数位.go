package leetcode

func separateDigits(nums []int) []int {
	result := make([]int, 0, len(nums))
	for i := range nums {
		num := nums[len(nums)-1-i]
		for num > 0 {
			result = append(result, num%10)
			num = num / 10
		}
	}

	for i := range len(result) / 2 {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}

	return result
}
