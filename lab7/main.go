package main

import (
	"fmt"
)

func cube(num int) int {
	return num * num * num
}

func main() {
	var a, sum int
	fmt.Scan(&a)
	arr := make([]int, a)
	for i := 0; i < a; i++ {
		fmt.Scan(&arr[i])
		sum += cube(arr[i])
	}
	fmt.Print(sum)
}
