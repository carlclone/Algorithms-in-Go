package math

import "strconv"

//X is a good number if after rotating each digit individually by 180 degrees, we get a valid number that is different from X.  Each digit must be rotated - we cannot choose to leave it alone.
//
//A number is valid if each digit remains a digit after rotation. 0, 1, and 8 rotate to themselves; 2 and 5 rotate to each other; 6 and 9 rotate to each other, and the rest of the numbers do not rotate to any other number and become invalid.
//
//Now given a positive number N, how many numbers X from 1 to N are good?
//
//Example:
//Input: 10
//Output: 4
//Explanation:
//There are four good numbers in the range [1, 10] : 2, 5, 6, 9.
//Note that 1 and 10 are not good numbers, since they remain unchanged after rotating.
//Note:
//
//N  will be in range [1, 10000].

func rotatedDigits(N int) int {
	/*
		 0 1 8 good
		 2 5 good
		 6 9 good

		if remain!=N , then *10 , if in map , res++ , if not continue
	*/

	res := 0

	for i := 1; i <= N; i++ {
		if isValid(i) {
			res++
		}
	}

	return res
}

func isValid(n int) bool {
	/*
		if the number has allnum in 102569 and has some 2569 then is valid
	*/
	map1 := map[string]int{
		"2": 1,
		"5": 1,
		"6": 1,
		"9": 1,
	}
	map2 := map[string]int{
		"1": 1,
		"0": 1,
		"8": 1,
		"2": 1,
		"5": 1,
		"6": 1,
		"9": 1,
	}
	str := strconv.Itoa(n)
	//fmt.Println(str[0:1])
	res := 0
	for i := 0; i < len(str); i++ {
		if _, ok := map2[str[i:i+1]]; ok {
			res++
		}
	}
	if res == len(str) {
		for i := 0; i < len(str); i++ {
			if _, ok := map1[str[i:i+1]]; ok {
				return true
			}
		}
	}
	return false

}
