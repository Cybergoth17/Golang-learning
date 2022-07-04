package shape

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}
type Square struct {
	sidelength float32
}
type Circle struct {
	radius float32
}

func (s Square) Area() float32 {
	return s.sidelength * s.sidelength
}
func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}
func main() {
	square := Square{5}
	circle := Circle{8}
	PrintShapeArea(square)
	PrintShapeArea(circle)
}
func PrintShapeArea(shape Shape) {
	fmt.Println(shape.Area())
}
