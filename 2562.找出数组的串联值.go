func findTheArrayConcVal(nums []int) (result int64) {
	n := len(nums)
	for i, j := 0, n-1; i <= j; {
		if i == j {
			result += int64(nums[i])
		} else {
			num, _ := strconv.Atoi(strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]))
			result += int64(num)
		}
		j--
		i++
	}
	return result
}
