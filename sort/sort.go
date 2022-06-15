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

/****************************************************
*MergeSort 归并排序
*采用分治的思想，将数组分成左右两部分，两部分排好序后，合并
*时间复杂度：O(NlogN)
*空间复杂度：O(N)
*可以变成稳定排序
****************************************************/
func MergeSort(arr []int, left, right int) {
	//这里是错误写法，要判断left和right,而不是arr的长度
	/*if len(arr) < 2 {
		return arr
	}*/
	if left == right {
		return
	}

	//求中点的标准写法，不要写成（left+right)/2
	//也不要写成 left + (right-left)/2
	mid := left + (right-left)>>2
	MergeSort(arr, left, mid)
	MergeSort(arr, mid+1, right)
	MergeArray(arr, left, mid, right)
}

/****************************************************
*MergeArray 合并数组中的两个子有序数组
*
****************************************************/
func MergeArray(arr []int, L, M, R int) {
	help := make([]int, R-L+1, R-L+1)
	i := 0
	p1 := L
	p2 := M + 1

	//如果左右两个子数组都不为空
	for ; p1 <= M && p2 <= R; {
		if arr[p1] <= arr[p2] {
			help[i] = arr[p1]
			i++
			p1++
		} else {
			help[i] = arr[p2]
			i++
			p2++
		}
	}

	//左子数组不为空
	for ; p1 <= M; {
		help[i] = arr[p1]
		i++
		p1++
	}

	//右子数组不为空
	for ; p2 <= R; {
		help[i] = arr[p2]
		i++
		p2++
	}

	//最后拷贝到原数组中
	for j := 0; j < len(help); j++ {
		arr[L+j] = help[j]
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

/******************************************************
*1.ValidFunc
*验证算法是否正确
******************************************************/
func ValidFunc(count uint32) bool {
	if count == 0 {
		return false
	}
	success := true
	maxSize := 50
	maxValue := 100
	for i := 0; i < int(count); i++ {
		arr := GenerateRandomArr(maxSize, maxValue)
		fmt.Println(arr)
		arr2 := make([]int, len(arr), maxSize)
		copy(arr2, arr)
		fmt.Println("------------")

		/*NetherlandsFlag2(arr, 0, len(arr)-1, 30)
		fmt.Println(arr)*/
		QuickSort(arr, 0, len(arr)-1)
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

/******************************************************
*HLNationalFlag1  荷兰国旗问题1（自己的写法）
*给定一个数组arr，和一个数num，把小于num的数放在数组的左边
*大于num的数放在数组的右边
******************************************************/
func HLNationalFlag1(arr []int, num int) {
	smallIndex := -1
	bigIndex := len(arr)
	for i := 0; i < len(arr); i++ {
		if smallIndex >= bigIndex {
			return
		}

		if arr[i] <= num {
			smallIndex++
		} else {
			for ; arr[i] > num; {
				if bigIndex >= 0 && smallIndex <= bigIndex {
					Swap(arr, i, bigIndex-1)
					bigIndex--
				} else {
					break
				}

			}
		}
	}
}

/******************************************************
*NetherlandsFlag1  荷兰国旗问题1
*<=num的放左边     >num放右边
*思路：
******************************************************/
func NetherlandsFlag1(arr []int, l, r, num int) {
	less := l - 1
	more := r + 1
	cur := l //当前数更好理解

	for ; cur < more; {
		if arr[cur] < num {
			less++
			Swap(arr, less, cur)
			cur++
		} else if arr[cur] > num {
			more--
			Swap(arr, more, cur)
		} else {
			cur++
		}
	}
}

/******************************************************
*NetherlandsFlag2  荷兰国旗问题2
*正确的写法
*小于num的放左边   =num放中间    >num放右边
******************************************************/
func NetherlandsFlag2(arr []int, l, r, num int) {
	less := l - 1
	more := r + 1
	cur := l //当前数更好理解

	for ; cur < more; {
		if arr[cur] < num {
			less++
			Swap(arr, less, cur)
			cur++
		} else if arr[cur] > num {
			more--
			Swap(arr, more, cur)
		} else {
			cur++
		}
	}
}

/****************************************************************************
*快排1.0
*选最后一个数，当pivot(主元)
*<=pivot 在左边   >pivot在右边
*
*快排2.0
*选取最后一个数，当pivot(主元)
*<pivot 放左边   =pivot的放中间   >pivot放右边   pivot和右边子数组的第一个元素交换
*优点：搞定了一批 =pivot
*
*快排3.0
*随机选一个数，交换到最后一个数，然后在按照快排2.0去做
*优点：随机后，如果次数足够多，长期期望  时间复杂度  O(NlogN)
****************************************************************************/

/***************************************************************************
*QuickSort  快排
*时间复杂度 O(NlogN)
*空间复杂度 O(logN)   最差O(N)   这里的空间是partition返回的数组？
*空间复杂度，主要是记录中点位置的消耗
***************************************************************************/
func QuickSort(arr []int, L, R int) {
	if L < R {
		rand := rand.Intn(R - L + 1)
		Swap(arr, L+rand, R)
		tmp := partition(arr, L, R)
		QuickSort(arr, L, tmp[0]-1)
		QuickSort(arr, tmp[1]+1, R)
	}
}

func partition(arr []int, L, R int) []int {
	less := L - 1
	more := R
	cur := L //当前数更好理解

	for ; cur < more; {
		if arr[cur] < arr[R] {
			less++
			Swap(arr, less, cur)
			cur++
		} else if arr[cur] > arr[R] {
			more--
			Swap(arr, more, cur)
		} else {
			cur++
		}
	}
	Swap(arr, more, R)
	return []int{less + 1, more}
}

/***************************************************************************
*HeapInsert  大根堆插入元素
*某个数处在index位置，能否往上移动
*与其父节点比较，如果大于父节点，交换，一直往上
*index元素的位置，就是新插入的元素位置
*这里需要注意:  -1/2 =0
***************************************************************************/
func HeapInsert(arr []int, index int) {
	for ; arr[index] > arr[(index-1)/2]; {
		Swap(arr, index, (index-1)/2)
		index = (index - 1) / 2
	}
}

/***************************************************************************
*HeapIfy  大根堆插入元素
*某个数处在index位置，能否往下移动
*找到index位置的左右孩子，index元素和最大孩子比较
***************************************************************************/
func HeapIfy(arr []int, index, heapSize int) {
	left := index*2 + 1 //左孩子的下标

	//循环条件：判断是否有左孩子，其实也就判断是否有孩子
	for ; left < heapSize; {
		//左右孩子中，最大一个孩子的下标
		largestIndex := left
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largestIndex = left + 1
		}

		//index元素和最大孩子的元素比较大小
		//如果index元素大
		if arr[largestIndex] <= arr[index] {
			largestIndex = index
		}

		if largestIndex == index {
			break
		}

		Swap(arr, largestIndex, index)
		index = largestIndex
		left = index*2 + 1
	}
}
