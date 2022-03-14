package other

import "fmt"

type A struct {
	Age int
}

type B struct {
	Age int
}

type C struct {
	Name  string
	AData *A
}

/*********************************************
*struct能不能比较？
*
*相同struct类型的可以比较
*不同struct类型的不可以比较,编译都不过，类型不匹配
*********************************************/
func TestEqualStruct() {
	a1 := A{1}
	a2 := A{1}
	//b1 := B{1}

	if a1 == a2 {
		fmt.Println("a1 == a2")
	} else {
		fmt.Println("a1 != a2")
	}

	//c1 != c2
	//c1 == c3
	str := "liucz"
	c1 := C{str, &a1}
	c2 := C{str, &a2}
	c3 := C{str, &a1}
	if c1 == c2 {
		fmt.Println("c1 == c2")
	} else {
		fmt.Println("c1 != c2")
	}

	if c1 == c3 {
		fmt.Println("c1 == c3")
	} else {
		fmt.Println("c1 != c3")
	}
	//invalid operation: a1 == b1 (mismatched types A and B)
	/*if a1 == b1 {
		fmt.Println("a1 == b1")
	}else{
		fmt.Println("a1 != b1")
	}*/
}

/*****************************************************
*结构体匿名字段
*
*结构体中每种类型可以有一个匿名字段
*相同类型只能有一个，如果再写一个会报错，无法识别
*
*通过类型来访问匿名字段
****************************************************/
type Worker struct {
	string //匿名字段
	int
	//string  再次定义一个 string 匿名字段时编译报错
}

func TestAnonymous() {
	bob := Worker{"Bob", 22}
	fmt.Println(bob)
	fmt.Println(bob.string)
	fmt.Println(bob.int)
}
