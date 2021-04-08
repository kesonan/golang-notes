package main

import (
	"fmt"
)

type Do func() error

func ExampleChain() {
	list := []Do{
		func() error {
			fmt.Println("1")
			return nil
		},
		func() error {
			fmt.Println("2")
			return nil
		},
		func() error {
			fmt.Println("3")
			return nil
		},
	}

	fn := func() error {
		for _, fn := range list {
			err := fn()
			if err != nil {
				return err
			}
		}
		return nil
	}

	_ = fn()
	// Output:
	// 1
	// 2
	// 3
}
