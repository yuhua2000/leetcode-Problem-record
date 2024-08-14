package leetcode

import (
	"bytes"
)

func isArraySpecial(nums []int, queries [][]int) []bool {
	n := len(nums)
	numParities := make([]byte, n)
	specialBytes := make([]byte, n)
	for i := range nums {
		numParities[i] = byte(nums[i] % 2)
		specialBytes[i] = byte(i % 2)
	}

	evenPrefix := []byte{0, 1}
	notEvenPrefix := []byte{1, 0}
	result := make([]bool, len(queries))
	for i, query := range queries {
		subNumParities := numParities[query[0] : query[1]+1]
		if len(subNumParities) == 1 {
			result[i] = true
		} else if bytes.HasPrefix(subNumParities, evenPrefix) {
			result[i] = bytes.HasPrefix(specialBytes, subNumParities)
		} else if bytes.HasPrefix(subNumParities, notEvenPrefix) {
			result[i] = bytes.HasPrefix(specialBytes, subNumParities[1:])
		}
	}
	return result
}
