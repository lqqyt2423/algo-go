package solve1472

/**
@index 1472
@title 上升下降字符串
@difficulty 简单
@tags sort,string
@draft false
@link https://leetcode-cn.com/problems/increasing-decreasing-string/
@frontendId 1370
*/
func sortString(s string) string {
	if len(s) == 0 {
		return ""
	}

	var nums [26]int
	cs := []byte(s)
	for _, c := range cs {
		nums[c-'a']++
	}

	resp := make([]byte, 0, len(cs))
	for len(resp) < len(cs) {
		for i := 0; i < 26; i++ {
			if nums[i] > 0 {
				nums[i]--
				resp = append(resp, byte(i)+'a')
			}
		}
		for i := 25; i >= 0; i-- {
			if nums[i] > 0 {
				nums[i]--
				resp = append(resp, byte(i)+'a')
			}
		}
	}

	return string(resp)
}
