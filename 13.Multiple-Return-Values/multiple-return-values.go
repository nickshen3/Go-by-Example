// Go has built-in support for _multiple return values_.
// This feature is used often in idiomatic Go, for example
// to return both result and error values from a function.
/*
Go 内建_多返回值_支持。这个特性在 Go 语言中经常被用到， 
例如用来同时返回一个函数的结果和错误信息。
*/
package main

import "fmt"

// The `(int, int)` in this function signature shows that
// the function returns 2 `int`s.
//(int, int) 在这个函数中标志着这个函数返回 2 个 int。
func vals() (int, int) {
	return 3, 7
}

func main() {

	// Here we use the 2 different return values from the
	// call with _multiple assignment_.
	//这里我们通过_多赋值_操作来使用这两个不同的返回值。
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// If you only want a subset of the returned values,
	// use the blank identifier `_`.
	//如果你仅仅需要返回值的一部分的话，你可以使用空白标识符_。
	_, c := vals()
	fmt.Println(c)
}
