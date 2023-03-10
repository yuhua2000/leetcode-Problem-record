package leetcode

/*
 * @lc app=leetcode.cn id=1590 lang=golang
 *
 * [1590] 使数组和能被 P 整除
 */

// @lc code=start
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minSubarray(nums []int, p int) int {
	n := len(nums)
	ans, sum := n, 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	rem := sum % p
	if rem == 0 {
		return 0
	}
	sum = 0
	m := make(map[int]int, n)
	m[0] = -1
	for i := 0; i < n; i++ {
		sum += nums[i]
		tRem := sum % p
		k := (tRem - rem + p) % p
		if _, ok := m[k]; ok {
			ans = min(ans, i-m[k])
		}
		m[tRem] = i
	}
	if ans == n {
		return -1
	}
	return ans
}

// @lc code=end

/*
sum % p =x
 
如果 sum-(fj-fi+1)==0
则

fi+1-fj %p = x
fj %d =(fi+1 -x) % d
如果 fi+1 % p =z
则 fj %d(z-x)%d
*/
