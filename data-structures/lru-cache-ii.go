package data_structures

type Node struct {
	next *Node
	prev *Node
	key  int
	val  int
}

type LRUCache struct {
	table map[int]*Node
	head  *Node
	tail  *Node
	cap   int
	size  int
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return LRUCache{make(map[int]*Node), head, tail, capacity, 0}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.table[key]; ok {
		this.delNode(node)
		this.offerNode(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.table[key]; ok {
		node.val = value
		this.delNode(node)
		this.offerNode(node)
	} else {
		if this.size+1 > this.cap {
			//定义记错
			evict := this.tail.prev
			delete(this.table, evict.key)
			this.delNode(evict)
			this.size--
		}

		node := &Node{}
		node.key = key
		node.val = value
		this.table[key] = node
		this.offerNode(node)
		this.size++
	}
}

func (this *LRUCache) delNode(node *Node) {
	prev := node.prev
	next := node.next
	prev.next = next
	//遗漏维护前指针
	next.prev = prev
}

func (this *LRUCache) offerNode(node *Node) {
	next := this.head.next
	this.head.next = node
	node.next = next
	node.prev = this.head
	//遗漏维护前指针
	next.prev = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
