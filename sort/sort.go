package sort

import (
	"fmt"
	"math/rand"
	"sort"
)

/****************************************************
*SelectionSort 选择排序
*每次遍历数组，选取最小的元素，放在起始位置
*时间复杂度： O(N^2)   空间复杂度：O(1)
*不稳定排序，交换的时候，有可能打乱相同值的顺序
****************************************************/
func SelectionSort(arr []int) {
	arrLen := len(arr)
	//这里可以写成 arrLen-1，前面的排好了，最后就是有序的
	for i := 0; i < arrLen-1; i++ {
		minIndex := i
		//可以不用minValue
		//minValue := int(math.MaxInt64)
		for j := i + 1; j < arrLen; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		if minIndex != i {
			Swap(arr, minIndex, i)
		}
	}
}

func Swap(arr []int, i, j int) {
	if i == j {
		return
	}

	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}

/****************************************************
*BubbleSort 冒泡排序
*每次遍历数组，如果前面的数字大于后面的，交换
*时间复杂度： O(N^2)   空间复杂度：O(1)
*稳定排序，相等的时候，可以不交换
****************************************************/
func BubbleSort(arr []int) {
	arrLen := len(arr)
	for i := 0; i < arrLen-1; i++ {
		for j := 0; j < arrLen-i-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j, j+1)
			}
		}
	}
}

/****************************************************
*BubbleSortV1 冒泡排序优化1
*每次遍历数组，如果没有交换，则停止
****************************************************/
func BubbleSortV1(arr []int) {
	arrLen := len(arr)
	for i := 0; i < arrLen-1; i++ {
		swapped := false
		for j := 0; j < arrLen-i-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j, j+1)
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}

/****************************************************
*InsertSort 插入排序
*从第一个元素开始，后面的元素安装顺序插入到前面。就想扑克抓牌一样
*这里注意，不是每次都需要交换，是往后移动，然后在插入
*时间复杂度： O(N^2)   空间复杂度：O(1)
*稳定排序，相等的时候，可以不交换
****************************************************/
func InsertSort(arr []int) {
	lenArr := len(arr)
	for i := 1; i < lenArr; i++ {
		j := i
		target := arr[i]
		for ; j > 0; j-- {
			if target < arr[j-1] {
				arr[j] = arr[j-1]
			} else {
				break
			}
		}

		if j != i {
			arr[j] = target
		}
	}
}

/******************************************************
*1.GenerateRandomArr
*随机生成随机长度，随机数值的数组
******************************************************/
func GenerateRandomArr(maxLen, maxValue int) []int {
	lenTmp := rand.Intn(maxLen)
	if lenTmp <= 0 {
		return []int{}
	}

	arr := make([]int, 0, lenTmp)
	for i := 0; i < lenTmp; i++ {
		tmp := rand.Intn(maxValue) - rand.Intn(maxValue+1)
		arr = append(arr, int(tmp))
	}

	return arr
}

func ValidFunc(count uint32) bool {
	if count == 0 {
		return false
	}
	success := true
	maxSize := 100
	maxValue := 100
	for i := 0; i < int(count); i++ {
		arr := GenerateRandomArr(maxSize, maxValue)
		fmt.Println(arr)
		arr2 := make([]int, len(arr), maxSize)
		copy(arr2, arr)
		fmt.Println("------------")

		InsertSort(arr)
		sort.Ints(arr2)
		if !isEqual(arr, arr2) {
			success = false
		}
		fmt.Println(arr)
		if !success {
			fmt.Println(arr2)
			fmt.Println("false")
		}
		fmt.Println("---------------------------------------------------------------------------------------------")
	}

	return success
}

//isEqual  判断两个数组的值是否相当
func isEqual(arr1, arr2 []int) bool {
	len1 := len(arr1)
	len2 := len(arr2)

	if len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
