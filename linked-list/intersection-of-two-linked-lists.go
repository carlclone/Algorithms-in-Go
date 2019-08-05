package linked_list

//https://leetcode.com/problems/intersection-of-two-linked-lists/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	aLen := 1
	node := headA
	for node != nil {
		node = node.Next
		aLen++
	}

	bLen := 1
	node = headB
	for node != nil {
		node = node.Next
		bLen++
	}

	var node1 *ListNode
	var node2 *ListNode
	var gap int

	if aLen > bLen {
		gap = aLen - bLen
		node1 = headA
		for i := 1; i <= gap; i++ {
			node1 = node1.Next
		}
		node2 = headB
	} else {
		gap = bLen - aLen
		node1 = headB
		for i := 1; i <= gap; i++ {
			node1 = node1.Next
		}
		node2 = headA
	}

	for node1 != nil {
		if node1 == node2 {
			return node1
		}
		node1 = node1.Next
		node2 = node2.Next

	}

	return nil

}
