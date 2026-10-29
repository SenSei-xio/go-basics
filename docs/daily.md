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

## 2026-10-19

- 学到：**W2 第 1 天，起锚 8 题清账后 7/8 = 88%** —— Q1 `a=b`/`b=c` 顺序变体全对（换序 `fibo(9)` = 256）；Q4 map 改 struct（`m.Name` 变体）全对；**Q7 `chmod 600` 第 5 次终于过**（⚠️→✅）；Q2 `v, ok` 说对"ok 是判定"；清题时 **指针接收者 `(s *Student)` 自己写出来了**（昨天 Q8 只知结果不知写法）；`[]Student{}` vs `[]Student`（图纸 ≠ 盒子）当堂补回
- 卡住：`append` 必须接住**又忘一次**（`d := append(b, 1)` 以为 `b` 会变长，实跑 `b` 还是 `len 0`）；`*Student{"小明",40,1}` 把 `*` 放错地方（报错 `cannot indirect Student{…}`）；`var b struct`（`struct` 是关键字不是类型名）；`_ = append(...)` 当成"改对"（`_` 是丢弃，不是接住）；切片字面量漏内层花括号两次（`[]Student{"小明",70,1}` 报 `cannot use "小明" ... as Student value`）
- 明天：W2 正式开工 —— Go 语法二（接口/包/错误处理）+ 单元测试；补考欠账：`append` 接住（D+1 必考）、`[]Student{}` vs `[]Student`、`*` 只能作用在地址上；读代码第 2 篇（W2 起每周 2×45min，周二/周五）
- 数字：commit 待提交（本条复盘）｜ 自己写的代码约 **92 行**（`LianXi/lianxi1019.go` 三组填空，最终全绿）

## 2026-10-20

- 学到：**起锚 10 题 4.75/10 = 48%，清题后重答 Q1–Q6 得 4.5/6** —— Q1 `append` 不接住终于是自己答出来了（「b=[],lenb=0，因为b空了只会返回一个空盒子」）；Q3 `*` 不能给 struct 字面量加星，答「把星号删掉就行」全对；Q7 `Marshal` 不要 `&` / `Unmarshal` 要 `&` 满分（唯一把「读取/写入」和「值传递」连起来的答案）；Q9 `chmod 600` 第 3 次复考仍对；Q10 `a=b`/`b=c` 顺序第 3 次复考仍对。**切片删除「造·走·判·塞·换」五字诀练到第 2 遍闭卷全对**（`del_lianxi.go` 填 5 行全中 → `del2.go` 只给口诀闭卷 5 行全对）
- 卡住：`del1.go` 第一次闭卷 **7 处错**（`[]Student{` 后面接了 `for` → `syntax error: unexpected keyword for, expected expression`；`for i:=0;i<len(list),i++` 逗号 → `unexpected ++, expected {`；`list[i].Id` 小写 d；`新list=list` 左右反）—— **根因是没分清「`[]Student{}` 是一个完整的词」**（`del1.go` 原始错误版存档在 `D:\Past2\练习\del1_10-20初版7处错.go`）；`v.score` 小写（跟 10-17 `fenshu` 同一个坑）；`m[2] := v` 用 `:=` 存进已有位置 → `non-name m[2] on left side of :=`；**「stu 是函数/参数」第 3 次错**；指针接收者「给 score 前一个星」第 2 次不会写；`新list := list` 以为是造空盒子（实跑证明是复制名字，`append` 后「一个人都没删掉，反而多了一个」）
- 明天：**按 `W2周计划.md` 走**（interface 2h + slice 底层 2h 已提前到今晚做完，明天份额往前挪）—— 10-21 起锚 10 题（含 8 条必考池）+ 九九乘法表第③级复验（欠 11 天）→ testing 表驱动测试 2h（W2 核心②）→ 岗位认知 5 条（欠 7 天）→ `del3.go` 第 3 遍闭卷限时 5–10min
- 数字：commit 1 条｜ 自己写的代码约 **85 行**（三份切片删除练习）
- 注：10-20 22:48 起练习文件改放 `D:\Past2\练习\`（不进 git）；`LianXi/` 只放「要传到 GitHub 的成型练习」。详见 `protocol.md` §10.10

- 补记（22:00–22:43，W2 核心① interface + 核心② slice 底层）：
  - **interface 2h**：从「没有接口时备份函数得写两个」开始，到「接口 = 一份必须会什么的清单」，最后写 `LianXi/iface1.go` —— 三题（接口签名 / 文件存储方法 / 内存存储方法）**一次全对**。实跑「传 文件存储」和「传 内存存储」都能用同一个 `备份` 函数，一个字没改。
  - **暴露的乱**：第一次交上来的代码 4 个错 —— `type 文件存储 struct{ 保存(Student) }`（能编译但那是**字段**不是方法）、`struct{ 保存(s Student) }` → `syntax error: unexpected name Student, expected )`、接口名 `s` 跟变量 `s` 撞车 → `s redeclared in this block`、`内存存储` 完全没写方法。**根因是没分清三种花括号各装什么**（struct 装数据 / interface 装要求 / func 装动作）。
  - **隐式实现验证题答错**：问「方法叫 `Save` 不叫 `保存` 能不能塞进接口」，答「能，Go 会自动认」→ 实跑 `硬盘 does not implement 存储 (missing method 保存)`。讲透「Go 只认字面不认意思」。
  - **slice 底层 2h**：cap 实跑 `4 → 8 → 16 → 32 → 64`（**我第一次给的 1→2→4→8 是凭记忆编的，错了，已当面认错纠正**）。用户观察到「cap 呈指数上涨但卡在 8 了」，是被我错误输出误导。讲清「slice = 指针 + len + cap 三样」、「翻倍是用空间浪费换搬家次数（不翻倍塞 1000 个要复制 49 万次，翻倍只约 1000 次）」、「append 有空位就共享、没空位就搬家」。
  - **验收 2 题第一遍全错**：Q1 `make([]int,2,5)` 后 append → 答「不会互相影响」，实际 `a[0] = 999`（cap 有空位，共享）；Q2 `[]int{1,2}`（cap 满）后 append → 答「也不会」，实际**确实不会**（这次蒙对了结论）但理由是 cap 满了搬家。第二遍复述「append 没位置才会搬家」✅ 完全正确。
  - **数字**：commit 待提交｜今天自己写的代码约 **150 行**（切片删除三份 85 行 + interface 练习 60 行 + 填空 5 行）

## 2026-10-21（周三 · W2 第 3 天）

### 学到
- **testing 入门**：测试文件三规矩（`_test.go` / `Test` 开头 / `t *testing.T`）、最小骨架 5 行（`got` / `want` / `if got != want` / `t.Errorf`）、`go test -v` / `-cover`。
- **表驱动测试**：`cases := []struct{...}{...}` + `for _, c := range cases`，把「写四遍」变成「写一遍跑四次」。
- **`for ... range` 彻底讲透**：三段式 `for i := 0; i < len(x); i++` 与 `for i, v := range x` 完全等价；`for v := range x` 拿到的是**下标**不是值；只要值必须 `for _, v := range x`。
- **`t.Fatal` vs `t.Errorf`**：Fatal 立刻停（后面一行不跑），Errorf 记账继续（表驱动必须用它）。
- **`go run` 不跑 Test，`go test` 不跑 main** —— 所以 main 里的 `Scanln` / `Println` 让测试无法进行。
- **重构**：把 `stu-manager/main.go` 里的查找/删除/修改拆成 `funcs.go` 三个函数，`main` 只负责输入输出。跑 `go build ./...` / `go vet ./...` 全干净，功能实跑没变。

### 卡住
- **`t.Fatal` 忘了删**：练习第 2 题 cases 表和 for 循环全写对了，但函数第一行还留着 `t.Fatal("第 2 题还没写")`，后面一行都没跑到 → `--- FAIL: 第 2 题还没写`。
- **`=` 和 `:=` 选错**：`ok2, list = 修改(...)` → `.\main.go:82:2: undefined: ok2`。`ok2` 第一次出现必须用 `:=`；`:=` 左边只要有一个新变量就合法（`ok2` 新、`list` 老也行）。
- **参数之间写成句号**：`func 修改(m map[int]Student,list []Student,id int.新成绩 int)` → `syntax error: unexpected name int in parameter list; possibly missing comma or )`。
- **旧代码忘记删**：重构时 `main.go` 里原来手写的删除段（71-77 行）和修改段（91-100 行）没删，被新函数调用又做了一遍。
- **花括号套错**：把「打印全部 + 保存 students.json」包进了 `if ok2` 里面，导致「修改的学号不存在」时**不保存数据**。已改成 if 只包成功/失败两行。
- **`for ... range` 一开始完全没懂**：用户原话「我大概清楚是用 if 判定，但是 for 的遍历我更是一点没看懂」。用「先手动写三遍 `和 = 和 + 分数[0/1/2]` → 看出只有下标在变 → 合成一行」讲通。

### 明天（10-22）
- 先补 `stu-manager` 的 `Test删除` 表驱动测试（今天剩的最后 10 分钟）
- W2 核心：map 底层（哈希/冲突/扩容/nil map panic）+ interface 深化（多态/any/类型断言）
- 补 `students.json` 累加 bug（连跑两次 2→5→8 人）
- 欠账：岗位认知 5 条（拖 8 天）、`del3.go` 第 3 遍

### 数字
- 起锚 10 题 **4.2/10 = 42%**，清题后 5 道全通
- `D:\Past2\练习\testing_1021\` 覆盖率 **77.8%**（`及格` 100% / `平均分` 100% / `main` 0%）
- 今天自己写的代码：`funcs.go` 约 40 行 + `练习_test.go` 约 35 行 + `funcs_test.go` 骨架 30 行 ≈ **105 行**
- commit：`go-basics` 1 条 `docs: 新增 testing 测试一节（§15）+ 10-21 复盘`；`stu-manager` 2 条 `feat: 拆分可测函数（查找/删除/修改）+ go.mod`、`test: Test删除 表驱动测试骨架（4 用例，暂 Skip）` —— 待用户手动 Push

## 2026-10-22（周四 · W2 第 4 天）

- 学到：**休息日**（整天忙别的，0 学习时长）。晚上做了三件复盘：① 把 W2 剩余工作量逐项核算出来 —— 3 天要 **≈20.6h**，可用只有 14–16h，**差 5–6h** ② 排了 10-23 冲刺版：本职 **6h15m**（打字 15 + 起锚 40 + map 底层 90 + interface 深化 90 + testing 收尾 60 + 累加 bug/平均分 60 + 收锚 20）+ 加速包 **5h45m**（读代码 45 + 拆包 150 + fileStore 150，酌情匀量）③ 备好 10-23 三份资料（起锚题 8 题 / 答案 / 手打练习 6 组）
- 卡住：没有卡住的技术点。唯一的坑是**我自己的排期账算错过一次**（先说 13.5h，实际 12h；先说本职 6.5h，实测 6h15m），已当场重算修正并写进 `W2周计划.md` 第十一节
- 明天：**按 10-23 冲刺版走** —— 本职 6h15m 是硬承诺（用户自己确认做得到），加速包不绑死、状态好就抢一块。**硬线「不连休两天」明天必须破掉，至少 1 条 commit**
- 数字：commit **0 条**（休息日）｜ 自己写的代码 0 行 ｜ 今日产出全在 `D:\Past2\学习档案\` 的规划文件里（`W2周计划.md` 445 行、`10-23起锚题.md` 103 行、`10-23起锚答案.md` 183 行、`10-23手打练习.md` 106 行、`protocol.md` §12.3 新增「砍单只在当晚定」）

## 10-23（周五 · W2 第 5 天）

**学到**
- **map 底层**（19:24–22:34）：8 个桶、哈希三步（算哈希 → %8 定桶 → 桶内存键值）、遍历随机（随机起始桶，防依赖实现）、扩容规则（桶数×6.5 触发，翻倍重排）、**map 元素不可寻址的唯一真原因 = 位置会变**、map 没有 cap 只有 len、nil map 读 OK 写 panic、map 排序三步
- **interface 深化**：`any` = `interface{}`，能装任何东西但装进去就忘了类型；类型断言 `n, ok := x.(int)` 两步写法；不加 ok 直接断言会 panic；`%T` 只能放 Printf；多态（形状/矩形/圆，一个 `报告` 函数收所有会 `面积()` 的东西）
- **testing 收尾**：`Test删除`（4 用例）+ `Test查找`（3 用例）+ `Test修改`（3 用例）三张表驱动测试全部 PASS，`coverage: 11.3%`（原 3.5%）
- **为什么测 func 不测 main**：main 里 7 个 `fmt.Scanln` 无人敲键盘、`fmt.Println` 结果拿不回来；要把逻辑搬出来才能测
- **返回值顺序**：先念 `func` 那一行，括号里什么顺序，左边就什么顺序。`查找` → `(Student, bool)`；`修改` → `(bool, []Student)`

**卡住**
- 起锚 8 题只有 **31%**（2.5/8）：`t.Fatal` vs `t.Errorf` 全忘、`for v := range` 单变量拿下标不知道、覆盖率不知道、`map 改 struct` 第 3 次错（顺序全反 + 用 `:=`）
- map 底层**我一开始说"不用手打"，用户连续三次纠正才改**（他说「我更想每节知识点都敲好几次来熟悉」）
- 手打练习反复踩格式错：`import` 写在 main 内、`for_` 粘一起、`krys` 拼错、`fmt.println` 小写、`x.("%T")`、`:= 重复用`
- 写测试时两个错：`t.Errorf((` 多一个左括号、`%d` 接了 `bool`（应该 `%v`）

**明天**
- 累加 bug + 平均分/排名（今天顺延的）
- 拆包 + 把覆盖率从 11.3% 提到 60%（W2 任务③）
- 读代码 `net/http` 第 2 篇（加速包，酌情）

**数字**
- 起锚 2.5/8 = 31%（昨天不算，10-22 休息日 → 这是 D+2）
- 手打练习 4 个文件各 2–3 遍（A/B/C/D）
- 测试 3 个函数 10 个用例全 PASS，覆盖率 3.5% → **11.3%**
- 今天自己写的代码约 200 行（A~D 练习 120 行 + 三个测试 80 行）
- 时间：17:14 开工 → 23:50+ 收工



## 10-24（周六 · W2 第 6 天）

**学到**
- **起锚 5.3/10 = 53%**（比 10-23 的 31% 回升 22pp）。唯一满分题：`range` 三种写法 + 主动指出 `for v := range` 最易踩坑
- **重构：从 main 里剥逻辑成函数**（用户 10-23 主动要求先复习）——三遍全过：第 1 遍 ≈35min（带修）、第 2 遍 ≈6min、第 3 遍 **≈5min 一字不差**，三遍 `go run` 都是 `平均分是： 66`
- 核心口诀：**「`fmt.` 开头的一律留在 main，剩下的才考虑搬」**；配套声明（`var del int` 只服务 `fmt.Scanln`）也跟着留
- 参数/返回值判定法：看变量第一次出现在哪 —— `x := 0` / `var x int` 是函数自己造的（不进参数）；函数里用了但函数里没造过 → 外面造的（必须传进来）；函数里改过且要让外面看到 → 必须走返回值
- **拆包（接口分层）第 1–3 步全部完成并提交**：`model/model.go`（`Student` 迁出）+ `store/store.go`（`Storage` 接口 + `FileStore` 的 `Load`/`Save`）
- **Go 导出规则讲透**：首字母必须大写英文字母才导出；汉字没有大小写 → 汉字开头的类型/函数/字段/方法永远导不出去。跨包要用的：`Student`/`Name`/`Score`/`ID`/`Storage`/`FileStore`/`Load`/`Save`/`Path`
- **`store` 的 33 行逐行讲了两遍**（用户原话「这三十三行代码我几乎全不认识」），再让他闭卷重打

**卡住**
- **字段类型填成「一个 json 方法」**：问 `FileStore` 里装什么，答「一个 json 的方法，可以把东西转换成 json 格式的方法」→ 把 `json.Marshal` 当成字段类型。真实报错 `json.Marshal (value of type func(v any) ([]byte, error)) is not a type`。**struct 花括号里只能填类型，函数/方法填不进去**
- **`if err == nil` 写反**（致命但编译能过）：存档里有 2 人，实跑读回 `len=0 err=<nil>` → 程序不报错、数据没了。改成 `if err != nil` 后 `len=2`
- **`Save` 签名写成 Load 的样子**：`func (f FileStore) Save() ([]model.Student, error)` → 三连报 `undefined: list` + `not enough return values`（×2）。正解 `Save(list []model.Student) error`（1 进 1 出）
- **`err = json, json.Unmarshal(...)` 多一个逗号** → `use of package json not in selector`
- **`path` 小写**（第 7 个拼写类错误，前六个：`fenshu`/`score`/`Id`/`Save`/`mian`/`Scote`）→ `unknown field Path ... but does have unexported path`
- **「四个函数我总乱」**：用户原话「os.ReadFile, json.Unmarshal, json.Marshal, os.WriteFile，哪怕有选择我也总是乱」→ 换讲法：不是四选一，是**一条流水线上的四个站点**；`Un` = 拆包（要 `&`），不带 = 打包（不要 `&`）；「变出字节」的两个还 2 个值、「只搬运」的两个还 1 个值
- **用户自评**：「写完了你看看，虽然还是看了几眼，看了 data 那一行和 if 那一行，out 和 if 那一行的写法，剩下的都还好」→ 闭卷重打 `store.go` 与提交版 **diff 只有 1 个空行差异**，实质零错误

**明天**
- 拆包第 4 步：`main.go` 改成只依赖 `store.Storage`（删掉老读写 20 行）——这是 10-25 核心① 的主体
- 拆包第 5 步：`go vet` 23 条 unkeyed fields 修正 + `gofmt -w .` 统一换行
- 修累加 bug + 平均分/排名（10-23、10-24 连续两天顺延）
- 覆盖率从 11.3% 提到 60%（拆包本身不解决，需抽 IO 层）
- 周检测：口述 slice 扩容 + interface 多态

**数字**
- 起锚 **5.3/10 = 53%**
- 重构练习 **3 遍**（35min → 6min → 5min）
- 拆包 **5 步走了 3 步（50%）**，用时 ≈85min / 计划 150min
- commit **1 条**：`db8a418 feat: 拆出 model 包与 store 包（Student 迁到 model，新增 Storage 接口 + FileStore 实现）`（5 files, +70 −29）
- 时间：19:16 开工 → 00:20 收工

---

*最后更新：2026-10-27*

## 10-26（周一 · 巩固日 · 不学新知识）

**性质**：用户 m06366 决定的巩固日（protocol §13），全天不学新知识、不出起锚卷、不赶进度，只把旧东西打到「闭卷 + 限时 + 连对 3 遍」。

**打到哪了**：A1–A9 九段全部 3 连对（A4 多打 1 遍，共 4 遍），A10 只算熟悉、明天从零重过。用时 17:45–23:35，约 5 小时 50 分。

**九个硬骨头全部当场打穿**
- A1 切片删除五字诀：造·走·判·塞·换，第 1 遍漏函数壳、第 2 遍漏返回值，第 3 遍闭卷全对
- A2 map 改 struct 三步：第 1 遍三步全对但写在函数外（`syntax error: non-declaration statement outside function body`），第 2/3 遍全对
- A3 做卡三种造法 + 字段名：`Student{"小明",70,1}` / `[]Student{{...}}` / `map[int]Student{1:{...}}`，一次全对
- A4 表驱动测试：`fmt.Println` 不算失败、必须 `t.Errorf`；最简版 `t.Errorf("%s 错了", c.名字)` 就够用；故意改错能被抓到
- A5 指针接收者：只改一个字符 `func (s *Student)`，三次全对（10-19/10-20 两次都答错位置）
- A6 `v, ok` + `delete`：第 1 遍踩 `:=` 重复，第 2/3 遍零报错
- A7 for 三段 + 三种 range：口诀「循环体里用了哪些变量，for 那一行就得接住哪些」
- A8 四种造切片 + nil 判断：`var` / `[]int{}` / `[]int(nil)` / `make([]int,0)`，len cap 都是 0 但 `== nil` 不同
- A9 `:=` vs `=` 三种情况：判断 5/5 一遍全对（两次起锚错过之后首次），手打改错全对，严格版 `x, z := 200, 3` 通过

**卡住 / 教训**
- A10 `os.WriteFile` 漏了两个参数（要 `(文件名, 内容, 权限)`），写成 `os.WriteFile(路径 & list)`
- A8 我第一次出题出错（用 `Student` 造装数字的切片，`cannot use 0 (untyped int constant) as Student value`）
- A3 第 3 遍卷面没给数据，我误判用户「手滑」，已查证改判并道歉
- 我自己的规则：出卷必须给全数据，不能让用户背数据

**明天（10-27）**
- A10 从零重过 3 遍
- A11–A15（Store 接口 / MemStore / 从 main 剥函数 / 四函数流水线 / 包与跨包类型）
- B1–B10 十个原理口述
- W3 并发入门（goroutine / channel）正式开课

**数字**：9 段通过 / 约 5h50m / 0 commit（巩固日不产出代码提交）

---

## 10-27（周二 · W3 第 2 天 · 巩固日续做 + 并发开课）

**学到**

- goroutine：`go 函数()` 派后台任务；`main` 一结束全部消失；顺序不固定（同一程序跑两次甲乙谁先都可能）
- 时间单位：`time.Second` / `Millisecond` / `Microsecond`；`1200 * time.Microsecond` = 1.2 毫秒（想睡 1.2 秒会一行输出都没有）；浮点常量乘 Duration 报 `truncated to int64`
- `sync.WaitGroup`：`Add(n)` + `defer wg.Done()` + `Wait()`；必须传 `*sync.WaitGroup` 和 `&wg`
- 数据竞争实证：`总数++` 各 1000 次，跑 8 次结果 `24232 / -35063 / -19139 / -66785 / -33253 / -49609 / 30103 / -53394`，一次都不是 0
- `x++` 是三步（读、算、写）；加 `sync.Mutex` 的 `Lock/Unlock` 后跑 5 次稳定 0
- channel：`ch <- 42` 送、`x := <-ch` 接；无缓冲要手递手；`make(chan int, 3)` 有缓冲（实测 `len = 3 cap = 3`，读出来是 1 先进先出）
- `close(ch)` + `for v := range ch` 自动收完退出；关了之后再读是 `0 false`
- `select` 每次只走一个 case、多个就绪时随机选、没 default 且全空就阻塞
- `select` + `time.After(100 * time.Millisecond)` 做超时；`select` + `default` 做不阻塞版本
- 巩固日续做：A10 从零 3 遍（第 1 遍 4 分 20 秒全对）、A11 Store 接口 3 遍、A12 MemStore 3 遍（含 CounterStore 变体）

**卡住 / 教训**

- `go` 忘了写 → 输出完全串行（G1 第 1 遍）
- `time.Sleep(1200 * time.Microsecond)` → 一行输出都没有（把毫秒写成微秒）
- WaitGroup：`go 干活(i, &wg)` → `undefined: 干活`（把别的演示里的函数名搬过来了）；`go 中($wg)` → `invalid character U+0024 '$'`
- WaitGroup 第 2 遍漏 `wg.Add(3)` → `Wait()` 立刻返回，输出只有两句 main
- `var wg sync.WaitGroup{` 多花括号 → `syntax error: non-declaration statement outside function body`
- channel 方向写反 `C <- ch` → `undefined: C`
- `select` 只写一次没套 for → 只处理一个就结束（C4 第 2 遍）
- 练习文件建了没真生成 / 漏占位行（C2、C3、C4b 各一次），被用户连续抓到 → 之后每次都 `cat -n` 复核

**明天（10-28）**

- C4c 第 3 遍（无缓冲双生产者）→ 5 生产者 2 消费者（W3 第 1 项验收）
- sync 6h（W3 第 2 项）
- 巩固日剩余 A13–A15 + B1–B10
- 今天净学约 2h（16:21–19:48），计划 5.5h，缺口顺延

**数字**

- 今天 0 代码 commit（练习都在 `Past2`，按 §10.10 不进 git），只有文档提交
- 两仓库 commit：本文档一条
- 巩固日进度 12 / 15
- 净学时约 2h / 计划 5h30m

*最后更新：2026-10-27*

## 10-28（周三 · W3 第 3 天 · 巩固日收官 + P1 手写通过）

**学到**

- A13 / A14 / A15 三块全部 3 连对，巩固日 **15 / 15 收官**：从 main 剥函数（16min→9min→4min）、四函数流水线（6min）、包 + 跨包类型（1min）—— **A15 第 3 遍建 model + store 两个包，汉字函数名导不出踩了一次，自己补齐 store.go**
- `error` 到底是什么：`err == nil` 是成功、`err != nil` 是失败；`bool` 只说对错、`error` 能装「为什么错」（实跑 `os.WriteFile` 失败返回 `*fs.PathError`）；`return` 不是「检查」是「交报告 + 走人」
- 面对空函数的三步法：① 用中文说这函数干嘛（一个动词）② 只有「拆/装」同时出现才用 Marshal/Unmarshal ③ 按顺序排（读→拆 / 装→写）
- **P1 写出来了：5 生产者 × 10 + 2 消费者 = 50，连跑 3 次都是 50**，两个消费者都确认在干活

**卡住**

- P1 第一遍 4 个错：消费者少一个 `}`（致命，报错却指在 `func main` 那行）、`i<=3`（应为 `每个生产量`）、`编号+100+i`（少了乘号，编号重复）、`i.ch`（点号应为逗号）—— 除第 1 处外全是手滑
- 「知道怎么写但不知道往哪写」（A15 时说的）：不是忘了名字，是「看到 Save 不知道这函数体该长什么样」，用四函数流水线三步法解决
- 我（AI）又犯了时间推算的错误：18:37 跑过一次 `date`，之后 18:44 / 18:47 全靠记忆加分钟，实际已 22:45，差 3 小时 58 分 —— 违反 §14 时间纪律，已当面认错

**明天**

- W3 主线：5 生产者 2 消费者验收已完成 → 并发下载器开工（14h 大块）
- 巩固日剩余 B1–B10 十个原理口述（10 分钟一个，可分散到每天）
- 两仓库今日 0 commit → 本文档一条补上

**数字**

- 今天净学约 7h40m（15:14 → 22:52，中间吃饭约 40min）
- 巩固日 **15 / 15**（A1–A15 全部 3 连对）
- P1 一次验收通过（改 4 处后），实跑 3 次都是 50
- 代码行数：`Past2/练习/1028/P1.go` 89 行
- 今日 commit：1（本文档）


## 10-29（周四 · W3 第 4 天 · 补起锚 + 清账 + STEP1/2/3）

- 学到：W3 起锚补齐（8 题 3.3/8）；清账 `for _, v := range` 拿值 / 四函数方向 / `func` 不能嵌套 / `os.WriteFile` 第一参数是路径；STEP1 单文件下载 17 分钟；STEP2 并发下载（WaitGroup + go + 三个文件 1.03 秒）；STEP3 自定义 io.Writer 做进度 + 3 次失败重试
- 卡住：STEP1/2/3 基本都是看答案才会，自写不出来；STEP3 掌握最差（写的时候结构全错：语句飞出函数、接收者大小写、Println 带占位符、全角逗号）；`-race` 装好但 sync 没讲
- 明天：起锚（含 15 道补考：10 填空 + 5 自写）+ STEP1 无提示重打 + Q2 WaitGroup 漏 Add / Q4 select 超时补考 + B1–B10 十个原理 + sync 深入（用已装的 race.sh）
- 数字：commit 3 条（go-basics 本篇 + go-downloader 首提交 + stu-manager 无改动）｜ 自写代码约 300 行（STEP1 46 行 / STEP2 78 行 / STEP3 130 行 / 清账 45 行）｜ 起锚 3.3/8 = 41%

### 今天的三个事实（诚实记录）

1. **起锚断档 3 天被我忘掉**（10-27/28/29），是用户 m08195 抓到的。已写入 `protocol.md §15`。
2. **卷子里写了答案式注释**（STEP1 的 TODO 里带着「resp, err := ...」），按 §15.3 等于填空题。已写入 `protocol.md §16`，STEP1 已作废，10-30 重打。
3. **STEP3 是「讲完才写」而不是「自己写出来」**：讲解六步（第 0–5 步）给了完整代码，用户照着写。真实独立水平待 10-30 用无提示卷验证。

*最后更新：2026-10-29*

