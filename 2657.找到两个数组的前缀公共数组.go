package leetcode

func findThePrefixCommonArray(A []int, B []int) []int {
	seen := make(map[int]bool)

	C := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		seen[A[i]] = true
		seen[B[i]] = true
		C[i] = 2*(i+1) - len(seen)
	}

	return C
}
