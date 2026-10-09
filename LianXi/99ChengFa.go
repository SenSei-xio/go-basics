package main

import "fmt"

func main() {
	for b := 1; b <= 9; b++ {
		for a := 1; a <= b; a++ {
			fmt.Printf("%d*%d=%d\t", a, b, a*b)
		}
		fmt.Println()
	}
}
