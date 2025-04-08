func minimumOperations(nums []int) int {
	exitsMap := make(map[int]struct{}, len(nums)/3)
	for i := len(nums) - 1; i >= 0; i-- {
		if _, ok := exitsMap[nums[i]]; ok {
			return (i + 1 + 2) / 3
		} else {
			exitsMap[nums[i]] = struct{}{}
		}
	}

	return 0
}
