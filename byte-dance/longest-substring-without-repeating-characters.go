package byte_dance

/*
int left = 0, right = 0;

while (right < s.size()) {`
    // 增大窗口
    window.add(s[right]);
    right++;

    while (window needs shrink) {
        // 缩小窗口
        window.remove(s[left]);
        left++;
    }
}
 */
func lengthOfLongestSubstring(s string) int {
	l,r,maxLength,nowLength:=0,0,0,0
	window:=make(map[byte]int)

	for r<len(s) {
		//增大窗口 window.add(right)
		window[s[r]]++
		r++
		nowLength++
		maxLength=mx(nowLength,maxLength)

		//window need shrink , 缩小窗口 window.move(s[left])
		for r<len(s) && window[s[r]]==1 {
			delete(window,s[l])
			l++
			nowLength--
		}
	}
	return maxLength
}

func mx(l int,r int) int{
	if l>r {
		return l
	}
	return r
}
