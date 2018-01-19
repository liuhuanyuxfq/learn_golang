# hello-world
按照惯例，学习一门语言都要从hello-world开始，那么我也不要打破惯例了。
## 1.编辑源代码helloworld.go
首先选定你的项目目录，项目目录一般是在`GOPATH/src/.../<porject_name>`下。

我是windows环境，这个例子的helloworld.go是在`%GOPATH%\src\github.com\liuhuanyuxfq\learn_golang\sourcecode`下创建的。

下面是helloworld.go的代码：

	package main //Go程序都组织成包。
	
	import "fmt"
	
	/*import语句用于导入外部代码。
	标准库中的fmt包用于格式化并输出数据。*/
	func main() { //像C语言一样，main函数是程序执行的入口。
		fmt.Println("Hello world.")
	}
让我们简单解释一下各部分：

- 第一行`package main`表明这是一个可独立执行的程序；`//Go程序都组织成包。`是注释，Go语言的单行注释使用`//`，多行注释使用`/*...*/`。
- 下一行` import "fmt"` 告诉Go编译器这个程序需要使用fmt包（的函数，或其他元素），fmt包实现了格式化IO（输入/输出）的函数。
- 下一行`func main()`是程序执行的入口。main函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。
- 下一行`fmt.Println("Hello world.")`将字符串打印至标准输出，并在最后自动增加换行字符。
## 2.编译执行
- 首先，Go提供了一个用来格式化代码的工具fmt，为了保持代码格式化风格的一致性，我们先采用该命令来格式化代码：

		$ go fmt helloworld.go
- 然后，编译这个程序，生产可执行的二进制文件helloworld.exe（如果是linux环境，生成的是helloworld）：

		$ go build helloworld.go
- 运行可执行文件：

		$ ./helloworld.exe
		Hello world.     //这是运行结果
- 如果你只是想临时运行，看看程序执行效果而不产生可执行文件，可以用下面的命令：

		$ go run helloworld.go