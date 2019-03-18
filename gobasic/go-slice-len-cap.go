package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSliceInfo(s)

	s = s[:0]
	printSliceInfo(s)

	s = s[:4]
	printSliceInfo(s)

	s = s[2:]
	printSliceInfo(s)

	s = s[1:2]
	printSliceInfo(s)

	var a []int
	printSliceInfo(a)

	if a == nil {
		fmt.Println("nil!")
	}

	b := make([]int, 5)
	printSliceInfo(b)

	c := make([]int, 0, 5)
	printSliceInfo(c)

	c = c[:cap(c)]
	printSliceInfo(c)

	c = c[1:]
	printSliceInfo(c)
}

func printSliceInfo(x []int) {
	fmt.Printf("len=%d, cap=%d, %v\n", len(x), cap(x), x)

}
