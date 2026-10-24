# Go 语法笔记（W1–W2 · 2026-10-12 ~ 2026-10-25）

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

## 13. 接口（interface）

### 没有接口时会怎样

两种存储，各有一个 `保存` 方法：

```go
type 文件存储 struct{ 路径 string }
type 内存存储 struct{ 容量 int }

func (f 文件存储) 保存(s Student) { fmt.Println("文件存了", s.Name) }
func (f 内存存储) 保存(s Student) { fmt.Println("内存存了", s.Name) }
```

想写一个「备份」函数让两种存储都能用，只能写两个：

```go
func 备份A(f 文件存储, s Student) { f.保存(s) }
func 备份B(m 内存存储, s Student) { m.保存(s) }
```

**每加一种存储，就得多写一个函数。**

### 加接口之后

```go
type 存储 interface {
    保存(s Student)      // 只写「长什么样」，不写 func，不写实现
}

func 备份(s 存储, stu Student) {   // 参数类型写成接口
    s.保存(stu)
}
```

实跑：

```
=== 传 文件存储 ===
【备份开始】
文件存了 小明
【备份结束】
=== 传 内存存储（备份函数一个字都没改）===
【备份开始】
内存存了 小明
【备份结束】
```

### 三种花括号，装三种东西（最容易混）

| 写法 | 花括号里装 | 圆括号 | 要 `func` 吗 |
|---|---|---|---|
| `type X struct { ... }` | **数据**（字段：名字 + 类型） | ❌ 不许有 | ❌ |
| `type X interface { ... }` | **要求**（方法签名：名字 + 圆括号） | ✅ 必须有 | ❌ 不许有 |
| `func (r X) 方法(参数) { ... }` | **动作**（真代码） | ✅ 有 | ✅ 必须有 |

写错时的真实报错：

```go
type 内存存储 struct{ 保存(s Student) }   // struct 里不许写括号
// syntax error: unexpected name Student, expected )
```

```go
type 文件存储 struct{ 保存(Student) }     // 能编译，但那是【字段】不是【方法】
// 打印出来是 {{}}，调用不了
```

方法拆 4 块：

```go
func (f 文件存储) 保存(s Student) {
// ①   ②         ③    ④
// ① func 关键字  ② 接收者（谁的方法）  ③ 方法名  ④ 参数
```

### 隐式实现：Go 不写 implements

**只要方法名 + 参数类型对得上，就算实现了接口，不用声明。**

名字不对（`保存` 写成 `Save`）：

```
cannot use 硬盘{} (value of struct type 硬盘) as 存储 value in assignment:
    硬盘 does not implement 存储 (missing method 保存)
```

参数类型不对（`保存(s Student)` 写成 `保存(s string)`）：

```
硬盘 does not implement 存储 (wrong type for method 保存)
    have 保存(string)
    want 保存(Student)
```

**Go 只认字面，不认意思。** 大小写、方法名、参数类型三处都极严（跟 `Id` ≠ `ID`、`score` ≠ `Score` 是同一个道理）。

### 指针接收者塞不进接口

```go
func (p *手机) 保存(s Student) { ... }   // 注意是 *手机

var x 存储 = 手机{}      // ❌
// cannot use 手机{} ... 手机 does not implement 存储 (method 保存 has pointer receiver)

var x 存储 = &手机{}     // ✅
```

### 一句话记住

> **接口 = 一份「必须会什么」的清单。谁的方法对得上，谁就自动算数，不用报名。**

---

## 14. 切片底层（`len` 和 `cap`）

### 切片不是数组，是三样东西

| 组成 | 意思 |
|---|---|
| 指针 | 底层数组在哪 |
| `len` | **现在已经用了几个** |
| `cap` | **总共能装几个** |

```go
s := make([]int, 3, 10)
fmt.Println(len(s), cap(s), s)   // 3 10 [0 0 0]
```

### cap 从 4 起，装满就翻倍（实跑）

```go
var s []int
for i := 0; i < 40; i++ {
    s = append(s, i)
    fmt.Println("len =", len(s), "cap =", cap(s))
}
```

| len 到 | cap |
|---|---|
| 1 | 4 |
| 5 | 8 |
| 9 | 16 |
| 17 | 32 |
| 33 | 64 |

**不是从 1 开始 —— 第一次 append 一口气给 4 格。**

### 为什么翻倍，不是每次 +1

塞 1000 个元素：

| 策略 | 要复制多少元素 |
|---|---|
| 每次 +1 格 | 1+2+3+…+999 ≈ **49 万** |
| 每次翻倍 | 4+8+…+1024 ≈ **1000 出头**，只搬约 10 次 |

**用空间浪费换搬家次数**，这叫「摊还 O(1)」。

### 搬家 vs 共享（最关键）

`append` 的时候看 **cap 还有没有空位**：

| cap 情况 | 行为 | 后果 |
|---|---|---|
| 还有空位 | **原地写** | 新旧切片**共享**同一块数组 |
| 已经装满 | **搬家**（造新数组，cap 翻倍） | 新旧切片**分开** |

实跑 A —— `cap` 还有空位：

```go
a := make([]int, 2, 5)     // len=2 cap=5，还有 3 格空位
b := append(a, 3)
b[0] = 999
// b = [999 2 3]   a = [999 2]      ← 改了 b，a 跟着变
```

实跑 B —— `cap` 正好装满：

```go
a := []int{1, 2}           // len=2 cap=2，没空位了
b := append(a, 3)          // 搬家，b 的 cap 变成 4
b[0] = 999
// b = [999 2 3]   a = [1 2]        ← 改了 b，a 不变
```

### 判据一句话

> **`append` 有空位就原地写（共享），没空位就搬家（分开）。跟是不是同一个变量、是不是 `:=` 都无关，只看 cap 够不够。**

### 这条解释了 10-18 周检测那道题

「切片赋值会联动，数组不会」—— 因为切片赋值只复制「指针 + len + cap」三样，**数组本体没复制**，两个切片指向同一块底层数组。

```go
s1 := []int{1, 2, 3}
s2 := s1          // 只复制三样指针，不复制数组
s2[0] = 999
// s1 = [999 2 3]    两个都变
```

---

## 15. testing 测试

### 15.1 测试文件三规矩

| 规矩 | 说明 |
|---|---|
| 文件名 `xxx_test.go` | `go test` 只认这个后缀 |
| 函数名 `TestXxx` | 必须大写 T 开头 |
| 参数 `t *testing.T` | `go test` 自动传进来 |

**同目录、同包**（`funcs.go` 是 `package main` → `funcs_test.go` 也 `package main`）。

### 15.2 最小骨架

```go
func Test平均分(t *testing.T) {
	got := 平均分([]int{70, 90, 40})
	want := 66
	if got != want {
		t.Errorf("平均分 = %d, 想要 %d", got, want)
	}
}
```

`got` / `want` 是惯例名字。`t.Errorf` 失败时自动带「文件名:行号」。

### 15.3 `t.Errorf` vs `t.Fatal`

| | 干什么 | 后面代码 |
|---|---|---|
| `t.Fatal` | 立刻失败 | **一行都不执行** |
| `t.Errorf` | 记一笔，继续 | **照样执行** |

**表驱动必须用 `t.Errorf`** —— 用 `Fatal` 第一个用例失败后面就全不测了。

实跑对照：

```
=== RUN   Test用Fatal
    a_test.go:6: 第一行就 Fatal
--- FAIL: Test用Fatal (0.00s)
=== RUN   Test用Errorf
    a_test.go:11: 第一行 Errorf，但继续跑
    a_test.go:12: 这行照样执行
--- FAIL: Test用Errorf (0.00s)
```

### 15.4 表驱动测试

```go
func Test平均分_表驱动(t *testing.T) {
	cases := []struct {
		名字 string
		输入 []int
		想要 int
	}{
		{"三个数", []int{70, 90, 40}, 66},
		{"一个数", []int{88}, 88},
		{"空切片", []int{}, 0},
		{"全零", []int{0, 0}, 0},
	}

	for _, c := range cases {
		got := 平均分(c.输入)
		if got != c.想要 {
			t.Errorf("%s：平均分(%v) = %d, 想要 %d", c.名字, c.输入, got, c.想要)
		}
	}
}
```

**三段：定义卡型 → 填数据 → `for range` 遍历。**
和 `type Student struct{}` + `[]Student{{...},{...}}` + `for` 遍历是同一套。

### 15.5 `for ... range` 遍历

```go
分数 := []int{70, 90, 40}

// 三段式
for i := 0; i < len(分数); i++ {
	fmt.Println(分数[i])
}

// range（等价）
for i, v := range 分数 {
	fmt.Println(i, v)
}
```

| 写法 | 拿到什么 |
|---|---|
| `for i, v := range 分数` | 下标 + 值 |
| `for i := range 分数` | **只有下标** |
| `for _, v := range 分数` | **只有值**（必须用 `_` 丢下标） |

> ⚠️ `for v := range 分数` 拿到的是**下标**，不是值。

### 15.6 覆盖率

```
go test -cover ./...                 # 总覆盖率
go test -coverprofile=c.out ./...    # 存明细
go tool cover -func=c.out            # 每个函数多少
```

实跑结果：

```
lianxi/s.go:6:   及格     100.0%
lianxi/s.go:11:  平均分   100.0%
lianxi/s.go:22:  main     0.0%
total:            77.8%
```

**`main` 永远是 0%** —— `go test` 不执行 main。

### 15.7 为什么要把逻辑拆出 `main`

`go run` 不执行 Test 函数，`go test` 不执行 main 函数。`main` 里的 `fmt.Scanln` 会让测试卡死、`fmt.Println` 结果拿不到。

**所以 `main` 只放「输入 → 调函数 → 输出」，逻辑搬进独立函数。**

`stu-manager` 拆出来的三个（`funcs.go`）：

```go
func 查找(m map[int]Student, want int) (Student, bool)
func 删除(list []Student, del int) []Student
func 修改(m map[int]Student, list []Student, id int, 新成绩 int) (bool, []Student)
```

---


## 16. map 底层（10-23）

### 16.1 存进去之后，它怎么找到的

```
m[2] = "小明"
  ↓
① 算哈希：把 2 变成一个数字，比如 38
  ↓
② 定桶：38 % 8 = 6  →  放进 6 号桶
  ↓
③ 桶里存：{键 2, 值 "小明"}
```

Go 的 map 内部有 **8 个桶（bucket）**，每个桶有 **8 个槽**。取的时候同样三步：算哈希 → 定位桶 → 在桶里找键。

### 16.2 为什么遍历顺序每次都不一样

Go 每次遍历前会**随机挑一个起始桶**，然后绕 8 个桶走一圈。同一个 map 连跑三次，输出顺序可以全不同。

**这不是 bug，是设计**：故意打乱，防止你写出「依赖遍历顺序」的代码。要顺序，就自己排序（见 16.4）。

### 16.3 为什么 `m[2].Score = 100` 不让写（唯一真原因）

**不是**「桶里没有顺序」，**是元素的位置会变**。

- map 扩容时，所有元素要**重新算哈希、搬到新桶**，位置全变
- 如果允许 `&m[2]`，你拿到的地址过一会儿就失效了 → 变成**野指针**
- 所以 Go 干脆不让 map 元素可取地址

实跑报错：

```
m[2].Score = 100   → cannot assign to struct field m[2].Score in map
fmt.Println(&m[2]) → cannot take address of m[2] (map index expression of struct type Student)
```

对照：**切片可以** `list[0].Score = 100`、`&list[0]` 都能跑，因为切片是货架，格子位置固定。

### 16.4 扩容规则

- 一开始 **1 个桶**，每个桶 8 个槽
- 装到 **「桶数 × 6.5」** 个元素时触发扩容（负载因子 6.5）
- 扩容时**桶数翻倍**，**所有元素重新哈希**

### 16.5 map 排序输出（三步）

```go
keys := []int{}                       // ① 造一个只装键的空切片
for k := range m {
	keys = append(keys, k)            // ② 把所有键捞出来（接住！）
}
sort.Ints(keys)                       // ③ 排序（要 import "sort"）

for _, k := range keys {
	fmt.Println(k, m[k])              // 按排好的键一个个取
}
```

### 16.6 map 的几条铁律

| 问题 | 答案 |
|---|---|
| map 有 `cap` 吗 | **没有**，`cap(m)` 编译报错，只有 `len` |
| 边遍历边删 | **允许** |
| 删不存在的键 | 不崩，静默跳过 |
| nil map 读 | 返回零值，不崩 |
| nil map 写 | **panic: assignment to entry in nil map** |
| 能对 map 元素取地址吗 | **不能**（见 16.3） |

---

## 17. 接口深化（10-23）

### 17.1 `any` 是什么

`any` 就是 `interface{}` 的别名，**能装任何东西**：

```go
var x any = 100          // 装整数
var y any = "hello"      // 装字符串
var z any = Student{}    // 装 struct
```

**代价：装进去就忘了它是什么。**

```go
var x any = 100
y := x + 1    // invalid operation: x + 1 (mismatched types any and untyped int)
```

编译器只知道「x 是个 any」，不知道里面是 int，所以不敢做加法。

### 17.2 类型断言：把东西取回来用

```go
var x any = 100

n, ok := x.(int)      // 意思是"我认为 x 里面是 int，帮我取出来"
fmt.Println(n, ok)    // 100 true

s, ok2 := x.(string)  // 猜错了
fmt.Println(s, ok2)   // "" false   ← 不 panic，只是 ok 是 false
```

**不加 ok 直接断言，猜错就 panic**：

```go
x.(string)   // panic: interface conversion: interface {} is int, not string
```

**所以永远写 `n, ok := x.(类型)` 这个两步形式。**

### 17.3 `%T` 和断言的三个坑

| 写法 | 结果 |
|---|---|
| `fmt.Printf("%T", x)` | ✅ 打印 x 的类型名 |
| `x.("%T")` | ❌ `syntax error: unexpected literal "%T", expected type` |
| `x.(abc)` | ❌ `undefined: abc` |

- `%T` **只能出现在 `Printf`/`Sprintf` 的格式字符串里**
- `x.(...)` 括号里**只能写类型名**（`int`、`string`、`Student`），不能写字符串字面量、不能写变量

### 17.4 多态：一个函数收所有「会这个方法」的东西

```go
type 形状 interface {
	面积() float64
}

type 矩形 struct{ 宽, 高 float64 }
func (r 矩形) 面积() float64 { return r.宽 * r.高 }

type 圆 struct{ 半径 float64 }
func (c 圆) 面积() float64 { return 3.14159 * c.半径 * c.半径 }

func 报告(s 形状) {
	fmt.Println("面积是", s.面积())
}

报告(矩形{3, 4})   // 面积是 12
报告(圆{1})        // 面积是 3.141592653589793
```

**`报告` 这个函数只写了一次，却能收矩形和圆** —— 因为它只依赖「有 `面积()` 这个方法」这个约定。

---


## 18. 包与拆包（10-24）

> 这一节是 10-24 晚上拆 `stu-manager` 时踩出来的。**「包」是 W1 顺延到 W2 的最后一块语法，也是 W2 任务③「接口分层重构」的地基。**

### 18.1 为什么必须拆包

拆包之前，`stu-manager` 里的 `main.go` 有 99 行，全是「打印提示 → 等键盘 → 打印结果」。问题有三个：

1. **测试测不到**：`go test` 不执行 `main`，`fmt.Scanln` 还会让测试卡死
2. **换存储要动全身**：读写存档的代码和界面代码混在一起
3. **名字全挤在一个包**：`Student`、三个函数、读写逻辑，谁都能碰谁

拆完的目标结构：

```
stu-manager/
├── main.go          只管界面：打印提示、读键盘、显示结果
├── funcs.go         业务逻辑：查找 / 删除 / 修改
├── funcs_test.go    测试
├── model/
│   └── model.go     package model —— 只放 Student 这个数据类型
└── store/
    └── store.go     package store —— 只放「怎么存、怎么取」
```

### 18.2 一个文件夹只能有一个包

```go
// 如果在同一个文件夹里放 a.go 和 b.go，一个写 package main、一个写 package model
found packages main (a.go) and model (b.go) in C:\Users\...\pk1
```

**要分两个包，就必须建两个文件夹。** 这是 Go 的硬规则，不是风格建议。

### 18.3 import 路径怎么拼

公式：**`go.mod` 里的 module 名 + 子目录名**

```go
// go.mod 第一行是 module stu-manager
import "stu-manager/model"   // ✅
import "stu-manager/store"   // ✅
import "./model"             // ❌
```

写成相对路径的真实报错：

```
main.go:5:2: "./model" is relative, but relative import paths are not supported in module mode
local import "./model" in non-local package
```

### 18.4 Go 导出规则：首字母大写才导出（汉字永远导不出）

这是拆包时最容易翻车的一条。规则只有一句：

> **名字的第一个字符是「大写英文字母」→ 其他包能用；否则 → 只有自己这个包能用。**

| 名字 | 能跨包用吗 | 原因 |
|---|---|---|
| `Student` / `FileStore` / `Storage` | ✅ | 首字母大写英文 |
| `Load` / `Save` | ✅ | 首字母大写英文 |
| `Name` / `Score` / `ID` / `Path` | ✅ | 首字母大写英文 |
| `save` / `path` | ❌ | 小写 |
| `保存` / `查找` / `删除` / `修改` | ❌ | **汉字没有大小写 → 永远导不出去** |

真实报错（跨包调小写方法）：

```
.\main.go:7:4: s.保存 undefined (cannot refer to unexported method 保存)
.\main2.go:10:5: st.save undefined (cannot refer to unexported method save)
```

汉字包名也不行：

```
malformed import path "baodemo/甲": invalid char '甲'
```

**实践推论**：
- 留在 `main` 包里的 `查找`/`删除`/`修改` **不用改名**（同一个包内随便叫）
- 一旦搬进 `model/` 或 `store/`，就必须是英文大写：`Student`/`Storage`/`FileStore`/`Load`/`Save`/`Path`
- 局部变量（`list`/`s`/`st`/`新list`）不跨包，随便叫

### 18.5 `Storage` 接口 + `FileStore` 实现

```go
package store

import (
	"encoding/json"
	"os"
	"stu-manager/model"
)

// 接口：一份清单，写着「想当我，必须会这几个方法」
type Storage interface {
	Load() ([]model.Student, error)
	Save(list []model.Student) error
}

// 具体实现：真的去读写硬盘上的 JSON 文件
type FileStore struct {
	Path string
}

func (f FileStore) Load() ([]model.Student, error) {
	data, err := os.ReadFile(f.Path)
	if err != nil {
		return []model.Student{}, err
	}
	var list []model.Student
	err = json.Unmarshal(data, &list)
	return list, err
}

func (f FileStore) Save(list []model.Student) error {
	out, err := json.Marshal(list)
	if err != nil {
		return err
	}
	return os.WriteFile(f.Path, out, 0644)
}
```

**四个设计要点**

1. **`Path` 必须大写**，否则 `main` 里 `store.FileStore{Path: "students.json"}` 会报 `unknown field Path in struct literal of type store.FileStore, but does have unexported path`
2. **`Load` 里用 `f.Path` 不用写死字符串**，否则换个存档路径就失效
3. **`if err != nil`（不是 `== nil`）** —— 写反了编译能过，但会「永远读回 0 人」：文件正常 `err` 是 `nil`，`err == nil` 为 true 就提前 return 空切片了。**这是本项目最隐蔽的一个 bug**
4. **错误要 `return` 出去，不要在存储层 `Println`** —— 存储层的职责只有「读到给数据 / 读不到给错误」，怎么告诉用户是 `main` 的事

### 18.6 读写存档四函数流水线（本节最重要的一张图）

初学者最容易把 `os.ReadFile` / `json.Unmarshal` / `json.Marshal` / `os.WriteFile` 记成一团。它们其实是一条流水线上的四个站点：

```
读：硬盘 --①os.ReadFile--> 字节 --②json.Unmarshal(&list)--> 名单
写：名单 --③json.Marshal--> 字节 --④os.WriteFile--> 硬盘
```

| 站点 | 干什么 | 参数 | 返回几个值 |
|---|---|---|---|
| ① `os.ReadFile(路径)` | 硬盘 → 字节 | 路径 | 2 个：`data []byte`, `err error` |
| ② `json.Unmarshal(字节, &名单)` | 字节 → 名单 | **要 `&`** | **1 个**：只有 `err` |
| ③ `json.Marshal(名单)` | 名单 → 字节 | 名单 | 2 个：`data []byte`, `err error` |
| ④ `os.WriteFile(路径, 字节, 0644)` | 字节 → 硬盘 | 路径、字节、权限 | **1 个**：只有 `err` |

**三条记法**
- **`Un` = 拆包（要 `&`），不带 `Un` = 打包（不要 `&`）**
- **"变出字节"的两个（ReadFile / Marshal）还 2 个值；"只搬运"的两个（Unmarshal / WriteFile）还 1 个值**
- `json.Unmarshal` 写成两个变量接 → `assignment mismatch: 2 variables but json.Unmarshal returns 1 value`

真实类型（实跑 `%T`）：`data` 是 `[]uint8`（就是 `[]byte`）、`list` 是 `[]main.Student`、`out` 是 `[]uint8`

**`0644` 是八进制**：`0644` = 420 十进制（八进制 6 = 二进制 110 = 读+写）。写成 `644` 会变成十进制 644，虽然不报错但语义错

### 18.7 函数签名里的括号位置

拆包时最容易懵的一点：**哪些括号是参数、哪个是返回值**

```go
func 甲() ([]model.Student, error)              // 第一个括号空 = 不要参数；第二个括号 = 返回 2 个
func 乙(list []model.Student) error             // 第一个括号有名字+类型 = 1 个参数；括号外 error = 返回 1 个
func 丙() (error)                               // 1 个返回值时括号可加可不加（实测都行）
func 丁(list []model.Student) ([]model.Student, error)  // 1 进 2 出
```

- **第一个括号 = 参数**（要进来什么）
- **第一个括号后面的那一坨 = 返回值**（要交出去什么）
- 返回值 **2 个以上必须加括号**；1 个可以加也可以不加
- 函数声明了几个返回值，**每个 `return` 就必须凑够几个**：`not enough return values have (error) want ([]model.Student, error)`

### 18.8 拆包踩过的 6 个真实报错

| 报错原文 | 病因 |
|---|---|
| `found packages main (a.go) and model (b.go) in ...` | 同一文件夹写了两个包名 |
| `"./model" is relative, but relative import paths are not supported in module mode` | import 写成相对路径 |
| `package stu-manager/mode1 is not in std` | `model` 打成了 `mode1`（数字 1） |
| `cannot refer to unexported method 保存` / `save` | 跨包的方法名首字母没大写 |
| `malformed import path "baodemo/甲": invalid char '甲'` | 汉字包名 |
| `json.Marshal (value of type func(v any) ([]byte, error)) is not a type` | 把函数当成字段类型写进 struct |

## 19. 报错原文速查

| 报错 | 意思 | 修法 |
|---|---|---|
| `no new variables on left side of :=` | `:=` 左边没有新变量 | 换成 `=`，或者换个新名字 |
| `undefined: s` | 这个变量没声明过 | 补 `:=` 或 `var` |
| `append (built-in function) is not a type` | 把 `append` 当类型名用了 | 做卡用 `Student{...}` |
| `cannot assign to struct field m[1].Score in map` | map 里的卡不能直接改字段 | 取出→改→存回 |
| `json: Unmarshal(non-pointer ...)` | `Unmarshal` 没给 `&` | 加 `&` |
| `index out of range` | 下标超了 | 检查 `< len` |
| `missing function body` | 只写了函数签名没写体 | 补 `{}` |
| `s.保存 undefined (type 存储 has no field or method 保存)` | 接口定义里没写这个方法 | 接口花括号里补 `保存(s Student)` |
| `X does not implement 存储 (missing method 保存)` | 方法名对不上（隐式实现只认字面） | 方法名改成接口里写的那个 |
| `X does not implement 存储 (wrong type for method 保存)` | 方法参数类型对不上 | 参数类型跟接口里一致 |
| `X does not implement 存储 (method 保存 has pointer receiver)` | 方法用的是指针接收者 | 传 `&X{}` 而不是 `X{}` |
| `syntax error: unexpected name Student, expected )` | struct 花括号里写了括号（当成方法写了） | struct 里只写「名字 + 类型」，方法写到外面 |
| `s redeclared in this block` | 名字撞了（`type s` 和 `var s`） | 换一个名字 |
| `syntax error: unexpected keyword for, expected expression` | `[]Student{` 后面直接接了 `for` | `[]Student{}` 是一个完整的词，第二个 `}` 必须有 |
| `syntax error: unexpected ++, expected {` | `for` 三段用了逗号 | 三段之间必须是**分号** |
| `non-name m[2] on left side of :=` | `:=` 左边不是新变量名 | 存进已有位置用 `=` |
| `invalid argument: list (variable of type []Student) is not a map` | 对切片用了 `delete` | 切片删除只能过滤重建 |
| `cannot indirect Student{…} (value of struct type Student)` | 给 struct 字面量加了 `*` | `*` 只能作用在地址上 |
| `syntax error: unexpected ], expected operand` | `[]Student` 是类型不是值 | 造盒子用 `[]Student{}` |
| `undefined: ok2` | `ok2` 从没声明过却用了 `=` | 第一次出现的变量用 `:=` |
| `no new variables on left side of :=`（又一例） | `:=` 左边全是老变量 | 换成 `=`，或加一个新名字 |
| `syntax error: unexpected name int in parameter list` | 参数之间写了句号 `int.新成绩` | 参数之间全部用逗号 |
| `--- FAIL: TestXxx` + `第 X 题还没写` | 忘了删 `t.Fatal` | 删掉骨架里的 `t.Fatal` 那行 |
| `pattern ./...: directory prefix . does not contain main module` | 目录里没有 `go.mod` | 在该目录 `go mod init 名字` |
| `panic: assignment to entry in nil map` | nil map 直接写 | 先 `make(map[...]...)` |
| `invalid operation: x + 1 (mismatched types any and untyped int)` | any 装进去就忘了类型 | 先类型断言 `n, ok := x.(int)` |
| `syntax error: unexpected literal "%T", expected type` | `x.("%T")` 括号里写了字符串 | `%T` 只能放 `Printf` 里，断言括号里只写类型名 |
| `panic: interface conversion: interface {} is int, not string` | 断言猜错且没写 ok | 永远用 `n, ok := x.(类型)` |
| `expected ')', found ','` | 括号不配对（多写了一个左括号） | 数一数 `(` 和 `)` 是不是一样多 |
| `format %d has arg ok of wrong type bool` | `%d` 接了 `bool` | bool 用 `%v`，数字才用 `%d` |
| `go: cannot run *_test.go files` | 用 `go run` 跑测试文件 | 测试只能用 `go test` |
| `assignment mismatch: 1 variable but 删除 returns 2 values` | 变量个数和返回值个数不一致 | 先念 `func` 那一行有几个返回值 |
| `found packages main (a.go) and model (b.go) in ...` | 同一文件夹写了两个包名 | 建子文件夹，一个文件夹一个包 |
| `"./model" is relative, but relative import paths are not supported in module mode` | import 写成相对路径 | 用 `module名/子目录`，如 `stu-manager/model` |
| `package stu-manager/mode1 is not in std` | `model` 打成 `mode1`（数字 1） | 改回字母 l |
| `cannot refer to unexported method 保存` / `save` | 跨包的方法名首字母没大写 | 跨包的名字首字母必须大写英文 |
| `json.Marshal (value of type func(v any) ([]byte, error)) is not a type` | 把函数当成字段类型写进 struct | struct 里只能填类型，字段类型是 `string` |

---

## 20. 本周实际写过的文件

| 文件 | 内容 |
|---|---|
| `LianXi/card.go` | 5 张卡 + 及格判断（`s.Score >= 60`） |
| `LianXi/slice3.go` | 切片塞 5 张卡 + `for` 遍历打印 |
| `LianXi/mymap.go` | map 按学号存取 |
| `LianXi/lookup.go` | map 查找 + `v, ok` + `if` 分支 + `Scanln` |
| `LianXi/lianxi1019.go` | 三组填空：图纸vs盒子 / 值→指针接收者 / append 接住 |
| `LianXi/iface1.go` | 接口 + 隐式实现 + 文件/内存两种存储共用同一个备份函数 |
| `LianXi/FeiBoNaQi.go` | 斐波那契（读代码第 1 篇，40 行逐行注释在 `reading-notes.md`） |
| `stu-manager/main.go` | 增删改查 + JSON 持久化（7 条 commit） |
| `stu-manager/funcs.go` | 从 main 拆出的三个可测函数：`查找` / `删除` / `修改`（10-21） |
| `stu-manager/funcs_test.go` | 第一个项目内测试文件：`删除`(4 用例) / `查找`(3 用例) / `修改`(3 用例) 三张表驱动测试，全 PASS，覆盖率 11.3%（10-23） |
| `Past2/练习/testing_1021/`（非仓库） | testing 练习场：`及格` / `平均分` + 普通测试 + 表驱动测试（覆盖率 77.8%） |
| `Past2/练习/1023/`（非仓库） | map/interface 手打练习：A nil map / B v,ok+delete / C map 排序 / D 接口+any+类型断言，各打 2–3 遍 |
| `model/model.go` | `package model`：`Student` 结构体从 `main.go` 迁出（10-24 拆包第 1 步） |
| `store/store.go` | `package store`：`Storage` 接口 + `FileStore` 实现（`Load`/`Save`），10-24 拆包第 2–3 步 |

---

*最后更新：2026-10-25*
