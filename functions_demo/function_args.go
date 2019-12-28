package functions_demo

import "fmt"

/*
函数可以接收0或多个参数
当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略
*/
func FuncArgs(x, y int) {
	fmt.Println(x + y)
}
