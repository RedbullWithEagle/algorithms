package array

import "fmt"

/*********************************************************
*model1 不会死循环
for range其实是golang的语法糖
*在循环开始前会获取切片的长度 len(切片)，然后再执行len(切片)次数的循环。
*
*model2 for循环，每次会重新获取数组的长度，会陷入死循环
***********************************************************/
func AppendData() {
	//model1
	arr := []int{1, 2, 3}
	for _, v := range arr {
		arr = append(arr, v)
		fmt.Println("the array len is ", len(arr))
	}

	//model2
	/*for i := 0; i < len(arr); i++ {
		arr = append(arr, arr[i])

		fmt.Println("sb the array len is ", len(arr))
	}*/
}
