// Sometimes we'll want to sort a collection by something
// other than its natural order. For example, suppose we
// wanted to sort strings by their length instead of
// alphabetically. Here's an example of custom sorts
// in Go.
/*
有时候我们想使用和集合的自然排序不同的方法对集合进行排序。
例如，我们想按照字母的长度而不是首字母顺序对字符串排序。
这里是一个 Go 自定义排序的例子。
*/

package main

import "sort"
import "fmt"

// In order to sort by a custom function in Go, we need a
// corresponding type. Here we've created a `byLength`
// type that is just an alias for the builtin `[]string`
// type.
//为了在 Go 中使用自定义函数进行排序，我们需要一个对应的类型。
//这里我们创建一个为内置 []string 类型的别名的ByLength 类型，
type byLength []string

// We implement `sort.Interface` - `Len`, `Less`, and
// `Swap` - on our type so we can use the `sort` package's
// generic `Sort` function. `Len` and `Swap`
// will usually be similar across types and `Less` will
// hold the actual custom sorting logic. In our case we
// want to sort in order of increasing string length, so
// we use `len(s[i])` and `len(s[j])` here.
/*
我们在类型中实现了 sort.Interface 的 Len，Less和 Swap 方法，
这样我们就可以使用 sort 包的通用Sort 方法了，Len 和 Swap 通常在各个类型中都差不多，
Less 将控制实际的自定义排序逻辑。在我们的例子中，我们想按字符串长度增加的顺序来排序，
所以这里使用了 len(s[i]) 和 len(s[j])。
*/
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// With all of this in place, we can now implement our
// custom sort by converting the original `fruits` slice
// to `byLength`, and then use `sort.Sort` on that typed
// slice.
/*
一切都准备好了，我们现在可以通过将原始的 fruits 切片转型成 ByLength 来实现我们的自定排序了。
然后对这个转型的切片使用 sort.Sort 方法。
*/
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
/*
运行这个程序，和预期的一样，显示了一个按照字符串长度排序的列表。
类似的，参照这个创建一个自定义类型的方法，
实现这个类型的这三个接口方法，然后在一个这个自定义类型的集合上调用 sort.Sort 方法，
我们就可以使用任意的函数来排序 Go 切片了。
*/