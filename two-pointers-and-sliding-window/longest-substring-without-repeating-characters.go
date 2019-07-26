package two_pointers_and_sliding_window

//标签 #@string

//Given a string, find the length of the longest substring without repeating characters.
//
//Example 1:
//
//Input: "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
//Example 2:
//
//Input: "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
//Example 3:
//
//Input: "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

//参考刘宇波老师的写法
func lengthOfLongestSubstring1(s string) int {
	//定义[i,j]窗口

	//优化,由于是ascii , 可以用freq[256]大小的数组代替map , index代表ascii码,value代表freq

	//边界问题讨论
	// 初始化l=0 r=-1 ,即没有窗口 . 窗口极限情况是 i=j=len(s)-1

	//情况讨论
	//j+1不是重复,则直接j++ , 维护map , 维护maxWindow   , j++不能超过len-1
	//j+1是重复的,i移动到重复字符index+1的位置,j++ , 维护map , 维护maxWindow      j++不能超过len-1 , index+1不能超过len-1

	//情况2优化 , 只移动 index+1 , 则变成情况1 , 在下次循环的时候处理
	//上面的优化有bug , 没有维护map , 改为l每次+1 , 并且每次都把l所在的元素从map中-1

	maxWindow, l, r := 0, 0, -1
	freq := [256]int{}

	for l < len(s) {
		if r+1 < len(s) && freq[s[r+1]] == 0 {
			r++
			freq[s[r]]++
			maxWindow = maxc(maxWindow, r-l+1)
		} else {
			freq[s[l]]--
			l++
		}
	}

	return maxWindow

}

//勉强通过的第一版
func lengthOfLongestSubstring(s string) int {
	max := 0
	j := 0
	k := -1
	curWmap := make(map[uint8]int)
	curWsize := 0
	endIndex := len(s) - 1

	for k < endIndex {
		k++
		if _, ok := curWmap[s[k]]; !ok {
			curWmap[s[k]] = k
			curWsize++
			max = maxc(curWsize, max)
		} else {
			curWsize++
			orij := j
			j = curWmap[s[k]] + 1
			curWsize = curWsize - (j - orij)
			for ; orij < j; orij++ {
				delete(curWmap, s[orij])
			}
			curWmap[s[k]] = k
			//sub:=curWmap[s[k]]
			//curWsize=curWsize - (sub-j)
			//orij:=j
			//curWmap[s[k]] = j
			//j=sub+1
			//
		}
	}
	return max

}
