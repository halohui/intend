package syntax

import "fmt"

//声明主要有变量，常量，类型以及函数

func VariableDeclare() {
	fmt.Println(f() == f()) //false
}

func f() *int {
	v := 1
	return &v //可以返回局部变量的指针，Go会做逃逸分析
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y  //使用并行赋值来求菲波那契数列
	}
	return x
}

func DecVar()  {
	var x int  = 23  //都不省略
	var y = 23    //省略类型
	var z1,z2 int   //省略表达式

	fmt.Println(x,y,z1,z2)

	var(
		t1,t2 = 23,78
		a1,a2=100,"abc"  //不同类型

	)
	fmt.Println(t1,t2,a1,a2)
}