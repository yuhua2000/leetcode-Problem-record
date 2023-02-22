/*
 * @lc app=leetcode.cn id=2341 lang=golang
 *
 * [2341] 数组能形成多少数对
 */

// @lc code=start
func numberOfPairs(nums []int) []int {
	// var m =make(map[int]int,len(nums)/2)
	//   res := 0
	// for	_,v:=range nums{
	// 	m[v]=m[v]+1
	// 	if m[v]==2{
	// 		res++
	// 		m[v]=0
	// 	}
	// }
	var cnt=make(map[int]bool,len(nums)/2)
	res:=0
	  for _, num := range nums {
        cnt[num] = !cnt[num]
        if !cnt[num] {
            res++
        }
    }

	return []int{res,len(nums)-res*2}
}
// @lc code=end

