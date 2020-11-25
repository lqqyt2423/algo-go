package solve0007

import "math"

/**
@index 7
@title 整数反转
@difficulty 简单
@tags math
@draft false
@link https://leetcode-cn.com/problems/reverse-integer/
@frontendId 7
*/
func reverse(x int) int {
	res := 0

	for {
		a := x % 10

		if res > math.MaxInt32/10 || (res == math.MaxInt32 && a > 7) {
			return 0
		}

		if res < math.MinInt32/10 || (res == math.MinInt32 && a < -8) {
			return 0
		}

		res = res*10 + a
		x = x / 10
		if x == 0 {
			break
		}
	}

	return res
}
