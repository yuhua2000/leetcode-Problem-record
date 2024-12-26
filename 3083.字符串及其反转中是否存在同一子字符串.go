func isSubstringPresent1(s string) bool {
	m := make(map[string]bool)
	for i := 0; i < len(s)-1; i++ {
		m[s[i:i+2]] = true
	}
	sBytes := []byte(s)
	slices.Reverse(sBytes)
	s = string(sBytes)

	for i := 0; i < len(s)-1; i++ {
		if m[s[i:i+2]] {
			return true
		}
	}

	return false
}

func isSubstringPresent(s string) bool {
	h := make([]int, 26)
	for i := 0; i < len(s)-1; i++ {
		x, y := s[i]-'a', s[i+1]-'a'
		h[x] |= 1 << y
		if h[y]&(1<<x) != 0 {
			return true
		}
	}
	return false
}
