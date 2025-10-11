package leetcode

/*
 * @lc app=leetcode.cn id=1656 lang=golang
 *
 * [1656] 设计有序流
 */

// @lc code=start
type OrderedStream struct {
	value []string
	ptr   int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		value: make([]string, n+1),
		ptr:   1,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) (result []string) {
	this.value[idKey] = value

	for i := this.ptr; i < len(this.value) && this.value[i] != ""; i++ {
		result = append(result, this.value[i])
	}

	this.ptr += len(result)
	return
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
// @lc code=end
