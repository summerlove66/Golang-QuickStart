package basic

import "fmt"

func fibonacci0() func() int {
	s0, s1 := 0, 1
	return func() int {
		s00 := s0
		s0, s1 = s1, s1+s0
		return s00
	}
}

func fibonacci1(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y

	}
	close(c)
}

func FibnoacciExample() {

	//第一种
	f := fibonacci0()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	//第二种
	fmt.Println("---------------")

	c := make(chan int, 10) // range只能迭代已经关闭的 信道
	go fibonacci1(10, c)
	for i := range c {
		fmt.Println(i)
	}
	//for i:=0 ;i <10;i++{
	//	fmt.Println(<-c)
	//}

}
