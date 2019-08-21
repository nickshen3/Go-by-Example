// [Environment variables](http://en.wikipedia.org/wiki/Environment_variable)
// are a universal mechanism for [conveying configuration
// information to Unix programs](http://www.12factor.net/config).
// Let's look at how to set, get, and list environment variables.
//环境变量 是一个为 Unix 程序传递配置信息的普遍方式。 让我们来看看如何设置，获取并列举环境变量。

package main

import "os"
import "strings"
import "fmt"

func main() {

	// To set a key/value pair, use `os.Setenv`. To get a
	// value for a key, use `os.Getenv`. This will return
	// an empty string if the key isn't present in the
	// environment.
	//使用 os.Setenv 来设置一个键值对。使用 os.Getenv 获取一个键对应的值。如果键不存在，将会返回一个空字符 串。
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// Use `os.Environ` to list all key/value pairs in the
	// environment. This returns a slice of strings in the
	// form `KEY=value`. You can `strings.Split` them to
	// get the key and value. Here we print all the keys.
	//使用 os.Environ 来列出所有环境变量键值对。这个函数 会返回一个 KEY=value 形式的字符串切片。
	//你可以使用 strings.Split 来得到键和值。这里我们打印所有的键。
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
