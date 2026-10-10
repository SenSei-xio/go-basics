package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x, y := 1, 2
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}

//都是值传递，只是*是传的地址，而不带*传的是数值
