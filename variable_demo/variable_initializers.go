package variable_demo

import "fmt"

/*
变量初始化
变量声明可以包含初始值，每个变量对应一个。
如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。
:= 表示变量创建且初始化（但是此运算符不能在函数外使用）
*/
func VarInit() {
	var a string = "hello"
	var b = "world"
	c := "hi"
	fmt.Println(a, b, c)
}
