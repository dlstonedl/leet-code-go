package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs1(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Scale1(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (f MyFloat) Abs2() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs1(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs2())

	v = Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	v = Vertex{3, 4}
	v.Scale1(10)
	fmt.Println(v.Abs())
}
