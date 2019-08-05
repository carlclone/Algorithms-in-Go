package two_pointers_and_sliding_window

func twoSum(numbers []int, target int) []int {
	p1 := 0
	p2 := len(numbers) - 1

	for numbers[p1]+numbers[p2] != target {
		if numbers[p1]+numbers[p2] > target {
			p2--
			continue
		}
		if numbers[p1]+numbers[p2] < target {
			p1++
			continue
		}
	}
	return []int{p1 + 1, p2 + 1}
}
