package main

import (
	"errors"
	"fmt"
)

// 自定义错误：type NotFoundError struct{...} 挂 Error() 方法
type NotFoundError struct {
	What string
}

func (n *NotFoundError) Error() string {
	return n.What
}

func loadFirst(test string) error {
	return &NotFoundError{What: test}
}

func loadSecond(test string) error {
	if err := loadFirst(test); err != nil {
		return fmt.Errorf("test %s：%v", test, err) //fmt.Errorf format %w has arg msg of wrong type string ???
	}
	return nil
}

func loadThird(test string) error {
	if err := loadSecond(test); err != nil {
		return fmt.Errorf("test %s：%v", test, err) //fmt.Errorf format %w has arg msg of wrong type string ???
	}
	return nil
}

func main() {

	err := loadThird("test") //test test：test test：test
	fmt.Println(err)

	var nf *NotFoundError
	fmt.Println("nf:", nf)                     //nf: <nil>
	fmt.Println("errors", errors.As(err, &nf)) //errors true
	fmt.Println("nf:", nf.What)                //nf: test

	fmt.Println("errors.Is:", errors.Is(err, &NotFoundError{})) //errors.Is: false

	//%v 打印错误信息
	// test test：test test：test
	// nf: <nil>
	// errors false
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal 0xc0000005 code=0x0 addr=0x0 pc=0x7ff7bf64afad]

	// goroutine 1 [running]:
	// main.main()
	// 	f:/MY/study/go/code/07-error/main.go:43 +0x1ad
	// exit status 2

	//为什么 Go 不把 err 设计成 throw 异常？
	//能预见到的失败不是异常，是业务流程的一部分 ，用返回值处理；只有程序员写错了的 bug 才 panic。
	//值可以被：传递、包装、组合、忽略、记录、转换。异常是 控制流机制 ，只能 throw/catch，不能像值一样操作。Go 选择把错误降级为值，换取最大的灵活性

}

/* ===== 验收评价（2026-09-28）=====
覆盖度：5/5 任务全覆盖
  ✅ ① 自定义错误：NotFoundError struct + Error() 方法
  ✅ ② 错误链：fmt.Errorf %w 包装 + errors.Is/As 双判定
  ✅ ③ 模拟分层：loadFirst → loadSecond → loadThird 三层包装，main 里 As 穿透取出 What 字段
  ✅ ④ 断链对比：%v 改跑，记录 As=false + nil 指针 panic（完整 goroutine 栈留存）
  ✅ ⑤ 思考题：为什么不用 throw——要点正确（业务流程 vs bug、值可组合 vs 控制流）

亮点：
  1. 断链实验不只记 true/false，还完整记录了 panic 现象 + goroutine 栈——调试习惯好
  2. errors.Is 和 errors.As 同时验证，直观体现"Is 比值、As 比类型"的分工
  3. 思考题用自己的话总结，非抄结论

问题与建议：
  1. 第 23、30 行过时注释 `//fmt.Errorf format %w has arg msg of wrong type string ???` 应删——当前代码用 %v 非 %w，注释从旧版本遗留
  2. Error() 返回 n.What 过于简单，fmt.Println(err) 输出只有 "test"——看不出错误类型。
     建议改为 return "not found: " + n.What，对照 back 的 ValidationError.Error() 带上下文
  3. 当前代码是 %v 断链版（go run 会 panic 退出），应恢复为 %w 版本作为最终版，%v 结果以注释保留即可

补全设计层答案（思考题）：
  并发安全——异常的栈冒泡在多 goroutine 下无法工作（每个 goroutine 有独立栈），
  error 作为返回值天然是每个 goroutine 的局部问题，不跨栈传播。
============================== */
