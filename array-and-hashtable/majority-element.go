package array_and_hashtable

//Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.
//
//You may assume that the array is non-empty and the majority element always exist in the array.
//
//Example 1:
//
//Input: [3,2,3]
//Output: 3
//Example 2:
//
//Input: [2,2,1,1,1,2,2]
//Output: 2

func majorityElement(nums []int) int {
	var (
		freqMap  map[int]int
		majority int
		maxFreq  int
		num      int
		freq     int
	)
	freqMap = make(map[int]int)

	for _, num = range nums {
		freqMap[num]++
	}
	for num, freq = range freqMap {
		if freq > maxFreq {
			maxFreq = freq
			majority = num
		}
	}
	return majority
}
