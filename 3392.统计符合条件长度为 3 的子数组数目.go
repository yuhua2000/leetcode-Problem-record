func countSubarrays(nums []int) (result int) {
	for i := 0; i < len(nums)-2; i++ {
		if 2*(nums[i]+nums[i+2]) == nums[i+1] {
			result++
		}
	}
	return result
}
