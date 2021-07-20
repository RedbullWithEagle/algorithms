package other

import "fmt"

/***********************************************
1.计算机无法精确表示浮点数，大多数采用IEEE计数
2.经过运算后，浮点数是不相等的，不要用==判断
3.涉及到金额和计算的，都不用浮点数表示，用最小单位，例如货币中的分
************************************************/
func TestFloat() {
	var a float64 = 1.7
	var b float64 = 1.4
	var result float64 = a - b

	var d float64 = 0.3
	if d == result { //d != result
		fmt.Println("d == result")
	} else {
		fmt.Println("d != result")
	}
	//0.30000000000000004441
	fmt.Printf("%.20f\n ", result)

	//浮点数运算后，转换成int64位
	var c float64 = 78.6
	// 7859
	fmt.Println(int64(c * 100))

	f1 := 0.3
	f2 := 0.3
	if f1 == f2 { //f1==f2
		fmt.Println("f1==f2")
	} else {
		fmt.Println("f1!=f2")
	}
}
