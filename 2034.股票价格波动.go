package leetcode

import "container/heap"

/*
 * @lc app=leetcode.cn id=2034 lang=golang
 *
 * [2034] 股票价格波动
 */

// @lc code=start
type StockPrice struct {
	maxPrice, minPrice hp
	timePriceMap       map[int]int
	maxTimestamp       int
}

func Constructor() StockPrice {
	return StockPrice{timePriceMap: make(map[int]int)}
}

func (this *StockPrice) Update(timestamp int, price int) {
	heap.Push(&this.maxPrice, pair{-price, timestamp})
	heap.Push(&this.minPrice, pair{price, timestamp})
	this.timePriceMap[timestamp] = price
	if timestamp > this.maxTimestamp {
		this.maxTimestamp = timestamp
	}
}

func (this *StockPrice) Current() int {
	return this.timePriceMap[this.maxTimestamp]
}

func (this *StockPrice) Maximum() int {
	for {
		if p := this.maxPrice[0]; -p.price == this.timePriceMap[p.timestamp] {
			return -p.price
		}
		heap.Pop(&this.maxPrice)
	}
}

func (this *StockPrice) Minimum() int {
	for {
		if p := this.minPrice[0]; p.price == this.timePriceMap[p.timestamp] {
			return p.price
		}
		heap.Pop(&this.minPrice)
	}
}

type pair struct{ price, timestamp int }

type hp []pair

func (h hp) Len() int { return len(h) }

func (h hp) Less(i, j int) bool { return h[i].price < h[j].price }

func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }

func (h *hp) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
// @lc code=end
