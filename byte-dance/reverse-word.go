package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	stack := []int{}
	res := []int{}
	for head.Next != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}

	for len(stack) != 0 {
		res = append(res, stack[0])
		stack = stack[1:]
	}

	return res

}

func replaceSpace(s string) string {
	byteS := []byte(s)
	length := len(s)
	spaceCount := 0
	for i := 0; i < length; i++ {
		if s[i] == ' ' {
			spaceCount++
		}
	}

	newLength := length + spaceCount*2 // %20 多两个字符
	newBytes := make([]byte, newLength)

	for k, v := range byteS {
		newBytes[k] = v
	}

	myReplaceSpace(newBytes, length, newLength)
	return string(newBytes)
}

//由于 string 是不可变的,转成字节数组来体现该题的考察点
func myReplaceSpace(str []byte, length int, newLength int) {

	//从后往前替换 , 不需要额外存储
	// i am happy
	//          cl    nl
	//          %20happy
	for cl, nl := length-1, newLength-1; cl >= 0 && nl >= 0; {
		if str[cl] != ' ' {
			str[nl] = str[cl]
			nl--
			cl--
		} else {
			str[nl] = '0'
			nl--
			str[nl] = '2'
			nl--
			str[nl] = '%'
			nl--
			cl--
		}
	}

}

func main() {
	fmt.Println(replaceSpace("i am happy"))
}
