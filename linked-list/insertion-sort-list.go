package linked_list

//https://leetcode.com/problems/insertion-sort-list/

// 情况讨论 0 1 2
// 插入位置情况讨论 最前面 中间 最后面
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// list[head,..,sortedEnd] is sorted
	sortedEnd := head
	cur := sortedEnd.Next
	var insertAfter *ListNode
	var insertBefore *ListNode
	for cur != nil {
		// don't need to insert
		if cur.Val >= sortedEnd.Val {
			sortedEnd = cur
			cur = cur.Next
		} else { // insert into sorted list
			// expand the sorted list
			sortedEnd.Next = cur.Next
			// initial search
			insertBefore = head
			insertAfter = nil
			for insertBefore != sortedEnd.Next && insertBefore.Val <= cur.Val {
				insertAfter = insertBefore
				insertBefore = insertBefore.Next
			}
			// insert now!
			cur.Next = insertBefore
			// if the cur as the head of list, head must be change
			if insertAfter == nil {
				head = cur
			} else {
				insertAfter.Next = cur
			}
			// next element
			cur = sortedEnd.Next
		}
	}
	return head
}
