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

//TODO;堆解法

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
