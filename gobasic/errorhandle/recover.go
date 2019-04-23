package main

import (
	"errors"
	"fmt"
)

func practiceRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Printf("err is %v", err)
		} else {
			panic(err)
		}

	}()

	panic(errors.New("panic error"))
	//panic(123)
}

func main() {
	practiceRecover()
}
