package main

import "fmt"

type Student struct {
	Name  string
	Score int
	ID    int
}

func main() {
	list := []Student{
		{"小明", 70, 1},
		{"小刚", 90, 2},
		{"小紫", 40, 3},
	}
	del := 2

	// ============================================
	// ↓ 只写这一块（5 步：造·走·判·塞·换）
	//   写完把下面这行 fmt.Println 保留，看输出
	// ============================================

	// TODO: 在这里写 5 行
	新list := []Student{}
	for i := 0; i < len(list); i++ {
		if list[i].ID != del {
			新list = append(新list, list[i])
		}
	}
	list = 新list
	// ============================================
	// ↑ 写到这里为止，下面别动
	// ============================================

	fmt.Println("删除后：", list)
	fmt.Println("应该输出：[{小明 70 1} {小紫 40 3}]")
}
