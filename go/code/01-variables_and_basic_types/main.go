package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 第一课 · 变量与基本类型
	// var a = 1
	// fmt.Printf("类型%T\n", a)
	// fmt.Println("字节数", unsafe.Sizeof(a))
	// fmt.Println("地址", &a)

	// a := 1
	// fmt.Println(a)

	// var name string = "smb"
	// var name2 = "smb2"
	// name3:= "smb3"
	// fmt.Println(name)
	// fmt.Printf("%s\n", name2)
	// fmt.Printf("%s", "name3 "+name3)

	// var i int
	// var s string
	// var b bool

	// name  = 1

	fmt.Println(len("中")) //3
	fmt.Println(len("中1")) //4
	fmt.Println(len("中abc")) //6
	fmt.Println(len([]rune("中abc"))) //6
	// ① Go 为什么坚持「声明不用就报错」？
	// 因为 Go 语言的变量是静态类型，必须在使用前声明，否则会报错。
	// 这是为了防止变量的未使用，导致程序运行时错误。

	//② len("中国abc") 是多少？len([]rune("中国abc")) 又是多少？为什么不等？
	// len("中国abc") 是 6，len([]rune("中国abc")) 是 12。
	// 因为 Go 语言的字符串是 Unicode 编码，每个中文字符占用 3 个字节，而每个 ASCII 字符占用 1 个字节。
	// 所以，len("中国abc") 是 6，len([]rune("中国abc")) 是 12。

	//第二课 · 常量、iota 与运算符
	const pi = 3.14
	// pi = 3 
	// const t  = time.Now()

	const (
		mon = iota //1
		tue
		wed
		thu
		fri
		sat
		sun
	)
	fmt.Println(mon, tue, wed, thu, fri, sat, sun)

	fmt.Println("string(65):", string(65))
	fmt.Println("string(20013):", string(20013))
	fmt.Println("strconv.Itoa(65):", strconv.Itoa(65))
	x := 1
	x ++ 
	fmt.Println(x)
	// y:= x++
	// fmt.Println(y)
	fmt.Println(x)
}
/*
===== 验收评价（2026-09-04 验收 · 2026-09-09 补记细评）=====
覆盖度：第一课 6 任务 + 第二课 4 任务基本覆盖（var 三姿势/零值/铁笼/洁癖/长度/truthy、const/iota/转换/++）。
亮点：① 自发探索 unsafe.Sizeof 与 &a（打印变量地址、看类型字节数）——为指针课埋了好底子；② 实验痕迹全注释保留，方便回看。
问题与建议（重要，逐条核对修正）：
① 第 34 行注释错误：len([]rune("中abc")) 实际输出是 4（1 个中文码点 + 3 个 ASCII），不是 6；
② 思考题②答案错误：len("中国abc") = 9（2 个中文 ×3 字节 + 3 个 ASCII），len([]rune("中国abc")) = 5（码点数）；且大小方向写反了——字节数永远 ≥ 码点数，不可能反过来；
③ 第 50 行注释 //1 不对：iota 从 0 起，mon = 0（实际运行输出 0 1 2 3 4 5 6 才算验证过）；
④ 思考题①答偏：问的是「声明了但不用为什么报错」，答的却是「用前必须声明」——正确要点：杜绝死变量，保证读到的每个变量必然有用，编译期就拦截；
⑤ x ++ 中间的空格 gofmt 会修成 x++（保存时自动格式化，以后不用手管）。
======================================
*/
