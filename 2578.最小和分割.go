func splitNum(num int) int {
	nums := []int{}
	for num > 0 {
		nums = append(nums, num%10)
		num = num / 10
	}
	sort.Ints(nums)
	num1, num2 := 0, 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			num1 = num1*10 + nums[i]
		} else {
			num2 = num2*10 + nums[i]
		}
	}
	return num1 + num2
}
