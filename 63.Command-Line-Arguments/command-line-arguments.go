// [_Command-line arguments_](http://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// are a common way to parameterize execution of programs.
// For example, `go run hello.go` uses `run` and
// `hello.go` arguments to the `go` program.
//命令行参数 是指定程序运行参数的一个常见方式。例如，go run hello.go， 
//程序 go 使用了 run 和 hello.go 两个参数。
package main

import "os"
import "fmt"

func main() {

	// `os.Args` provides access to raw command-line
	// arguments. Note that the first value in this slice
	// is the path to the program, and `os.Args[1:]`
	// holds the arguments to the program.
	//os.Args 提供原始命令行参数访问功能。注意，切片中 的第一个参数是该程序的路径，
	//并且 os.Args[1:]保存 所有程序的参数。
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// You can get individual args with normal indexing.
	//你可以使用标准的索引位置方式取得单个参数的值。
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
