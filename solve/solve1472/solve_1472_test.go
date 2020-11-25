package solve1472

import (
	"testing"
)

func TestSortString(t *testing.T) {
	arr := [][]string{
		{"aaaabbbbcccc", "abccbaabccba"},
		{"rat", "art"},
		{"leetcode", "cdelotee"},
		{"ggggggg", "ggggggg"},
		{"spo", "ops"},
	}

	for _, item := range arr {
		if res := sortString(item[0]); res != item[1] {
			t.Errorf("sortString(%v) should get %v, but get %v\n", item[0], item[1], res)
		}
	}
}
