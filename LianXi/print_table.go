package main

import "fmt"

func main() {
	for i := 1; i <= 4; i++ {
		fmt.Printf("%d\t%d\n", i, i*i)
	}
}
