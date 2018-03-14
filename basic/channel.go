package basic

import "fmt"

func addChannel(c chan int, it []int) {
	for _, i := range it {
		c <- i
	}
	close(c)
}

func MakeSlice(x int) []int {
	var it []int

	for i := 0; i < x; i++ {
		it = append(it, i)
	}
	fmt.Println("ok了") //先打印

	return it

}

func ChannelExample() {

	c := make(chan int)
	go addChannel(c, MakeSlice(10))

	for i := range c {

		fmt.Println(i)
	}

}
