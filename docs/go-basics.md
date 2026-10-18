# Go 语法笔记（W1 · 2026-10-12 ~ 2026-10-18）

> 这份笔记按「我实际写过、跑通过的」整理，不是抄文档。
> 每条都附可运行的代码；凡是标了「实跑输出」的，都是真的跑过。

---

## 1. 程序骨架

```go
package main

import "fmt"

func main() {
    fmt.Println("hello")
}
```

- `package main`：声明这是一个可执行程序（不是给人 import 的库）
- `import "fmt"`：要用哪个包就列哪个，漏了报 `undefined: fmt`
- `func main()`：程序从这里开始跑，**必须有、必须没有参数和返回值**

### 圆括号 vs 花括号

| 用在哪 | 用什么 |
|---|---|
| 函数调用 / 定义参数 | **圆括号** `()` |
| struct 定义、代码块 | **花括号** `{}` |

踩过的坑：`append{list, s}` 是错的，`append(list, s)` 才对。

---

## 2. 变量：`var` 和 `:=`

```go
var x int = 3      // 完整写法：var 名字 类型 = 值
var y int          // 只声明：自动给零值（int → 0）
z := 10            // 短声明：类型自动推断，只能在函数内用
```

### `:=` 是「出生」，`=` 是「改值」

```go
a := 1     // 第一次出生，用 :=
a = 2      // 已经存在了，改它用 =
```

同一个变量**不能出生两次**：

```go
a := 1
a := 2   // 报错：no new variables on left side of :=
```

> **报错原文**：`no new variables on left side of :=`
> 意思是 `:=` 左边没有一个是新变量。

如果确实要重新赋值，把 `:` 删掉：`a = 2`。

### 零值表

| 类型 | 零值 |
|---|---|
| `int` | `0` |
| `string` | `""`（空字符串） |
| `bool` | `false` |
| 切片 / map / 指针 | `nil` |

---

## 3. 控制流

### for 三段式

```go
for i := 0; i < len(list); i++ {
    fmt.Println(list[i])
}
```

三段：**起点 `i := 0`** / **继续条件 `i < len`** / **每轮变化 `i++`**

### `i < len(list)` 跑几次

从 **0** 数起，最大下标是 **len-1**，所以跑 **len 次**。

```
list 有 5 个元素 → i = 0,1,2,3,4 → 跑 5 次
```

### if / else

```go
if ok {
    fmt.Println("找到了")
} else {
    fmt.Println("没有")
}
```

`==` 等于，`!=` 不等于（`!` 就是把判断反过来）。

---

## 4. 函数

```go
func add(a int, b int) int {
    return a + b
}
```

- 类型写在**参数名后面**
- 返回值类型写在**参数括号后面**
- 多返回值：`func f() (int, error)`

### 命名返回值能被 defer 改

```go
func f() (n int) {
    defer func() { n = 20 }()
    n = 10
    return n     // 实际返回 20，因为 defer 在真正返回前跑
}
```

---

## 5. 指针：`&` 和 `*`

```go
x := 10
p := &x      // & = 取地址，p 装上 x 的门牌号
*p = 100     // * = 顺着门牌号找过去，改那个盒子
fmt.Println(x)   // 100
```

**关键理解**：`p` 和 `x` 指向**同一块内存**。改 `*p` 就是改 `x`。

### 为什么函数要用指针

Go 的函数**拿的都是副本**：

```go
func 改复印件(v int) { v = 999 }
改复印件(x)
fmt.Println(x)     // 还是原来的值，没变

func 改原件(p *int) { *p = 999 }
改原件(&x)
fmt.Println(x)     // 变成 999 了
```

这就是 `fmt.Scanln(&want)` 必须带 `&` 的原因：不带的话键盘输入装进副本，你的变量还是零值。

---

## 6. struct

```go
type Student struct {    // 图纸
    Name  string
    Score int
    ID    int
}

s := Student{"小明", 88, 1}     // 做卡（按字段顺序）
s2 := Student{Name: "小刚"}      // 做卡（带字段名，没写的字段为零值）
var s3 Student                   // 做一张空卡
s3.Name = "小紫"                 // 往卡上填
```

### 花括号紧贴类型名

```go
s := Student{"小明", 88, 1}    // ✅ 做卡
s := append{"小明", 88, 1}     // ❌ append 不是类型
```

报错原文：`append (built-in function) is not a type`

### 点号取字段

```go
s.Name    // 读
s.Score = 90   // 写
```

字段名是**英文**：`Name` / `Score` / `ID`，不是 `名字` / `fenshu`。

---

## 7. 方法

```go
func (s Student) IsPass() bool {
    return s.Score >= 60
}
```

- `(s Student)` 叫**接收者**：**`s` 就是调用这个方法的那张卡本身**
- 调用：`s.IsPass()`

```go
s := Student{"小明", 88, 1}
fmt.Println(s.IsPass())    // true
```

`s.IsPass()` 里的 `s` 会被自动装进接收者位置。

**方法体里不能调自己**：`IsPass` 里面不能再写 `s.IsPass()`，会无限递归。判断逻辑直接写 `return s.Score >= 60`。

---

## 8. 切片

```go
var list []Student                  // 声明，长度 0
list = append(list, s)              // 塞卡，必须接住返回值
len(list)                           // 几张卡
list[0]                             // 第 0 张（编号从 0 起）
```

**`append` 必须接住返回值**：

```go
append(list, s)        // ❌ 白塞，list 没变
list = append(list, s) // ✅
```

---

## 9. 数组 vs 切片

```go
var 数组 [3]int     // 方括号里有数字 → 数组，长度固定
var 切片 []int      // 方括号里空的   → 切片，可增长
```

| | 数组 `[3]int` | 切片 `[]int` |
|---|---|---|
| 方括号里 | 有数字 | 空的 |
| 长度 | 固定，不能变 | `append` 可增长 |
| 赋值给别的变量 | **复制一份** | **共享同一块数据** |
| 传进函数 | 改不动外面 | **改得动外面** |
| 日常用哪个 | 很少 | **基本都用切片** |

实跑对照：

```go
a := [3]int{1, 2, 3}
b := a
b[0] = 999
// a = [1 2 3]   b = [999 2 3]      ← 数组：互不影响

s1 := []int{1, 2, 3}
s2 := s1
s2[0] = 999
// s1 = [999 2 3]  s2 = [999 2 3]   ← 切片：互相影响
```

---

## 10. map

```go
m := make(map[int]Student)     // 造一张电话本，键是 int，值是 Student
m[1] = Student{"小明", 70, 1}  // 存
m[1]                           // 取，得到整张卡
m[1].Name                      // 取卡再点字段
len(m)                         // 存了几个人
delete(m, 1)                   // 删
```

### `v, ok` 两步写法

```go
stu, ok := m[want]
if ok {
    fmt.Println("找到了", stu.Name)
} else {
    fmt.Println("没有这个学号")
}
```

- `stu`：接住那张卡（名字随便起）
- `ok`：接住「有没有这个人」（`true` / `false`，名字也随便起）

为什么需要 `ok`：查一个不存在的键会返回**全空的卡**（`{ 0 0}`）且不报错，光看卡分不清「空名字的人」和「根本没这人」。

### map 不能直接改字段

```go
m[1].Score = 100    // ❌ 报错：cannot assign to struct field m[1].Score in map
```

必须先整张取出、改完、再整张塞回：

```go
stu, ok := m[id]
if ok {
    stu.Score = 100
    m[id] = stu      // ← 不能漏，漏了改动丢
}
```

---

## 11. 两个容器怎么配合（stu-manager 里的做法）

```go
var list []Student             // 按录入顺序存，用来「列出全部」
m := make(map[int]Student)     // 按学号存，用来「查找」

// 录入时两边都写
list = append(list, s)
m[s.ID] = s
```

**坑：两份数据是独立的。**

- `delete(m, 2)` **只删电话本，切片里那张还在**
- 切片删除只能「过滤重建」：

```go
新list := []Student{}
for i := 0; i < len(list); i++ {
    if list[i].ID != del {
        新list = append(新list, list[i])
    }
}
list = 新list
```

- 修改时两边都要改：`m[id] = stu` + `list[i].Score = 新成绩`

---

## 12. 文件读写 + JSON

### import

```go
import (
    "encoding/json"
    "fmt"
    "os"
)
```

### 保存（切片 → 磁盘）

```go
data, err := json.Marshal(list)          // 切片 → JSON 字节
if err == nil {
    os.WriteFile("students.json", data, 0644)   // 字节 → 文件
}
```

### 读取（磁盘 → 切片）

```go
data, err := os.ReadFile("students.json")   // 文件 → 字节
if err == nil {
    json.Unmarshal(data, &list)              // 字节 → 切片（必须有 &）
}
```

### `&` 为什么不能少

不给 `&` 时实跑报错：

```
json: Unmarshal(non-pointer []main.Student)
```

而且 `list` 还是空的。原因和 `Scanln(&want)` 一样：**函数拿副本，不给地址就改不到你的变量**。

### `err`

每个可能失败的操作都返回 `err`：

```go
if err != nil {
    fmt.Println("出错了：", err)
    return
}
```

`err == nil` 表示**没出错**（`nil` = 空）。

### 相对路径的坑

`os.ReadFile("students.json")` 里的相对路径，**相对于你运行程序时所在的目录**，不是相对于 `main.go`。

- 在 `D:\Past3\stu-manager\` 里跑 → 找到同目录的文件
- 在 `D:\Past2\` 里跑 → 找 `D:\Past2\students.json`，找不到

VS Code 里用 `.vscode/launch.json` 的 `"cwd": "${fileDirname}"` 解决。

---

## 13. 报错原文速查

| 报错 | 意思 | 修法 |
|---|---|---|
| `no new variables on left side of :=` | `:=` 左边没有新变量 | 换成 `=`，或者换个新名字 |
| `undefined: s` | 这个变量没声明过 | 补 `:=` 或 `var` |
| `append (built-in function) is not a type` | 把 `append` 当类型名用了 | 做卡用 `Student{...}` |
| `cannot assign to struct field m[1].Score in map` | map 里的卡不能直接改字段 | 取出→改→存回 |
| `json: Unmarshal(non-pointer ...)` | `Unmarshal` 没给 `&` | 加 `&` |
| `index out of range` | 下标超了 | 检查 `< len` |
| `missing function body` | 只写了函数签名没写体 | 补 `{}` |

---

## 14. 本周实际写过的文件

| 文件 | 内容 |
|---|---|
| `LianXi/card.go` | 5 张卡 + 及格判断（`s.Score >= 60`） |
| `LianXi/slice3.go` | 切片塞 5 张卡 + `for` 遍历打印 |
| `LianXi/mymap.go` | map 按学号存取 |
| `LianXi/lookup.go` | map 查找 + `v, ok` + `if` 分支 + `Scanln` |
| `stu-manager/main.go` | 增删改查 + JSON 持久化（7 条 commit） |

---

*最后更新：2026-10-18*
