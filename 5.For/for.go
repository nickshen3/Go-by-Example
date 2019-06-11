// `for` is Go's only looping construct. Here are
// three basic types of `for` loops.
//for 是 Go 中唯一的循环结构。这里有 for 循环 的三个基本使用方式。
package main

import "fmt"

func main() {

	// The most basic type, with a single condition.
	//最基础的方式，单个循环条件。
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after `for` loop.
	//经典的初始/条件/后续 for 循环。
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// `for` without a condition will loop repeatedly
	// until you `break` out of the loop or `return` from
	// the enclosing function.
	//不带条件的 for 循环将一直重复执行，直到在循环体内使用 了 break 或者 return 来跳出循环。
	for {
		fmt.Println("loop")
		break
	}


	// You can also `continue` to the next iteration of
	// the loop.
	//你也可以使用 continue 来跳到下一个循环迭代
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
