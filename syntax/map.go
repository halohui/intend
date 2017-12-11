package syntax

import (
	"fmt"
)

func AboutMap() {
	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals { //只使用key,因此不需要_
		fmt.Println(key, capitals[key])
	}

	var a1 map[int]string //空map，只读不写
	a1 = map[int]string{5: "chen"}
	fmt.Println(a1)

	var a2 = make(map[[5]int]string) //数组也可以为key，只要支持==,！=
	fmt.Println(a2)
	b1 := [5]int{23, 6, 9, 10}
	a2[b1] = "xiao"
	fmt.Println(a2)

	a3 := map[int]string{1: "xiao", 63: "chen", 45: "hui"}
	fmt.Println(a3)

	for key := range a3 { //key是随机的排序
		fmt.Print(key, "\t", a3[key])
	}
	fmt.Println()

	a3[1] += "xiao" //可以直接赋值

	for key := range a3 { //key是随机的排序
		fmt.Print(key, "\t", a3[key], "\t")
	}
	fmt.Println()

	a4 := map[int][4]int{1: {2, 3}} //value是数组的初始化
	fmt.Println(a4)

	/*a4[1][2]=34  //值是数组或者结构体时，不能直接赋值
	fmt.Println(a4)*/

	type T struct {
		x int
		y float32
	}

	var t1 T
	var a5 = make(map[T]string)

	a5[t1] = "xiao"

	fmt.Println(a5)

	a6 := map[int]func() int{
		1: func() int { return 1 },
		2: func() int { return 2 },
		3: func() int { return 3 },
	}

	fmt.Println(a6) //map[1:0x10982f0 2:0x1098300 3:0x1098310]

	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)

		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)
}
