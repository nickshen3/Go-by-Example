// Go supports [_anonymous functions_](http://en.wikipedia.org/wiki/Anonymous_function),
// which can form <a href="http://en.wikipedia.org/wiki/Closure_(computer_science)"><em>closures</em></a>.
// Anonymous functions are useful when you want to define
// a function inline without having to name it.
/*
Go 支持匿名函数，并能用其构造 闭包。 匿名函数在你想定义一个不需要命名的内联函数时是很实用的。
*/
package main

import "fmt"

// This function `intSeq` returns another function, which
// we define anonymously in the body of `intSeq`. The
// returned function _closes over_ the variable `i` to
// form a closure.
/*
这个 intSeq 函数返回另一个在 intSeq 函数体内定义的 匿名函数。
这个返回的函数使用闭包的方式 隐藏 变量 i。
*/
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// We call `intSeq`, assigning the result (a function)
	// to `nextInt`. This function value captures its
	// own `i` value, which will be updated each time
	// we call `nextInt`.
	/*
	我们调用 intSeq 函数，将返回值（一个函数）赋给 nextInt。
	这个函数的值包含了自己的值 i，这样在每 次调用 nextInt 时都会更新 i 的值。
	*/
	nextInt := intSeq()

	// See the effect of the closure by calling `nextInt`
	// a few times.
	//通过多次调用 nextInt 来看看闭包的效果。
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// To confirm that the state is unique to that
	// particular function, create and test a new one.
	//为了确认这个状态对于这个特定的函数是唯一的，我们 重新创建并测试一下。
	newInts := intSeq()
	fmt.Println(newInts())
}
