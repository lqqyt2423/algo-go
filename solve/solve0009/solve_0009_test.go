package solve0009

import "testing"

func TestIsPalindrome(t *testing.T) {
	arr := []struct {
		val int
		is  bool
	}{
		{val: 121, is: true},
		{val: -121, is: false},
		{val: 10, is: false},
	}

	for _, item := range arr {
		if res := isPalindrome(item.val); item.is != res {
			t.Errorf("isPalindrome(%v) should get %v, but get %v\n", item.val, item.is, res)
		}
	}
}
