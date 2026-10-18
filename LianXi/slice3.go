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
	list = append(list, Student{"小明", 70, 1})
	list = append(list, Student{"小紫", 40, 2})
	list = append(list, Student{"小冰", 55, 4})
	list = append(list, Student{"小火", 88, 5})

	fmt.Println(len(list))
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i].Name, list[i].ID, list[i].Score)
	}
}
