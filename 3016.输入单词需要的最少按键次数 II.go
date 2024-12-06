func minimumPushes(word string) int {
	ans := make(map[byte]int)
	for i := range word {
		ans[word[i]]++
	}
	maps.DeleteFunc(ans, func(k byte, v int) bool { return v == 0 })

	letters := make([]int, 0, len(ans))
	for k := range maps.Values(ans) {
		letters = append(letters, k)
	}
	sort.Slice(letters, func(i, j int) bool {
		return letters[i] > letters[j]
	})
	reuslt := 0
	for i, k := range letters {
		reuslt += ((i + 8) / 8) * k
	}
	return reuslt
}
