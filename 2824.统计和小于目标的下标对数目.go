func countPairs(nums []int, target int) (result int) {
	sort.Ints(nums)
	for i, j := 0, len(nums)-1; i < j; i++ {
		for i < j && nums[i]+nums[j] >= target {
			j--
		}
		result += j - i
	}
    return
}
