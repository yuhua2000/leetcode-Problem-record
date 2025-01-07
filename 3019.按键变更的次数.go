func countKeyChanges(s string) int {
	result := 0
	for i := 1; i < len(s); i++ {
		if s[i] > 'Z' {
			if s[i] != s[i-1] && s[i]-32 != s[i-1] {
				result++
			}
		} else {
			if s[i] != s[i-1] && s[i]+32 != s[i-1] {
				result++
			}
		}

	}

	return result
}
