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
	// ↓ 闭卷写这 5 行（不许看 del_lianxi.go）
	//   口诀：造 · 走 · 判 · 塞 · 换
	// ============================================
	新list := []Student{}
	for i := 0; i < len(list); i++ {
		if list[i].ID != del {
			新list = append(新list, list[i])
		}
	}
	list = 新list

	// ============================================

	fmt.Println("删除后：", list)
}
