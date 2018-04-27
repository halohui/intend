package syntax

import (
	"fmt"
	"math"
	"strconv"
)

func UseStrconv() {
	a := 45
	fmt.Printf("0b%b,0b%#b\n", a, a)

	fmt.Printf("0b%#b,%#o,%#x\n", a, a, a)

	fmt.Println(math.MinInt8, math.MaxInt8)

	b, _ := strconv.ParseInt("010001", 2, 32)
	fmt.Println(b)

	c1, _ := strconv.ParseInt("0144", 8, 32)
	fmt.Println(c1)

	c2, _ := strconv.ParseInt("0144", 8, 32)
	fmt.Println(c2)

	c3, _ := strconv.ParseInt("0144", 0, 32)
	fmt.Println(c3)

	d1, _ := strconv.ParseInt("64", 16, 32)
	fmt.Println(d1)

	//字符串带有0x前缀的话，那么base应该设置为0
	d2, _ := strconv.ParseInt("0x64", 0, 32)
	fmt.Println(d2)

	d3, _ := strconv.ParseInt("0x64", 16, 32)
	fmt.Println(d3) //0
}
