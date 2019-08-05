package data_structures

//Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.
//
//get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
//put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.
//
//The cache is initialized with a positive capacity.
//
//Follow up:
//Could you do both operations in O(1) time complexity?
//
//Example:
//
//LRUCache cache = new LRUCache( 2 /* capacity */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // returns 1
//cache.put(3, 3);    // evicts key 2
//cache.get(2);       // returns -1 (not found)
//cache.put(4, 4);    // evicts key 1
//cache.get(1);       // returns -1 (not found)
//cache.get(3);       // returns 3
//cache.get(4);       // returns 4

type DLinkedNode struct {
	Key  int
	Val  int
	Next *DLinkedNode
	Last *DLinkedNode
}

type LRUCache struct {
	DMap     map[int]*DLinkedNode
	Count    int
	Capacity int
	Head     *DLinkedNode //dummy head
	Tail     *DLinkedNode //dummy tail
}

func Constructor(capacity int) LRUCache {
	head := DLinkedNode{}
	tail := DLinkedNode{}
	head.Next = &tail
	tail.Last = &head

	return LRUCache{DMap: make(map[int]*DLinkedNode), Count: 0, Capacity: capacity, Head: &head, Tail: &tail}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.DMap[key]; ok {
		node := this.DMap[key]
		this.moveToHead(node)
		return node.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.DMap[key]; ok {
		node := this.DMap[key]
		node.Val = value
		this.moveToHead(node)
	} else {
		if this.Count >= this.Capacity {
			delete(this.DMap, this.Tail.Last.Key)
			this.removeNode(this.Tail.Last)

			new := DLinkedNode{Val: value, Key: key}
			this.addNode(&new)
			this.DMap[key] = &new
		} else {
			new := DLinkedNode{Val: value, Key: key}
			this.addNode(&new)
			this.DMap[key] = &new
			this.Count++
		}
	}

}

func (t *LRUCache) moveToHead(node *DLinkedNode) {
	t.removeNode(node)
	t.addNode(node)
}

func (t *LRUCache) removeNode(node *DLinkedNode) {
	next := node.Next
	last := node.Last

	last.Next = next
	next.Last = last
}

func (t *LRUCache) addNode(node *DLinkedNode) {
	headNext := t.Head.Next
	t.Head.Next = node
	node.Last = t.Head
	node.Next = headNext
	headNext.Last = node
}
