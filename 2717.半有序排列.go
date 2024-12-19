func semiOrderedPermutation(nums []int) int {
	first, last := 0, 0
	n := len(nums)
	for i, num := range nums {
		if num == 1 {
			first = i
		}
		if num == n {
			last = i
		}
	}

	if first > last {
		return first + n - 1 - last - 1
	}
	return first + n - 1 - last
}
