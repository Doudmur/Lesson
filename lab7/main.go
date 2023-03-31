package main

import "fmt"

func main() {
	a := 0
	fmt.Scan(&a)
	if a > 0 {
		fmt.Println("Число положительное")
	} else if a < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Ноль")
	}
}
