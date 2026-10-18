package main

import "fmt"

type Student struct {
	Name  string
	Score int
	ID    int
}

func main() {
	m := make(map[int]Student)
	m[1] = Student{"小明", 70, 1}
	m[2] = Student{"小刚", 90, 2}

	var want int
	fmt.Println("请输入学号")
	fmt.Scanln(&want)

	stu, ok := m[want]
	fmt.Println(ok, stu.Name)

	if ok {
		fmt.Println("找到了", stu.Name, stu.Score)
	} else {
		fmt.Println("没有这个学号")
	}
}
