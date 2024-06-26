package array

import (
	"fmt"
	"sort"
)

// ReverseArray 反转数组
func ReverseArray(source []int) []int {
	len := len(source)
	if len <= 1 {
		return source
	}
	right := 0
	left := len - 1

	fmt.Println("9/2:", 9/2)
	fmt.Println("8/2:", 8/2)

	for right < left {
		tmp := source[left]
		source[left] = source[right]
		source[right] = tmp
		left--
		right++
	}

	return source
}

// BSearch 基本的二分查找
func BSearch(nums []int, target int) int {
	len := len(nums)
	left := 0
	right := len - 1

	//如果这里是<,查找不是在闭区间内，很多值查不到，第一个和最后一个元素取不到
	for left < right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

/*********************************************
*No.1 两数之和
*给定一个整数数组 nums 和一个整数目标值 target，
*请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
********************************************/
func twoSum(nums []int, target int) []int {
	ret := make([]int, 0, 2)
	mapNums := make(map[int]int, 0)
	for index, value := range nums {
		mapNums[value] = index
	}

	for index, value := range nums {
		tmp := target - value
		if v, ok := mapNums[tmp]; ok {
			if v == index {
				continue
			}
			ret = append(ret, index)
			ret = append(ret, v)
			return ret
		}
	}

	return ret
}

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0, 1024)
	nLen := len(nums)
	if nLen <= 2 {
		return ret
	}

	for i := 0; i < nLen-1; i++ {
		twoSum := nums[i] + nums[i+1]
		if i+2 == nLen {
			return ret
		}
		for j := i + 2; j < nLen; j++ {
			if twoSum+nums[j] == 0 {
				tmp := []int{i, i + 1, j}
				ret = append(ret, tmp)
			}
		}
	}

	return ret
}

/*********************************************
*No.26 删除有序数组中的重复项
*给你一个有序数组 nums ，请你 原地 删除重复出现的元素，
*使每个元素 只出现一次 ，返回删除后数组的新长度。
*不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
*
********************************************/
func removeDuplicates(nums []int) []int {
	nLen := len(nums)
	if nLen <= 1 {
		return nums
	}
	sort.Ints(nums)
	for i := nLen - 1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i-1], nums[i:]...)
		}
	}

	return nums
}

/*********************************************
*No.26 删除有序数组中的重复项
*采用双指针算法
*双指针的前提是数组有序
********************************************/
func removeDuplicatesGF(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func RangeMap(fields map[string]string) error {
	//如果fields是nil，for不会进入，直接运行后面的代码
	for k, v := range fields {
		fmt.Println(k, v)
	}
	return nil
}

func TestSliceMake() {
	//slice 声明后不make，也可以正常使用
	var abc []int
	abc = append(abc, 3)
	abc = append(abc, 4)
	fmt.Println(abc) //可正常打印

	//map 声明后不make，会panic
	var tmp map[string]string
	//assignment to entry map may panic because of 'nil' map
	tmp["abc"] = "123"
	fmt.Println(tmp)
}

/*
********************************************************
数组作为函数参数
*1.数组作为函数参数时，是值传递，底层数组共享
*2.底层数组共享：A。如果修改了某个元素的值外面是看不到的

	B。如果Append元素或者delete元素，也对切片的len修改了，所以外面的切片看不到

*3.如果实现，里面的修改外面看不到   A.使用Copy函数，这样底层数组不会共享

	如果想实现增加元素，外面能看到，传递数组指针  *[]int

********************************************************
*/
func TestSliceFunc(arr []int) {
	arr[0] = 88
	fmt.Println("func in arr :")
	fmt.Println(arr)

	arr = append(arr, 99)
	fmt.Println("func append arr :")
	fmt.Println(arr)
}

func TestArray() {
	//TestSliceMake()
	arr := []int{1, 2, 3, 4, 5}
	/*TestSliceFunc(arr)
	fmt.Println("func out arr:")
	fmt.Println(arr)*/
	//var arr2 []int
	arr2 := make([]int, 5)
	copy(arr2, arr)
	TestSliceFunc(arr2)
	fmt.Println("func out arr:")
	fmt.Println(arr)
	//RangeMap(nil)

	/*fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))
	fmt.Println(countAndSay(6))
	fmt.Println(countAndSay(7))*/
	nums := []int{386987186, 364389196, 50977443, 50977443, 399358080, 50977443, 409259588, 409259588, 409259588,
		409259588, 409259588, 409259588, 409259588, 409259588, 386281270, 386281270, 386281270, 25298682, 25298682,
		25298682, 25298682, 25298682, 25298682, 25298682, 25298682, 409259588, 409259588, 409259588, 409259588,
		409259588, 409259588, 409259588, 409259588, 409259588, 409259588, 409259588, 409259588, 409259588,
		223542312, 223542312, 409259588, 386281270, 409259588, 406869766, 332315507, 410071207, 399309372,
		406869766, 406869766, 406869766, 406869766, 406869766, 406869766, 406869766, 406869766, 406869766,
		406869766, 406869766, 406869766, 406869766, 406869766, 406869766, 406869766, 406869766, 406869766,
		406869766, 375799397, 333034168, 164687972, 405973939, 178781576, 372718266, 80992258, 351287809,
		334266133, 426676632, 426676632, 388437659, 419133949, 109356078, 431761173, 429719767, 343820848,
		91087337, 343820848, 343820848, 343820848, 40457729, 400459887, 400459887, 400459887, 400459887,
		400459887, 400459887, 400459887, 400459887, 400459887, 331430611, 165357821, 215053565, 179283002,
		219597372, 219597372, 219597372, 219597372, 219597372, 219597372, 219597372, 219597372, 343820848,
		427562206, 427562206, 427562206, 427562206, 427562206, 427562206, 427562206, 136808852, 136808852,
		136808852, 136808852, 136808852, 136808852, 136808852, 136808852, 136808852, 136808852, 136808852}

	fmt.Println(removeDuplicates(nums))
	/*l1 := makeListNode([]int{1,2,3,4,5})
	tmp := removeNthFromEnd(l1, 2)
	Traversal(tmp)
	l2 := makeListNode([]int{1, 3, 4})
	l3 := mergeTwoLists(l1, l2)
	Traversal(l3)*/
}

/*********************************************
*No.38 外观数列
*1.   1
*2.   11
*3.   21
*4.   12 11
********************************************/
func countAndSay(n int) string {
	if n < 1 || n > 30 {
		return ""
	}

	if n == 1 {
		return "1"
	}
	retStr := ""
	strBefore := countAndSay(n - 1)
	nLen := len(strBefore)
	if nLen == 0 {
		return ""
	}
	var cChar uint8
	nCount := 0
	for i := 0; i < nLen; i++ {
		if cChar == 0 {
			cChar = strBefore[i]
		}

		if cChar != strBefore[i] && cChar != 0 {
			tmp := fmt.Sprintf("%d%c", nCount, cChar)
			retStr += tmp
			cChar = strBefore[i]
			nCount = 0
		}

		nCount++

		if i == nLen-1 {
			tmp := fmt.Sprintf("%d%c", nCount, cChar)
			retStr += tmp
			return retStr
		}
	}

	return retStr
}

/****************************************************
*No.4  寻找两个正序数组的中位数
*给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
*请你找出并返回这两个正序数组的 中位数 。
****************************************************/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen%2 == 1 {
		midIndex := totalLen / 2
		return float64(getKthElement(nums1, nums2, midIndex))
	} else {
		midIndex1, midIndex2 := totalLen/2-1, totalLen/2
		num1 := float64(getKthElement(nums1, nums2, midIndex1))
		num2 := float64(getKthElement(nums1, nums2, midIndex2))
		return (num1 + num2) / 2
	}
	return 0
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}

		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
	}

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
********************************************************
No.136 只出现一次的数字
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
找出那个只出现了一次的元素。
*********************************************************
*/
func singleNumber(nums []int) int {
	/*for i:=0;i<len(nums);i++{

	}*/
	return 0
}
