package point_to_offer

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		if num%2 == 1 {
			count++
		}
		num /= 2
	}
	return count
}

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {

		count++

		num = num & (num - 1)

	}
	return count
}
