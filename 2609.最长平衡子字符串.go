func findTheLongestBalancedSubstring(s string) (result int) {
	count0, count1 := 0, 0
	for i, c := range s {
		if c == '1' {
			count1++
			result = max(result, 2*min(count0, count1))
		} else if i == 0 || s[i-1] == '1' {
			count0, count1 = 1, 0
		} else {
			count0++
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
