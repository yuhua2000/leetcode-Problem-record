package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=2080 lang=golang
 *
 * [2080] 区间内查询数字的频率
 */
// @lc code=start
type RangeFreqQuery struct {
	numIndex map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	m := make(map[int][]int, len(arr)/4)
	for i, num := range arr {
		m[num] = append(m[num], i)
	}
	return RangeFreqQuery{
		numIndex: m,
	}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	numIndex := this.numIndex[value]
	l := sort.SearchInts(numIndex, left)
	r := sort.SearchInts(numIndex, right)
	if r < len(numIndex) && numIndex[r] <= right {
		return r - l + 1
	}
	return r - l
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
// @lc code=end
