package syntax

import "fmt"

//声明主要有变量，常量，类型以及函数

func VariableDeclare(){
	fmt.Println(f()==f()) //false
}

func f()*int{
	v:=1
	return &v  //可以返回局部变量的指针，Go会做逃逸分析
}


