## 2026-10-09

- 学到：双层 for 写乘法表；函数多返回值；GitHub Desktop 从零跑通建仓库 / commit / push / 改仓库名
- 卡住：swap 的注释把"值变了"写成了"内存地址改变了"；& 和 * 分不清；defer 只能抄、合上写不出来
- 明天：起床默写 swap；把 defer / slice / map 三题做完
- 数字：commit 1 条 ｜ 自己写的代码约 40 行

## 2026-10-10

- 学到：defer 后进先出、函数返回前才执行；命名返回值能被 defer 改掉（跑出来 20）；Go 只有值传递
- 卡住：defer2.go 那份是抄的，合上写不出来；Linux 和岗位认知一天没动
- 明天：Linux 20 命令 4h + docs/linux.md → Git 收尾（merge 冲突、Profile README、第二仓库）→ 岗位认知录音 1h
- 数字：commit 0 条 ｜ 自己写的代码约 25 行

## 2026-10-11

- 学到：Linux 20 个命令；Go 的 for 三段式 + "碗"（累加器）；`&` 取地址 / `*` 进门取值
- 卡住：逻辑错编译器不报，只能自己对预期结果；`echo` 后面要空格
- 明天：起锚默写 sumTo + C/D 两题 + 抽考 Linux 命令；W1 第一天，`stu-manager` 开工
- 数字：commit 2 条（linux-notes 的 Linux 命令笔记 + go-basics 的每日复盘）｜ 自己写的代码约 50 行（Go 练习 5 个文件：7/8/9/11/12.go）+ Linux 笔记 77 行

## 2026-10-12

- 学到：休息日，没有新知识（按计划允许天任务顺延 1–2 天，W1 从周二起算）
- 卡住：无
- 明天：起锚闭卷默写 sumTo + C/D 两题 + 抽考 Linux 命令 → merge 冲突 20min → `stu-manager` 开工（最小版：只增 + 查）｜弹性：打字 10min、岗位认知 20min
- 数字：commit 1 条 ｜ 自己写的代码约 0 行

## 2026-10-13

- 学到：循环三格（起点 / 继续条件 / 每轮变化）+ 循环体里 `i` 只准"读"不准"写"（改计数器 → 死循环，1.go 亲身体会）；碗的用法 `sum += i`、`p *= i`（碗在左、数在右，乘法碗初始值 1）；跳跃式 `i += 2` 和踏步式 `i++` 怎么选（中间那些数要不要单独看）；先写中文伪代码再逐行翻译；`Printf` 两列表格（`%d` 占位、`\t` 对齐、`\n` 换行）；merge 冲突完整流程（建分支 → 两边改同一行 → 合并 → `<<<<<<<` 三行标记 → 决定留哪边 → 删标记 → 提交合并 → 删分支）；项目开张 `stu-manager`（README + 菜单骨架已推上 GitHub）
- 卡住：`:=`（出生）和 `=`（改值）老混，反复 `undefined: xxx`；循环体里忍不住去改计数器/条件变量；函数名对不上（定义 `sumOdd`、调用写 `sum`）；起锚考试只过 1.5/5 —— `&` 和 `*` 第 4 次搞反，`chmod 600` 和 `df -h` 说错
- 明天：起锚复考 `sumOfEven` / `sumOdd` / `fact` 闭卷各写一遍 + `&`/`*` 口述 → 新内容 **struct（结构体）+ 方法** → `stu-manager` 实现"录入 + 查看"，目标 ≥3 条 commit
- 数字：commit 7 条（stu-manager 2 条：开张 + 骨架；git-practice 5 条：分支/合并/冲突练习）｜ 自己写的代码约 60 行（1.go / 2.go / 4.go / 3.go / main.go）

## 2026-10-14

- 学到：**struct**（`type Student struct{}` 是模板、`var s Student` 做卡、`s.ID` 点号取栏位、`Println(s)` 打整张卡）；**方法**（`func (s Student) IsPass() bool`，接收者在本体里叫 `s`，调用变 `s1.IsPass()`）；**读键盘输入**（`fmt.Print("提示：")` 不换行 + `fmt.Scanln(&s.ID)` 读一行写进那一栏 —— `&` 是门牌号，要"写进去"就必须带 `&`）；把"录入一个学生"做进了 `stu-manager`
- 卡住：`&` 和 `*` **第 5 次搞反**（最后靠"亲眼看一遍 `&x` 打印地址、`*p` 打印值"才通）；`countOdd` 写成 `sum = i`（覆盖）而不是 `count++`（计数）；`type` 和 `var` 混用；`fmt.scanln` 小写踩坑；只抄了函数签名没写函数体（`missing function body`）
- 明天：**切片 = 一叠卡片**（`var list []Student` + `list = append(list, s)` + `for` 遍历）→ `stu-manager` 实现"查看所有学生"；补考 `&`/`*` 口述 + `countOdd`
- 数字：commit 2 条（stu-manager：实现录入；go-basics：struct 与方法练习）｜ 自己写的代码约 50 行（`struct.go` + `main.go`）
