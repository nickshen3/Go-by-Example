// Go supports _constants_ of character, string, boolean,
// and numeric values.
//Go 支持字符、字符串、布尔和数值 常量 。
package main

import "fmt"
import "math"

// `const` declares a constant value.
//const 用于声明一个常量。
const s string = "constant"

func main() {
	fmt.Println(s)

	// A `const` statement can appear anywhere a `var`
	// statement can.
	//const 语句可以出现在任何 var 语句可以出现 的地方
	const n = 500000000

	// Constant expressions perform arithmetic with
	// arbitrary precision.
	//常数表达式可以执行任意精度的运算
	const d = 3e20 / n
	fmt.Println(d)

	// A numeric constant has no type until it's given
	// one, such as by an explicit conversion.
	//数值型常量没有确定的类型，直到被给定 ，比如一次显示的类型转化。
	fmt.Println(int64(d))

	// A number can be given a type by using it in a
	// context that requires one, such as a variable
	// assignment or function call. For example, here
	// `math.Sin` expects a `float64`.
	/*
	当上下文需要时，比如变量赋值或者函数调用， 一个数可以被给定一个类型。
	举个例子，这里的 math.Sin 函数需要一个 float64 的参数。
	*/
	fmt.Println(math.Sin(n))
}
