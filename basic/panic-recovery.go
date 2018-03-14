package basic

import "fmt"

func main() {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常

		if err := recover(); err != nil {
			fmt.Println("报错了吧 小伙子") // 这里的err其实就是panic传入的内容，55
		}

	}()
	items := [4]int{10, 20, 30, 25}
	for _, i := range items {

		f(i)
	}
}

type empty struct{}

func (e empty) Error() string {
	return "divsion error because of 0"
}

func f(x int) error {
	if x != 30 {

		fmt.Println(100 / (30 - x))
		return nil
	} else {

		return empty{}
	}
}
