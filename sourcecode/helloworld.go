package main //Go程序都组织成包。

import "fmt"

/*import语句用于导入外部代码。
标准库中的fmt包用于格式化并输出数据。*/
func main() { //像C语言一样，main函数是程序执行的入口。
	fmt.Println("Hello world.")
}
