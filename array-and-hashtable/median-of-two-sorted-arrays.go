package array_and_hashtable

//There are two sorted arrays nums1 and nums2 of size m and n respectively.
//
//Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
//
//You may assume nums1 and nums2 cannot be both empty.
//
//Example 1:
//
//nums1 = [1, 3]
//nums2 = [2]
//
//The median is 2.0
//Example 2:
//
//nums1 = [1, 2]
//nums2 = [3, 4]
//
//The median is (2 + 3)/2 = 2.5
//https://leetcode.com/problems/median-of-two-sorted-arrays/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//思路 , 合并后取中间数返回
	//复杂度分析 , 合并操作 每个元素访问一次 , O(n+m)
	//合并操作情况讨论
	//两个指针,分别指向两个数组
	//边界讨论 , 两个指针都等于他们对应数组的长度 (遍历完全部元素)

	p1, p2 := 0, 0
	len1, len2 := len(nums1), len(nums2)
	combined := []int{}

	for p1 < len1 && p2 < len2 {
		if nums1[p1] > nums2[p2] {
			combined = append(combined, nums2[p2])
			p2++
		} else {
			combined = append(combined, nums1[p1])
			p1++
		}
	}

	for p1 < len1 {
		combined = append(combined, nums1[p1])
		p1++
	}

	for p2 < len2 {
		combined = append(combined, nums2[p2])
		p2++
	}

	if len(combined)%2 == 0 {
		return float64((combined[len(combined)/2] + combined[len(combined)/2-1])) / 2.0
	} else {
		return float64(combined[len(combined)/2])
	}
}
