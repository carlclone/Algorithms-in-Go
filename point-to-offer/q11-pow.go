package point_to_offer

import "fmt"

func main() {
	fmt.Println(pow1(2, 2))

}

//没有考虑负数和 0 的解法
func pow1(base int, exponent int) float64 {
	res := base
	for i := 1; i < exponent; i++ {
		res *= base
	}
	return float64(res)
}

//负数:转为正数处理然后倒数
func pow2(base int, exponent int) float64 {

}

//递归, 有重叠子问题
//todo;画图
func pow3(base float64, exp int) float64 {
	if exp == 0 {
		return 1
	}
	if exp == 1 {
		return base
	}
	res := pow3(base, exp>>1) // exp>>1 == exp/2
	res *= res

	//奇数 exp , 再乘一次
	if exp&0x1 == 1 {
		res *= base
	}
	return res
}
