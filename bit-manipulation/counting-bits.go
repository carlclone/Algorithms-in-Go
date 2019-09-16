package bit_manipulation

func countBits(num int) []int {
	res := []int{}
	for i := 0; i <= num; i++ {
		res = append(res, hammingWeight(i))
	}
	return res
}

func hammingWeight(num int) int {
	var (
		res int
	)
	for num != 0 {
		if num%2 == 1 {
			res++
		}
		num = num / 2
	}
	return res
}

func countBits(num int) []int {

	// count[i] = count[ i & i-1 ] +1
}
