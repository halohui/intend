package main

import (
	"intend/syntax"
	"fmt"
	"time"
	"intend/task"
)

var tt = 345 //未使用的全局变量不会提示编译错误

func main() {

	start := time.Now()
	//data := []int{89, 34, 2, 3, 8, 11, 89, 233, 5, 8, 12, 3, 6}
	//data:=[]int{9,3,2,3,8,1,8,3,5,8,2,3,6}
	//fmt.Println(data)

	//arith.InsertSort(data)
	//arith.BinInsertSort(data)
	//arith.ShellSort(data)
	//arith.BubbleSortBig(data)
	//arith.BubbleSortSmall(data)
	//arith.QuickSort(data,0,len(data)-1)
	//arith.SelectSort(data)
	//arith.HeapSort(data)
	//arith.MergeSort(data)
	//arith.CountSort(data,10)
	//arith.RadixSort(data[:], 10)
	//fmt.Println(data)
	//syntax.DeclareArray()
	//syntax.DeclareSlice()
	//syntax.AboutMap()
	//syntax.DeclareStruct()
	//syntax.DeclareConst()
	//syntax.UseIfElse()
	syntax.UseSwitch(1)

	a := []int{1, 2, 4}
	syntax.VarLength(a...)
	syntax.AboutMap()

	data := [3] int{10, 20, 30}
	for i, x := range data {
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		fmt.Println(x, data[i])
	} //10 110 20 220 30 330
	for i, x := range data[:] {
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		fmt.Println(x, data[i])
	} //110 210 420 420 630 630吗

	var xxx HT
	xxx.testHT()
	//task.PrintNTo1(10)
	//task.Print1ToN(10)
	task.Factorial(30)

	end := time.Now()
	delta := end.Sub(start)
	syntax.ShaperAndSquare()
	syntax.RectangleAndSquare()
	syntax.ShowValuable()
	syntax.BaseUse()

	const x = 100
	const y byte = x //直接展开x，相当于const y byte = 100

	const (
		t1, t2 int  = 99, -999
		t3     byte = byte(t1) //t1 指定为int型，需显式转换为byte类型
		//t4 = uint8(t2)  // 常量-999,超出了unit8的范围
	)

	/*const y1 int =3
	const y2 int =y1
	fmt.Println(&y2)*/
	data1 := [3]string{"a", "b", "c"};
	fmt.Println(data1)
	for range data {

	}
	fmt.Println(data1)

	syntax.VariableDeclare()
	fmt.Println(^3,^uint8(3))

	const yyu int32 =45.8
	fmt.Printf("从main函数执行开始到结束的时间耗费: %s\n", delta)
}

type HT struct {
	x int
}

/*func init() {
	fmt.Println("This init function in main!")
}*/

func tt1(a ...int) {
	fmt.Println(a)
}

func (a HT) testHT() {
	fmt.Println("HT")
}
