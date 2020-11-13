package solve0001

/**
@index 1
@title 两数之和
@difficulty 简单
@tags array,hash-table
@draft false
@link https://leetcode-cn.com/problems/two-sum/
@frontendId 1
*/

func twoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
