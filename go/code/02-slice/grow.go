package main

import "fmt"

func growFunc() {
	s := make([]int, 0, 3)
	for i := 1; i <= 3; i++ {
		s = append(s, i)
		fmt.Printf("第%d次: len=%d cap=%d 底层数组地址=%p\n", i, len(s), cap(s), &s[0])
	}

	old := s // 此刻 cap 满、未搬家：old 和 s 共享同一块底层数组
	fmt.Printf("old 抓住的地址: %p（和上面三次相同）\n", &old[0])

	s = append(s, 4) // 超过 cap 3 → 搬家！
	fmt.Printf("第4次: len=%d cap=%d 底层数组地址=%p（跳变了！）\n", len(s), cap(s), &s[0])

	s[0] = 999 // 在"新房子"里改第一个元素
	fmt.Println("改 s[0]=999 后：")
	fmt.Println("  s:  ", s,   "地址:", &s[0])
	fmt.Println("  old:", old, "地址:", &old[0])
}
/*
===== 验收评价（2026-09-07 验收 · 2026-09-09 补记落档）=====
亮点（超纲加分）：用 %p 打印 &s[0] 底层数组地址，一次实验验证三件事——① cap 未满时地址不变（没搬家）；② 第 4 次 append 地址跳变（搬家）；③ 搬家后改 s[0] 不影响 old（共享断开）。old 别名手法教科书级，此文件值得复习时重读。
问题与建议：无。
======================================
*/
