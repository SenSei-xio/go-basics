package main

import "fmt"

type Student struct {
	Name  string
	Score int
	Keys  int
}

func (s Student) IsPass() bool {
	return s.Score >= 60
}
func (s Student) Show() {
	fmt.Println(s.Keys, "：", s.Name, " 成绩", s.Score)
}

func main() {
	var s1 Student
	s1.Name = "小明"
	s1.Score = 01
	s1.Keys = 1000001

	fmt.Println(s1.IsPass())
	s1.Show()

	fmt.Println(s1)
	fmt.Println(s1.Name)
	fmt.Println(s1.Score)
	fmt.Println(s1.Keys)
}
