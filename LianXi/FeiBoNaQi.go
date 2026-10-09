package main

import "fmt"

/*func fibo(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}
func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(fibo(i))
	}
}
*/
func fibo(n int) int {
	/*if n == 0 || n == 1 {
		return 1
	}
	*/
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 1, 1
	for i := 2; i <= n; i++ {
		c := a + b
		a = b
		b = c
	}
	return b
}
func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(fibo(i))
	}
}
