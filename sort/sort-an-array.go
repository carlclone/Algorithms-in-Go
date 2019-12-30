package sort

//Given an array of integers nums, sort the array in ascending order.
//
//
//
//Example 1:
//
//Input: [5,2,3,1]
//Output: [1,2,3,5]
//Example 2:
//
//Input: [5,1,1,2,0,0]
//Output: [0,0,1,1,2,5]
//
//
//Note:
//
//1 <= A.length <= 10000
//-50000 <= A[i] <= 50000

// bubble
func sortArray(nums []int) []int {
	n := len(nums) - 1
	for i := 0; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if nums[j-1] > nums[j] {
				tmp := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = tmp
			}
		}
	}
	return nums
}

//bubble optimization
func sortArrayBubble(nums []int) []int {
	var (
		endIndex     int
		index        int
		iterateCount int
		tmp          int
	)
	endIndex = len(nums) - 1

	if endIndex == 0 {
		return nums
	}

	for index = 0; iterateCount <= endIndex; iterateCount++ {
		for ; index <= endIndex-iterateCount-1; index++ {
			if nums[index] > nums[index+1] {
				tmp = nums[index]
				nums[index] = nums[index+1]
				nums[index+1] = tmp
			}
		}
	}
	return nums
}

// select
func sortArray(nums []int) []int {
	n := len(nums) - 1
	for i := 0; i <= n; i++ {
		max := nums[0]
		index := 0
		limit := n - i
		for j := 0; j <= limit; j++ {
			if nums[j] > max {
				max = nums[j]
				index = j
			}
		}
		tmp := nums[index]
		nums[index] = nums[limit]
		nums[limit] = tmp
	}

	return nums
}

//insert
func sortArray(nums []int) []int {
	n := len(nums) - 1
	for i := 1; i <= n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[j] > nums[j+1] {
				tmp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = tmp
			}
		}
	}
	return nums
}

// merge

func sortArray(nums []int) []int {
	endIndex := len(nums) - 1
	return mergeSort(nums, 0, endIndex)
}

func mergeSort(nums []int, i1 int, i2 int) []int {
	if i1 == i2 {
		return []int{nums[i1]}
	}
	if i1 > i2 {
		return []int{}
	}

	//return merge(mergeSort(nums, i1, i2/2), mergeSort(nums, i2/2+1 , i2))

	mid := (i1 + i2) / 2
	return merge(mergeSort(nums, i1, mid), mergeSort(nums, mid+1, i2))
}

func merge(sort1 []int, sort2 []int) []int {
	newArr := make([]int, len(sort1)+len(sort2))

	ls := 0
	le := len(sort1) - 1

	rs := 0
	re := len(sort2) - 1

	np := 0

	for ls <= le && rs <= re {
		if sort2[rs] < sort1[ls] {
			newArr[np] = sort2[rs]
			rs++
		} else {
			newArr[np] = sort1[ls]
			ls++
		}
		np++
	}

	for ls <= le {
		newArr[np] = sort1[ls]
		ls++
		np++
	}

	for rs <= re {
		newArr[np] = sort2[rs]
		rs++
		np++
	}
	return newArr
}

//quickSort
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}
func quickSort(nums []int, start int, end int) {
	if start >= end { //warning
		return
	}

	p := partition2(nums, start, end)
	quickSort(nums, start, p-1)
	quickSort(nums, p+1, end)
}

func partition2(nums []int, l int, r int) int {
	var (
		j, i int
	)
	// [l+1,j] <v , [j+1,i)>v

	v := nums[l]
	j = l
	for i = l + 1; i <= r; i++ {
		if nums[i] < v {
			swap(nums, i, j)
			j++
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

func swap(nums []int, index1 int, index2 int) {
	tmp := nums[index1]
	nums[index1] = nums[index2]
	nums[index2] = tmp

}

//TODO; quick sort optimization
func sortArrayQuick(nums []int) []int {

}
