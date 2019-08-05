package data_structures

//Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
//
//push(x) -- Push element x onto stack.
//pop() -- Removes the element on top of the stack.
//top() -- Get the top element.
//getMin() -- Retrieve the minimum element in the stack.
//
//
//Example:
//
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> Returns -3.
//minStack.pop();
//minStack.top();      --> Returns 0.
//minStack.getMin();   --> Returns -2.

type MinStack struct {
	dataStack *Stack
	minStack  *Stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{dataStack: &Stack{data: make([]int, 9999), head: 0}, minStack: &Stack{data: make([]int, 9999), head: 0}}
}

func (this *MinStack) Push(x int) {
	this.dataStack.Push(x)
	if this.minStack.isEmpty() || x <= this.minStack.Top() {
		this.minStack.Push(x)
	}
}

func (this *MinStack) Pop() {
	v := this.dataStack.Top()
	if v == this.minStack.Top() {
		this.minStack.Pop()
	}
	this.dataStack.Pop()
}

func (this *MinStack) Top() int {
	return this.dataStack.Top()
}

func (this *MinStack) GetMin() int {
	return this.minStack.Top()
}

type Stack struct {
	data []int
	head int
}

func (this *Stack) Push(x int) {

	this.data[this.head] = x
	this.head++
}

func (this *Stack) Pop() {
	this.head--
}

func (this *Stack) isEmpty() bool {
	if this.head == 0 {
		return true

	}
	return false
}

func (this *Stack) Top() int {
	return this.data[this.head-1]
}
