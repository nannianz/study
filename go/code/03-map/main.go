package main

import (
	"fmt"
)



func main() {
	//  words :=	strings.Fields("go is fun go is fast")
	//   fmt.Println("words:", words)
	//  m := make(map[string]int)
	//  for _,w := range words {
	// 	fmt.Println("w:", w)
	// 	m[w]++
	//  }
	//  fmt.Println("m:", m)
	// 输出：map[go:2 is:2 fun:1 fast:1]
	//  fmt.Println("m[go]:", m["go1"])

	//  m := map[string]int{
	// 	"go": 0,
	//  }
	//  v,ok := m["php"]

	//  fmt.Println("v:", v) // 0
	//  fmt.Println("ok:", ok)// false

	// v1,ok1 := m["go"]
	// fmt.Println("v1:", v1) // 0
	// fmt.Println("ok1:", ok1)// true

	// if m["go"]
	// {

	//  }

	// var m map[string]int

	//  fmt.Println("m:", m["a"])
	//  fmt.Println("m1:", m["a"] = 1)
	// m["a"] = 1 //nil dereference in map update

	// m := make(map[string]int)
	// m["go"] = 0
	// m["php"] = 1
	// m["js"] = 2

	// for k,v := range m {
	// 	fmt.Println("k:", k, "v:", v)
	// }

	// for k,v := range m {
	// 	fmt.Println("k1:", k, "v1:", v)
	// }

	for i,c := range "中国go" {
		fmt.Println("i:", i, "c:", c,"s:", string(c))
	}
	//为什么 Go 故意把 map 遍历顺序随机化
	// 因为 map 是无序的，所以遍历顺序是随机的
	// 如果需要遍历顺序，需要使用 sort 包
	// sort.Strings(words)
	// fmt.Println("words:", words)
	
}

/*
===== 验收评价（2026-09-09）=====
覆盖度：6/6 任务全部动手（任务 3 的 if m["go"] 写了但没记报错——结论：non-boolean condition，int 不能当 if 条件，这就是「必须用 ok」的编译器侧论证）。
亮点：① ok 实验做了正反用例（php→0/false 配 go→0/true），同一个零值 0 靠 ok 区分存在与否——v, ok := m[k] 存在的全部理由被钉死；② nil map 读安全写 panic 亲手实测并留了注释。
问题与建议：① 第 19 行 m["go1"] 应是 m["go"]（typo：go1 不存在，所以那行打印的其实是 0/false）；② 思考题答案停在实用层（无序→用 sort），设计层补全：故意随机化是防止代码依赖遍历顺序——顺序若稳定但依赖底层实现，换实现会静默 break；随机化把隐藏的顺序依赖变成立即暴露的 bug。Go 1.0 引入此设计正是看了 JS/Python 的顺序踩坑史。
思考题：已补全（见 ②）。
======================================
*/
