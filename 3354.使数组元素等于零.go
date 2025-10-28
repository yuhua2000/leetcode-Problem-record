package leetcode

func countValidSelections(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	ans := 0
	result := 0
	for i := 0; i < len(nums); i++ {
		ans += nums[i]
		if nums[i] == 0 && ans*2 == sum {
			result += 2

		} else if nums[i] == 0 && (ans*2 == sum-1 || ans*2 == sum+1) {
			result += 1
		}
	}

	return result
}
