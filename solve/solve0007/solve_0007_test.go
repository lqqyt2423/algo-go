package solve0007

import (
	"math"
	"testing"
)

func TestReverse(t *testing.T) {
	arr := [][]int{
		{123, 321},
		{-123, -321},
		{-1234, -4321},
		{120, 21},
		{math.MaxInt32, 0},
		{math.MinInt32, 0},
	}

	for _, item := range arr {
		if res := reverse(item[0]); item[1] != res {
			t.Errorf("reverse(%d) should get %d, but get %d\n", item[0], item[1], res)
		}
	}
}
