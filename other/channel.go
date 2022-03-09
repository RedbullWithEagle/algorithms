package other

import (
	"fmt"
	"sync"
	"time"
)

const PrintCount = 10

func dog(dogChan, catChan chan int, wg *sync.WaitGroup) {
	counter := 0
	for {
		if counter >= PrintCount {
			wg.Done()
			return
		}

		<-dogChan
		fmt.Print("dog ")
		counter++
		catChan <- 1
	}
}

func cat(fishChan, catChan chan int, wg *sync.WaitGroup) {
	counter := 0
	for {
		if counter >= PrintCount {
			wg.Done()
			return
		}

		<-catChan
		fmt.Print("cat ")
		counter++
		fishChan <- 1
	}
}

func fish(dogChan, fishChan chan int, wg *sync.WaitGroup) {
	counter := 0
	for {
		if counter >= PrintCount {
			wg.Done()
			return
		}

		<-fishChan
		fmt.Print("fish ")
		fmt.Println()
		counter++

		//如果使用无缓冲的channel，必须要添加下面的判断
		//否则dog函数已经退出了，dogChan不会接受，会报错
		if counter <= 9 {
			dogChan <- 1
		}
	}
}

func TestPrintAnimals() {
	var wg sync.WaitGroup
	//注意这里使用有缓存的
	//无缓冲的，发送到管道就会阻塞，直到有人取，赋值需要写在下面
	//有缓冲的，发送到管道返回，除非缓冲满了
	dogChan := make(chan int, 1)
	catChan := make(chan int, 1)
	fishChan := make(chan int, 1)
	wg.Add(3)
	dogChan <- 1
	go dog(dogChan, catChan, &wg)
	go cat(fishChan, catChan, &wg)
	go fish(dogChan, fishChan, &wg)
	wg.Wait()
}

func TestPrintAnimalsNoBuffer() {
	var wg sync.WaitGroup
	//使用缓存的channel写法
	dogChan := make(chan int)
	catChan := make(chan int)
	fishChan := make(chan int)
	wg.Add(3)

	//下面这句话放在这里，会报错：
	//fatal error: all goroutines are asleep - deadlock!
	//因为没有协程来接受dogChan
	//dogChan <- 1
	go dog(dogChan, catChan, &wg)
	go cat(fishChan, catChan, &wg)
	go fish(dogChan, fishChan, &wg)

	dogChan <- 1
	wg.Wait()
}

//channel 设置接受超时时间
func ChannelTimeout() {
	ch := make(chan int)
	quit := make(chan bool)
	//新开一个协程
	go func() {
		for {
			select {
			case num := <-ch:
				str := fmt.Sprintf("time:%d, num = %d", time.Now().Unix(), num)
				fmt.Println(str)
			case <-time.After(3 * time.Second):
				fmt.Println("超时,time:%d", time.Now().Unix())
				quit <- true
			}
		}
	}() //别忘了()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	<-quit
	fmt.Println("程序结束")
}
