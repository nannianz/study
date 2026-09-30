package main

import (
	"fmt"
	"strconv"
)

type Site struct {
	ID int `json:"id"`
}

func modifySite(s Site) {
	s.ID = 1000
	fmt.Println("modifySite:", s) //{1000}
}
func modifySite2(s *Site) {
	s.ID = 2000
}
func main() {
	//  a:=42
	//  p:= &a
	//  fmt.Println("a:",a)
	//  fmt.Println("p:",p)
	//  fmt.Println("*p:",*p)
	//  *p = 100
	//  fmt.Println("a2:",a)
	//  fmt.Println("*p2:",*p)

	//swap 双版本：swapByValue(a, b int)（改不了外部）和 swapByPtr(a, b *int)（改得了），打印对比
	//  x,y:= 1,2
	//  swapByValue(x,y)
	//  fmt.Println("x:",x)
	//  fmt.Println("y:",y)

	//  swapByPtr(&x,&y)
	//  fmt.Println("x2:",x)
	//  fmt.Println("y2:",y)

	site := Site{
		ID: 1,
	}
	fmt.Println("site:", site) //{1}
	modifySite(site)
	fmt.Println("site:", site) //{1}

	modifySite2(&site)
	fmt.Println("site:", site) //{2000}

	//  i,j := int(3.14),int(4.56) //cannot convert 3.14 (untyped float constant) to type int
	//  fmt.Println("i:",i)//3
	//  fmt.Println("j:",j)//4

	k, err := strconv.Atoi("abc")
	fmt.Println("k:", k, err)       //k: 0 strconv.Atoi: parsing "abc": invalid syntax
	fmt.Println("err:", err != nil) //true

	//：back 里 req := &service.UpdateSiteRequest{} 为什么加 &，不加会发生什么？
	//&service.UpdateSiteRequest{} &本质是为了取得site的引用地址，这样可以改到原始的site，不加就不能修改到原site的地址
}

// func swapByValue(a,b int){
// 	a,b = b,a
// 	fmt.Println("a:",a)
// 	fmt.Println("b:",b)
// }
// func swapByPtr(a,b *int){
// 	*a,*b = *b,*a
// 	fmt.Println("a2:",*a)
// 	fmt.Println("b2:",*b)
// }

/*
===== 验收评价（2026-09-11）=====
覆盖度：4.5/5。
✅ 任务1 解引用：&a / *p / *p=100 全链路，输出注释齐全；
✅ 任务2 swap 双版本：*a,*b = *b,*a 多重赋值利落，值版改副本 / 指针版改原值对比清晰；
✅ 任务3 struct 值 vs 指针传参（live code）：modifySite(site) 外部纹丝不动、modifySite2(&site) 外部变 2000，预期输出全注释；
⚠️ 任务4 转换实验：strconv.Atoi("abc") 完整拿到 (0, err)；int(3.14) 撞上编译错误后中断——但这个错撞得有价值，见亮点②；
✅ 任务5 思考题：本质答对（取地址让 BindJSON 改到原对象）。
亮点：
① 类型定义放包级了！两次点评后当场养成，Site 和两个函数都在包级——习惯闭环；
② 撞见的「cannot convert 3.14 (untyped float constant) to type int」是条真知识：常量转换必须精确可表示（3.14 不是整数，编译器直接拒），变量转换才是运行时截断——f := 3.14; int(f) 得 3。JS 没有这对区分。补一行变量版实验即可收口。
问题与建议：
① modifySite → modifySite（2 处拼写）——工作代码里函数拼错会一路污染调用点；
② 思考题措辞：不加 & 时 BindJSON 把数据灌进「副本」、解析结果全丢（不是「修改不到地址」）——正是你 modifySite 实验的翻版；
③ gofmt 小项照旧保存自动修。
结论：通过，ch16/ch22 勾选，阶段 2 结业。
======================================
*/
