package sync

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
	wg  sync.WaitGroup
)

func TestSync(){
	for i:=0;i<10;i++{
		wg.Add(1)
		// 这里要注意讲i显示的当参数传入内部的匿名函数
		go func(i int) {
			defer wg.Done()
			fmt.Println("once", i)
			/*once.Do(func() {
				fmt.Println("once", i)
			})*/
		}(i)
	}

	wg.Wait()
	fmt.Printf("over")
}