# Hot 100 Go 速通 · 内容规范

单一事实来源：`chapters/NN.json`（元数据+讲解）和 `go/`（可编译、已测试的 Go 代码）。
`build.mjs` 把两者拼成一个 HTML 页面。页面上展示的代码 = Go 文件去掉 `package` 行之后的内容，所以 Go 文件里写的就是读者看到的。

## 目录

```
go/go.mod                      module hot100, go 1.24
go/pNNNN/solution.go           package pNNNN，只放题解（LeetCode 签名，函数名与 LeetCode Go 模板一致）
go/pNNNN/types.go              （可选）ListNode / TreeNode / Node 等 LeetCode 预置类型及测试辅助函数，不展示
go/pNNNN/solution_test.go      测试，不展示
go/tpl_<slug>/template.go      章节模板代码，package tpl_<slug>，必须能编译，最好也有测试
```
NNNN 为四位补零题号（p0042）。

## chapters/NN.json

```json
{
  "id": "02",
  "slug": "two-pointers",
  "title": "双指针",
  "day": 1,
  "stars": 2,
  "when": "一两句：什么信号出现时用这个套路，它为什么成立。",
  "pitfalls": ["章节级易错点，1–3 条，每条一句"],
  "template": { "title": "模板名，如：相向双指针 / 同向快慢指针", "dir": "tpl_two_pointers" },
  "problems": [
    {
      "id": 42,
      "title": "接雨水",
      "difficulty": "困难",
      "hot": true,
      "insight": "核心思路 1–2 句，可用「」强调关键词。讲清楚为什么对，而非复述代码。",
      "pitfall": "这题最容易错的一点，或面试官会追问的变体。一句。",
      "time": "O(n)",
      "space": "O(1)"
    }
  ]
}
```

- `difficulty` 只能是 `简单` / `中等` / `困难`，与 LeetCode 一致。
- `hot`：面试真实高频（每章约 1/3 题标 true，宁缺毋滥）。
- `stars`：章节面试权重 1–3。
- 题目顺序 = 学习顺序（由易到难、前题为后题铺垫），不必照官方顺序。
- `template` 可省略（若这章没有统一模板）。
- 语气：简洁、技术、中文，像高手写给自己的速记卡。不要套话、不要 emoji。

## Go 代码要求

- 写面试里会写出来的最优/标准解；短、地道、可读（Go 1.24：可用内置 `min`/`max`、`slices`、`container/heap`）。
- 变量命名清楚，必要处一两行中文注释，不写长注释。
- 设计类题（LRU、Trie、MinStack、MedianFinder）照 LeetCode 的 `Constructor` + 方法签名。一个包一道题，所以 `Constructor` 不冲突。

## 测试要求（重点）

- 每题必须有 `solution_test.go`，`go test ./...` 全绿，`go vet ./...` 干净。
- 包含 LeetCode 官方样例。
- **能写暴力解的题必须做性质测试/对拍**：用固定种子的 `math/rand/v2` 生成数百组小规模随机输入，和一个朴素暴力解（写在 test 文件里）比较结果。或者检查不变量（如：排序结果有序且为原数组的排列；返回的组合各自合法且无重复且数量等于暴力枚举）。
- 答案不唯一的题，检查合法性而不是比较具体值。
- 测试要能真的抓出错误：写完后可以临时改坏一处题解确认测试会红，再改回来。
