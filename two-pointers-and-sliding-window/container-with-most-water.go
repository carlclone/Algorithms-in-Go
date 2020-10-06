package two_pointers_and_sliding_window

/*
Example:

Input: [1,8,6,2,5,4,8,3,7]
Output: 49
*/

//解法 1 ,填坑大法 , 减去面积
func maxArea(height []int) int {
	maxArea := 0
	l := 0
	r := len(height) - 1

	curArea := 0

	for l < r {
		min := min(height[l], height[r])
		curArea = (r - l) * min
		maxArea = maxC(maxArea, curArea)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxArea
}

func maxArea(height []int) int {
	ans := 0
	l := 0
	r := len(height) - 1

	for l < r {
		ans = maxx(ans, (r-l)*minn(height[l], height[r]))

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

func maxx(l int, r int) int {
	if l > r {
		return l
	} else {
		return r
	}
}

func minn(l int, r int) int {
	if l > r {
		return r
	}
	return l
}
