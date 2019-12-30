package tree

//写的时候Wrong了很多次, 原因是直接在原来的代码上改,但函数的定义却已经发生了变化,下次直接推翻重写
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	insertInto(root, val)
	return root
}
func insertInto(root *TreeNode, val int) {
	if val < root.Val && root.Left != nil {
		insertInto(root.Left, val)
	}
	if val > root.Val && root.Right != nil {
		insertInto(root.Right, val)
	}
	if val < root.Val && root.Left == nil {
		root.Left = &TreeNode{Val: val}
	}

	if val > root.Val && root.Right == nil {
		root.Right = &TreeNode{Val: val}
	}
}

// 另一种常用到的技巧,插入某棵树并返回树的根节点,重新赋值 , 在堆结构里也有相似的案例
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
