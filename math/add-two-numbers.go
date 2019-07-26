package math

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	guard := ListNode{Val: 0}
	sum := 0
	d := &guard

	for l1 != nil || l2 != nil {
		sum /= 10
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		d.Next = &ListNode{Val: sum % 10}
		d = d.Next
	}

	if sum/10 == 1 {
		d.Next = &ListNode{Val: 1}
	}
	return guard.Next
}

//第二次写
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	dummy := &ListNode{}
	current := dummy

	currentCarry := 0

	for l1 != nil || l2 != nil {
		if l1 != nil {
			currentCarry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			currentCarry += l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: currentCarry % 10}
		current = current.Next
		currentCarry /= 10
	}

	if currentCarry == 1 {
		current.Next = &ListNode{Val: 1}
	}

	return dummy.Next
}
