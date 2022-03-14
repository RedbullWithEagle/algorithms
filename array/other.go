package array

/**************************************
No46.全排列
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。
你可以 按任意顺序 返回答案
***************************************/
func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	numsLen := len(nums)
	if numsLen <= 1 {
		ret = append(ret, nums)
	}

	for i := 0; i < numsLen; i++ {
		//tmp := make([]int, 0)
	}
	return ret
}

/**************************************
*No28.实现 strStr() 函数。
*给你两个字符串 haystack 和 needle
*请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。
*如果不存在，则返回 -1 。
***************************************/
func strStr(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}

	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}