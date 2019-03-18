package main

import "fmt"

func main() {
	i := []int{0, 2, 4, 8, 16}
	for k, v := range i {
		fmt.Printf("2*%d == %d\n", k, v)
	}

	for e := range i {
		fmt.Printf("index is %d\n", e)
	}

	for _, v := range i {
		fmt.Printf("value is %d\n", v)
	}

	a := 1 << 0
	fmt.Printf("a=%d\n", a)

}
