package array

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*************GO指针******************************
*1.指针访问成员变量和方法都是用'.'，不支持'->'
*2.Go的指针不能进行数学运算
*3.不同类型的指针不能相互转换
*************************************************/

/**************unSafe******************************
*1.任何类型的指针和unsafe.Pointer可以相互转换
*2.uintptr和unsafe.Pointer可以相互转换

	uintptr<=>unsafe.Pointer<=>*T

*************************************************/

func testSlice() {
	var z1 []int //z1 is nil
	z2 := make([]int, 5, 12)
	z3 := new([]int)

	z1Len := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&z1)) + uintptr(8)))
	z2Len := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&z2)) + uintptr(8)))
	z3Len := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(z3)) + uintptr(8)))
	z1Cap := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&z1)) + uintptr(16)))
	z2Cap := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&z2)) + uintptr(16)))
	z3Cap := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(z3)) + uintptr(16)))
	fmt.Printf("z1:%v,len:%d,cap:%d\n", z1, z1Len, z1Cap)
	fmt.Printf("z2:%v,len:%d,cap:%d\n", z2, z2Len, z2Cap)
	fmt.Printf("z3:%v,len:%d,cap:%d\n", *z3, z3Len, z3Cap)

	//指针类型之前不能比较和复制，转换，所以要用unsafe.Pointer
	/*as1 := (*reflect.SliceHeader)(unsafe.Pointer(&z1))
	as2 := (*reflect.SliceHeader)(unsafe.Pointer(&z2))
	as3 := (*reflect.SliceHeader)(unsafe.Pointer(z3))
	fmt.Printf("z1:%v,len:%d,cap:%d,data:%d\n", z1, len(z1), cap(z1), as1.Data)
	fmt.Printf("z2:%v,len:%d,cap:%d,data:%d\n", z2, len(z2), cap(z2), as2.Data)
	fmt.Printf("z3:%v,len:%d,cap:%d,data:%d\n", z3, len(*z3), cap(*z3), as3.Data)*/
}

/***********************共享数组问题总结*****************************
1.append追加的元素没有超过底层数组的容量，修改共享的底层数组
2.append追加的元素超过底层数组的容量，重新申请新数组，并将原来的数组值复制到新数组
******************************************************************/
func shareSlice() {
	a := []int{0, 1, 2, 3, 4, 5, 6}
	b := a[0:4]
	as := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	bs := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	fmt.Printf("a:%v,len:%d,cap:%d,data:%d\n", a, len(a), cap(a), as.Data)
	fmt.Printf("b:%v,len:%d,cap:%d,data:%d\n", b, len(b), cap(b), bs.Data)

	b = append(b, 10, 11, 12)
	fmt.Println("----------append未发生扩容时--------------")
	fmt.Printf("a:%v,len:%d,cap:%d,data:%d\n", a, len(a), cap(a), as.Data)
	fmt.Printf("b:%v,len:%d,cap:%d,data:%d\n", b, len(b), cap(b), bs.Data)

	b = append(b, 100, 200)
	fmt.Println("----------append未发生扩容后--------------")
	fmt.Printf("a:%v,len:%d,cap:%d,data:%d\n", a, len(a), cap(a), as.Data)
	fmt.Printf("b:%v,len:%d,cap:%d,data:%d\n", b, len(b), cap(b), bs.Data)
}
