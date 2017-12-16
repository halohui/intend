package syntax

import (
	"fmt"
	"reflect"
)

type shape interface {
	Area() float64 //不用func关键字
}

type Square struct {
	side float64
}

func (sq *Square) Area() float64 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.length
}
func ShapeAndSquare() {
	sq1 := new(Square)
	sq1.side = 5

	areaInf := sq1
	fmt.Println("the type of Square is", reflect.TypeOf(areaInf))
	fmt.Printf("The square has area:%f\n", areaInf.Area())
}

func RectangleAndSquare() {
	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	// shapes := []Shaper{Shaper(r), Shaper(q)} // or shorter
	shapes := []shape{r, q}

	for i := range shapes {
		fmt.Println("形状信息", shapes[i])
		fmt.Println("这个形状的面积: ", shapes[i].Area())
	}
}

type stockPosition struct {
	ticker     string
	sharePrice float64
	count      float64
}

func (s *stockPosition) GetVaue() float64 {
	return s.sharePrice * s.count //直接计算价值时，如果先将count定义为int时会报错，类型不兼容
}

type car struct {
	make  string
	model string
	price float32
}
