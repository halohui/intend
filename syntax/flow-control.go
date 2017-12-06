package syntax

import (
	"fmt"
	"math/rand"
)

func UseIfElse() {
	x := rand.Intn(100)
	const JUDGE = 8
	//第一种形式
	if x > JUDGE { //左大括号要和if 关键字在同一行
		fmt.Printf("%d > %d is: %t\n", x, JUDGE, x > JUDGE)
	}
}

func UseSwitch(x int) {
	switch x {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	default:
		fmt.Println("default")
	}

	switch {
	case x < 1:
		fmt.Printf("%d <1\n", x)
	case x >= 1:
		fmt.Printf("%d >= 1\n", x)
	default:
		fmt.Println("default")
	}

	switch x := rand.Intn(3); x {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

	a, b, c, d := 1, 2, 3, 4
	switch x {
	case a, b: //case中变量可以重复，常量不可以重复
		fmt.Println(a, b)
	case a, c, d:
		fmt.Println(a, d)
	default:
		fmt.Println("default")

	}
}
