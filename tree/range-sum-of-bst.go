package tree

func rangeSumBST(root *TreeNode, L int, R int) int {
	//中序遍历 , 遍历的时候如果>= or <=  则add
	sum := 0
	if root == nil {
		return sum
	}
	midOrderIterate(root, L, R, &sum)
	return sum
}

func midOrderIterate(node *TreeNode, L int, R int, sum *int) {
	if node.Val >= L && node.Val <= R {
		*sum += node.Val
	}
	if node.Left != nil {
		midOrderIterate(node.Left, L, R, sum)
	}
	if node.Right != nil {
		midOrderIterate(node.Right, L, R, sum)
	}
}
