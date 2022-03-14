package array

import (
	"fmt"
	"reflect"
	"unsafe"
)

/**********************************************************************************
*nil切片和空切片的区别？
*
*nil切片和空切片指向的地址不一样。
*nil空切片引用数组指针地址为0（无指向任何实际地址）
*空切片的引用数组指针地址是有的，且固定为一个值
**********************************************************************************/
func TestNil() {
	var s1 []int
	s2 := make([]int, 0)
	s4 := make([]int, 0)

	fmt.Printf("s1 pointer:%+v, s2 pointer:%+v, s4 pointer:%+v, \n",
		*(*reflect.SliceHeader)(unsafe.Pointer(&s1)),
		*(*reflect.SliceHeader)(unsafe.Pointer(&s2)),
		*(*reflect.SliceHeader)(unsafe.Pointer(&s4)))
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s1))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data)
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s4))).Data)
}

/*************************************************************************************
*s1 pointer:{Data:0 Len:0 Cap:0}, s2 pointer:{Data:824634834528 Len:0 Cap:0}, s4 pointer:{Data:824634834528 Len:0 Cap:0},
*false
*true
**************************************************************************************/