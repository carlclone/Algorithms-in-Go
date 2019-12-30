package main

func main() {
	s := [][]int{
		{1, 3},
		{1, 2},
	}
}

//quickSort
func sortArray(nums [][]int) [][]int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}
func quickSort(nums [][]int, start int, end int) {
	if start >= end { //warning
		return
	}

	p := partition2(nums, start, end)
	quickSort(nums, start, p-1)
	quickSort(nums, p+1, end)
}

func partition2(nums [][]int, l int, r int) int {
	var (
		j, i int
	)
	// [l+1,j] <v , [j+1,i)>v

	v := nums[l]
	j = l
	for i = l + 1; i <= r; i++ {
		if nums[i][0] < v[0] {
			if nums[i][1] > v[1] {
				swap(nums, i, j+1)
				j++
			}

		}

	}
	swap(nums, l, j)
	return j
}

func partition(nums []int, start int, end int) int {
	//           orange    yellowEnd             whiteStart
	// [orange = pivot ][yellow < v][purple > v][white]
	var (
		orange     int
		yellowEnd  int
		whiteStart int
		pivot      int
	)
	orange = start
	yellowEnd = start
	whiteStart = start + 1
	pivot = nums[orange]
	for whiteStart <= end {
		if nums[whiteStart] < pivot {
			swap(nums, yellowEnd+1, whiteStart)
			whiteStart++
			yellowEnd++
		} else if nums[whiteStart] >= pivot { //warning <= , not <
			whiteStart++
		}
	}
	swap(nums, orange, yellowEnd)
	return yellowEnd
}

func swap(nums [][]int, index1 int, index2 int) {
	tmp := nums[index1]
	nums[index1] = nums[index2]
	nums[index2] = tmp

}
