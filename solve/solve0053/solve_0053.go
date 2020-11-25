package solve0053

/**
@index 53
@title 最大子序和
@difficulty 简单
@tags array,divide-and-conquer,dynamic-programming
@draft false
@link https://leetcode-cn.com/problems/maximum-subarray/
@frontendId 53
*/
func maxSubArray(nums []int) int {
	pre := 0
	res := nums[0]

	for _, v := range nums {
		max := v
		if n := pre + v; n > v {
			max = n
		}
		pre = max

		if pre > res {
			res = pre
		}
	}

	return res
}
