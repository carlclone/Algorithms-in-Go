package array

func twoSum1(nums []int, target int) []int {
	numMap := make(map[int]int)
	for k, v := range nums {
		numMap[v] = k
	}
	for k, v := range nums {
		if _, ok := numMap[target-v]; ok && (k != numMap[target-v]) {
			return []int{k, numMap[target-v]}
		}
	}

	return []int{}

}

//改进 ， 只遍历一次
func twoSum2(nums []int, target int) []int {
	numMap := make(map[int]int)

	for k, v := range nums {
		if _, ok := numMap[target-v]; ok {
			return []int{k, numMap[target-v]}
		} else {
			numMap[v] = k
		}
	}

	return []int{}
}
