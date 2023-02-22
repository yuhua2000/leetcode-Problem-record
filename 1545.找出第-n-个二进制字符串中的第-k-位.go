/*
 * @lc app=leetcode.cn id=1545 lang=golang
 *
 * [1545] 找出第 N 个二进制字符串中的第 K 位
 */

// @lc code=start
func findKthBit1(n int, k int) byte {
	s := bytes.NewBufferString("0")
	last := s.String()
	for i := 2; i <= n; i++ {
		s.Reset()
		s.WriteString(last)
		s.WriteString("1")
		for i := len(last) - 1; i >= 0; i-- {
			if last[i]-'0' == 0 {
				s.WriteString("1")
			} else {
				s.WriteString("0")
			}
		}
		last = s.String()
	}
	return last[k-1]
}


func findKthBit(n int, k int) byte {
	if k == 1 {
		return '0'
	}
	mid := 1 << (n - 1) 
	if k == mid {
		return '1'
	} else if k < mid {
		return findKthBit(n-1, k)
	} else if k > mid {
		k = mid*2 - k
		if findKthBit(n-1, k) == '1' {
			return '0'
		} else {
			return '1'
		}
	}
	return '0'
}

// @lc code=end

