package main

import (
	"fmt"
	"github.com/dlstonedl/leet-code-go/gobasic/rpc"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":12345")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)
	div(client, 10, 3)
	div(client, 10, 0)
}

func div(client *rpc.Client, a int, b int) {
	var result float64
	err := client.Call("DemoService.Div",
		rpcdemo.Args{a, b}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
