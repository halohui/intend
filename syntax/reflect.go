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
	fmt.Println("Type:",v.Type())
	fmt.Println("Kind:",v.Kind())
	fmt.Println("value:",v.Float()) //返回实际值
	fmt.Println(v.Interface())

	y:=v.Interface().(float64)
	fmt.Println(y)
}

