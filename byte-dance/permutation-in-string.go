package main

import "fmt"

//可行解 1: s2每个子串都和 s1 对比一下 , 配合 hashmap
//可行解 2: 同上,但no hash map , 而是排序后对比
//可行解 3: 滑动窗口, 窗口可行的定义: 窗口大小和 s1 大小一致,并且match 数量和 s1 大小一致 ,
// 扩大窗口:找可行解 , 缩小窗口:排除不可行解
// 因此缩小窗口的条件是窗口大小和 s1 大小不一致或者 match 数量和 s1不一致,注意越界

//可行解 4:分析：对s2全排再一一跟s1对比肯定会超时。所以我们可以维护两个窗口，一个窗口装s1，另一个窗口装s2。假设s1长度为len1,s2长度为len2。开始先分别装s1和s2的前lne1个字符进各自的窗口。如果此时两个窗口相等则直接返回true，如果不等则s2的窗口从len1开始装s2的字符，同时窗口的左边要删除一个元素，因为两个窗口要保持大小，期间如果两个窗口相等则返回true
//
//class Solution {
//public:
//    bool checkInclusion(string s1, string s2) {
//        //v1、v2维护一个大小相同的窗口，先计算出len1前的字符出现的次数，如果相等直接返回TURE，如果不等则操作v2继续往后走，后面的字符添上，窗口左边的字符删除
//        int len1 = s1.size(),len2 = s2.size();
//        vector<int> v1(128,0),v2(128,0);
//        for(int i = 0;i<len1;i++)
//        {
//            v1[s1[i]]++;
//            v2[s2[i]]++;
//        }
//        if(v1==v2) return true;
//        for(int i = len1;i<len2;i++)  //v2从len1位置开始装
//        {
//            v2[s2[i]]++;            //装新的字符
//            v2[s2[i-len1]]--;       //删除早装入的字符
//            if(v1 == v2)
//                return true;
//        }
//        return false;
//    }

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}


	countsS1 := [26]int{}
	for _, c := range s1 {
		countsS1[c-'a']++
	}

	i := 0
	countsWin := [26]int{}

	for j := 0; j < len(s2); j++ {
		countsWin[s2[j]-'a']++

		if j-i+1 > len(s1) {
			countsWin[s2[i]-'a']--
			i++
		}

		if countsS1 == countsWin {
			return true
		}

	}
	return false
}

func main() {

	fmt.Println(checkInclusion("ab","eidboaoo"))
	fmt.Println(checkInclusion("ab","eidbaooo"))
}