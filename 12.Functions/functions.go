// _Functions_ are central in Go. We'll learn about
// functions with a few different examples.
//函数 是 Go 的中心。我们将通过一些不同的例子来 进行学习。
package main

import "fmt"

// Here's a function that takes two `int`s and returns
// their sum as an `int`.
//这里是一个函数，接受两个 int 并且以 int 返回它们的和
func plus(a int, b int) int {

	// Go requires explicit returns, i.e. it won't
	// automatically return the value of the last
	// expression.
	//Go 需要明确的返回，不会自动返回最 后一个表达式的值
	return a + b
}

// When you have multiple consecutive parameters of
// the same type, you may omit the type name for the
// like-typed parameters up to the final parameter that
// declares the type.
//当多个连续的参数为同样类型时，最多可以仅声明最后一个参数类型 而忽略之前相同类型参数的类型声明。
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// Call a function just as you'd expect, with
	// `name(args)`.
	//通过 name(args) 来调用函数，
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
