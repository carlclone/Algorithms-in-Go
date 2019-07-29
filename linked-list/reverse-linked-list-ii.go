package linked_list

/*
Reverse a linked list from position m to n. Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.

Example:

Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL

https://leetcode.com/problems/reverse-linked-list-ii/
*/

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var pre *ListNode
	cur := head
	next := head.Next

	for next != nil {
		cur.Next = pre
		tmp := next.Next
		next.Next = cur

		pre = cur
		cur = next
		next = tmp
	}
	return cur
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head.Next == nil {
		return head
	}
	if m == n {
		return head

	}
	node := head
	counter := 1
	var p1 *ListNode
	var p2 *ListNode
	var p3 *ListNode
	var p4 *ListNode

	for node != nil {
		if counter < m-1 {
			node = node.Next
		}
		if counter == m-1 {
			p1 = node
			node = node.Next
			p1.Next = nil
		}
		if counter == m {
			p2 = node
			node = node.Next
		}
		if counter > m && counter < n {
			node = node.Next
		}
		if counter == n {
			p3 = node
			node = node.Next
			p3.Next = nil
		}
		if counter == n+1 {
			p4 = node
			node = node.Next
		}
		if counter > n+1 {
			break
		}

		counter++
	}

	reverseList(p2)
	if m == 1 {
		p2.Next = p4
		return p3
	} else {
		p1.Next = p3
		p2.Next = p4
		return head
	}

}
