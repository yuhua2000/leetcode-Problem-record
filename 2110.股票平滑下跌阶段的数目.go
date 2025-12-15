package leetcode

func getDescentPeriods(prices []int) int64 {
	prePos := 0
	var result int64
	for i := 1; i < len(prices); i++ {
		if prices[i] != prices[i-1]-1 {
			length := i - prePos
			prePos = i
			if length%2 == 0 {
				result += int64((length + 1) * length / 2)
			} else {
				result += int64((length) * (length + 1) / 2)
			}
		}
	}

	length := len(prices) - prePos

	if length > 0 && length%2 == 0 {
		result += int64((length + 1) * length / 2)
	} else if length > 0 {
		result += int64((length) * (length + 1) / 2)
	}

	return result
}
