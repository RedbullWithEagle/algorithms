package other

import "fmt"

//Int322bin 自己写的
func Int322bin(source int) string {
	result :=make([]rune, 32, 32)
	isF := false
	if source < 0 {
		source = -source
		isF = true
	}
	for i := 0; i <= 31; i++ {
		re := source % 2
		if re > 0 {
			result[31-i] = '1'
		}else{
			result[31-i] = '0'
		}

		source = source >> 1
	}

	if isF {
		//原码求反码
		fanMa := make([]rune, 32, 32)
		for index, v := range result {
			if v == '1' {
				fanMa[index] = '0'
			} else {
				fanMa[index] = '1'
			}
		}

		//反码求 补码
		for i := 0; i < 32; i++ {
			if fanMa[31-i] == '0' {
				fanMa[31-i] = '1'
				break
			}
			fanMa[31-i] = '0'
		}
		result = fanMa
	}
	return string(result)
}

//打印2进制
//简洁版本，
//golang不支持三目运算符 exp ? a : b
//使用if，代码可读性更强
func PrintB(num int32) {
	for i:=31;i>=0;i--{
	    if num &(1<<i) == 0 {
	    	fmt.Print("0")
		}else{
			fmt.Print("1")
		}
	}
}
