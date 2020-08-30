package main

import (
	"fmt"
	"strconv"
)

/*
https://github.com/labuladong/fucking-algorithm/blob/master/%E7%AE%97%E6%B3%95%E6%80%9D%E7%BB%B4%E7%B3%BB%E5%88%97/%E5%AD%97%E7%AC%A6%E4%B8%B2%E4%B9%98%E6%B3%95.md
https://leetcode.com/problems/multiply-strings/
*/
func multiply(num1 string, num2 string) string {
	//生成一个 m+n 长度,默认值为 0 的数组 因为结果最多为 m+n位
	m := len(num1)
	n := len(num2)

	slots := make([]int, m+n)

	//从个位数开始每位相乘 , 然后把结果叠加放置在正确的位置上
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := (num1[i] - '0') * (num2[j] - '0')

			p1 := i + j
			p2 := i + j + 1

			sum := int(mul) + slots[p2]
			slots[p2] = sum % 10
			slots[p1] += sum / 10
		}
	}

	//需要把字符(字节)转成数字类型
	i := 0
	//跳过前面未使用的位 0
	for i < len(slots) && slots[i] == 0 {
		i++
	}
	res := ""

	//将数字转回字节(字符) , 字节数组 (字符串)
	for ; i < len(slots); i++ {
		res += strconv.Itoa(slots[i])
	}

	//如果str 为空 , 则输出 0 (字符串)
	if len(res) == 0 {
		return "0"
	} else {
		return res
	}
}

func main() {
	fmt.Println(multiply("100", "200"))
}
