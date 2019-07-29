package sort

/*
Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note: You are not suppose to use the library's sort function for this problem.

Example:

Input: [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
Follow up:

A rather straight forward solution is a two-pass algorithm using counting sort.
First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.
Could you come up with a one-pass algorithm using only constant space?
*/

//three way partition
func sortColors(nums []int) {
	j := -1
	i := j + 1
	k := len(nums) - 1 + 1

	for i <= k-1 {
		if nums[i] == 0 {
			tmp1 := nums[i]
			nums[i] = nums[j+1]
			nums[j+1] = tmp1
			i++
			j++
			continue
		}
		if nums[i] == 1 {
			i++
			continue
		}
		if nums[i] == 2 {
			k--
			tmp2 := nums[i]
			nums[i] = nums[k]
			nums[k] = tmp2
			continue
		}
	}
}

// 计数排序
func sortColors(nums []int) {
	//make a map store each color number
	//iterate get numbers
	//iterate with 0 1 2 order , fill in arr

	map1 := make(map[int]int)

	for _, v := range nums {
		map1[v]++
	}

	pointer := 0
	for _, v := range []int{0, 1, 2} {
		for i := 1; i <= map1[v]; i++ {
			nums[pointer] = v
			pointer++
		}
	}
}
