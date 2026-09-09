package main

func main() {
	// var arr [3]int

	// fmt.Println(len(arr))

	// s := []int{1,2,3}
	// fmt.Println("s的长度为:", len(s))
	// s  = append(s, 4)
	// fmt.Println("s的长度为:", len(s))
	// fmt.Println("s:", s)
	// sub := s[1:3]
	// fmt.Println("sub:", sub)

	// s:= make([]int,0,3)
	// s = append(s,1)
	// fmt.Println("s1:", s)
	// fmt.Println("s1-len",len(s))
	// fmt.Println("s1-cap:", cap(s))

	// s = append(s,2)
	// fmt.Println("s2:", s)
	// fmt.Println("s2-len",len(s))
	// fmt.Println("s2-cap:", cap(s))

	// s = append(s,3)
	// fmt.Println("s3:", s)
	// fmt.Println("s3-len",len(s))
	// fmt.Println("s3-cap:", cap(s))

	// s = append(s,4)
	// fmt.Println("s4:", s)
	// fmt.Println("s4-len",len(s))
	// fmt.Println("s4-cap:", cap(s))

	// s = append(s,5)
	// fmt.Println("s5:", s)
	// fmt.Println("s5-len",len(s))
	// fmt.Println("s5-cap:", cap(s))

	// append(s,100)
	// fmt.Println("s6:", s)
	// fmt.Println("s6-len",len(s))
	// fmt.Println("s6-cap:", cap(s))

	//  s:= []int{1,2,3,4,5}
	//  sub := s[1:3]
	//  fmt.Println("sub:", sub)
	//  fmt.Println("sub-len",len(sub))
	//  fmt.Println("sub-cap:", cap(sub))
	//  sub [0] = 99
	//  fmt.Println("sub:", sub)
	//  fmt.Println("s:", s)

	//用一句话说清 JS 的 arr.slice(1,3) 与 Go 的 s[1:3] 本质区别
	//slice是新数组，不会改变原数组。s是原数组。
	growFunc()
}

/*
===== 验收评价（2026-09-07 验收 · 2026-09-09 补记落档）=====
覆盖度：4.5/5——任务 2/3/4/5 全做；任务 1 的「对数组调 append 看报错」漏做（结论：cannot use arr (variable of type [3]int) as []int value——数组不是 append 的合法入参，想补随时跑）。
亮点：见 grow.go（本课加分项全在那份超纲实验里）。
问题与建议：思考题措辞收紧——Go 的 s[1:3] 不是「原数组」，而是共享原数组底层内存的新切片（视图），sub 改的是同一块内存里的 s[1]；你顺手打印的 cap(sub)=4 是个彩蛋：截取切片的 cap = 从截取起点到底层数组尽头。
思考题：JS arr.slice(1,3) 返回元素拷贝的新数组；Go s[1:3] 返回共享底层的新切片视图——判断正确，措辞已修正。
======================================
*/
