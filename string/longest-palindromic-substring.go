package string

/*
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
*/

func longestPalindrome(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}

	longest := ""
	longestLen := 0

	for i := 0; i <= l-1; i++ {
		len1 := calLP(s, i, i)
		len2 := calLP(s, i, i+1)
		var len int
		var st string
		if len1 > len2 {
			len = len1
			st = s[i-len/2 : i+len/2+1]

		} else {
			len = len2
			st = s[i-len/2+1 : i+len/2+1]
		}
		if len > longestLen {
			longestLen = len
			longest = st
		}

	}
	return longest
}

// [1,2,3]  以1和2中间为中心点扩散(l=0,r=1) ,  以2为中心点扩散(l=1,r=1)
func calLP(s string, l int, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l + 1 - 2
}
