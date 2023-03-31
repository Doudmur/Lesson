package main

import (
	"fmt"
)

func main() {
	var a, sum int
	fmt.Scan(&a)
	arr := make([]int, a)
	for i := 0; i < a; i++ {
		fmt.Scan(&arr[i])
		sum += arr[i] * arr[i] * arr[i]
	}
	fmt.Print(sum)
}
