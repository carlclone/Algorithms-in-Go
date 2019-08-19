package tree

func bfsIteration(root *TreeNode) []int {
	var (
		queue []*TreeNode
		pop   *TreeNode
		res   []int
	)

	queue = append(queue, root)

	for len(queue) != 0 {
		pop = queue[0]
		queue = queue[1:]

		res = append(res, pop.Val)
		if pop.Left != nil {
			queue = append(queue, pop.Left)
		}
		if pop.Right != nil {
			queue = append(queue, pop.Right)
		}
	}

	return res
}
