func countSeniors(details []string) int {
	result := 0
	for _, d := range details {
		if d[11:13] > "60" {
			result++
		}
	}
	return result
}
