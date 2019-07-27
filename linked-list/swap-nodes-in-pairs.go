package linked_list

/*
Given a linked list, swap every two adjacent nodes and return its head.

You may not modify the values in the list's nodes, only nodes itself may be changed.



Example:

Given 1->2->3->4, you should return the list as 2->1->4->3.
*/

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	p1 := head
	dummyHead := ListNode{Val: -1}
	dummyHead.Next = p1
	p2 := head.Next
	p4 := &dummyHead
	var p3 *ListNode

	for p1 != nil && p2 != nil {
		p3 = p2.Next
		p4.Next = p2
		p1.Next = p3
		p2.Next = p1

		p4 = p1
		p1 = p3
		if p3 == nil {
			p2 = nil
		} else {
			p2 = p3.Next
		}
	}

	return dummyHead.Next
}
