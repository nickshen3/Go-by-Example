// In Go, an _array_ is a numbered sequence of elements of a
// specific length.
//在 Go 中，数组 是一个具有固定长度且编号的元素序列。
package main

import "fmt"

func main() {

	// Here we create an array `a` that will hold exactly
	// 5 `int`s. The type of elements and length are both
	// part of the array's type. By default an array is
	// zero-valued, which for `int`s means `0`s.
	/*
	这里我们创建了一个数组 a 来存放刚好 5 个 int。 元素的类型和长度都是数组类型的一部分。
	数组默认是 零值的，对于 int 数组来说也就是 0。
	*/
	var a [5]int
	fmt.Println("emp:", a)

	// We can set a value at an index using the
	// `array[index] = value` syntax, and get a value with
	// `array[index]`.
	//我们可以使用 array[index] = value 语法来设置数组 指定位置的值，或者用 array[index] 得到值。
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// The builtin `len` returns the length of an array.
	//使用内置函数 len 返回数组的长度。
	fmt.Println("len:", len(a))

	// Use this syntax to declare and initialize an array
	// in one line.
	//使用这个语法在一行内声明并初始化一个数组。
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Array types are one-dimensional, but you can
	// compose types to build multi-dimensional data
	// structures.
	//数组类型是一维的，但是你可以组合 构造多维的数据结构。
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
