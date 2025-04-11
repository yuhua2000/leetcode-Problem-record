func countSymmetricIntegers(low int, high int) (res int) {
	for i := low; i <= high; i++ {
		if i < 100 && i%11 == 0 {
			res++
		} else if i > 1000 && i < 10000 {
			left := i/1000 + (i%1000)/100
			right := (i%100)/10 + i%10

			if left == right {
				res++
			}
		}
	}

	return
}
