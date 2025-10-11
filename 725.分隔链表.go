package leetcode

/*
 * @lc app=leetcode.cn id=725 lang=golang
 *
 * [725] 分隔链表
 */

// @lc code=start

//Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func splitListToParts(head *ListNode, k int) []*ListNode {
	var count int

	for root := head; root != nil; root = root.Next {
		count++
	}
	segments := count / k
	remainder := count % k

	var res []*ListNode = make([]*ListNode, k)

	for i, root := 0, head; i < k && root != nil; i++ {
		res[i] = root
		for j := 1; j < segments; j++ {
			root = root.Next
		}
		if i < remainder {
			root = root.Next
		}
		root, root.Next = root.Next, nil
	}
	return res
}

// @lc code=end
