package linked_list

func isPalindrome(head *ListNode) bool {
	var (
		newHead *ListNode
		node2   *ListNode
		arr     []int
	)
	if head == nil {
		return true
	}
	newHead, arr = reverseList(head)
	node2 = newHead
	for _, v := range arr {
		if v != node2.Val {
			return false
		}
		node2 = node2.Next
	}
	return true
}

func reverseList(head *ListNode) (*ListNode, []int) {
	arr := []int{}
	if head == nil {
		return head, arr
	}
	var pre *ListNode
	cur := head
	next := head.Next

	for next != nil {
		arr = append(arr, cur.Val)
		cur.Next = pre
		tmp := next.Next
		next.Next = cur

		pre = cur
		cur = next
		next = tmp
	}
	return cur, arr
}
