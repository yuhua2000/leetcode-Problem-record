package leetcode

func numberOfBeams(bank []string) int {

	prveFacilities := 0

	result := 0
	for i := 0; i < len(bank); i++ {
		count := 0
		for _, b := range bank[i] {
			if b == '1' {
				count++
			}
		}
		if count != 0 {
			result += count * prveFacilities
			prveFacilities = count
		}
	}
	return result
}
