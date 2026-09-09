package main

import "fmt"

func main() {
	//定义 Device struct（Name / Price / Tags），三种初始化各写一遍并打印
	// type Device struct {
	// 	Name string
	// 	Price int
	// 	Tags  []string
	// }
	// //初始化
	// device := Device {
	// 	Name :"go1",
	// 	Price: 100,
	// 	Tags:  []string{"go","2"},
	// }
	// device2 := Device {"22",200,[]string{"2"}}
	// device3 := &Device {
	// 	Name: "go3",
	// 	Price: 300,
	// 	Tags:  []string{"go","3"},
	// 	}
	// fmt.Println("device:", device)//device: {go1 100 [go 2]}
	// fmt.Println("device2:", device2)//device2: {22 200 [2]}
	// fmt.Println("device3:", *device3)//device3: {go3 300 [go 3]}

	// b := device
	// b.Name = "go2"
	// fmt.Println("device:", device)//device: {go1 100 [go 2]}
	// fmt.Println("b:", b)//b: {go2 100 [go 2]}

	//tag 实验：定义带 json:"device_name" tag 的 struct，用 json.Marshal 序列化打印（import "encoding/json"），看键名按 tag 走
	// type Device struct {
	// 	Name string `json:"device_name"`
	// }
	// device := Device {
	// 	Name: "go1",
	// }
	// jsonStr, _ := json.Marshal(device)
	// fmt.Println(string(jsonStr))//{"device_name":"go1"}

	type Base struct {
		ID int `json:"id"`
	}
	
	type Device struct {
		Base
		Name string `json:"device_name"`
	}
	device := Device {
		Name: "go1",
	}
	device.ID = 1
	fmt.Println("device",device) //device {{1} go1}

	//思考题：smb 的 Site 为什么用 int64 存 CreatedAt 而不是 time.Time？
// 	a) 跨语言契约最简——前端 JS 拿到毫秒数直接 new Date(ms)；time.Time 序列化成字符串要协商格式和时区；
// b) 数据库 bigint 直接存、可比较可索引，无时区歧义；
// c) gorm 的 autoCreateTime:milli 原生按这个设计；
// d) JS 的 Date 底层本来就是毫秒数——int64 时间戳是前后端的共同语言。
}

/*
===== 验收评价（2026-09-09）=====
覆盖度：2.5/5。
✅ 任务2 值语义实验：b := device 改名不影响原值，结论与注释双正确；
✅ 任务3 tag 实验：json.Marshal 输出 {"device_name":"go1"}，键名按 tag 走验证到位；
⚠️ 任务1 三种初始化只写了字段名式（Device{Name:...}），缺 ②按顺序 Device{"go1", 100, ...} 和 ③指针式 d := &Device{...}（④ var d Device 零值也值得一试）；
⚠️ 任务4 嵌入方向做反（任务是 Base 嵌入 Device → device.ID 提升；你写的是 Device 嵌入 Base），且没演示提升访问——试试 base.Name 不带 .Device 直接访问；
❌ 任务5 思考题「int64 可以直接用」未答到点，标准答案见下。
亮点：① 每个实验注释都保留了预期输出（device 不变 / b 变），验证意识好；② 嵌套初始化用了 Device: device 字段名写法（正确姿势）；③ jsonStr, _ := json.Marshal(...) 用 _ 丢 err——与 smb 源码 sitecontroller.go 的 reqJson, _ := json.Marshal(req) 完全一致，工程上可接受。
问题与建议：① struct 定义写在 main 函数内部——合法但只能函数内使用，工程中类型一律放包级（entity/site.go 全在包级）；② 第 7 行 func 前有个多余空格（gofmt 保存时自动修，知道即可）。
思考题标准答案（为什么 int64 而不是 time.Time）：
a) 跨语言契约最简——前端 JS 拿到毫秒数直接 new Date(ms)；time.Time 序列化成字符串要协商格式和时区；
b) 数据库 bigint 直接存、可比较可索引，无时区歧义；
c) gorm 的 autoCreateTime:milli 原生按这个设计；
d) JS 的 Date 底层本来就是毫秒数——int64 时间戳是前后端的共同语言。
补做清单（补完勾 ch17）：① 三种初始化补齐各打印一次；② 改成 Base 嵌入 Device，演示 device.ID 与 device.Name 的提升访问；③ 思考题用自己的话写一遍注释。
======================================
*/

/*
===== 复验评价（2026-09-09 第二轮）=====
补做核对：3/3 ✅——
① 三种初始化齐全：顺序式 Device{"22",200,...}、指针式 &Device{...}，还顺手用 *device3 解引用打印（指针预习加分）；
② 嵌入方向已改对（Base 嵌入 Device），device.ID = 1 正是提升访问的现场演示，输出 {{1} go1} 里能看见嵌套结构；
③ 思考题答案四点齐全。
结论：通过，ch17 勾选。
遗留提醒（不阻塞）：
a) Base/Device 仍定义在 main 函数内——上轮已提、本轮未改，工程中类型放包级；05-pointer 练习请养成（下一课就是指针，包级定义 + 方法正好一起练）；
b) 思考题是标准答案誊抄，复习时试着用自己的话重述一遍，检验是否真理解；
c) gofmt 小项（逗号后空格、tab 缩进混用）保存时自动修，不用手管。
======================================
*/
