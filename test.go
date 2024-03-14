package main

import "fmt"

func sumInt(args ...int) {
	var sum int
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	fmt.Println(len(args), sum)
}
