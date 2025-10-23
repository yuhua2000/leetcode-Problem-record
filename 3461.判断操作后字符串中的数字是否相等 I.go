package leetcode

func hasSameDigits(s string) bool {
	ans := []byte(s)
	for len(ans) > 2 {
		for i := 0; i < len(ans)-1; i++ {
			ans[i] = byte((int(ans[i]-'0')+int(ans[i+1]-'0'))%10) + '0'
		}
		ans = ans[:len(ans)-1]
	}
	return ans[0] == ans[1]
}
