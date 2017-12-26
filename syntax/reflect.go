package syntax

import (
	"fmt"
	"reflect"
)

func BaseUse() {
	var x float64= 4.6
	fmt.Println("type:",reflect.TypeOf(x))
	v:=reflect.ValueOf(x)
	fmt.Println("value:",v)
	fmt.Println("Type:",v.Type())   //获取值的类型
	fmt.Println("Kind:",v.Kind())
	fmt.Println("value:",v.Float()) //返回实际值
	fmt.Println(v.Interface())  //Interface()还原接口值
	fmt.Printf("value is: %5.2e\n",v.Interface())
	y:=v.Interface().(float64)  //检测目的类型是否是float64
	fmt.Println(y)
}

