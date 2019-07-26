package recursion_and_backtracking

/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.
*/

//第一版
func letterCombinations(digits string) []string {
	res := []string{}
	var keyMap = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	totalIndex := len(digits) - 1
	if totalIndex == -1 {
		return res
	}
	/*
		foreach first digit
		pass the res and digit index and currStr to next level
	*/
	for _, v := range keyMap[digits[0:1]] {
		currStr := string(v)
		letterSearch(&res, 1, totalIndex, currStr, keyMap, digits)
	}
	return res
}

func letterSearch(res *[]string, index int, totalIndex int, currStr string, keyMap map[string]string, digits string) {
	if index == totalIndex+1 {
		*res = append(*res, currStr)
		return
	}
	for _, v := range keyMap[digits[index:index+1]] {
		tmp := currStr + string(v)
		letterSearch(res, index+1, totalIndex, tmp, keyMap, digits)
	}
}
