package math

/*
参考discuss中的一个简明的解法 ,  /10和*10就像位运算里的左移 , 右移操作 , %10是取最低位
class Solution {
    public int reverse(int x) {
        long res = 0;
        while (x != 0) {
            res *= 10;
            res += x % 10;
            x /= 10;
        }
        return (int)res == res ? (int)res : 0;
    }
}
*/

func reverse(x int) int {
	res := 0
	for x != 0 {
		res *= 10
		res += x % 10
		x /= 10
	}
	if res == int(int32(x)) { //溢出判断
		return res
	}
	return 0
}
