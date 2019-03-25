package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	point := Point{1, 2}
	p := &point
	p.X = 3
	fmt.Println(p)
}
