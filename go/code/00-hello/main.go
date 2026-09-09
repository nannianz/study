package main

import "fmt"

func main() {
	/* 这是我的第一个简单的程序 */
	fmt.Println("Hello, World!")

	var a = 1
	fmt.Println(a)
}

/*
===== 验收评价（2026-09-01 验收 · 2026-09-09 补记落档）=====
覆盖度：达标。package main / import "fmt" / func main 三件套 + Println，阶段 0 目标全部命中；额外加的 var a = 1 算第一课预热。
亮点：第一个程序就该这么小，无多余内容。
问题与建议：无。
======================================
*/
