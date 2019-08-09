package heap_related

//Given a non-empty array of integers, return the k most frequent elements.
//
//Example 1:
//
//Input: nums = [1,1,1,2,2,3], k = 2
//Output: [1,2]
//Example 2:
//
//Input: nums = [1], k = 1
//Output: [1]
//Note:
//
//You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
//Your algorithm's time complexity must be better than O(n log n), where n is the array's size.

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

//解法2
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	h := Constructor()

	for _, v := range nums {
		m[v]++
	}

	for k, v := range m {
		h.add([]int{k, v})
	}

	res := []int{}
	for i := 1; i <= k; i++ {
		res = append(res, h.extractMax()[0])
	}

	return res
}

type MaxHeap struct {
	data [][]int
	size int
}

func Constructor() *MaxHeap {
	return &MaxHeap{}
}

func (t *MaxHeap) sizes() int {
	return t.size
}

func (t *MaxHeap) isEmpty() bool {
	return t.size == 0
}

func (t *MaxHeap) parent(index int) int {
	if index == 0 {
		panic("index 0 dont have parent")
	}
	return (index - 1) / 2
}

func (t *MaxHeap) leftChild(index int) int {
	return 2*index + 1
}

func (t *MaxHeap) rightChild(index int) int {
	return 2*index + 2
}

func (t *MaxHeap) add(v []int) {
	t.data = append(t.data, v)
	t.size++
	newValueIndex := t.size - 1

	if newValueIndex != 0 {
		t.shiftUp(newValueIndex)
	}
}

func (t *MaxHeap) shiftUp(index int) {

	parent := t.parent(index)

	for parent >= 0 && t.data[index][1] > t.data[parent][1] {
		tmp := t.data[index]
		t.data[index] = t.data[parent]
		t.data[parent] = tmp

		index = parent

		if parent == 0 {
			break
		}
		parent = t.parent(index)
	}
}

func (t *MaxHeap) extractMax() []int {
	max := t.data[0]

	t.data[0] = t.data[t.size-1]
	t.size--
	t.data = t.data[:t.size] //移除最后一个元素,其实也可以不用删吧
	t.shiftDown(0)
	return max
}

func (t *MaxHeap) shiftDown(index int) {

	for t.leftChild(index) <= t.size-1 || t.rightChild(index) <= t.size-1 {
		leftIndex := t.leftChild(index)
		rightIndex := t.rightChild(index)

		var bV []int
		var bI int
		if t.rightChild(index) > t.size-1 {
			bV = t.data[leftIndex]
			bI = leftIndex
		} else if t.data[leftIndex][1] > t.data[rightIndex][1] {
			bV = t.data[leftIndex]
			bI = leftIndex
		} else {
			bV = t.data[rightIndex]
			bI = rightIndex
		}
		if t.data[index][1] > bV[1] {
			break
		}

		tmp := t.data[index]
		t.data[index] = bV
		t.data[bI] = tmp
		index = bI
	}

}

//解法1
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	res := []int{}
	for i := 1; i <= k; i++ {
		maxK := INT_MIN
		indexK := INT_MIN

		for k, v := range m {
			if v > maxK {
				maxK = v
				indexK = k
			}
		}
		res = append(res, indexK)
		delete(m, indexK)
	}
	return res
}
