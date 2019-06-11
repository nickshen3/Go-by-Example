// Go supports
// <a href="http://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>recursive functions</em></a>.
// Here's a classic factorial example.
//Go 支持 递归。 这里是一个经典的阶乘示例。
package main

import "fmt"

// This `fact` function calls itself until it reaches the
// base case of `fact(0)`.
//fact 函数在到达 fact(0) 前一直调用自身。
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
