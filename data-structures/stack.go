package data_structures

type StackNode struct {
	data [1000]*TreeNode
	pos  int
}

func NewStackNode() *StackNode {
	return &StackNode{}
}
func (s *StackNode) IsEmpty() bool {
	return s.pos == 0
}
func (s *StackNode) Push(n *TreeNode) {
	s.data[s.pos] = n
	s.pos++
}
func (s *StackNode) Pop() (n *TreeNode) {
	s.pos--
	return s.data[s.pos]
}

//stack implementation in go slice
type SliceStack struct {
}

//recursion stack emulation

type Command struct {
	Type string // go , print
	Val  *TreeNode
}

type Stack struct {
	data [1000]*Command
	pos  int
}

func (s *Stack) pop() *Command {
	s.pos--
	return s.data[s.pos]

}

func (s *Stack) push(t *Command) {
	s.data[s.pos] = t
	s.pos++
}

func (s *Stack) isEmpty() bool {
	return s.pos == 0
}

func NewStack() *Stack {
	return &Stack{}
}
