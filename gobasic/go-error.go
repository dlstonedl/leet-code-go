package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (error *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", error.When, error.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
