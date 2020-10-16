package heap_related

//Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.
//
//Example 1:
//
//Input: [3,2,1,5,6,4] and k = 2
//Output: 5
//Example 2:
//
//Input: [3,2,3,1,2,4,5,5,6] and k = 4
//Output: 4
//Note:
//You may assume k is always valid, 1 ≤ k ≤ array's length.

//堆解法
import (
	"container/heap"
	"fmt"
)

// min heap https://play.golang.org/p/rJvJCqw73ql
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() (val interface{}) {
	old := *h
	val = old[len(old)-1]
	*h = old[0 : len(old)-1]
	return
}

func findKthLargestHeap(nums []int, k int) int {

	if len(nums) < k {
		fmt.Println("in edge case")
		return 0
	}

	h := &IntHeap{}
	heap.Init(h)

	for index := 0; index < k; index++ {
		heap.Push(h, nums[index])
	}

	fmt.Println("intial h ", h)

	for i := k; i < len(nums); i++ {
		if nums[i] < (*h)[0] {
			continue
		}
		heap.Pop(h)
		heap.Push(h, nums[i])
	}

	return (*h)[0]
}

//快排分区解法nlogn
func partition(nums []int, start int, end int) int {
	p := start
	j := start + 1
	i := start

	v := nums[p]

	for j <= end {
		if nums[j] < v {
			tmp1 := nums[j]
			nums[j] = nums[i+1]
			nums[i+1] = tmp1

			i++
			j++
		} else if nums[j] >= v {
			j++
		}
	}

	tmp2 := nums[i]
	nums[i] = nums[p]
	nums[p] = tmp2

	p = i
	return p
}

func findKthLargest(nums []int, k int) int {

	start := 0
	end := len(nums) - 1
	return findKthLargestIn(nums, start, end, end+1-k-1+1) //第k大的数=len-k+1大的数

}

func findKthLargestIn(nums []int, start int, end int, n int) int {
	k := partition(nums, start, end)

	if k == n {
		return nums[k]
	}
	if k < n {
		return findKthLargestIn(nums, k+1, end, n)
	}
	if k > n {
		return findKthLargestIn(nums, start, k-1, n)
	}
	return 0
}

// heapify 解法

func siftUp(nums []int, i, end int) {
	l := 2*i + 1
	r := 2*i + 2
	largest := i
	if l <= end && nums[largest] < nums[l] {
		largest = l
	}
	if r <= end && nums[largest] < nums[r] {
		largest = r
	}
	if i != largest {
		nums[i], nums[largest] = nums[largest], nums[i]
		siftUp(nums, largest, end)
	}
}

func heapify(nums []int) {
	// first non-leaf node
	base := len(nums)/2 - 1
	for i := base; i >= 0; i-- {
		siftUp(nums, i, len(nums)-1)
	}
}

func findKthLargestHeapify(nums []int, k int) int {
	heapify(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		k--
		if k == 0 {
			return nums[0]
		}
		nums[0], nums[i] = nums[i], nums[0]
		siftUp(nums, 0, i-1)
	}
	return nums[0]
}
