package array_and_hashtable

//Given two arrays, write a function to compute their intersection.
//
//Example 1:
//
//Input: nums1 = [1,2,2,1], nums2 = [2,2]
//Output: [2]
//Example 2:
//
//Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//Output: [9,4]
//Note:
//
//Each element in the result must be unique.
//The result can be in any order.

func intersection(nums1 []int, nums2 []int) []int {
	/*
		make a first_set
		make a second_set
		put nums1 all into first_set
		forloop nums2
			if can find nums2[i] in first_set
			put into second set

		return second set
	*/
	fset := make(map[int]int)
	sset := make(map[int]int)
	for _, v := range nums1 {
		fset[v] = 1
	}
	for _, v := range nums2 {
		if _, ok := fset[v]; ok {
			sset[v] = 1
		}
	}
	i := 0
	var result []int
	for k := range sset {
		result = append(result, k)
		i++
	}
	return result
}
