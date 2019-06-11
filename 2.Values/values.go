// Go has various value types including strings,
// integers, floats, booleans, etc. Here are a few
// basic examples.
/*
Go 拥有多种值类型，包括字符串，整型，浮点型，布尔 型等。下面是一些基本的例子。
*/
package main

import "fmt"

func main() {

	// Strings, which can be added together with `+`.
	//字符串可以通过 + 连接。
	fmt.Println("go" + "lang")

	// Integers and floats.
	//整数和浮点数
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Booleans, with boolean operators as you'd expect.
	//布尔型，以及常见的布尔操作。
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
