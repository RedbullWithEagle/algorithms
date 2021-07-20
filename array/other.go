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
