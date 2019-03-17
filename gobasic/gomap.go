package main

import (
	"fmt"
)

func main() {
	var intMap map[int]int
	//intMap[0] = 0
	if intMap == nil {
		fmt.Println("empty map")
	}

	intMap = make(map[int]int)
	if intMap == nil {
		fmt.Println("empty map")
	} else {
		fmt.Println("not empty map")
		len := len(intMap)
		fmt.Println(len)
	}

	intMap[1] = 1
	intMap[2] = 2

	for k, v := range intMap {
		fmt.Println(k, "对应", v)
	}

	v, ok := intMap[1]
	if ok {
		fmt.Println("1对应值为", v)
	} else {
		fmt.Println("1不存在")
	}

	delete(intMap, 1)

	for k, v := range intMap {
		fmt.Println(k, "对应", v)
	}

}
