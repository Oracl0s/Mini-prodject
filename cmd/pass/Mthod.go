package p

import "fmt"

type Square struct {
	Side int
}

func (s Square) Perimeter() {
	fmt.Printf("%T, %v\n", s, s)
	fmt.Printf("%T, %v\n", s.Side, s.Side)
}
func (s Square) Scale(multiplier int) {
	fmt.Printf("%T, %v\n", s, s)
	s.Side = multiplier
}
func (s Square) WrongScale(multiplier int) {
	fmt.Printf("%T, %v\n", s, s)
	s.Side *= multiplier
	fmt.Printf("%T, %v\n", s, s)
}
func Definition() {
	square := Square{Side: 4}
	pSquare := &square
	square2 := Square{Side: 2}
	square.Perimeter()
	square2.Perimeter()
	pSquare.Scale(2)
	pSquare.Perimeter()
	pSquare.Scale(2)
	pSquare.Perimeter()
}
