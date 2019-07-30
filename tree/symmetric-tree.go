package tree

/*
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3


But the following [1,2,2,null,3,null,3] is not:

    1
   / \
  2   2
   \   \
   3    3


Note:
Bonus points if you could solve it both recursively and iteratively.
*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	p := root.Left
	q := root.Right

	res := isS(p, q)
	if res {
		return true
	}

	return false

}

func isS(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val {
		res1 := isS(p.Left, q.Right)
		res2 := isS(p.Right, q.Left)
		if res1 && res2 {
			return true
		}
	}

	return false
}
