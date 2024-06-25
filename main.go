package main

import "fmt"

func main() {
	var total int

	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Printf("Total Value in For Loop is %d\n", total)
	}

	fmt.Printf("Total Value is %d\n", total)
}
