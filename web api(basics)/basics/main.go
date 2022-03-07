package main

import "fmt"

func main() {
	sum1 := 1
	sum2 := 1
	var temp int
	fmt.Println("\n", sum1, "\n", sum2)
	for i := 1; i < 10; i++ {
		temp = sum2
		sum2 = sum1 + sum2
		sum1 = temp
		fmt.Println("\n", sum2)

	}
}
