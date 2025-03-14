func isBalanced(num string) bool {
	sum := 0
	for i, digit := range num {
		value := int(digit - '0')
		if i%2 == 0 {
			sum += value
		} else {
			sum -= value
		}
	}

	return sum == 0
}
