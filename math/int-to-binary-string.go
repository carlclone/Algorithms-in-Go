package main

import (
	"fmt"
	"strconv"
	"strings"
)

//时间复杂度为从(最右一位)到(最左一位 1 的个数) , 假设是 n 位 , 大致为  O(n)
//空间复杂度O(2n)  , builder 为 O(n) , rune[]为 O(n) ,
func tobinAryString(n int) string {
	if n == 0 {
		return "0"
	}
	r := uint32(n)
	var sb strings.Builder
	for r != 0 {
		sb.WriteString(strconv.Itoa(int(r & 1)))
		r = r >> 1
	}

	//strings.Builder 的底层 byte 数组是私有的...执行不了原地翻转
	res := []rune{}
	for i := len(sb.String()) - 1; i >= 0; i-- {
		res = append(res, rune(sb.String()[i]))
	}
	return string(res)
}

//空间复杂度优化到 O(n) , 时间复杂度 O(n)
func tobinAryString2(n int) string {
	if n == 0 {
		return "0"
	}
	r := uint32(n) //go 里没有>>> , 负数得先强转后当正数来操作
	var runeArr []rune
	m := make(map[uint32]rune)
	m[0] = '0'
	m[1] = '1'
	for r != 0 {
		runeArr = append(runeArr, m[r&1])
		r = r >> 1
	}

	//原地翻转
	p1, p2 := 0, len(runeArr)-1
	for p1 < p2 {
		tmp := runeArr[p1]
		runeArr[p1] = runeArr[p2]
		runeArr[p2] = tmp
		p1++
		p2--
	}
	return string(runeArr)
}

//以前的我会这样写,问题在哪?
func tobinAryString3(n int) string {
	if n == 0 {
		return "0"
	}
	r := uint32(n) //go 里没有>>> , 负数得先强转后当正数来操作
	var runeArr string
	m := make(map[uint32]string)
	m[0] = "0"
	m[1] = "1"
	for r != 0 {
		runeArr = m[r&1] + runeArr
		r = r >> 1
	}
	return runeArr
}

func main() {
	fmt.Println(tobinAryString(0))
	fmt.Println(tobinAryString2(-20))
	fmt.Println(tobinAryString3(8))
}
