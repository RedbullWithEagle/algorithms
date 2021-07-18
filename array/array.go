package array

import "fmt"

//ReverseArray 反转数组
func ReverseArray(source []int) []int {
	len := len(source)
	if len <= 1 {
		return source
	}
	right := 0
	left := len - 1

	for ; right < left; {
		tmp := source[left]
		source[left] = source[right]
		source[right] = tmp
		left--
		right++
	}

	return source
}

func TestDivide() {
	//结论： 整数除法取 floor
	a := 3 / 2 //1.6 result=1
	b := 9 / 5 //1.8 result=1
	c := 8 / 7 //1.15 result=1

	fmt.Println("3/2=", a)
	fmt.Println("9/5=", b)
	fmt.Println("8/7=", c)
}

//BSearch 基本的二分查找
func BSearch(nums []int, target int) int {
	len := len(nums)
	left := 0
	right := len - 1

	//如果这里是<,查找不是在闭区间内，很多值差不到，第一个和最后一个元素取不到
	for ; left < right; {
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

func testBinarySearch() {
	nums1 := []int{1, 5, 8, 12, 34, 45, 46, 78, 88}
	//nums2:= []int{2,2}
	//nums3:= []int{}

	fmt.Println(nums1)
	for i := 0; i < len(nums1); i++ {
		tmp := BSearch(nums1, nums1[i])
		fmt.Println(nums1[i], "  result is ", tmp)
	}
	/*result1 := bs.BSearch(nums1,23)
	result11 := bs.BSearch(nums1,78)
	result88 := bs.BSearch(nums1,88)
	result101 := bs.BSearch(nums1,1)
	fmt.Println(nums1)
	fmt.Println("23 result:",result1)
	fmt.Println("78 result:",result11)
	fmt.Println("88 result:",result88)
	fmt.Println("1 result:",result101)
	fmt.Println("----------------------------")

	result2 := bs.BSearch(nums2,8)
	result22 := bs.BSearch(nums2,2)
	fmt.Println("8 result:",result2)
	fmt.Println("2 result:",result22)
	fmt.Println("----------------------------")

	result3 := bs.BSearch(nums3,8)
	result33 := bs.BSearch(nums3,2)
	fmt.Println("8 result:",result3)
	fmt.Println("2 result:",result33)
	fmt.Println("----------------------------")*/
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
func removeDuplicates(nums []int) int {
	nLen := len(nums)
	if nLen <= 1 {
		return 0
	}

	for i := nLen - 1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i-1], nums[i:]...)
		}
	}

	return len(nums)
}

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

func TestArray() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(removeDuplicatesGF(nums))
	fmt.Println(nums)
	/*l1 := makeListNode([]int{1,2,3,4,5})
	tmp := removeNthFromEnd(l1, 2)
	Traversal(tmp)
	l2 := makeListNode([]int{1, 3, 4})
	l3 := mergeTwoLists(l1, l2)
	Traversal(l3)*/
}
