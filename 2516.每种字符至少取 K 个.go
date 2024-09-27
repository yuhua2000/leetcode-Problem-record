package leetcode

func min(...int) int

func takeCharacters(s string, k int) int {
	n := len(s)
	prefixSum := [3]int{}

	for i := 0; i < len(s); i++ {
		prefixSum[s[i]-'a']++
	}

	if min(prefixSum[0], prefixSum[1], prefixSum[2]) < k {
		return -1
	}

	result := n
	l := 0
	for r := 0; r < n; r++ {
		prefixSum[s[r]-'a']--
		for l < r && (prefixSum[0] < k || prefixSum[1] < k || prefixSum[2] < k) {
			prefixSum[s[l]-'a']++
			l++
		}

		if min(prefixSum[0], prefixSum[1], prefixSum[2]) >= k {
			result = min(result, n-(r-l+1))
		}
	}

	return result
}
