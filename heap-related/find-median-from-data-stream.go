package main

import "container/heap"

type MedianFinder struct {
	lowerHeap  *MinHeap
	higherHeap *MaxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	return MedianFinder{minHeap, maxHeap}
}

func (this *MedianFinder) AddNum(num int) {
	if this.lowerHeap.Peek() > num || this.lowerHeap.Len() == 0 {
		this.lowerHeap.Push(num)
	} else {
		this.higherHeap.Push(num)
	}

	//re-balance
	this.reBalance()
}

func (this *MedianFinder) reBalance() {
	var biggerHeap interface{}
	var smallerHeap interface{}
	if this.lowerHeap.Len() > this.higherHeap.Len() {
		biggerHeap = this.lowerHeap
		smallerHeap = this.higherHeap
	} else {
		biggerHeap = this.higherHeap
		smallerHeap = this.lowerHeap
	}
	for biggerHeap.Len()-smallerHeap.Len() >= 2 {
		peek := biggerHeap.Pop()
		smallerHeap.Push(peek)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.lowerHeap.Len() == this.higherHeap.Len() {

	} else {

	}
}

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() (val interface{}) {
	old := *h
	val = old[len(old)-1]
	*h = old[0 : len(old)-1]
	return
}
func (h *MinHeap) Peek() (val int) {
	old := *h
	val = old[len(old)-1]
	return
}

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() (val int) {
	old := *h
	val = old[len(old)-1]
	*h = old[0 : len(old)-1]
	return
}
func (h *MaxHeap) Peek() (val interface{}) {
	old := *h
	val = old[len(old)-1]
	return
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
