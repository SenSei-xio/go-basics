package main

import "fmt"

func sumOdd(n int) int {
	sum := 0
	for i := 1; i <= n; i += 2 {
		sum += i
	}
	return sum
}
func main() {
	fmt.Println(sumOdd(10))
}
