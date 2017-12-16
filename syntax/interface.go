package syntax

import (
	"fmt"
	"reflect"
)

type shape interface {
	area() float64 //不用func关键字
}

type Square struct {
	side float64
}

func (sq *Square) Area() float64 {
	return sq.side * sq.side
}

func ShapeAndSquare() {
	sq1:=new(Square)
	sq1.side=5

	areaInf:=sq1
	fmt.Println("the type of Square is",reflect.TypeOf(areaInf))
	fmt.Printf("The square has area:%f\n",areaInf.Area())

}
