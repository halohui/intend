package syntax

import (
	"math"
	"fmt"
)

//当如果函数返回一个无名变量或者
//没有返回值，返回值列表的括号是可以省略的
func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func UseFunc() {
	fmt.Printf("%T\n", hypot)   //打印
	fmt.Printf("%T\n", UseFunc) //打印
}
