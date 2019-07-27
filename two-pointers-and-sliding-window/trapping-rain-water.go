package two_pointers_and_sliding_window

/*
https://leetcode.com/problems/trapping-rain-water/
*/

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	max := height[0]
	maxIndex := 0
	for k, v := range height {
		if v > max {
			max = v
			maxIndex = k
		}
	}

	filled := make([]int, len(height))
	max1 := height[0]

	for i := 0; i < maxIndex; i++ {
		if height[i] < max1 {
			filled[i] = max1
		} else {
			max1 = height[i]
			filled[i] = max1
		}
	}

	endIndex := len(height) - 1
	max2 := height[endIndex]
	for j := endIndex; j >= maxIndex; j-- {
		if height[j] < max2 {
			filled[j] = max2
		} else {
			max2 = height[j]
			filled[j] = max2
		}
	}

	sumFilled := 0
	for _, v := range filled {
		sumFilled += v
	}

	sumOrigin := 0
	for _, v := range height {
		sumOrigin += v
	}

	return sumFilled - sumOrigin

}
