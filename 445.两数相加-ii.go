/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
 */

// @lc code=start

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var stack1, stack2 []int
	for root := l1; root != nil; root = root.Next {
		stack1 = append(stack1, root.Val)
	}
	for root := l2; root != nil; root = root.Next {
		stack2 = append(stack2, root.Val)
	}

	carry := 0
	sum := 0
	var root *ListNode
	for i, j := len(stack1)-1, len(stack2)-1; i >= 0 || j >= 0; {
		if i >= 0 && j >= 0 {
			sum = stack1[i] + stack2[j] + carry
		} else if i >= 0 {
			sum = stack1[i] + carry
		} else if j >= 0 {
			sum = stack2[j] + carry
		}
		carry = sum / 10
		root = &ListNode{
			Next: root,
			Val:  sum % 10,
		}
		i--
		j--

	}
	if carry != 0 {
		root = &ListNode{
			Next: root,
			Val:  carry,
		}
	}
	return root
}

// @lc code=end
