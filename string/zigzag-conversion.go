package string

func convert(s string, numRows int) string {
	//先计算出需要的列数 or 直接使用动态数组?
	arr := [][]rune{}

	runes := rune(s)

	currentRow := 0
	counter := 0

	for counter <= len(runes) {

		if notYetTouchTheBottom(currentRow) {
			arr[currentRow] = append(arr[currentRow], runes[counter])
			counter++
		} else {
			for notYetTouchTheTop(currentRow) {
				moveUpRightAndPush()
			}
		}

	}

	str := ""
	for _, v := range arr {
		for _, j := range v {
			str += j
		}
	}
	return str
}
