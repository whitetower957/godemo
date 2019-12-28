package packages_demo

import "fmt"

/*
方法，变量小写只能在包范围内被引用，大写才能被其他包范围引用
类比Java， 小写--默认，大写--public

在packages_demo/packages_import.go中引用正常，在/main/main.go中引用报错
*/

func pacPrint(args interface{}) {
	fmt.Println(args)
}
