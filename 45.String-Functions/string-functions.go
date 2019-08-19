// The standard library's `strings` package provides many
// useful string-related functions. Here are some examples
// to give you a sense of the package.
//标准库的 strings 包提供了很多有用的字符串相关的函数。这里是一些用来让你对这个包有个初步了解的例子。
package main

import s "strings"
import "fmt"

// We alias `fmt.Println` to a shorter name as we'll use
// it a lot below.
//我们给 fmt.Println 一个短名字的别名，我们随后将会经常用到。
var p = fmt.Println

func main() {

	// Here's a sample of the functions available in
	// `strings`. Since these are functions from the
	// package, not methods on the string object itself,
	// we need pass the string in question as the first
	// argument to the function. You can find more
	// functions in the [`strings`](http://golang.org/pkg/strings/)
	// package docs.
	/*
	这是一些 strings 中的函数例子。注意他们都是包中的函数，不是字符串对象自身的方法，
	这意味着我们需要考虑在调用时传递字符作为第一个参数进行传递。
	你可以在 strings包文档中找到更多的函数
	*/
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	// Not part of `strings`, but worth mentioning here, are
	// the mechanisms for getting the length of a string in
	// bytes and getting a byte by index.
	/*
	虽然不是 strings 的一部分，但是仍然值得一提的是获取字符串长度和通过索引获取一个字符的机制。
	*/
	
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}

// Note that `len` and indexing above work at the byte level.
// Go uses UTF-8 encoded strings, so this is often useful
// as-is. If you're working with potentially multi-byte
// characters you'll want to use encoding-aware operations.
// See [strings, bytes, runes and characters in Go](https://blog.golang.org/strings)
// for more information.

