package sort

import "fmt"

/***************************************************************************
*No.136 只出现一次的数字
*select1Num
*一个数组中，有一种数出现了奇数次，其他数都出现了偶数次
*思路：遍历数组，与每个数做^,由于异或的性质(相同为0，不同为1)，最后剩下的就是出现奇数次的数
***************************************************************************/
func select1Num(arr []int) int {
	tmp := 0
	for _, val := range arr {
		tmp = tmp ^ val
	}

	return tmp
}

func TestSelect1Num() {
	arr := []int{23, 46, 52, 35, 52, 78, 46, 23, 23, 35, 78, 11}
	fmt.Println(Select2Num(arr))
}

/***************************************************************************
*select2Num
*一个数组中，有两种数出现了奇数次，其他数都出现了偶数次
*思路：1.遍历数组，与每个数做^,得出两个奇数次数的异或结果tmp1
*2.tmp1必大于0，然后找到不为0的那位，遍历数组，与每个该位不为0的数字做异或，得出结果tmp2
*3.tmp2与tmp1做异或，得出第二个数
***************************************************************************/
func Select2Num(arr []int) []int {
	tmp1 := 0
	for _, val := range arr {
		tmp1 = tmp1 ^ val
	}

	tmp := 0
	for i := 0; i < 32; i++ {
		tmp2 := 1 << i
		if (tmp1 & tmp2) > 0 {
			tmp = tmp2
			break
		}
	}

	tmp2 := tmp1
	for _, val := range arr {
		if (val & tmp) > 0 {
			tmp2 = tmp2 ^ val
		}
	}

	result := make([]int, 2, 2)
	result[0] = tmp2
	result[1] = tmp2 ^ tmp1

	return result
}
