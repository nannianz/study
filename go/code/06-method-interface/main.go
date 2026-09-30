package main

import (
	"fmt"
)

// type Device struct {
// 	Name string
// 	Base
// }

// type Base struct {
// 	ID int
// }

// func (d Device) String() string {
// 	return d.Name
// }

// func (d Device) Rename(name string) {
// 	d.Name = name
// }
// func (d *Device) Rename2(name string) {
// 	d.Name = name
// }

// func (b Base) Describe() string {
// 	return fmt.Sprintf("ID: %d", b.ID)
// }

// type Describer interface {
// 	Describe() string
// }

type Device struct {
	Name string
}

// func (d Device) Describe() string {
// 	return "设备" + d.Name
// }

// type Org struct {
// 	City string
// }

// func (o Org) Describe() string {
// 	return "组织" + o.City
// }

// type Temp int

// func (t Temp) Describe() string {
// 	return fmt.Sprintf("温度: %d", int(t))
// }

// func (d Device) Describe() string {
// 	return fmt.Sprintf("设备: %s", d.Name)
// }
// func (o *Org) Describe() string {
// 	return fmt.Sprintf("组织: %s", o.City)
// }

func show(v any) {
	if s, ok := v.(string); ok {
		fmt.Println("字符串:", s, ok)
	}
	if s, ok := v.(int); ok {
		fmt.Println("整数:", s, ok)
	}
	if s, ok := v.(float64); ok {
		fmt.Println("浮点数:", s, ok)
	}
	if s, ok := v.(bool); ok {
		fmt.Println("布尔值:", s, ok)
	}
	if s, ok := v.(Device); ok {
		fmt.Println("设备:", s, ok)
	}

}
func describe(v any) {
	switch x := v.(type) {
	case string:
		fmt.Println("字符串:", x)
	case int:
		fmt.Println("整数:", x)
	case float64:
		fmt.Println("浮点数:", x)
	case bool:
		fmt.Println("布尔值:", x)
	case Device:
		fmt.Println("设备:", x)
	default:
		fmt.Println("未知类型")
	}
}

func main() {
	// d := Device{Name: "phone"}
	// fmt.Println(d)          // phone
	// fmt.Println(d.String()) // phone

	// d.Rename("new phone")
	// fmt.Println("Rename(d):", d) //  phone
	// d.Rename2("new phone 2")
	// fmt.Println("Rename2(d):", d) // new phone 2

	// d := Device{Name: "phone", Base: Base{ID: 100}}
	// fmt.Println(d)            // phone
	// fmt.Println(d.Describe()) // ID: 100

	// d := Device{Name: "phone"}
	// fmt.Println("设备描述:", d, d.Describe())

	// o := Org{City: "org"}
	// fmt.Println("组织描述:", o, o.Describe())

	// list := []Describer{
	// 	Device{Name: "11"},
	// 	Org{City: "org"},
	// 	Temp(20),
	// }
	// for _, v := range list {
	// 	fmt.Println(v.Describe())
	// }

	// var x Describer
	// x = Device{Name: "phone"}
	// fmt.Println("x", x) //x {phone}
	// x = &Device{Name: "phone2"}
	// fmt.Println("x1", x) //x1 &{phone2}

	// //x = Org{City: "org"} // cannot use Org{…} (value of struct type Org) as Describer value in assignment: Org does not implement Describer (method Describe has pointer receiver)
	// // fmt.Println("x2", x)
	// x = &Org{City: "org2"}
	// fmt.Println("x3", x) //x3 &{org2}

	// show("hello") //字符串: hello true
	// show(123)     //整数: 123 true
	// show(123.456) //浮点数: 123.456 true
	// show(true)    //布尔值: true true

	arr := []any{"hello", 123, 123.456, true, nil, Device{Name: "111"}}

	for k, v := range arr {
		fmt.Println("索引:", k, "值:", v)
		show(v)
		describe(v)
	}
	// 索引: 0 值: hello
	// 字符串: hello
	// 索引: 1 值: 123
	// 整数: 123
	// 索引: 2 值: 123.456
	// 浮点数: 123.456
	// 索引: 3 值: true
	// 布尔值: true
	// 索引: 4 值: <nil>
	// 未知类型
}

/*
===== 验收评价（方法部分 2026-09-14 · 接口部分 2026-09-16 补记合并）=====
方法部分 3/3 ✅：String() 自动调用（Println(d) 直出 phone）、Rename 值/指针双版本（改副本 vs 改原值，输出全留）、嵌入提升（d.Describe() 直达 Base 方法）。
接口部分 4/4 ✅：Describe 接口 + Device/Org 双实现、[]Describer 混装多态 range、类型断言 show()、类型 switch describe()（live code，nil 落 default 分支的边界用例也覆盖）。
亮点：
① 123~132 行是本卷黄金：亲手撞出并原样保留「Org does not implement Describe (method Describe has pointer receiver)」——方法集规则：方法挂在 *Org（指针接收者）上时，Org 值不满足接口、只有 *Org 满足；值接收者则值和指针都满足。Go 接口第一经典坑，被你用实验钉死；
② Temp int（int 别名）也实现接口装进 []Describer——吃透了「任意类型都能挂方法」；
③ nil 塞进 []any 验证 default 分支——边界意识在线；
④ 文件夹拼写自己修正为 06-method-interface ✓。
问题与建议：
① 接口名 Describe 与方法名 Describe 同名——合法但易混，惯例 -er 后缀（Describer、Stringer、Controller 都是这条路数）【已改：同日应用户委托同步改为 Describer】；
② 任务 3 断言对象原要求含 Device，用 string/int/float64/bool 替代——机制等价，补一行 v.(Device) 更完整（可选）。
结论：通过，方法/接口/类型断言三项勾选；ch30「嵌入 interface」5 分钟示例下节课开场补。
======================================
*/
