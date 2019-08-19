// Go's `sort` package implements sorting for builtins
// and user-defined types. We'll look at sorting for
// builtins first.
/*
Go 的 sort 包实现了内置和用户自定义数据类型的排序功能。我们首先关注内置数据类型的排序。
*/

package main

import "fmt"
import "sort"

func main() {

	// Sort methods are specific to the builtin type;
	// here's an example for strings. Note that sorting is
	// in-place, so it changes the given slice and doesn't
	// return a new one.
	/*
	排序方法是正对内置数据类型的；这里是一个字符串的例子。注意排序是原地更新的，
	所以他会改变给定的序列并且不返回一个新值。
	*/
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// An example of sorting `int`s.
	//一个 int 排序的例子。
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// We can also use `sort` to check if a slice is
	// already in sorted order.
	//我们也可以使用 sort 来检查一个序列是不是已经是排好序的。
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
/*
运行程序，打印排序好的字符串和整形序列以及我们 AreSorted测试的结构 true。
*/