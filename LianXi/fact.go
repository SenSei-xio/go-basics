package main

import "fmt"

func fact(n int) int {
	sum := 1
	for i := 1; i <= n; i++ {
		sum *= i
	}
	return sum
}
func main() {
	fmt.Println(fact(5))
}
