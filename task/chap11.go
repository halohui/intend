package task

import "fmt"

type simpler interface {
	get() float64
	set(x float64)
}

type simple struct {
	key float64
}

func (s simple) get() float64 {
	return s.key
}

func (s *simple) set(x float64) {
	(*s).key = x
}

func UseSimple(s simpler) {
	fmt.Println(s.get())
	s.set(34.89)
	fmt.Println(s.get())
}
