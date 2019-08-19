package tree

// using recursion
func genBST(nums []int) (root *TreeNode) {
	var (
		val int
	)

	for _, val = range nums {
		root = _genBST(root, val)
	}

	return
}

func _genBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
		return root
	}
	if val < root.Val {
		root.Left = _genBST(root.Left, val)
	}
	if val > root.Val {
		root.Right = _genBST(root.Right, val)
	}
	return root
}

//using iteration

//test
func main() {

}
