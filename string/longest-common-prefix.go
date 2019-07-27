package string

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.


*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	//思路 遍历最短的长度
	minLenth := len(strs[0])
	for _, v := range strs {
		if len(v) < minLenth {
			minLenth = len(v)
		}
	}
	if minLenth == 0 {
		return ""
	}

	equal := true
	index := -1
	for i := 0; i < minLenth; i++ {
		curCharacter := strs[0][i]
		for _, v := range strs {
			equal = equal && (curCharacter == v[i])
		}
		if equal == false {
			index = i - 1
			break
		} else {
			index = i
		}
	}
	return strs[0][:index+1]
}
