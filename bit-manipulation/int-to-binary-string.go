package bit_manipulation

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
