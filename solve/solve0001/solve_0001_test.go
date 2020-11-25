package solve0001

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	resp := twoSum([]int{2, 7, 11, 15}, 9)

	if !reflect.DeepEqual(resp, []int{0, 1}) {
		t.Error("twoSum error")
	}
}
