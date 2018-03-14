package basic

import (
	"fmt"
	"time"
)

var m chan int = make(chan int, 5)

func Product(x []int) {
	for _, i := range x {
		fmt.Println("product", i)
		m <- i
		time.Sleep(10 * time.Second)

	}
	fmt.Println("finsh product ...")
}

func Consumer() {
	fmt.Println("consumer", <-m)
}

func GoroutinesExample() {
	data := MakeSlice(20)
	go Product(data)
	fmt.Println("start to consume....")
	for i := 0; i < len(data); i++ {
		//time.Sleep(1 * time.Second)
		Consumer()
	}

}
