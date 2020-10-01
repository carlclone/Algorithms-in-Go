package point_to_offer

// 1

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	type void struct{}
	var emptyValue void
	fatherMap := make(map[int]*TreeNode)
	//用map实现set，map的值为struct{}则不额外占用空间
	visit := make(map[int]void)
	var dfs func(*TreeNode)
	// 构造子节点指向祖先节点的指针
	dfs = func(root *TreeNode) {
		if root.Left != nil {
			fatherMap[root.Left.Val] = root
			dfs(root.Left)
		}
		if root.Right != nil {
			fatherMap[root.Right.Val] = root
			dfs(root.Right)
		}
	}
	dfs(root)
	// 从p节点开始不断往它的祖先移动，并用数据结构记录已经访问过的祖先节点。
	for p != nil {
		visit[p.Val] = emptyValue
		p = fatherMap[p.Val]
	}
	// 从q节点开始不断往它的祖先移动，如果有祖先已经被访问过，即意味着这是 p 和 q 的深度最深的公共祖先，即 LCA 节点。
	for q != nil {
		if _, exist := visit[q.Val]; exist {
			return q
		}
		q = fatherMap[q.Val]
	}
	return nil
}

//2

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}
