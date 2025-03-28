func minimizedStringLength(s string) int {
	m := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		m[s[i]] = true
	}

	return len(m)
}
