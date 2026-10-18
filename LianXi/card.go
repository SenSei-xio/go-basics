package main

import "fmt"

type Student struct {
	Name  string
	Score int
	ID    int
}

func main() {
	s1 := Student{"小刚", 60, 3}
	fmt.Println(s1.Name)
	fmt.Println(s1.Score >= 60)

	s2 := Student{"小明", 70, 1}
	fmt.Println(s2.ID)
	fmt.Println(s2.Score >= 60)

	s3 := Student{"小紫", 40, 2}
	fmt.Println(s3.Score)
	fmt.Println(s3.Score >= 60)

	s4 := Student{"小冰", 55, 4}
	fmt.Println(s4.Name)
	fmt.Println(s4.Score >= 60)

	s5 := Student{"小火", 88, 5}
	fmt.Println(s5.Name)
	fmt.Println(s5.Score >= 60)
}
