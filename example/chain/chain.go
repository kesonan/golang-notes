package main

import (
	"fmt"
)

type Invoker func()
type Interceptor func(invoker Invoker)

func main() {
	list := []Interceptor{func(invoker Invoker) {
		fmt.Println(1)
		invoker()
	}, func(invoker Invoker) {
		fmt.Println(2)
		invoker()
	}, func(invoker Invoker) {
		fmt.Println(3)
		invoker()
	}}

	interceptor := func(invoker Invoker) {
		newInvoker := getInvoker(0, invoker, list)
		list[0](newInvoker)
	}

	interceptor(func() {
		fmt.Println("---")
	})
}

func getInvoker(cur int, invoker Invoker, list []Interceptor) Invoker {
	if cur == len(list)-1 {
		return invoker
	}

	return func() {
		newInvoker := getInvoker(cur+1, invoker, list)
		list[cur+1](newInvoker)
	}
}
