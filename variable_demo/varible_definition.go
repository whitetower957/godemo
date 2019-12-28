package variable_demo

import "fmt"

/*
定义变量
var关键字表示定义变量，参数名，参数类型
可以出现在包或函数级别

Java定义变量 int a = 4
*/
var a int = 4

func VarDefinition() {
	var b string
	b = "hi"
	fmt.Println(a, b)
}
