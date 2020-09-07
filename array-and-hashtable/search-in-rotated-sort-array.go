package array_and_hashtable

func search(nums []int, target int) int {
	//0+6=6 6/2=3 nums[3]=7  7>2 , 在第一个有序数组里, 4~7 >0 , 3+6 / 2 = 4 , nums[4]=0 , return 4;
	// 7 , 7>2 , 在第一个有序数组里, 4~7>3 , 3+6/2=4 , nums[4]=0 , 0<4 , 在第二个有序数组里 , 0~2<3 , 返回-1
	//先找到第一个有序数组的右边界, 然后再判断在哪个有序数组中, 然后再进行bs
	//case : 画图 , 看left right mid的各种落点

	b := findBoundary(nums)
	left := nums[0]
	right := nums[len(nums)-1]

	res := -1
	if target >= left && target <= b {
		res = findIn(nums, left, b, target)
	}
	if target > b && target <= right {
		res = findIn(nums, b+1, right, target)
	}
	return res
}

/*
   |
[4,5,6,7,0,1,2]
           |
[4,5,6,7,0,1,2]

*/
func findBoundary(nums []int) int {
	l := 0
	r := len(nums) - 1
	eIdx := len(nums) - 1
	if nums[r] >= nums[l] {
		return r
	}
	var mid int
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] > nums[r] && mid+1 <= eIdx && nums[mid+1] < nums[mid] {
			return mid
		}
		if nums[mid] > nums[l] && nums[mid] > nums[r] {
			l = mid + 1
		}
		if nums[mid] > nums[r] && nums[mid] < nums[l] {
			r = mid - 1
		}
	}
	return -1
}

// bs
func findIn(nums []int, l int, r int, target int) int {

	var mid int
	for l <= r {
		mid = (l + r) / 2

		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
		}
		if nums[mid] < target {
			l = mid + 1
		}
	}
	return -1
}
