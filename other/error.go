package other

import (
	"errors"
	"fmt"
)

func getInteger() error {
	return errors.New("not integer err ")
}

func getString() error {
	return nil
}

func getABC() error {
	return errors.New("not abc err ")
}

/*************************************
*TestError 变量作用域
* :=声明一个新的变量，作用域在{}内
*结果如下：
*not integer err
*error test
*not abc err
*error test
************************************/
func TestError() error {
	var err error
	if err := getInteger(); err != nil {
		fmt.Println(err)
	}

	err = errors.New("error test")
	if err := getString(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(err)
	if err := getABC(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(err)
	return nil
}
