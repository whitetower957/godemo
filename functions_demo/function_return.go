package functions_demo

/*
函数可以返回任意数量的返回值
Go 的返回值可被命名，它们会被视作定义在函数顶部的变量
没有参数的 return 语句返回已命名的返回值。也就是 直接 返回
Java不能直接返回任意数量的返回值，而是将多个返回值进行封装，再返回一个返回值
*/
func FuncReturn() (cls string, num int, stu string) {
	cls = "班上有"
	num = 6
	stu = "名同学"
	return
}
