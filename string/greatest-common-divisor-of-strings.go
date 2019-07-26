package string

//标签 #@string

//For strings S and T, we say "T divides S" if and only if S = T + ... + T  (T concatenated with itself 1 or more times)
//
//Return the largest string X such that X divides str1 and X divides str2.
//
//
//
//Example 1:
//
//Input: str1 = "ABCABC", str2 = "ABC"
//Output: "ABC"
//Example 2:
//
//Input: str1 = "ABABAB", str2 = "ABAB"
//Output: "AB"
//Example 3:
//
//Input: str1 = "LEET", str2 = "CODE"
//Output: ""
//
//
//Note:
//
//1 <= str1.length <= 1000
//1 <= str2.length <= 1000
//str1[i] and str2[i] are English uppercase letters.
//参考: https://www.cnblogs.com/hwd9654/p/10971812.html
// lc1071 Greatest Common Divisor of Strings

// 找两个字符串的最长公共子串

// 假设：str1.length > str2.length

// 因为是公共子串，所以str2一定可以和str1前面一部分匹配上，否则不存在公共子串。

// 所以我们比较str2和str1的0~str2.length()-1部分,

// 若不同，则直接返回””，不存在公共子串。

// 若相同，继续比较str2和str1的剩下部分，这里就是递归了，调用原函数gcd(str2, str1.substring(str2.length))

// 还有一些细节:

// 比如保证str1永远>str2，可以用一个if(str1 < str2){swap(str1, str2)}

// str1 == str2 直接检查 str1.equals(str2)
// class Solution {
//     public String gcdOfStrings(String str1, String str2) {
//         return GCD(str1, str2);
//     }
//     public String GCD(String a, String b) {
//         if(a.length() > b.length()) {
//             for(int i = 0; i < b.length(); i++){
//                 if(b.charAt(i) != a.charAt(i)){
//                     return "";
//                 }
//             }
//             String temp = a.substring(b.length());
//             return GCD(temp,b);
//         }
//         else if(b.length() > a.length())
//             return GCD(b, a);
//         else
//             return a.equals(b) ? a : "";
//     }
// }

// 蕴含了子问题 , 递归 , 问题降级  等思想  , 字符串处理是弱项,多练
func gcdOfStrings(str1 string, str2 string) string {
	s1 := []rune(str1)
	s2 := []rune(str2)
	l1 := len(str1)
	l2 := len(str2)
	minl := l2
	if l1 > l2 {
		minl = l1
	}

	for i := 0; i <= minl; i++ {

	}
}

//go的string处理 , 重看go 极客时间上的相关课程 , 当用下标取访问的时候 , 是用什么方式 ,  ascii码吗 底层是个 []uint8{}的数组 ?
// 用range的时候底层是个[]unicode{}的数组
