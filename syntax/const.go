package syntax

import (
	"fmt"
	"math"
)

func DeclareConst() {
	const pi1 = math.Pi
	fmt.Printf("%T\n", pi1)     //float64
	const pi2 float64 = math.Pi //math.Pi默认是float64类型
	fmt.Printf("%T\n", pi2)     //float64

	//const p2 int32 =pi1 运行时错误，因为int32不可以存下pi1
	const p3 float32 = pi1
	const p4 float64 = pi1
	const p5 complex64 = pi1
	const p6 complex128 = pi1

	fmt.Println(p3)
	fmt.Println(p4)
	fmt.Println(p5)
	fmt.Println(p6)

	const (
		a = 1
		b
		c = iota //2
		d
		e = iota //4
	)

	fmt.Printf("a=%d,b=%d,c=%d,d=%d,e=%d\n", a, b, c, d, e)

}
