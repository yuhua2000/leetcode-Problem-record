func vowelStrings(words []string, left int, right int) (result int) {
	vowel := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	for i := left; i <= right; i++ {
		if vowel[words[i][0]] && vowel[words[i][len(words[i])-1]] {
			result++
		}
	}
	return result
}
