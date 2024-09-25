package leetcode

func maximumSubsequenceCount(text string, pattern string) int64 {
	var res, cnt1, cnt2 int64
	for i := range text {
		if text[i] == pattern[1] {
			res += cnt1
			cnt2++
		}
		if text[i] == pattern[0] {
			cnt1++
		}
	}
	return res + max(cnt1, cnt2)
}
