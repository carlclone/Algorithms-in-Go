package tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	left := &TreeNode{2, nil, nil}
	right := &TreeNode{2, nil, nil}
	root := &TreeNode{1, left, right}
	fmt.Println(levelOrder(root))
}

/**
 *
 * @param root TreeNode类
 * @return int整型二维数组
 */
func levelOrder(root *TreeNode) [][]int {
	// write code here
	if root == nil {
		return [][]int{}
	}
	queue := &Queue{}
	queue.Push(root)
	res := [][]int{}
	levelTable := make(map[*TreeNode]int)
	levelTable[root] = 0
	res = append(res, []int{})
	for queue.Size() != 0 {
		node := queue.Pop()
		newLevel := levelTable[node] + 1
		if node.Left != nil {
			levelTable[node.Left] = newLevel
			queue.Push(node.Left)
		}
		if node.Right != nil {
			levelTable[node.Right] = newLevel
			queue.Push(node.Right)
		}
		// 当有新的层 且尚未分配时,分配空间
		if (node.Left != nil || node.Right != nil) && len(res)-1 < newLevel {
			res = append(res, []int{})
		}
		level := levelTable[node]

		res[level] = append(res[level], node.Val)
	}
	return res
}

type Queue struct {
	list []*TreeNode
}

func (t *Queue) Push(node *TreeNode) {
	t.list = append(t.list, node)
}

func (t *Queue) Pop() *TreeNode {
	if len(t.list) == 0 {
		panic("queue is empty")
	}

	res := t.list[0]
	t.list = t.list[1:]
	return res
}

func (t *Queue) Size() int {
	return len(t.list)
}
