/*
 * @lc app=leetcode.cn id=1361 lang=golang
 *
 * [1361] 验证二叉树
 */

// @lc code=start
func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	var indeg = make([]int, n)
	for i := 0; i < n; i++ {
		if leftChild[i] != -1 {
			indeg[leftChild[i]] += 1
		}
		if rightChild[i] != -1 {
			indeg[rightChild[i]] += 1
		}
	}
	root := -1
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			root = i
			break
		}
	}
	if root == -1 {
		return false
	}
	vis := map[int]bool{root: true}
	quene := []int{root}
	for len(quene) > 0 {
		u := quene[0]
		quene = quene[1:]
		if leftChild[u] != -1 {
			if vis[leftChild[u]] {
				return false
			}
			vis[leftChild[u]] = true
			quene = append(quene, leftChild[u])
		}
		if rightChild[u] != -1 {
			if vis[rightChild[u]] {
				return false
			}
			vis[rightChild[u]] = true
			quene = append(quene, rightChild[u])
		}
	}
	return len(vis) == n
}

// @lc code=end

