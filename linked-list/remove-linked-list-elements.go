package linked_list

//Remove all elements from a linked list of integers that have value val.
//
//Example:
//
//Input:  1->2->6->3->4->5->6, val = 6
//Output: 1->2->3->4->5

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	dummyHead := ListNode{Val: -1}
	dummyHead.Next = head

	node := &dummyHead

	for node != nil {

		if node.Next != nil && node.Next.Val == val {
			before := node
			after := node.Next.Next

			node.Next.Next = nil
			before.Next = after
			node = before
		} else {
			node = node.Next
		}
	}
	return dummyHead.Next

}
