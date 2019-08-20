// Parsing numbers from strings is a basic but common task
// in many programs; here's how to do it in Go.
//从字符串中解析数字在很多程序中是一个基础常见的任务，在Go 中是这样处理的
package main

// The built-in package `strconv` provides the number
// parsing.
//内置的 strconv 包提供了数字解析功能。
import "strconv"
import "fmt"

func main() {

	// With `ParseFloat`, this `64` tells how many bits of
	// precision to parse.
	//使用 ParseFloat 解析浮点数，这里的 64 表示表示解析的数的位数。
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// For `ParseInt`, the `0` means infer the base from
	// the string. `64` requires that the result fit in 64
	// bits.
	//在使用 ParseInt 解析整形数时，例子中的参数 0 表示自动推断字符串所表示的数字的进制。
	//64 表示返回的整形数是以 64 位存储的。
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// `ParseInt` will recognize hex-formatted numbers.
	//ParseInt 会自动识别出十六进制数
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// A `ParseUint` is also available.
	//ParseInt 会自动识别出十六进制数
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// `Atoi` is a convenience function for basic base-10
	// `int` parsing.
	//Atoi 是一个基础的 10 进制整型数转换函数。
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Parse functions return an error on bad input.
	//在输入错误时，解析函数会返回一个错误
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
