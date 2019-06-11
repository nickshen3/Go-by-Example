// In Go, _variables_ are explicitly declared and used by
// the compiler to e.g. check type-correctness of function
// calls.
//在 Go 中，变量 被显式声明，并可以被编译器用来 检查函数调用时的类型正确性。

package main

import "fmt"

func main() {

	// `var` declares 1 or more variables.
	//var 声明 1 个或者多个变量。
	var a = "initial"
	fmt.Println(a)

	// You can declare multiple variables at once.
	//你可以一次性声明多个变量。
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables.
	//Go 将自动推断已经初始化的变量类型。
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding
	// initialization are _zero-valued_. For example, the
	// zero value for an `int` is `0`.
	//声明后却没有给出对应的初始值时，变量将会初始化为 零值 。例如，一个 int 的零值是 0。
	var e int
	fmt.Println(e)

	// The `:=` syntax is shorthand for declaring and
	// initializing a variable, e.g. for
	// `var f string = "apple"` in this case.
	//:= 语法是声明并初始化变量的简写，例如 这个例子中的 var f string = "short"。
	f := "apple"
	fmt.Println(f)
}
