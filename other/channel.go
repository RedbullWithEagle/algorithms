package other

import (
	"fmt"
	"sync"
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
		dogChan <- 1
	}
}

/***************************************************
*有三个方法，分别打印dog，cat,fish
*要求按顺序打印，100次
*****************************************************/
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

func PrintThree() {
	var wg sync.WaitGroup
	//var catCounter, dogCounter, fishCounter uint32
	var catChan, dogChan, fishChan chan int
	catChan = make(chan int)
	dogChan = make(chan int)
	fishChan = make(chan int)
	wg.Add(3)
	go dog(dogChan, catChan, &wg)
	go cat(fishChan, catChan, &wg)
	go fish(dogChan, fishChan, &wg)
	//无缓冲的channel，必须将cat初始化放在这，否则会因为没有接受者而阻塞
	catChan <- 1
	wg.Wait()
}