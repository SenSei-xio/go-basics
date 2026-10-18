package main

import "fmt"

type Student struct {
	Name  string
	Score int
	ID    int
}

func main() {
	var m map[int]Student
	m = make(map[int]Student)
	m[1] = Student{"小明", 70, 1}
	m[2] = Student{"小刚", 90, 2}
	m[3] = Student{"小玉", 100, 3}

	fmt.Println(m[2].Name)
	fmt.Println(len(m))
}
