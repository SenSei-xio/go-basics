package main

import "fmt"

type Student struct {
	Name  string
	Score int
	ID    int
}

func main() {
	var list []Student

	list = append(list, Student{"小刚", 60, 3})

	fmt.Println(len(list))
}
