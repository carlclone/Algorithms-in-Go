package math

/*
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:

Input: 121
Output: true
Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/

//第一种,直接反转后比较
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	rev := 0
	origin := x

	for x > 0 {
		rev = rev*10 + x%10
		x /= 10
	}
	return rev == origin
}

//第二种,反转一半
func isPalindrome2(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) { //排除负数和 10, 100 (最低位 0 怎么左移都是 0 , 无法参与下面的判断)
		return false
	}

	rev := 0

	for x > rev {
		rev = rev*10 + x%10 //左移一位, 再加上x 最右边一位
		x = x / 10          //右移一位
	}

	//x 长度为奇数的时候,如 12321 , rev 会多左移一位,x 会多右移一位, 变成12 , 123 , 把 rev 往回移动一位作比较
	return x == rev || x == rev/10
}
