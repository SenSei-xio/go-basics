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

## 2026-10-15

- 学到：**切片 = 一叠卡片**（`var list []Student` 造盒 → `list = append(list, s)` 塞卡**必须接住返回值** → `len(list)` 几张 → `list[i]` 第几张 → `list[i].Name` 哪一栏）；**方法**（`func (s Student) IsPass() bool` —— 接收者挂在类型身上，调用 `s1.IsPass()` 不用自己写 `&`；`return s.Score >= 60` 里比较符号本身就是 true/false，不用 `if`）；**接收者为什么必须写在括号里**（① 给方法上户口，同包不许两个同名函数 ② 调用时省掉 `&`，也就省掉"忘了写 `&`"那类事故 ③ 接口只认方法 —— W2 的地基）；`stu-manager` 接上切片，能录 3 人并全部列出
- 卡住：起锚补考 struct 那块只过 **1.5/7 = 21%**，根因是**分不清「图纸（type）」和「卡片（var）」** —— 这个根一断，`Student.Name` 为什么不能跑、接收者是什么、字段大小写，全都塌；指针 `&`/`*` 第 6 次搞反；`i < len(list)` 答成"5 张卡跑 6 次"；`:=` 和 `=` 又混了一次
- 明天：复考 **指针 `&`/`*`（第 7 次）**+ 图纸 vs 卡片 + `i < len(list)` + `:=`/`=` → 新内容 **map**（给 `stu-manager` 加"按学号查找"）→ 补做拖了 6 天的**岗位认知 5 条** + **打字训练**启动
- 数字：commit 2 条（go-basics 练习；stu-manager 查看全部学生）｜ 自己写的代码约 90 行（`slice.go` / `slice1.go` / `slice2.go` / `1.go` / `2.go` / `main.go` 改造）

## 2026-10-16

- 学到：**起锚 7 题 63%**（历史最高）—— `&` 取地址 / `*` 取值**口述第 7 次终于全对**；「图纸 vs 卡片」答对（10-15 的 ❌ 翻过来了）；`i < len(list)` 答对（跑 5 次、最大 4）；`Scanln` 要 `&` 答对
- 卡住：`*p = 99` 之后以为 `x` 还是 10 → **「p 和 x 是两个名字指向同一个盒子」这层没立住**；`IsPass` 函数体写成自己调自己（`s.IsPass("及格了吗")`）→ 5 个编译错误；做卡从 `Student{"小明", 88}` 退化成 `s = append{"小明","88"}`（把 `append` 当构造器）；`:=` 还说成"赋值"而不是"出生"；Linux `df -h` 第 3 次答错（说成删除文件）
- 明天：**指针「同一个盒子」复考**（今天已看 demo）→ 方法体补 `return` → 做卡动作立住 → `stu-manager` **加「按学号查找」+ map** → 岗位认知 5 条（拖 8 天）+ 打字训练启动
- 数字：commit 1 条（daily 复盘）｜ 自己写的代码约 10 行（`10.16\1.go`，未通过编译）

## 2026-10-17

- 学到：白天休息（作息反了），晚上 23:48 做了 **6 题起锚快考（40%）** —— **指针这条根补上了**：`*p = 100` 之后 `x` 是 100、`&x == p` 是 true，两个值全对（第 8 次考）；`:=` 说不成"赋值"了，改成"生成" ✅；方法体那一行 `return s.Score >= 60` 写对、知道要在 main 里打印 ✅
- 卡住：**做卡第 3 次写成 `s = append{"小明"，88}`**（append 是塞卡不是造卡；字段名还记成 `fenshu`，真名 `Score`）；接收者 `s` 第 3 次没答对（答"函数方法"）；Linux `df -h` 第 4 次错（这次说"搜索 h 盘"）、`chmod 600` 说成"仅自己可写、都能读"
- 明天：**W1 最后一天（截止 10-18 23:59）** —— ①补根 40min（做卡写 10 遍 + 接收者换讲法）②map+按学号查找 1h ③删除 45min ④修改 45min ⑤JSON 持久化 1.5h ⑥包/多文件 45min ⑦读代码 45min ＋ 周检测口述
- 数字：commit 1 条（本条复盘）｜ 自己写的代码 0 行

## 2026-10-18

- 学到：**W1 最后一天冲刺，CRUD 全通** —— 做卡/塞卡/遍历/map/按学号查找/**删除**（`delete(m, id)` + 切片过滤重建）/**修改**（取出改副本 `m[id] = stu` 存回 + 切片同步）/**JSON 持久化**（`json.Marshal`+`os.WriteFile` 存、`os.ReadFile`+`json.Unmarshal(data, &list)` 读、`&` 不能省）全部实跑通过；**周检测通过 4/4**（数组 vs 切片：方括号里有数字=数组、空的=切片；数组不能 append；赋值复制 vs 共享；传函数数组不变切片变）；**读代码第 1 篇完成**（`FeiBoNaQi.go` 40 行逐行注释 + 脱稿口述）；新学 **`a=b` / `b=c` 顺序不能换**（换顺序 `fibo(9)` 从 55 变 256）
- 卡住：`i++` 口述时打成 `1++`（报错 `cannot assign to 1 (neither addressable nor a map index expression)`）；把"赋值"说成"返回"；JSON 持久化时说"没反应"，根因是 **VS Code 工作目录在 `D:\Past2`** 而 `os.ReadFile("students.json")` 是相对路径 → 已建 `.vscode/launch.json` 的 `"cwd": "${fileDirName}"` 修好；`m := make(map[int]Student)` 写了两遍 → `no new variables on left side of :=`
- 明天：**W2（10.19~10.25）** —— Go 语法二（接口/包/错误处理）+ 单元测试 + `net/http` 读代码①；补考欠账：接收者 `s`（第 4 次）、`df -h`（第 5 次）、`chmod 600`（第 5 次）；包/多文件拆 `main.go` 从 W1 顺延
- 数字：commit **7 条**（go-basics 6 + stu-manager 4，实际另加 docs 笔记）｜ 自己写的代码约 **180 行**（`card.go` 31 + `slice3.go` 24 + `mymap.go` 20 + `lookup.go` 28 + `main.go` 新增 ~80）
