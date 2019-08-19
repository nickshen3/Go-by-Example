// We often need our programs to perform operations on
// collections of data, like selecting all items that
// satisfy a given predicate or mapping all items to a new
// collection with a custom function.
/*
我们经常需要程序在数据集上执行操作，比如选择满足给定条件的所有项，
或者将所有的项通过一个自定义函数映射到一个新的集合上。
*/

// In some languages it's idiomatic to use [generic](http://en.wikipedia.org/wiki/Generic_programming)
// data structures and algorithms. Go does not support
// generics; in Go it's common to provide collection
// functions if and when they are specifically needed for
// your program and data types.
/*
在某些语言中，会习惯使用泛型。Go 不支持泛型，在 Go 中，当你的程序或者数据类型需要时，
通常是通过组合的方式来提供操作函数。
*/
// Here are some example collection functions for slices
// of `strings`. You can use these examples to build your
// own functions. Note that in some cases it may be
// clearest to just inline the collection-manipulating
// code directly, instead of creating and calling a
// helper function.
/*
这是一些 strings 切片的组合函数示例。你可以使用这些例子来构建自己的函数。
注意有时候，直接使用内联组合操作代码会更清晰，而不是创建并调用一个帮助函数。
*/

package main

import "strings"
import "fmt"

// Index returns the first index of the target string `t`, or
// -1 if no match is found.
//返回目标字符串 t 出现的的第一个索引位置，或者在没有匹配值时返回 -1。
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include returns `true` if the target string t is in the
// slice.
//如果目标字符串 t 在这个切片中则返回 true。
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Any returns `true` if one of the strings in the slice
// satisfies the predicate `f`.
//如果这些切片中的字符串有一个满足条件 f 则返回true。
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns `true` if all of the strings in the slice
// satisfy the predicate `f`.
//如果切片中的所有字符串都满足条件 f 则返回 true。
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter returns a new slice containing all strings in the
// slice that satisfy the predicate `f`.
//返回一个包含所有切片中满足条件 f 的字符串的新切片。
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map returns a new slice containing the results of applying
// the function `f` to each string in the original slice.
//返回一个对原始切片中所有字符串执行函数 f 后的新切片。
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {

	// Here we try out our various collection functions.
	//这里试试这些组合函数。
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	// The above examples all used anonymous functions,
	// but you can also use named functions of the correct
	// type.
	//上面的例子都是用的匿名函数，但是你也可以使用类型正确的命名函数
	fmt.Println(Map(strs, strings.ToUpper))

}
