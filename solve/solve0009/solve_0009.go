package solve0009

/**
@index 9
@title 回文数
@difficulty 简单
@tags math
@draft false
@link https://leetcode-cn.com/problems/palindrome-number/
@frontendId 9
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	if x%10 == 0 {
		return false
	}

	a := 0
	for x > a {
		a = a*10 + x%10
		x /= 10
	}

	return x == a || x == a/10
}
