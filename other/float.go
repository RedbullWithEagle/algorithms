package other

import "fmt"

/***********************************************
1.计算机无法精确表示浮点数，大多数采用IEEE计数
2.经过运算后，浮点数是不相等的，不要用==判断
3.涉及到金额和计算的，都不用浮点数表示，用最小单位，例如货币中的分
************************************************/
func TestFloat() {
	aa := 389
	tmpA := float32(aa) / 100
	tmpB := 10.0 - tmpA
	fmt.Println(tmpB)
	tmpC := fmt.Sprintf("%0.2f", tmpB)
	fmt.Println(tmpC)
	for i := 1; i < 1000; i++ {
		tmp := float32(i) / 100
		tmp2 := 10.0 - tmp
		//tmp3:= fmt.Sprintf("%0.2f",tmp2)
		fmt.Println(tmp2)
	}

	aa = 111
	bb := float32(aa) / 100
	cc := 110
	dd := float32(cc) / 100
	fmt.Println(dd)
	fmt.Println(bb)
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

/***************************************************************
Math.Floor 地板
	Math.floor(11.46)=Math.floor(11.68)=Math.floor(11.5)=11
	Math.floor(-11.46)=Math.floor(-11.68)=Math.floor(-11.5)=-12

Math.Ceil 天花板
	Math.ceil(11.46)=Math.ceil(11.68)=Math.ceil(11.5)=12
	Math.ceil(-11.46)=Math.ceil(-11.68)=Math.ceil(-11.5)=-11

Math.Round 四舍五入 round附近、周围
	Math.round(11.46)=11
	Math.round(11.5) = Math.round(11.68)=12
***************************************************************/

func TestDivide() {
	//结论： 整数除法取 floor
	a := 3 / 2 //1.6 result=1
	b := 9 / 5 //1.8 result=1
	c := 8 / 7 //1.15 result=1

	fmt.Println("3/2=", a)
	fmt.Println("9/5=", b)
	fmt.Println("8/7=", c)
}
