package basic

import (
	"fmt"
	"math"
)

type Rangle struct {
	X, Y float64
}

func (rangle *Rangle) changeRangle(multiple int) {

	rangle.X = rangle.X * float64(multiple)
	rangle.Y = rangle.Y * float64((multiple))

}

func info(rangle Rangle) Rangle {
	rangle.Y = 44
	rangle.X = 33
	return rangle
}

//方法参数写法  等价于 Abs(v Rangle)
func (v Rangle) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func ObjectExample() {

	r := Rangle{3, 4}
	//	r = changeRangle( r ,10)
	//r = info(r)

	r.changeRangle(12)
	fmt.Print((&r).Abs())

}
