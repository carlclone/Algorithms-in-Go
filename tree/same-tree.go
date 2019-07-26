package tree

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val {
		res1 := isSameTree(p.Left, q.Left)
		res2 := isSameTree(p.Right, q.Right)
		if res1 && res2 {
			return true
		}
	}

	return false

}
