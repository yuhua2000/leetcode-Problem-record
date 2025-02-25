/*
 * @lc app=leetcode.cn id=2502 lang=golang
 *
 * [2502] 设计内存分配器
 */
// @lc code=start
type Allocator struct {
	memory []int
}

func Constructor(n int) Allocator {
	return Allocator{memory: make([]int, n)}
}

func (this *Allocator) Allocate(size int, mID int) int {
	ans := 0
	for i := range this.memory {
		if this.memory[i] == 0 {
			ans++
		} else {
			ans = 0
		}

		if ans == size {
			for j := 0; j < size; j++ {
				this.memory[i-j] = mID
			}
			return i - size + 1
		}
	}

	return -1
}

func (this *Allocator) FreeMemory(mID int) (result int) {
	for i := range this.memory {
		if this.memory[i] == mID {
			result++
			this.memory[i] = 0
		}
	}

	return
}

/**
 * Your Allocator object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Allocate(size,mID);
 * param_2 := obj.FreeMemory(mID);
 */
// @lc code=end
