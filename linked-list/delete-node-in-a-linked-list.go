package linked_list

//https://leetcode.com/problems/delete-node-in-a-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	/*
			if p1.next!=nil
		      p1.val=p1.next.val
		      p1.next=nil
		    else
		      p1=nil
	*/
	if node.Next != nil {
		node.Val = node.Next.Val
		node.Next = node.Next.Next
	} else {
		node = nil
	}
}
