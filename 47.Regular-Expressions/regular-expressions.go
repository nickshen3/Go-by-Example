// Go offers built-in support for [regular expressions](http://en.wikipedia.org/wiki/Regular_expression).
// Here are some examples of  common regexp-related tasks
// in Go.
//Go 提供内置的正则表达式。这里是 Go 中基本的正则相关功能的例子。

package main

import "bytes"
import "fmt"
import "regexp"

func main() {

	// This tests whether a pattern matches a string.
	//这个测试一个字符串是否符合一个表达式。
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Above we used a string pattern directly, but for
	// other regexp tasks you'll need to `Compile` an
	// optimized `Regexp` struct.
	//上面我们是直接使用字符串，但是对于一些其他的正则任务，你需要 Compile 一个优化的 Regexp 结构体。
	r, _ := regexp.Compile("p([a-z]+)ch")

	// Many methods are available on these structs. Here's
	// a match test like we saw earlier.
	//这个结构体有很多方法。这里是类似我们前面看到的一个匹配测试。
	fmt.Println(r.MatchString("peach"))

	// This finds the match for the regexp.
	//这是查找匹配字符串的。
	fmt.Println(r.FindString("peach punch"))

	// This also finds the first match but returns the
	// start and end indexes for the match instead of the
	// matching text.
	//这个也是查找第一次匹配的字符串的，但是返回的匹配开始和结束位置索引，而不是匹配的内容。
	fmt.Println(r.FindStringIndex("peach punch"))

	// The `Submatch` variants include information about
	// both the whole-pattern matches and the submatches
	// within those matches. For example this will return
	// information for both `p([a-z]+)ch` and `([a-z]+)`.
	/*
	Submatch 返回完全匹配和局部匹配的字符串。例如，这里会返回 p([a-z]+)ch 和 `([a-z]+) 的信息。
	*/
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// Similarly this will return information about the
	// indexes of matches and submatches.
	//类似的，这个会返回完全匹配和局部匹配的索引位置。
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// The `All` variants of these functions apply to all
	// matches in the input, not just the first. For
	// example to find all matches for a regexp.
	//带 All 的这个函数返回所有的匹配项，而不仅仅是首次匹配项。例如查找匹配表达式的所有项。
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// These `All` variants are available for the other
	// functions we saw above as well.
	//All 同样可以对应到上面的所有函数。
	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// Providing a non-negative integer as the second
	// argument to these functions will limit the number
	// of matches.
	//这个函数提供一个正整数来限制匹配次数。
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// Our examples above had string arguments and used
	// names like `MatchString`. We can also provide
	// `[]byte` arguments and drop `String` from the
	// function name.
	/*
	上面的例子中，我们使用了字符串作为参数，并使用了如 MatchString 这样的方法。
	我们也可以提供 []byte参数并将 String 从函数命中去掉。
	*/
	fmt.Println(r.Match([]byte("peach")))

	// When creating constants with regular expressions
	// you can use the `MustCompile` variation of
	// `Compile`. A plain `Compile` won't work for
	// constants because it has 2 return values.
	//创建正则表示式常量时，可以使用 Compile 的变体MustCompile 。因为 Compile 返回两个值，不能用于常量。
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// The `regexp` package can also be used to replace
	// subsets of strings with other values.
	//regexp 包也可以用来替换部分字符串为其他值。
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// The `Func` variant allows you to transform matched
	// text with a given function.
	//Func 变量允许传递匹配内容到一个给定的函数中
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
